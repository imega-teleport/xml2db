package v204

import (
	"encoding/xml"

	"github.com/imega-teleport/xml2db/commerceml"
)

type parser struct {
	storage commerceml.Storage
}

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
	Description
	Type            string          `xml:"ТипЗначений"`
	RequireProperty RequireProperty `xml:"Обязательное"`
	Multiple        bool            `xml:"Множественное"`
	ForDocument     bool            `xml:"ДляДокументов"`
	ForOffer        bool            `xml:"ДляПредложений"`
	ForCatalog      bool            `xml:"ДляТоваров"`
}

type RequireProperty int

const (
	CATALOG  RequireProperty = iota
	DOCUMENT
	OFFER
)

type Description struct {
	Value string `xml:"Описание"`
}

func NewParser204(storage commerceml.Storage) parser {
	return parser{
		storage: storage,
	}
}

func (p parser) Parse(data []byte) (err error) {
	cml := &CommerceML{}
	err = xml.Unmarshal(data, cml)

	for _, g := range cml.Classifier.Groups {
		p.CreateGroup("", g)
	}

	for _, i := range cml.Classifier.Properties {
		p.CreateProperty(i)
	}

	return
}

func (p parser) CreateGroup(parentId string, group Group) (err error) {
	err = p.storage.CreateGroup(parentId, commerceml.Group{
		IdName: commerceml.IdName{
			Id:   group.Id,
			Name: group.Name,
		},
	})

	if len(group.Groups) == 0 {
		return
	}

	for _, g := range group.Groups {
		err = p.CreateGroup(group.Id, g)
	}

	return
}

func (p parser) CreateProperty(property Property) (err error) {
	err = p.storage.CreateProperty(commerceml.Property{
		IdName: commerceml.IdName{
			Id:   property.Id,
			Name: property.Name,
		},
		Type: commerceml.Property{}.Type.Get(property.Type),
	})

	return
}
