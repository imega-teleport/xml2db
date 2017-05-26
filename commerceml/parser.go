package commerceml

type CommerceML struct {
	Classifier Classifier
}

type Group struct {
	IdName
	Groups     []Group
	Properties []Property
}

type Classifier struct {
	IdName
	Owner      Owner
	Groups     []Group
	Properties []Property
	PriceTypes []PriceType
}

type PriceType struct {
}

type Owner struct {
	IdName
}

type IdName struct {
	Id   string
	Name string
}

type Property struct {
	IdName
	Type            TypeProperty
	Description
	Variants        []Variant // @since 2.05
	RequireProperty RequireProperty
	Multiple        bool
	ForDocument     bool  // @only 2.04
	ForOffer        bool  // @only 2.04
	ForCatalog      bool  // @only 2.04
	Usage           Usage // @since 2.05
}

type TypeProperty int

const (
	NONE      TypeProperty = iota
	DIRECTORY
	DIGIT
	STRING
)

func (t TypeProperty) Get(v string) TypeProperty {
	switch v {
	case "Справочник":
		return DIRECTORY
	case "Число":
		return DIGIT
	case "Строка":
		return STRING
	}
	return NONE
}

// @since 2.05
type Usage struct {
	ForDocument bool
	ForOffer    bool
	ForCatalog  bool
}

type Variant struct {
	Value     string
	Directory IdValue // @since 2.05
}

// @since 2.05
type IdValue struct {
	Id    string
	Value string
}

type RequireProperty int

const (
	CATALOG  RequireProperty = iota
	DOCUMENT
	OFFER
)

type Description struct {
	Value string
}

type Parser interface {
	CreateGroup(group Group) (err error)
	CreateProperty(property Property) (err error)
	Parse(data []byte) (err error)
}
