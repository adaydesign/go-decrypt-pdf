package main

import (
	"fmt"
	"os"
	"time"

	"github.com/adaydesign/go-decrypt-pdf/lib"
)

func main() {

	input := "files/locked_doc.pdf"
	password := "@1234"
	output := "files/out_pdfcpu_pkg.pdf"
	output1 := "files/out_pdfcpu.pdf"
	output2 := "files/out_qpdfcpu.pdf"

	// 1. use pdfcpu package
	start := time.Now()
	result, _ := lib.PdfCpuDecryptPassword(input, password, output)
	fmt.Println(result)
	elapse := time.Since(start)
	fmt.Println("PDFCPU (pkg) took ", elapse)

	// 2. use pdfcpu (run excute file)
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	pdfcpu := path + "/lib/pdfcpu_0.3.13/pdfcpu.exe"

	start = time.Now()
	result, _ = lib.PdfCpuDecryptPasswordByCLI(pdfcpu, input, password, output1)
	fmt.Println(result)
	elapse = time.Since(start)
	fmt.Println("PDFCPU took ", elapse)

	// 3. use qpdf (run excute file)
	qpdf := path + "/lib/qpdf-10.5.0/bin/qpdf.exe"

	start = time.Now()
	result, _ = lib.QPdfDecryptPasswordByCLI(qpdf, input, password, output2)
	fmt.Println(result)
	elapse = time.Since(start)
	fmt.Println("QPDF took ", elapse)
}
