package controllers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ZiplEix/PDF-tools/services"
	"github.com/ZiplEix/PDF-tools/utils"
	"github.com/labstack/echo/v4"
)

const (
	maxFilesAllowed = 50
)

func MergePDF(c echo.Context) error {
	// check multipart
	form, err := c.MultipartForm()
	if err != nil || form == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "multipart/form-data attendu")
	}

	files := form.File["files"]
	if len(files) < 2 {
		return echo.NewHTTPError(http.StatusBadRequest, "fournir au moins 2 fichiers via files[]")
	}
	if len(files) > maxFilesAllowed {
		return echo.NewHTTPError(http.StatusRequestEntityTooLarge, "trop de fichiers")
	}

	// tmp folder for each request
	jobID := utils.ShortID()
	tmpDir, err := os.MkdirTemp("", "merge-"+jobID+"-")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "erreur dossier temporaire")
	}
	defer os.RemoveAll(tmpDir)

	// save and validate PDF files
	var paths []string
	for i, fh := range files {
		if err := quickValidatePDFHeader(fh); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "fichier invalide ou non-PDF: "+fh.Filename)
		}
		dst := filepath.Join(tmpDir, utils.SafeName(i, fh.Filename))
		if err := utils.SaveUploadedFile(fh, dst); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "échec sauvegarde: "+fh.Filename)
		}
		paths = append(paths, dst)
	}

	outPath := filepath.Join(tmpDir, "merged.pdf")

	// server timeout
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	// merge
	if err := services.MergePDF(ctx, paths, outPath); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// downlaod
	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="merged.pdf"`)
	return c.File(outPath)
}

// SplitPDF splits a PDF into single-page PDFs and returns a ZIP archive.
func SplitPDF(c echo.Context) error {
	jobID := utils.ShortID()
	tmpDir, err := os.MkdirTemp("", "split-"+jobID+"-")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "temp dir error")
	}
	defer os.RemoveAll(tmpDir)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	inPath, err := readSinglePDFFromMultipart(c, tmpDir)
	if err != nil {
		return err
	}

	outDir := filepath.Join(tmpDir, "pages")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "temp dir error")
	}

	// Delegate to service
	if err := services.SplitPDF(ctx, inPath, outDir); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Zip all page PDFs
	zipPath := filepath.Join(tmpDir, "split.zip")
	if err := utils.ZipDirOfPDFs(outDir, zipPath); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "zip error")
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/zip")
	c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="split.zip"`)
	return c.File(zipPath)
}

// ExtractPages extracts specified ranges (e.g., "1-3,5,7-") into a new PDF.
func ExtractPages(c echo.Context) error {
	ranges := c.QueryParam("ranges")
	if strings.TrimSpace(ranges) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "query param 'ranges' required, e.g. 1-3,5,7-")
	}

	jobID := utils.ShortID()
	tmpDir, err := os.MkdirTemp("", "extract-"+jobID+"-")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "temp dir error")
	}
	defer os.RemoveAll(tmpDir)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	inPath, err := readSinglePDFFromMultipart(c, tmpDir)
	if err != nil {
		return err
	}

	outPath := filepath.Join(tmpDir, "extracted.pdf")
	if err := services.ExtractPages(ctx, inPath, ranges, outPath); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="extracted.pdf"`)
	return c.File(outPath)
}

// ReorderPages reorders/duplicates/deletes pages using an order string like "3,1,1,4-7".
func ReorderPages(c echo.Context) error {
	order := c.QueryParam("order")
	if strings.TrimSpace(order) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "query param 'order' required, e.g. 3,1,1,4-7")
	}

	jobID := utils.ShortID()
	tmpDir, err := os.MkdirTemp("", "reorder-"+jobID+"-")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "temp dir error")
	}
	defer os.RemoveAll(tmpDir)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	inPath, err := readSinglePDFFromMultipart(c, tmpDir)
	if err != nil {
		return err
	}

	outPath := filepath.Join(tmpDir, "reordered.pdf")
	if err := services.ReorderPages(ctx, inPath, order, outPath); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="reordered.pdf"`)
	return c.File(outPath)
}

// RotatePages rotates selected pages by angle (90, 180, 270). pages="all" or ranges like "1,3-5".
func RotatePages(c echo.Context) error {
	angle := strings.TrimSpace(c.QueryParam("angle"))
	if angle == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "query param 'angle' required: 90|180|270 or +90/-90")
	}
	pages := strings.TrimSpace(c.QueryParam("pages"))
	if pages == "" {
		pages = "all"
	}

	jobID := utils.ShortID()
	tmpDir, err := os.MkdirTemp("", "rotate-"+jobID+"-")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "temp dir error")
	}
	defer os.RemoveAll(tmpDir)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Minute)
	defer cancel()

	inPath, err := readSinglePDFFromMultipart(c, tmpDir)
	if err != nil {
		return err
	}

	outPath := filepath.Join(tmpDir, "rotated.pdf")
	if err := services.RotatePages(ctx, inPath, angle, pages, outPath); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="rotated.pdf"`)
	return c.File(outPath)
}
