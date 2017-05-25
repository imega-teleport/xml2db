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
	Type string
	Description
}

type Description struct {
	Value string
}

type Parser interface {
	CreateGroup(group Group) (err error)
	Parse(data []byte) (err error)
}

func test() {
	g := Group{
		IdName: IdName{
			Id: "",
		},
	}
}