// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"database/sql"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"thumbnail-generator/services"
)

// Injectors from wire.go:

func InitializeImageService() services.ImageService {
	imageService := services.NewImageService()
	return imageService
}

func IntializeSpacesService(s3Client *s3.S3) services.Spaces {
	spaces := services.NewSpacesService(s3Client)
	return spaces
}

func InitializeDBService(dbConn *sql.DB) services.DB {
	db := services.NewDBService(dbConn)
	return db
}

// wire.go:

var imageSet = wire.NewSet(wire.Bind(new(services.ImageInterface), new(services.ImageService)), services.NewImageService)

var spacesServiceSet = wire.NewSet(services.NewSpacesService, wire.Bind(new(services.SpacesInterface), new(services.Spaces)))
