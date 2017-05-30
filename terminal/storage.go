package terminal

import (
	"fmt"

	"github.com/imega-teleport/xml2db/account"
	"github.com/imega-teleport/xml2db/commerceml"
)

type storage struct {
	account *account.Account
}

func NewStorage(account *account.Account) *storage {
	return &storage{
		account: account,
	}
}

func (s storage) CreateGroup(parentID string, group commerceml.Group) (err error) {
	fmt.Println(parentID, group)
	return
}

func (s storage) CreateProperty(property commerceml.Property) (err error) {
	fmt.Println(property)

	return
}

func (s storage) CreateProduct(product commerceml.Product) (err error) {
	fmt.Println(product)

	return
}
