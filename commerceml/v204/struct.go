package v204

type CommerceML struct {
	CommerceInfo CommerceInfo `xml:"КоммерческаяИнформация"`
}

type CommerceInfo struct {
	Classifier Classifier `xml:"Классификатор"`
}

type Classifier struct {
	Id   string
	Name string
}
