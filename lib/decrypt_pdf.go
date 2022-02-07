package lib

import (
	"fmt"
	"os/exec"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func PdfCpuDecryptPassword(input string, password string, output string) (result string, error string) {
	conf := pdfcpu.NewAESConfiguration(password, "", 256)
	err := api.DecryptFile(input, output, conf)

	if err != nil {
		result = fmt.Sprint("PDFCPU (pkg) Unlock Error : ", err.Error())
		error = err.Error()
		return
	} else {
		result = fmt.Sprint("PDFCPU (pkg) Unlock Finished : ", output)
		return
	}
}

func PdfCpuDecryptPasswordByCLI(pdfcpu string, input string, password string, output string) (result string, error string) {
	arg0 := "decrypt"
	arg1 := "-opw"

	cmd := exec.Command(pdfcpu, arg0, arg1, password, input, output)
	stdout, err := cmd.Output()
	if err != nil {
		result = fmt.Sprint("PDFCPU Unlock Error : ", string(stdout))
		error = err.Error()
		return
	} else {
		result = fmt.Sprint("PDFCPU Unlock Finished : ", output)
		return
	}
}

func QPdfDecryptPasswordByCLI(qpdf string, input string, password string, output string) (result string, error string) {
	arg0 := "--password=" + password
	arg1 := "--decrypt"

	cmd := exec.Command(qpdf, arg0, arg1, input, output)
	stdout, err := cmd.Output()

	if err != nil {
		result = fmt.Sprint("QPDF Unlock Error : ", string(stdout))
		error = err.Error()
		return
	} else {
		result = fmt.Sprint("QPDF Unlock Finished : ", output)
		return
	}
}
