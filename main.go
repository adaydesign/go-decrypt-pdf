package main

import (
	"fmt"
	"os"
	"time"

	"github.com/adaydesign/go-decrypt-pdf/lib"
	"github.com/unidoc/unipdf/v3/common/license"
)

func init() {
	// To get your free API key for metered license, sign up on: https://cloud.unidoc.io
	// Make sure to be using UniPDF v3.19.1 or newer for Metered API key support.
	err := license.SetMeteredKey(`### Your UniPDF Key Here ###`)
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		panic(err)
	}
}

func main() {

	input := "files/locked_doc.pdf"
	password := "@1234"
	output1 := "files/out_pdfcpu_pkg.pdf"
	output2 := "files/out_pdfcpu.pdf"
	output3 := "files/out_qpdfcpu.pdf"
	output4 := "files/out_unipdf_pkg.pdf"

	// 1. use pdfcpu package
	start := time.Now()
	err := lib.PdfCpuDecryptPassword(input, password, output1)
	if err != nil {
		fmt.Println("PdfCpuDecryptPassword - fail")
		fmt.Println(err)
	} else {
		fmt.Println("PdfCpuDecryptPassword - success")
	}
	elapse := time.Since(start)
	fmt.Println(">> PDFCPU (pkg) took ", elapse)

	// 2. use pdfcpu (run excute file)
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	pdfcpu := path + "/bin/pdfcpu_0.3.13/pdfcpu.exe"

	start = time.Now()
	err = lib.PdfCpuDecryptPasswordByCLI(pdfcpu, input, password, output2)
	if err != nil {
		fmt.Println("PdfCpuDecryptPasswordByCLI - fail")
		fmt.Println(err)
	} else {
		fmt.Println("PdfCpuDecryptPasswordByCLI - success")
	}
	elapse = time.Since(start)
	fmt.Println(">> PDFCPU took ", elapse)

	// 3. use qpdf (run excute file)
	qpdf := path + "/bin/qpdf-10.5.0/bin/qpdf.exe"

	start = time.Now()
	err = lib.QPdfDecryptPasswordByCLI(qpdf, input, password, output3)
	if err != nil {
		fmt.Println("QPdfDecryptPasswordByCLI - fail")
		fmt.Println(err)
	} else {
		fmt.Println("QPdfDecryptPasswordByCLI - success")
	}
	elapse = time.Since(start)
	fmt.Println(">> QPDF took ", elapse)

	// 4. use UniPDF
	lk := license.GetLicenseKey()
	if lk == nil {
		fmt.Printf("Failed retrieving license key")
		return
	}

	// GetMeteredState freshly checks the state, contacting the licensing server.
	state, err := license.GetMeteredState()
	if err != nil {
		fmt.Printf("ERROR getting metered state: %+v\n", err)
		panic(err)
	}
	if !state.OK {
		fmt.Printf("State is not OK\n")
	}

	start = time.Now()
	err = lib.UniPdfDecryptPassword(input, password, output4)
	if err != nil {
		fmt.Println("UniPdfDecryptPassword - fail")
		fmt.Println(err)
	} else {
		fmt.Println("UniPdfDecryptPassword - success")
	}
	elapse = time.Since(start)
	fmt.Println(">> UniPDF took ", elapse)

}
