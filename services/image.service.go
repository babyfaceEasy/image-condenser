package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
Handles all interaction with the image object.
*/

var (
	imagePath = filepath.Join("./images/")
)

type ImageService struct{}

func NewImageService() ImageService {
	return ImageService{}
}

func (is ImageService) getOutputName(imageName string) string {
	imageNameSlice := strings.Split(imageName, ".")
	extension := imageNameSlice[len(imageNameSlice)-1]

	outputName := ""
	for i := 0; i < len(imageNameSlice)-1; i++ {
		outputName += imageNameSlice[i]
	}

	outputName = fmt.Sprintf("thumbnail_%s.%s", outputName, extension)
	return outputName
}

func (is ImageService) GenerateThumbnail(imageName string) (string,error) {
	thumbnailWidth := os.Getenv("THUMBNAIL_WIDTH")
	thumbnailHeight := os.Getenv("THUMBNAIL_HEIGHT")

	outputName := is.getOutputName(imageName)
	outputPath := filepath.Join(imagePath, outputName)
	imageName = filepath.Join(imagePath, imageName)

	resizeArg := fmt.Sprintf("%sx%s", thumbnailWidth, thumbnailHeight)

	cmd := exec.Command("convert", imageName, "-resize", resizeArg, outputPath)
	//cmd := exec.Command("convert", "in.jpeg", "-resize", "128x128", "out.jpeg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return outputName, nil
}

func (is ImageService) DeleteImage(imageName string) error {
	err := os.Remove(filepath.Join(imagePath, imageName))
	if err != nil {
		return err
	}
	return nil
}
