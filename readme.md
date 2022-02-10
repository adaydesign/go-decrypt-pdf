# GO - Decrypt PDF file

## PDFCPU VS QPDF VS UniPDF

- example file size : 179kb
- example file pdf version : 1.6
- test with operation system : Windows 10

### Experimental Result
|lib|excution|time(ms)|output pdf version|out file size(kb)|
|---|---|---|---|---|
|PdfCpu|pkg|25.41486|1.7|178|
|PdfCpu|cli|121.29205|1.7|178|
|QPdf|cli|351.25808|1.6|179|
|UniPdf|pkg|6.83598|1.3|177|

### Limitation
- PdfCpu support all versions of PDF up to version 1.7
- QPdf is a low-level tool for working with the structure of PDF files
- UniPDF is not free. A free version can use 100 credits per billing cycles

<br>
<br>

# Usage library

#### Import
```
import "github.com/adaydesign/go-decrypt-pdf/lib"
```

#### PDFCPU
```
input := "files/locked_doc.pdf"
password := "@1234"
output := "files/out_pdfcpu_pkg.pdf"

err := lib.PdfCpuDecryptPassword(input, password, output1)
if err != nil {
    fmt.Println(err)
}
```

#### PDFCPU CLI
```
input := "files/locked_doc.pdf"
password := "@1234"
output1 := "files/out_pdfcpu.pdf"

path, err := os.Getwd()
if err != nil {
    fmt.Println(err)
}
pdfcpu := path + "/bin/pdfcpu_0.3.13/pdfcpu.exe"

err = lib.PdfCpuDecryptPasswordByCLI(pdfcpu, input, password, output2)
if err != nil {
    fmt.Println(err)
}
```

#### QPDF CLI
```
input := "files/locked_doc.pdf"
password := "@1234"
output2 := "files/out_qpdfcpu.pdf"

path, err := os.Getwd()
if err != nil {
    fmt.Println(err)
}
qpdf := path + "/bin/qpdf-10.5.0/bin/qpdf.exe"

err = lib.QPdfDecryptPasswordByCLI(qpdf, input, password, output3)
if err != nil {
    fmt.Println(err)
}
```

#### UniPDF
```
func init() {
	// To get your free API key for metered license, sign up on: https://cloud.unidoc.io
	// Make sure to be using UniPDF v3.19.1 or newer for Metered API key support.
	err := license.SetMeteredKey(`##### your unipdf key here #####`)
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		panic(err)
	}
}

func main(){
    input := "files/locked_doc.pdf"
	password := "@1234"
	output4 := "files/out_unipdf_pkg.pdf"

	err = lib.UniPdfDecryptPassword(input, password, output4)
    if err != nil {
        fmt.Println(err)
    }
}
```

### Reference
- https://github.com/pdfcpu/pdfcpu
- https://github.com/qpdf/qpdf
- https://github.com/unidoc/unipdf