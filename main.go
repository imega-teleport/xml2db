package main

import (
	"fmt"
	"os"

	"github.com/imega-teleport/xml2db/account"
	"github.com/imega-teleport/xml2db/commerceml/v204"
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

	b1 := make([]byte, stat.Size())
	_, err = xmlFile.Read(b1)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	acc := &account.Account{
		ID: "123",
	}
	storage := terminal.NewStorage(acc)
	parser := v204.NewParser204(storage)
	err = parser.Parse(b1)
    if err != nil {
        fmt.Printf("error: %v", err)
        os.Exit(1)
    }
}
