package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// SplitPDF uses qpdf to split a PDF into single-page PDFs inside outDir.
// It writes files like out-001.pdf, out-002.pdf, ...
func SplitPDF(ctx context.Context, inPath, outDir string) error {
	// qpdf --split-pages=1 in.pdf out-%d.pdf
	outPattern := filepath.Join(outDir, "page-%d.pdf")
	args := []string{"--split-pages=1", inPath, outPattern}
	cmd := exec.CommandContext(ctx, "qpdf", args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("qpdf split failed: %v\n%s", err, string(b))
	}
	// basic sanity: ensure at least one page file exists
	matches, _ := filepath.Glob(filepath.Join(outDir, "page-*.pdf"))
	if len(matches) == 0 {
		return errors.New("split produced no pages")
	}
	return nil
}

// ExtractPages uses qpdf to extract page ranges into a new PDF.
// ranges example: "1-3,5,7-"
func ExtractPages(ctx context.Context, inPath, ranges, outPath string) error {
	// qpdf in.pdf --pages . 1-3,5,7- -- out.pdf
	args := []string{inPath, "--pages", ".", ranges, "--", outPath}
	cmd := exec.CommandContext(ctx, "qpdf", args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("qpdf extract failed: %v\n%s", err, string(b))
	}
	return ensureNonEmpty(outPath, "extract")
}

// ReorderPages uses qpdf with an order string to reorder/duplicate/remove pages.
// order example: "3,1,1,4-7"
func ReorderPages(ctx context.Context, inPath, order, outPath string) error {
	// qpdf in.pdf --pages . 3,1,1,4-7 -- out.pdf
	args := []string{inPath, "--pages", ".", order, "--", outPath}
	cmd := exec.CommandContext(ctx, "qpdf", args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("qpdf reorder failed: %v\n%s", err, string(b))
	}
	return ensureNonEmpty(outPath, "reorder")
}

// RotatePages uses qpdf to rotate selected pages.
// angle: 90|180|270 or prefixed with +/-, pages: "all" or ranges like "1,3-5"
func RotatePages(ctx context.Context, inPath, angle, pages, outPath string) error {
	// Normalize angle
	ang := strings.TrimSpace(angle)

	// Apply De Morgan's law to avoid negated OR-chain (QF1001)
	if ang != "90" && ang != "180" && ang != "270" &&
		!strings.HasPrefix(ang, "+") && !strings.HasPrefix(ang, "-") {
		return errors.New("invalid angle: use 90|180|270 or +90/-90")
	}

	if pages == "all" || pages == "" {
		pages = "1-z"
	}

	// qpdf --rotate=+90:1-3 in.pdf out.pdf
	opt := fmt.Sprintf("--rotate=%s:%s", ang, pages)
	cmd := exec.CommandContext(ctx, "qpdf", opt, inPath, outPath)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("qpdf rotate failed: %v\n%s", err, string(b))
	}
	return ensureNonEmpty(outPath, "rotate")
}

// ensureNonEmpty ensures a non-empty output file exists.
func ensureNonEmpty(outPath, op string) error {
	st, err := os.Stat(outPath)
	if err != nil || st.Size() == 0 {
		if err == nil {
			return fmt.Errorf("%s produced empty output", op)
		}
		return fmt.Errorf("%s failed: %v", op, err)
	}
	return nil
}
