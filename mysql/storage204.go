package mysql

import (
	"database/sql"

	"github.com/imega-teleport/xml2db/account"
	"github.com/imega-teleport/xml2db/commerceml"
)

type storage struct {
	db      *sql.DB
	account *account.Account
}

func NewStorage204(db *sql.DB, account *account.Account) *storage {
	return &storage{
		db:      db,
		account: account,
	}
}

func (s storage) CreateGroup(group commerceml.Group) (err error) {

	return
}
