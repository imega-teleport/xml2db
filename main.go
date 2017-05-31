package main

import (
	"fmt"
	"os"

	"github.com/imega-teleport/xml2db/parser/v204"
	"github.com/imega-teleport/xml2db/terminal"
)

func main() {
	xmlFile, err := os.Open("import.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	defer xmlFile.Close()

	stat, err := xmlFile.Stat()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	data := make([]byte, stat.Size())
	_, err = xmlFile.Read(data)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	storage := terminal.NewStorage()
	parser := v204.NewParser204(storage)
	err = parser.Parse(data)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
