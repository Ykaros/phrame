package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func AddFrame(imgPath, outPath string, borderRatio float64) error {

	// TODO: Add image validation

	img, err := os.Open(imgPath)
	if err != nil {
		fmt.Printf("Error reading image from %s: %v\n", imgPath, err)
	}
	defer img.Close()
	photo, _, err := image.Decode(img)
	if err != nil {
		fmt.Printf("Error reading image from %s: %v\n", imgPath, err)
	}

	// Determine the size of the border and the square
	borderWidth := int(borderRatio * float64(min(photo.Bounds().Dx(), photo.Bounds().Dy())))
	squareSize := max(photo.Bounds().Dx(), photo.Bounds().Dy()) + 2*borderWidth

	// Create a squared RGBA canvas
	canvas := image.NewRGBA(image.Rect(0, 0, squareSize, squareSize))

	// Calculate the center point
	startX := squareSize / 2
	startY := squareSize / 2

	// Fill the entire canvas with white
	// More color options
	draw.Draw(canvas, canvas.Bounds(),
		&image.Uniform{color.White}, image.Point{}, draw.Over)

	// Place the photo at the center of the canvas
	draw.Draw(canvas,
		image.Rect(startX-photo.Bounds().Dx()/2, startY-photo.Bounds().Dy()/2,
			startX+photo.Bounds().Dx()/2, startY+photo.Bounds().Dy()/2),
		photo, image.Point{}, draw.Over)

	//TODO: Ask to create dir if there is none
	outIMG, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("Error creating new folder: %v\n", err)
	}
	defer outIMG.Close()

	// Encode the image and save it as JPEG
	err = jpeg.Encode(outIMG, canvas, nil)
	if err != nil {
		fmt.Printf("Error saving image: %v\n", err)
	}

	fmt.Printf("Image successfully saved to: %s\n", outPath)

	return nil
}

func AddFrames(sourceDir, destinationDir string, borderRatio float64) error {
	// Multi-threading support
	images, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Printf("Error reading files from %s: %v\n", sourceDir, err)
	}

	for _, file := range images {
		//fmt.Println(file)
		if file.IsDir() || filepath.Ext(file.Name()) == ".DS_Store" {
			continue
		}

		imgPath := filepath.Join(sourceDir, file.Name())
		outPath := filepath.Join(destinationDir, file.Name())

		err := AddFrame(imgPath, outPath, borderRatio)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", imgPath, err)
		}
	}

	return nil
}
