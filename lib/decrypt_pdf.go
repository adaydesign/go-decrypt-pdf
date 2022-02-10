package lib

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	unipdf "github.com/unidoc/unipdf/v3/model"
)

func PdfCpuDecryptPassword(input string, password string, output string) error {
	conf := pdfcpu.NewAESConfiguration(password, "", 256)
	err := api.DecryptFile(input, output, conf)

	if err != nil {
		return err
	}

	return nil
}

func PdfCpuDecryptPasswordByCLI(pdfcpu string, input string, password string, output string) error {
	arg0 := "decrypt"
	arg1 := "-opw"

	cmd := exec.Command(pdfcpu, arg0, arg1, password, input, output)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func QPdfDecryptPasswordByCLI(qpdf string, input string, password string, output string) error {
	arg0 := "--password=" + password
	arg1 := "--decrypt"

	cmd := exec.Command(qpdf, arg0, arg1, input, output)
	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func UniPdfDecryptPassword(input string, password string, output string) error {
	pdfWriter := unipdf.NewPdfWriter()

	f, err := os.Open(input)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := unipdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	// Try decrypting both with given password and an empty one if that fails.
	if isEncrypted {
		auth, err := pdfReader.Decrypt([]byte(password))
		if err != nil {
			return err
		}
		if !auth {
			return fmt.Errorf("Wrong password")
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		err = pdfWriter.AddPage(page)
		if err != nil {
			return err
		}
	}

	fWrite, err := os.Create(output)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}
