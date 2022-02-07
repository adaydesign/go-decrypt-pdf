# GO - Decrypt PDF file

## PDFCPU VS QPDF

- example file size : 179kb
- example file pdf version : 1.6
- test with operation system : Windows 10

### Time Excution
|lib|excution|time(ms)|output pdf version|
|---|---|---|---|
|PdfCpu|pkg|14.9964|1.7|
|PdfCpu|cli|71.3979|1.7|
|QPdf|cli|434.3817|1.6|

### Limitation
- PdfCpu support all versions of PDF up to version 1.7
- QPdf is a low-level tool for working with the structure of PDF files

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

result, _ := lib.PdfCpuDecryptPassword(input, password, output)
fmt.Println(result)
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

result, _ = lib.PdfCpuDecryptPasswordByCLI(pdfcpu, input, password, output1)
fmt.Println(result)
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

result, _ = lib.QPdfDecryptPasswordByCLI(qpdf, input, password, output2)
fmt.Println(result)
```

### Reference
- https://github.com/pdfcpu/pdfcpu
- https://github.com/qpdf/qpdf