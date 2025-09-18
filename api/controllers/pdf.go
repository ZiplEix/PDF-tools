package controllers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
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
