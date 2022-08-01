package main

import (
	"os"
	utils "pdfreader/utils"
)

func main() {

	content, err := utils.ReadPdfRow("test_files/0478_s16_qp_11.pdf")
	if err != nil {
		panic(err)
	}

	os.WriteFile("output.txt", []byte(content), 0644)

	utils.FormatLines("output.txt")
}
