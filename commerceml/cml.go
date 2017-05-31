package commerceml

import (
	"encoding/xml"
	"net/url"
	"strconv"
)

type Parser interface {
	Parse(data []byte) (err error)
}

type CommerceML struct {
	CommerceInfo xml.Name   `xml:"КоммерческаяИнформация"`
	Version      string     `xml:"ВерсияСхемы,attr"`
	Classifier   Classifier `xml:"Классификатор"`
	Catalog      Catalog    `xml:"Каталог"`
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
	Id   string `xml:"Ид";json:"id"`
	Name string `xml:"Наименование";json:"name"`
}

type Property struct {
	IdName
	Description
	Type            TypeProperty    `xml:"ТипЗначений"`
	Variants        []Variant       // @since 2.05
	RequireProperty RequireProperty `xml:"Обязательное"`
	Multiple        bool            `xml:"Множественное"`
	ForDocument     bool            `xml:"ДляДокументов"`  // @only 2.04
	ForOffer        bool            `xml:"ДляПредложений"` // @only 2.04
	ForProduct      bool            `xml:"ДляТоваров"`     // @only 2.04
	Usage           Usage           // @since 2.05
}

type TypeProperty int

const (
	NONE TypeProperty = iota
	DIRECTORY
	DIGIT
	STRING
)

func (t *TypeProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	*t = TypeProperty(t.Get(content))

	return nil
}

func (t TypeProperty) Get(v string) TypeProperty {
	switch v {
	case "Справочник":
		return DIRECTORY
	case "Число":
		return DIGIT
	case "Строка":
		return STRING
	}
	return NONE
}

// @since 2.05
type Usage struct {
	ForDocument bool `xml:"ДляДокументов"`
	ForOffer    bool `xml:"ДляПредложений"`
	ForProduct  bool `xml:"ДляТоваров"`
}

type Variant struct {
	Value     string
	Directory IdValue // @since 2.05
}

// @since 2.05
type IdValue struct {
	Id    string `xml:"Ид"`
	Value string `xml:"Значение"`
}

type RequireProperty int

const (
	CATALOG RequireProperty = iota
	DOCUMENT
	OFFER
)

type Description struct {
	Value string `xml:"Описание"`
}

type Catalog struct {
	IdName
	Classifier Classifier
	Owner      Owner
	Products   []Product `xml:"Товары>Товар"`
}

type Product struct {
	IdName
	Description
	BarCode      string      `xml:"Штрихкод"`
	Article      string      `xml:"Артикул"`
	Unit         Unit        `xml:"БазоваяЕдиница"`
	FullName     string      `xml:"Описание"`
	Groups       []Group     `xml:"Группы"`
	Images       []Image     `xml:"Картинка"`
	Properties   []IdValue   `xml:"ЗначенияСвойств>ЗначенияСвойства"`
	Taxes        []Tax       `xml:"СтавкиНалогов>СтавкаНалога"`
	Requisites   []Requisite `xml:"ЗначенияРеквизитов>ЗначениеРеквизита"`
	Country      string      `xml:"Страна"`
	Brand        string      `xml:"ТорговаяМарка"`
	OwnerBrand   Contractor  `xml:"ВладелецТорговойМарки"`
	Manufacturer Contractor  `xml:"Изготовитель"`
	Excises      []Excise    `xml:"Акцизы>Акциз"`
	Components   []Component `xml:"Комплектующие>Комплектующее"`
}

type Image struct {
	*url.URL
}

func (i *Image) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	u, err := url.Parse(content)
	if err != nil {
		return err
	}

	*i = Image{u}
	return nil
}

type Contractor struct {
	IdName
	Title    string `xml:"ОфициальноеНаименование"`
	FullName string `xml:"ПолноеНаименование"`
}

type Excise struct {
	Name     string  `xml:"Наименование"`
	Sum      float32 `xml:"СуммаЗаЕдиницу"`
	Currency string  `xml:"Валюта"`
}

type Requisite struct {
	Name  string `xml:"Наименование"`
	Value string `xml:"Значение"`
}

type Tax struct {
	Name string `xml:"Наименование"`
	Rate Rate   `xml:"Ставка"`
}

type Rate int

func (r *Rate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	if content == "Без налога" {
		*r = Rate(0)
	} else {
		i, err := strconv.Atoi(content)
		if err != nil {
			return err
		}
		*r = Rate(i)
	}
	return nil
}

type Unit struct {
	Name string `xml:"НаименованиеПолное"`
	Code int    `xml:"Код"`
}

type Component struct {
	Product      Product `xml:"Товар"`
	CatalogID    string  `xml:"ИдКаталога"`
	ClassifierID string  `xml:"ИдКлассификатора"`
	Amount       int     `xml:"КоличествоТип"`
}
