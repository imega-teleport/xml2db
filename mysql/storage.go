package mysql

import (
	"database/sql"
	"encoding/json"

	"github.com/imega-teleport/xml2db/commerceml"
)

type storage struct {
	db db
}

func NewStorage(sqlDB *sql.DB) *storage {
	return &storage{
		db: db{sqlDB},
	}
}

func (s storage) CreateGroup(parentID string, group commerceml.Group) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.CreateGroup(parentID, group)

	return
}

func (s storage) CreateProperty(property commerceml.Property) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.CreateProperty(property)

	return
}

func (s storage) CreateProduct(product commerceml.Product) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.CreateProduct(product)

	return
}

func (s storage) CreateProductGroup(parentID string, group commerceml.Group) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.CreateProductGroup(parentID, group)

	return
}

func (s storage) CreateProductImage(parentID string, image commerceml.Image) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.CreateProductImage(parentID, image)

	return
}

func (s storage) CreateProductProperty(parentID string, property commerceml.IdValue) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.CreateProductProperty(parentID, property)
	return
}

func (s storage) CreateProductTax(parentID string, tax commerceml.Tax) (err error) {
	return
}

func (s storage) CreateProductRequisite(parentID string, requisite commerceml.Requisite) (err error) {
	return
}

func (s storage) CreateProductContractor(parentID string, contractor commerceml.Contractor) (err error) {
	return
}

func (s storage) CreateProductExcise(parentID string, excise commerceml.Excise) (err error) {
	return
}

func (s storage) CreateProductComponent(parentID string, component commerceml.Component) (err error) {
	return
}

type db struct {
	*sql.DB
}

func (d db) Begin() (*Tx, error) {
	tx, err := d.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}

type Tx struct {
	*sql.Tx
}

func (tx *Tx) CreateGroup(parentID string, group commerceml.Group) (err error) {
	stmt, err := tx.Prepare("INSERT groups(id,parent_id,name) VALUES (?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(group.Id, parentID, group.Name)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProperty(property commerceml.Property) (err error) {
	stmt, err := tx.Prepare("INSERT properties(id,name,type) VALUES (?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(property.Id, property.Name, property.Type)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProduct(product commerceml.Product) (err error) {
	stmt, err := tx.Prepare("INSERT products(id,name,groups) VALUES (?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	var groups []struct {
		ID string `json:"id"`
	}
	for _, i := range product.Groups {
		s := struct {
			ID string `json:"id"`
		}{
			ID: i.Id,
		}
		groups = append(groups, s)
	}
	groupsJson, err := json.Marshal(groups)
	if err != nil {
		return
	}

	_, err = stmt.Exec(product.Id, product.Name, groupsJson)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductGroup(parentID string, group commerceml.Group) (err error) {
	stmt, err := tx.Prepare("INSERT products_groups(parent_id,id) VALUES (?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, group.Id)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductImage(parentID string, image commerceml.Image) (err error) {
	stmt, err := tx.Prepare("INSERT products_images(parent_id,url) VALUES (?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, image.String())
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductProperty(parentID string, property commerceml.IdValue) (err error) {
	stmt, err := tx.Prepare("INSERT products_properties(parent_id,id,value) VALUES (?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, property.Id, property.Value)
	if err != nil {
		return
	}

	return
}
