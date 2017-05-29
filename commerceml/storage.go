package commerceml

type Storage interface {
	CreateGroup(parentID string, group Group) (err error)
	CreateProperty(property Property) (err error)
	CreateProduct(product Product) (err error)
}
