package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
)

func Untar(src, dest string) error {
	if strings.HasSuffix(src, ".tar.gz") {
		return UntarGz(src, dest)
	} else if strings.HasSuffix(src, ".tar.xz") {
		return UntarXz(src, dest)
	}
	return fmt.Errorf("unsupported archive format: %s", src)
}

func UntarGz(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open tar.gz source %s: %w", src, err)
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("could not create gzip reader: %w", err)
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	return untar(tr, dest)
}

func UntarXz(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open tar.xz source %s: %w", src, err)
	}
	defer file.Close()

	xzr, err := xz.NewReader(file)
	if err != nil {
		return fmt.Errorf("could not create xz reader: %w", err)
	}

	tr := tar.NewReader(xzr)

	return untar(tr, dest)
}

func untar(tr *tar.Reader, dest string) error {
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading next tar header: %w", err)
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return fmt.Errorf("could not create directory %s: %w", target, err)
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return fmt.Errorf("could not create parent directory for %s: %w", target, err)
			}
			outFile, err := os.Create(target)
			if err != nil {
				return fmt.Errorf("could not create file %s: %w", target, err)
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return fmt.Errorf("could not write to file %s: %w", target, err)
			}
			outFile.Close()
		default:
			log.Printf("unsupported type: %s", header.Name)
		}
	}
	fmt.Println("Successfully extracted tar archive.")
	return nil
}
