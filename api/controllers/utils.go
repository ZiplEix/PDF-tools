package controllers

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/ZiplEix/PDF-tools/utils"
	"github.com/labstack/echo/v4"
)

// quickValidatePDFHeader reads a few bytes to check for the %PDF- signature
// it does not guarantee the file is a valid PDF, just that it looks like one
func quickValidatePDFHeader(fh *multipart.FileHeader) error {
	f, err := fh.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 5)
	n, _ := io.ReadFull(f, buf)
	if n < 5 {
		return errors.New("fichier trop court")
	}
	if string(buf) != "%PDF-" {
		return errors.New("signature PDF manquante")
	}
	return nil
}

// readSinglePDFFromMultipart reads a single 'file' field, quick-checks %PDF- header, and saves it into tmpDir.
func readSinglePDFFromMultipart(c echo.Context, tmpDir string) (string, error) {
	fh, err := c.FormFile("file")
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "multipart/form-data with field 'file' required")
	}
	if err := quickValidatePDFHeader(fh); err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "invalid or non-PDF file: "+fh.Filename)
	}
	inPath := filepath.Join(tmpDir, utils.SanitizeName(fh.Filename))
	if err := utils.SaveUploadedFile(fh, inPath); err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "failed to save uploaded file")
	}
	return inPath, nil
}
