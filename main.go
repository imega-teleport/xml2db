package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type CommerceML struct {
	CommerceInfo xml.Name   `xml:"КоммерческаяИнформация"`
	Version      string     `xml:"ВерсияСхемы,attr"`
	Classifier   Classifier `xml:"Классификатор"`
}

type Group struct {
	IdName
	Groups     []Group    `xml:"Группы>Группа"`
	Properties []Property `xml:"Свойства>Свойство"`
}

type Classifier struct {
	IdName
	Owner      Owner       `xml:"Владелец"`
	Groups     []Group     `xml:"Группы>Группа"`
	Properties []Property  `xml:"Свойства>Свойство"`
	PriceTypes []PriceType `xml:"ТипыЦен"`
}

type PriceType struct {
}

type Owner struct {
	IdName
}

type IdName struct {
	Id   string `xml:"Ид"`
	Name string `xml:"Наименование"`
}

type Property struct {
	IdName
	Type string `xml:"ТипЗначений"`
	Description
}

type Description struct {
	Value string `xml:"Описание"`
}

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
	cml := &CommerceML{}

	err = xml.Unmarshal(b1, cml)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("XMLName: %#v\n", cml.CommerceInfo)
	fmt.Printf("XMLName: %#v\n", cml.Version)
	fmt.Printf("XMLName: %#v\n", cml.Classifier.Id)
	fmt.Printf("XMLName: %#v\n", cml.Classifier.Groups)

	for _, g := range cml.Classifier.Groups {
		g.createGroup("")
	}
}

func (g Group) createGroup(parentId string) {
	fmt.Println(parentId, g.Name)

	if len(g.Groups) == 0 {
		return
	}

	for _, c := range g.Groups {
		c.createGroup(g.Id)
	}
}
