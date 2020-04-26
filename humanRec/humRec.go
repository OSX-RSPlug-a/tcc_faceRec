package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"time"

	"gocv.io/x/gocv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Needs to run:\n\tcrazy_detect [camera ID] [classifier XML file]")
		return
	}

	deviceID, _ := strconv.Atoi(os.Args[1])
	xmlFile := os.Args[2]

	webcam, err := gocv.VideoCaptureDevice(int(deviceID))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer webcam.Close()

	window := gocv.NewWindow("Human detected")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	colorSqr := color.RGBA{255, 0, 0, 1}

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading file: %v\n", xmlFile)
		return
	}

	fmt.Printf("starting camera device: %v\n", deviceID)

	for {

		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Cannot start device %d\n", deviceID)
			return
		}

		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("Found %d humans\n", len(rects))

		for _, r := range rects {

			gocv.Rectangle(&img, r, colorSqr, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)

			fmt.Println("In the time: ", time.Now())

			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)

			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, colorSqr, 2)
		}

		window.IMShow(img)

		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
