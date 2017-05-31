package commerceml

type Storage interface {
	CreateGroup(parentID string, group Group) (err error)
	CreateProperty(property Property) (err error)

	CreateProduct(product Product) (err error)
	CreateProductGroup(parentID string, group Group) (err error)
	CreateProductImage(parentID string, image Image) (err error)
	CreateProductProperty(parentID string, property IdValue) (err error)
	CreateProductTax(parentID string, tax Tax) (err error)
	CreateProductRequisite(parentID string, requisite Requisite) (err error)
	CreateProductContractor(parentID string, contractor Contractor) (err error)
	CreateProductExcise(parentID string, excise Excise) (err error)
	CreateProductComponent(parentID string, component Component) (err error)
}
