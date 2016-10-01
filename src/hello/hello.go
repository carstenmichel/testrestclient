package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
	"github.com/jung-kurt/gofpdf"
)

// JSONTime IS THE TIME
type JSONTime struct { // this is a time struct
	Time  string `json:"time"`
	Date  string `json:"date"`
	Epoch int    `json:"milliseconds_since_epoch"`
}

func main() {
	resp, err := resty.R().Get("http://time.jsontest.com/")

	if err != nil {
		fmt.Printf("\nError %v", err)
		return
	}

	var timeVar JSONTime

	//	fmt.Print(typeOf(resp.Body()))
	err2 := json.Unmarshal([]byte(resp.Body()), &timeVar)

	if err2 != nil {
		fmt.Printf("Unmarshalling error %v\n", err2)
	}
	fmt.Printf("Time: %v\n", timeVar.Time)
	fmt.Printf("Epoch: %v\n", timeVar.Epoch)
	fmt.Printf("Date: %v\n", timeVar.Date)

	writepdf(timeVar)
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T\n", v)
}

func writepdf(ti JSONTime) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d /{nb}", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	pdf.Cell(5, 20, fmt.Sprintf("Time is %v", ti.Time))
	pdf.Cell(5, 30, fmt.Sprintf("Epoch is %v", ti.Epoch))
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Printf("Error opening PDF for output %v\n", err)
	}

}
