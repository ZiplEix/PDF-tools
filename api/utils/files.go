package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func SaveUploadedFile(fh *multipart.FileHeader, dst string) error {
	src, err := fh.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func SafeName(idx int, name string) string {
	base := filepath.Base(name)
	// supprime chemins et espaces, force .pdf
	base = strings.ReplaceAll(base, " ", "_")
	if !strings.HasSuffix(strings.ToLower(base), ".pdf") {
		base += ".pdf"
	}
	return strings.TrimLeft(filepath.Clean(
		filepath.Join("", base),
	), "/\\")
}
