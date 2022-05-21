package services

type SpacesInterface interface {
	IsFileAnImage(string) bool
	GetObject(string) (string, error)
	UploadObject(string, string) (string, error)
	DeleteObject(string, string) error
	ListAllObjectsInAFolder(string) error
}
