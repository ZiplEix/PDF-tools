package controllers

import (
	"errors"
	"io"
	"mime/multipart"
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
