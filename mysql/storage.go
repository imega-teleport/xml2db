package mysql

import (
	"database/sql"

	"encoding/json"

	"github.com/imega-teleport/xml2db/account"
	"github.com/imega-teleport/xml2db/commerceml"
)

type storage struct {
	db      db
	account *account.Account
}

func NewStorage204(sqlDB *sql.DB, account *account.Account) *storage {
	return &storage{
		db:      db{sqlDB},
		account: account,
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

	err = tx.CreateGroup(s.account.ID, parentID, group)

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

	err = tx.CreateProperty(s.account.ID, property)

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

	err = tx.CreateProduct(s.account.ID, product)

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

func (tx *Tx) CreateGroup(account, parentID string, group commerceml.Group) (err error) {
	stmt, err := tx.Prepare("INSERT groups(client_id,id,parent_id,name) VALUES (?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(account, group.Id, parentID, group.Name)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProperty(account string, property commerceml.Property) (err error) {
	stmt, err := tx.Prepare("INSERT properties(client_id,id,name,type) VALUES (?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(account, property.Id, property.Name, property.Type)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProduct(account string, product commerceml.Product) (err error) {
	stmt, err := tx.Prepare("INSERT products(client_id,id,name,groups) VALUES (?,?,?,?)")
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

	_, err = stmt.Exec(account, product.Id, product.Name, groupsJson)
	if err != nil {
		return
	}

	return
}
