package v204

import (
	"encoding/xml"

	"github.com/imega-teleport/xml2db/commerceml"
)

type parser struct {
	storage commerceml.Storage
}

func NewParser204(storage commerceml.Storage) parser {
	return parser{
		storage: storage,
	}
}

func (p parser) Parse(data []byte) (err error) {
	cml := &commerceml.CommerceML{}
	err = xml.Unmarshal(data, cml)

	/*for _, g := range cml.Classifier.Groups {
		p.CreateGroup("", g)
	}

	for _, i := range cml.Classifier.Properties {
		p.CreateProperty(i)
	}*/

	for _, i := range cml.Catalog.Products {
		p.CreateProduct(i)
	}

	return
}

func (p parser) CreateGroup(parentId string, group commerceml.Group) (err error) {
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

func (p parser) CreateProperty(property commerceml.Property) (err error) {
	err = p.storage.CreateProperty(commerceml.Property{
		IdName: commerceml.IdName{
			Id:   property.Id,
			Name: property.Name,
		},
		Type: property.Type,
	})

	return
}

func (p parser) CreateProduct(product commerceml.Product) (err error) {
	var groups []commerceml.Group
	for _, i := range product.Groups {
		g := commerceml.Group{
			IdName: commerceml.IdName{
				Id: i.Id,
			},
		}
		groups = append(groups, g)
	}

	err = p.storage.CreateProduct(commerceml.Product{
		IdName: commerceml.IdName{
			Id:   product.Id,
			Name: product.Name,
		},
		Description: commerceml.Description{
			Value: product.Description.Value,
		},
		Groups: groups,
		Images: product.Images,
	})

	return
}
