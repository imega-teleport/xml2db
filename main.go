package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imega-teleport/xml2db/account"
	"github.com/imega-teleport/xml2db/commerceml/v204"
	"github.com/imega-teleport/xml2db/mysql"
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

	host := ""
	dsn := fmt.Sprintf("%s", host)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	account := &account.Account{
		ID: "123",
	}
	storage := mysql.NewStorage204(db, account)
	parser := v204.NewParser204(storage)
	parser.Parse(b1)

}

/*
func (g Group) createGroup(parentId string) {
	fmt.Println(parentId, g.Name)

	if len(g.Groups) == 0 {
		return
	}

	for _, c := range g.Groups {
		c.createGroup(g.Id)
	}
}
*/
