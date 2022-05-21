//go:build wireinject
// +build wireinject

package di

import (
	"thumbnail-generator/services"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"database/sql"
)

var imageSet = wire.NewSet(
	wire.Bind(new(services.ImageInterface), new(services.ImageService)),
	services.NewImageService,
)
func InitializeImageService() services.ImageService {
	wire.Build(imageSet)
	return services.ImageService{}
}

var spacesServiceSet = wire.NewSet(
	services.NewSpacesService,
	wire.Bind(new(services.SpacesInterface), new(services.Spaces)),
)
func IntializeSpacesService(s3Client *s3.S3) services.Spaces {
	wire.Build(spacesServiceSet)
	return services.Spaces{}
}

func InitializeDBService(dbConn *sql.DB) services.DB {
	wire.Build(services.NewDBService)
	return services.DB{}
}
