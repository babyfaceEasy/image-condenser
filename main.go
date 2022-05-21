package main

import (
	"fmt"
	"log"
	"thumbnail-generator/configs"
	"thumbnail-generator/di"
	"thumbnail-generator/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// get s3 client
	s3Client, err := configs.GetSpacesClient()
	if err != nil {
		log.Fatal("error getting s3Client")
	}

	// get db conn
	dbConn, err := configs.GetDBConnection()
	if err != nil {
		log.Fatal("error getting db conn")
	}
	defer dbConn.Close()

	dbService := di.InitializeDBService(dbConn)
	spacesService := di.IntializeSpacesService(s3Client)
	imageService := di.InitializeImageService()

	err = initThumbnailGeneration(dbService, spacesService, imageService)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func initThumbnailGeneration(dbService services.DB, spacesService services.Spaces, imageService services.ImageService) error {

	prdts, err := dbService.GetAllProduct("id, name")
	if err != nil {
		return err
	}

	for _, prdt := range prdts {
		fmt.Printf("%+v\n", prdt)

		// if no picture path continue
		if !prdt.PicturePath.Valid || len(prdt.PicturePath.String) <= 0 {
			continue
		}

		// no need to generate if thumbnail exists
		/*
			if prdt.ThumbnailPath.Valid && len(prdt.ThumbnailPath.String) > 0 {
				continue
			}
		*/

		// only work on images
		if !spacesService.IsFileAnImage(prdt.PicturePath.String) {
			continue
		}

		// download image
		fmt.Printf("Compressing image: %s\n", prdt.PicturePath.String)
		imageName, err := spacesService.GetObject(prdt.PicturePath.String)
		if err != nil {
			fmt.Println(err.Error())
		}

		// create thumbnail
		thumbnailName, err := imageService.GenerateThumbnail(imageName)
		if err != nil {
			fmt.Println(err.Error())
		}

		// delete thumbnail if one exist already
		if prdt.ThumbnailPath.Valid && len(prdt.ThumbnailPath.String) > 0 {
			//delete the image first
			spacesService.DeleteObject(thumbnailName, prdt.PicturePath.String)
		}

		// upload thumbnail
		thumbnailPath, err := spacesService.UploadObject(thumbnailName, prdt.PicturePath.String)
		if err != nil {
			fmt.Println(err.Error())
		}

		// update db
		err = dbService.UpdateStockThumbnailPath(thumbnailPath, prdt.ID)
		if err != nil {
			fmt.Println(err.Error())
		}

		// delete image both initial and thumbnail
		err = imageService.DeleteImage(imageName)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = imageService.DeleteImage(thumbnailName)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("Done compressing image: %s\n", prdt.PicturePath.String)
	}

	return nil
}
