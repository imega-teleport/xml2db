package commerceml

type Storage interface {
	CreateGroup(parentID string, group Group) (err error)
}