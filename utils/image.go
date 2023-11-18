package utils

import (
	"fmt"
	"github.com/fogleman/gg"
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

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		//fmt.Printf("Invalid file path at %s: %v\n", path, err)
		//fmt.Printf("Please check if the file exists\n")
		return false
	}
	return fileInfo.IsDir()
}

// TDDO: Image validation
func readImage(path string) (image.Image, error) {
	img, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to read image at %s: %v\n", path, err)
		return nil, err
	}
	defer img.Close()
	photo, _, err := image.Decode(img)
	if err != nil {
		fmt.Printf("Failed to load image: %v\n", err)
		return nil, err
	}
	return photo, nil
}

func saveImage(path string, photo image.Image) error {
	outIMG, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to save image at %s: %v\n", path, err)
		return err
	}
	defer outIMG.Close()
	err = jpeg.Encode(outIMG, photo, nil)
	if err != nil {
		fmt.Printf("Failed to convert image format: %v\n", err)
		return err
	}
	return nil
}

func createCanvas(photo image.Image, borderRatio float64, squared bool, c, fontColor color.RGBA, signature string, fontSize int) (canvas *image.RGBA) {
	// Determine the size of the border and the position of the photo
	width, height := photo.Bounds().Dx(), photo.Bounds().Dy()
	var startX, startY int

	// Squared or original ratio
	if squared {
		borderWidth := int(borderRatio * float64(min(width, height)))
		squareSize := max(width, height) + 2*borderWidth

		// Create a squared RGBA canvas
		canvas = image.NewRGBA(image.Rect(0, 0, squareSize, squareSize))

		// Calculate the starting point
		startX = (squareSize - width) / 2
		startY = (squareSize - height) / 2
	} else {
		borderWidth := int(borderRatio * float64(width))
		borderHeight := int(borderRatio * float64(height))

		// Create a squared RGBA canvas
		canvas = image.NewRGBA(image.Rect(0, 0,
			width+2*borderWidth, height+2*borderHeight))

		// Calculate the starting point
		startX = borderWidth
		startY = borderHeight
	}

	// Fill the entire canvas with frameColor (default color is white)
	draw.Draw(canvas, canvas.Bounds(),
		&image.Uniform{c}, image.Point{}, draw.Over)

	// Place the photo at the center of the canvas
	draw.Draw(canvas, image.Rect(startX, startY, startX+width, startY+height),
		photo, image.Point{}, draw.Over)

	if signature != "" {
		// Sign the framed image
		dc := gg.NewContextForRGBA(canvas)
		dc.SetColor(fontColor)
		// Default font: Inter-Regular
		dc.LoadFontFace("font.ttf", float64(fontSize))

		// Determine the signature position
		// TODO: Make the signature position customizable
		textWidth, _ := dc.MeasureString(signature)
		dc.DrawString(signature, (float64(canvas.Bounds().Dx())-textWidth)/2, float64(canvas.Bounds().Dy()-fontSize))
	}

	return
}

func AddFrames(sourcePath, outPath string, borderRatio float64, squared bool, c, fontColor color.RGBA, signature string, fontSize int) error {
	if outPath == "" || IsDir(outPath) {
		currentTime := time.Now()
		outPath += currentTime.Format("2112_01_02_03_04_05")
	}

	// Determine if the source is a directory or a file
	// If it's a directory, create a new directory to contain all the framed images
	if IsDir(sourcePath) {
		err := os.Mkdir(outPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create a new directory at %s: %v\n", outPath, err)
			return err
		}
		images, err := os.ReadDir(sourcePath)
		if err != nil {
			fmt.Printf("Failed to read images from %s: %v\n", sourcePath, err)
			return err
		}

		var wg sync.WaitGroup
		for _, file := range images {
			// Ignore subdirectories and .DS_Store files
			if file.IsDir() || filepath.Ext(file.Name()) == ".DS_Store" {
				continue
			}
			wg.Add(1)

			go func(imgPath, savePath string) {
				defer wg.Done()
				photo, err := readImage(imgPath)
				if err != nil {
					fmt.Printf("Failed to read image at %s: %v\n", imgPath, err)
					return
				}
				canvas := createCanvas(photo, borderRatio, squared, c, fontColor, signature, fontSize)
				err = saveImage(savePath, canvas)
				if err != nil {
					fmt.Printf("Failed to save image at %s: %v\n", savePath, err)
					return
				}
				fmt.Printf("Image successfully saved to: %s \u2713\n", savePath)
			}(filepath.Join(sourcePath, file.Name()), filepath.Join(outPath, file.Name()))
		}
		wg.Wait() // Wait for all goroutines to finish

	} else {
		// Just one image
		format := filepath.Ext(sourcePath)
		outPath = outPath + format

		photo, err := readImage(sourcePath)
		if err != nil {
			fmt.Printf("Failed to read image at %s: %v", sourcePath, err)
			return err
		}
		canvas := createCanvas(photo, borderRatio, squared, c, fontColor, signature, fontSize)
		err = saveImage(outPath, canvas)
		if err != nil {
			fmt.Printf("Failed to save image at %s: %v", outPath, err)
			return err
		}
		fmt.Printf("Image successfully saved to: %s \u2713\n", outPath)
	}
	return nil
}

func Cut(sourcePath, grid string) error {
	photo, err := readImage(sourcePath)
	if err != nil {
		fmt.Printf("Failed to read image at %s: %v", sourcePath, err)
		return err
	}
	// Split the image into a grid
	width, height := photo.Bounds().Dx(), photo.Bounds().Dy()
	err = os.Mkdir("GRID", os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create directory: %v", err)
		fmt.Printf("Please remove the existing directory 'GRID' and try again\n")
		return err
	}
	if grid == "4" {
		// Split the image into 4 parts
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				// Create a new RGBA canvas
				canvas := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
				// Obtain each part of the image
				draw.Draw(canvas, image.Rect(0, 0, width/2, height/2),
					photo, image.Point{width / 2 * i, height / 2 * j}, draw.Over)
				// Save the image
				err = saveImage(fmt.Sprintf("GRID/out_%d_%d.jpg", i, j), canvas)
				if err != nil {
					fmt.Printf("Failed to save image: %v", err)
					return err
				}
			}
		}
	} else if grid == "9" {
		// Split the image into 9 parts
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// Create a new RGBA canvas
				canvas := image.NewRGBA(image.Rect(0, 0, width/3, height/3))
				// Obtain each part of the image
				draw.Draw(canvas, image.Rect(0, 0, width/3, height/3),
					photo, image.Point{width / 3 * i, height / 3 * j}, draw.Over)
				// Save the image
				err = saveImage(fmt.Sprintf("GRID/out_%d_%d.jpg", i, j), canvas)
				if err != nil {
					fmt.Printf("Failed to save image: %v", err)
					return err
				}
			}
		}
	} else {
		return fmt.Errorf("Invalid grid option.. \nPlease use 4 or 9..")
	}
	return nil
}
