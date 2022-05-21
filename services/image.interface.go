package services

type ImageInterface interface {
	//NewImageService() services.ImageService
	GenerateThumbnail(string) (string, error)
	DeleteImage(string) error
}
