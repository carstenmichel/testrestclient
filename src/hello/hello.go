package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/go-resty/resty"
	"github.com/jung-kurt/gofpdf"
)

type CodeList struct {
	CodeListName  string `json:"codeListName"`
	VersionNumber int    `jdon:"versionNumber"`
	Codes         []struct {
		SenderCode   string `json:"senderCode"`
		ReceiverCode string `json:"receiverCode"`
		Description  string `json:"Description"`
	} `json:"codes"`
}

func main() {
	var username string
	var pword string

	var host string
	var port string
	var codelistname string
	var codelistversion string
	var allValuesSet bool = true

	flag.StringVar(&username, "username", "", "specify username for login")
	flag.StringVar(&pword, "password", "", "specify password for login")
	flag.StringVar(&host, "host", "", "specify host")
	flag.StringVar(&port, "port", "", "specify port")
	flag.StringVar(&codelistname, "codelistname", "", "specify codelistname")
	flag.StringVar(&codelistversion, "codelistversion", "", "specify codelistversion")

	flag.Parse()
	flag.VisitAll(func(arg1 *flag.Flag) {
		if len(arg1.Value.String()) == 0 {
			allValuesSet = false
		}
	})
	if allValuesSet == false {
		fmt.Printf("Please use with these Parameters\n")
		flag.PrintDefaults()
		return
	}
	var url string
	url = fmt.Sprintf("http://%v:%v/B2BAPIs/svc/codelists/%v:||%v", host, port, codelistname, codelistversion)
	fmt.Printf("Download from %v\n", url)
	resp, err := resty.R().
		SetBasicAuth(username, pword).
		Get(url)

	if err != nil {
		fmt.Printf("\nError %v", err)
		return
	}

	var clVar CodeList

	err2 := json.Unmarshal([]byte(resp.Body()), &clVar)

	if err2 != nil {
		fmt.Printf("Unmarshalling error %v\n", err2)
	}
	fmt.Printf("CodelistName : %v\n", clVar.CodeListName)
	fmt.Printf("VersionNumber : %v\n", clVar.VersionNumber)
	for zaehler := 0; zaehler < len(clVar.Codes); zaehler++ {
		fmt.Printf("----------------------\n")
		fmt.Printf("Sender : %v\n", clVar.Codes[zaehler].SenderCode)
		fmt.Printf("Receiver : %v\n", clVar.Codes[zaehler].ReceiverCode)
		fmt.Printf("Description : %v\n", clVar.Codes[zaehler].Description)
	}

	writepdf(clVar)
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T\n", v)
}

func writepdf(cl CodeList) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d /{nb}", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("Dump of codelist %v Version %v", cl.CodeListName, cl.VersionNumber))
	pdf.Ln(10)
	for zaehler := 0; zaehler < len(cl.Codes); zaehler++ {
		pdf.Cell(50, 20, fmt.Sprintf("Sender : %v\n", cl.Codes[zaehler].SenderCode))
		pdf.Cell(50, 20, fmt.Sprintf("Receiver : %v\n", cl.Codes[zaehler].ReceiverCode))
		pdf.Cell(50, 20, fmt.Sprintf("Description : %v\n", cl.Codes[zaehler].Description))
		pdf.Ln(10)
	}
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Printf("Error opening PDF for output %v\n", err)
	}

}
