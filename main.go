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

    user, pass, host, dbname := "root", "", "10.0.3.90:3306", "teleport"
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s", user, pass, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
    err = db.Ping()
    if err != nil {
        fmt.Printf("error: %v", err)
        os.Exit(1)
    }
    defer func() {
        err := db.Close()
        if err != nil {
            fmt.Printf("error: %v", err)
            return
        }
        fmt.Println("Closed db connection")
    }()

	acc := &account.Account{
		ID: "123",
	}
	storage := mysql.NewStorage204(db, acc)
	parser := v204.NewParser204(storage)
	err = parser.Parse(b1)
    if err != nil {
        fmt.Printf("error: %v", err)
        os.Exit(1)
    }
}
