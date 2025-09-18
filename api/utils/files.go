package utils

import (
	"archive/zip"
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

func SanitizeName(name string) string {
	base := filepath.Base(name)
	base = strings.ReplaceAll(base, " ", "_")
	if !strings.HasSuffix(strings.ToLower(base), ".pdf") {
		base += ".pdf"
	}
	return base
}

// zipDirOfPDFs zips all *.pdf in dir into zipPath.
func ZipDirOfPDFs(dir, zipPath string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	out, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()

	for _, e := range files {
		if e.IsDir() || !strings.HasSuffix(strings.ToLower(e.Name()), ".pdf") {
			continue
		}
		fp := filepath.Join(dir, e.Name())
		fw, err := w.Create(e.Name())
		if err != nil {
			return err
		}
		in, err := os.Open(fp)
		if err != nil {
			return err
		}
		if _, err := io.Copy(fw, in); err != nil {
			in.Close()
			return err
		}
		in.Close()
	}
	return nil
}
