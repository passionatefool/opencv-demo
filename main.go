package main

import (
	"fmt"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

var xmlFile = "/usr/local/Cellar/opencv/4.5.3_2/share/opencv4/haarcascades/haarcascade_eye_tree_eyeglasses.xml"

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening web cam: %v", err)
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			fmt.Printf("cannot read device")
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		color := color.RGBA{0, 255, 0, 0}
		for _, r := range rects {
			fmt.Println("detected", r)
			gocv.Rectangle(&img, r, color, 3)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(50)
	}
}
