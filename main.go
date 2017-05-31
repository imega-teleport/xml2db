package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imega-teleport/xml2db/mysql"
	"github.com/imega-teleport/xml2db/parser/v204"
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

    user, pass, host, dbname := "root", "", "10.0.3.94:3306", "teleport"
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

	storage := mysql.NewStorage(db)
	parser := v204.NewParser204(storage)
	err = parser.Parse(data)
    if err != nil {
        fmt.Printf("error: %v", err)
        os.Exit(1)
    }
}
