package mysql

import (
    "database/sql"

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
    defer tx.Rollback()

    err = tx.CreateGroup(s.account.ID, parentID, group)

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
    defer stmt.Close()

    _, err = stmt.Exec(account, group.Id, parentID, group.Name)
    if err != nil {
        return
    }
    err = tx.Commit()
    if err != nil {
        return
    }

    return
}
