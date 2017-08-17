package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imega-teleport/xml2db/mysql"
	"github.com/imega-teleport/xml2db/parser/v204"
)

func main() {
	var fileImport bool
	user, pass, host := os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST")
	dbname := flag.String("db", "", "Database name")
	f := flag.String("file", "", "File to parse")
	flag.Parse()
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s", user, pass, host, *dbname)
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

	xmlFile1, err := os.Open(*f)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	defer xmlFile1.Close()

	scanner := bufio.NewScanner(xmlFile1)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "Классификатор") {
			fileImport = true
			break
		}
		if strings.Contains(t, "ПакетПредложений") {
			fileImport = false
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("reading file: %v", err)
		os.Exit(1)
	}

	xmlFile, err := os.Open(*f)
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

	storage := mysql.NewStorage(db)
	parser := v204.NewParser204(storage)
	if fileImport == true {
		err = storage.FulfillTask("store", true);
		err = parser.Parse(data)
	} else {
		err = storage.FulfillTask("offer", true);
		err = parser.ParseBundling(data)
	}
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
