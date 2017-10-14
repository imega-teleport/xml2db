package mysql

import (
	"database/sql"

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

func (s storage) FulfillTask(name string, finish bool) (err error) {
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

	err = tx.FulfillTask(name, finish)

	return
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

func (s storage) CreateProductImage(parentID string, image commerceml.Image, idx int) (err error) {
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

	err = tx.CreateProductImage(parentID, image, idx)

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

	err = tx.CreateProductTax(parentID, tax)
	return
}

func (s storage) CreateProductRequisite(parentID string, requisite commerceml.Requisite) (err error) {
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

	err = tx.CreateProductRequisite(parentID, requisite)
	return
}

func (s storage) CreateProductContractor(parentID string, contractor commerceml.Contractor) (err error) {
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

	err = tx.CreateProductContractor(parentID, contractor)
	return
}

func (s storage) CreateProductExcise(parentID string, excise commerceml.Excise) (err error) {
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

	err = tx.CreateProductExcise(parentID, excise)
	return
}

func (s storage) CreateProductComponent(component commerceml.Component) (err error) {
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

	err = tx.CreateProductComponent(component)
	return
}

func (s storage) CreateBundling(bundling commerceml.Bundling) (err error) {
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

	err = tx.CreateBundling(bundling)
	return
}

func (s storage) CreateOffers(bundling commerceml.Bundling, offers []commerceml.Offer) (err error) {
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
	for _, i := range offers {
		err = tx.CreateOffer(bundling.Id, i)
	}
	return
}

func (s storage) CreatePricesTypes(bundling commerceml.Bundling, pricesTypes []commerceml.PriceType) (err error) {
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
	for _, i := range pricesTypes {
		err = tx.CreatePriceType(bundling.Id, i)
	}
	return
}

func (s storage) CreatePrices(offer commerceml.Offer, prices []commerceml.Price) (err error) {
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
	for _, i := range prices {
		err = tx.CreatePrice(offer.Id, i)
	}
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

func (tx *Tx) FulfillTask(name string, finish bool) (err error) {
	stmt, err := tx.Prepare("INSERT tasks(name,val) VALUES (?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(name, finish)
	if err != nil {
		return
	}

	return
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
	stmt, err := tx.Prepare("INSERT products(id, name, description, barcode, article, full_name, country, brand) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(product.Id, product.Name, product.Description.Value, product.BarCode, product.Article, product.FullName, product.Country, product.Brand)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductGroup(parentID string, group commerceml.Group) (err error) {
	stmt, err := tx.Prepare("INSERT products_groups(product_id,group_id) VALUES (?,?)")
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

func (tx *Tx) CreateProductImage(parentID string, image commerceml.Image, idx int) (err error) {
	stmt, err := tx.Prepare("INSERT products_images(parent_id,entity_id,url) VALUES (?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, idx, image.String())
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

func (tx *Tx) CreateProductTax(parentID string, tax commerceml.Tax) (err error) {
	stmt, err := tx.Prepare("INSERT products_taxes(parent_id,name,rate) VALUES (?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, tax.Name, tax.Rate)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductRequisite(parentID string, requisite commerceml.Requisite) (err error) {
	stmt, err := tx.Prepare("INSERT products_requisites(parent_id,name,value) VALUES (?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, requisite.Name, requisite.Value)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductContractor(parentID string, contractor commerceml.Contractor) (err error) {
	stmt, err := tx.Prepare("INSERT products_contractor(parent_id,id,name,title,full_name) VALUES (?,?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, contractor.Id, contractor.Name, contractor.Title, contractor.FullName)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductExcise(parentID string, excise commerceml.Excise) (err error) {
	stmt, err := tx.Prepare("INSERT products_excises(parent_id, name, sum, currency) VALUES (?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, excise.Name, excise.Sum, excise.Currency)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateProductComponent(component commerceml.Component) (err error) {
	stmt, err := tx.Prepare("INSERT products_component(parent_id,	catalog_id, classifier_id, quantity) VALUES (?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(component.Product.Id, component.CatalogID, component.ClassifierID, component.Quantity)
	if err != nil {
		return
	}

	return
}

func (s storage) CreateProducts(products []commerceml.Product) (err error) {
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

	for _, i := range products {
		err = tx.CreateProduct(i)
		if err != nil {
			return
		}
	}
	return
}

func (tx *Tx) CreateBundling(bundling commerceml.Bundling) (err error) {
	stmt, err := tx.Prepare("INSERT bundling(id, name, catalog_id, classifier_id) VALUES (?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(bundling.Id, bundling.Name, bundling.CatalogID, bundling.ClassifierID)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreateOffer(parentID string, offer commerceml.Offer) (err error) {
	stmt, err := tx.Prepare("INSERT bundling_offers(parent_id, id, name, base_unit, base_unit_name, base_unit_code, base_unit_global, quantity) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, offer.Id, offer.Name, offer.BaseUnit, offer.BaseUnitName, offer.BaseUnitCode, offer.BaseUnitGlobal, offer.Quantity)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreatePriceType(parentID string, priceType commerceml.PriceType) (err error) {
	stmt, err := tx.Prepare("INSERT bundling_prices_types(parent_id) VALUES (?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID)
	if err != nil {
		return
	}

	return
}

func (tx *Tx) CreatePrice(parentID string, price commerceml.Price) (err error) {
	stmt, err := tx.Prepare("INSERT bundling_offers_prices(parent_id, display, price_type_id, unit_price, currency, unit, coefficient) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.Exec(parentID, price.Display, price.PriceTypeID, price.UnitPrice, price.Currency, price.Unit, price.Coefficient)
	if err != nil {
		return
	}

	return
}
