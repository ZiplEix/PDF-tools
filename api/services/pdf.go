package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// MergePDF merge multiple PDF files into a single PDF using qpdf.
// inputPaths: paths of input PDFs (order preserved)
// outputPath: path of the output PDF
func MergePDF(ctx context.Context, inputPaths []string, outputPath string) error {
	if len(inputPaths) < 2 {
		return errors.New("au moins 2 fichiers nécessaires")
	}

	// qpdf --empty --pages in1.pdf in2.pdf ... -- out.pdf
	args := []string{"--empty", "--pages"}
	args = append(args, inputPaths...)
	args = append(args, "--", outputPath)

	cmd := exec.CommandContext(ctx, "qpdf", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("qpdf a échoué: %v\n%s", err, string(out))
	}

	// Sanity check
	stat, err := os.Stat(outputPath)
	if err != nil || stat.Size() == 0 {
		return errors.New("fusion échouée: sortie absente ou vide")
	}
	return nil
}
