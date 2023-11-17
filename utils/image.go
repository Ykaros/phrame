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
	"sync"
	"time"
)

// TDDO: Image validation
func readImage(path string) (image.Image, error) {
	img, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer img.Close()
	photo, _, err := image.Decode(img)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func saveImage(path string, photo image.Image) error {
	//fmt.Println(path)
	outIMG, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outIMG.Close()
	err = jpeg.Encode(outIMG, photo, nil)
	if err != nil {
		return err
	}
	return nil
}

func createCanvas(photo image.Image, borderRatio float64, squared bool, c color.RGBA) *image.RGBA {
	// Determine the size of the border and the square
	width, height := photo.Bounds().Dx(), photo.Bounds().Dy()
	if squared {
		borderWidth := int(borderRatio * float64(min(width, height)))
		squareSize := max(width, height) + 2*borderWidth

		// Create a squared RGBA canvas
		canvas := image.NewRGBA(image.Rect(0, 0, squareSize, squareSize))

		// Calculate the starting point
		startX := (squareSize - width) / 2
		startY := (squareSize - height) / 2

		// Fill the entire canvas with white
		// More color options
		draw.Draw(canvas, canvas.Bounds(),
			&image.Uniform{c}, image.Point{}, draw.Over)

		// Place the photo at the center of the canvas
		draw.Draw(canvas, image.Rect(startX, startY, startX+width, startY+height),
			photo, image.Point{}, draw.Over)
		return canvas
	} else {
		borderWidth := int(borderRatio * float64(width))
		borderHeight := int(borderRatio * float64(height))
		// Create a squared RGBA canvas
		canvas := image.NewRGBA(image.Rect(0, 0,
			width+2*borderWidth, height+2*borderHeight))

		// Calculate the starting point
		startX := borderWidth
		startY := borderHeight

		// Fill the entire canvas with white
		// More color options
		draw.Draw(canvas, canvas.Bounds(),
			&image.Uniform{c}, image.Point{}, draw.Over)

		// Place the photo at the center of the canvas
		draw.Draw(canvas, image.Rect(startX, startY, startX+width, startY+height),
			photo, image.Point{}, draw.Over)
		return canvas
	}

}

func AddFrames(sourcePath, outPath string, borderRatio float64, squared bool, c color.RGBA) error {
	if outPath == "" {
		currentTime := time.Now()
		outPath = currentTime.Format("2112_01_02_03_04_05")
	}

	// Check if the source exists and source type
	fileInfo, err := os.Stat(sourcePath)
	if err != nil {
		fmt.Printf("Error accessing file(s) at %s: %v\n", sourcePath, err)
	}

	// Determine if the source is a directory or a file
	if fileInfo.IsDir() {
		err := os.Mkdir(outPath, os.ModePerm)
		if err != nil {
			return err
		}
		images, err := os.ReadDir(sourcePath)
		if err != nil {
			fmt.Printf("Error reading files from %s: %v\n", sourcePath, err)
		}

		var wg sync.WaitGroup
		for _, file := range images {
			//fmt.Println(file)
			if file.IsDir() || filepath.Ext(file.Name()) == ".DS_Store" {
				continue
			}
			wg.Add(1)

			go func(imgPath, savePath string) {
				defer wg.Done()
				photo, err := readImage(imgPath)
				if err != nil {
					fmt.Printf("Error reading image %s: %v\n", imgPath, err)
					return
				}
				canvas := createCanvas(photo, borderRatio, squared, c)
				err = saveImage(savePath, canvas)
				if err != nil {
					fmt.Printf("Error saving image %s: %v\n", savePath, err)
					return
				}
				fmt.Printf("Image successfully saved to: %s\n", savePath)
			}(filepath.Join(sourcePath, file.Name()), filepath.Join(outPath, file.Name()))
		}
		wg.Wait() // Wait for all goroutines to finish

	} else {
		// Just one image
		format := filepath.Ext(sourcePath)
		outPath = sourcePath[:len(sourcePath)-len(format)] + outPath + format

		photo, err := readImage(sourcePath)
		if err != nil {
			return fmt.Errorf("Error reading image %s: %v", sourcePath, err)
		}
		canvas := createCanvas(photo, borderRatio, squared, c)
		err = saveImage(outPath, canvas)
		if err != nil {
			return fmt.Errorf("Error saving image %s: %v", outPath, err)
		}
		fmt.Printf("Image successfully saved to: %s\n", outPath)
	}
	return nil
}
