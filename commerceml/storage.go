package commerceml

type Storage interface {
	CreateGroup(parentID string, group Group) (err error)
	CreateProperty(property Property) (err error)

	CreateProduct(product Product) (err error)
	CreateProducts(products []Product) (err error)
	CreateProductGroup(parentID string, group Group) (err error)
	CreateProductImage(parentID string, image Image) (err error)
	CreateProductProperty(parentID string, property IdValue) (err error)
	CreateProductTax(parentID string, tax Tax) (err error)
	CreateProductRequisite(parentID string, requisite Requisite) (err error)
	CreateProductContractor(parentID string, contractor Contractor) (err error)
	CreateProductExcise(parentID string, excise Excise) (err error)
	CreateProductComponent(component Component) (err error)

	CreateBundling(bundling Bundling) (err error)
	CreateOffers(bundling Bundling, offers []Offer) (err error)
	CreatePricesTypes(bundling Bundling, pricesTypes []PriceType) (err error)
	CreatePrices(offer Offer, prices []Price) (err error)
}
