package traffic_counter

import (
	"errors"
	"image"
	"image/color"
	"time"

	cv "gocv.io/x/gocv"
)

const (
	UrlStreaming       = "https://58481b709ae2e.streamlock.net/strassennrw-rtplive/_definst_/mp4:10107898794605103925.stream/playlist.m3u8"
	UrlFile            = "cars2.ts"
	CarClassifierModel = "haarcascade_car.xml"
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
)

type Car struct {
	Speed     float32
	Direction struct {
		x, y float32
	}
}

func TrafficCounter(file string, cars chan Car) error {
	vcap, err := cv.VideoCaptureFile(file)
	if err != nil {
		return err
	}
	defer vcap.Close()

	// open display window
	window := cv.NewWindow("Face Detect")
	defer window.Close()

	fps := vcap.Get(cv.VideoCaptureFPS)
	wt := time.Second / time.Duration(fps)

	cars_cascade := cv.NewCascadeClassifier()
	if !cars_cascade.Load(CarClassifierModel) {
		return errors.New("failed to load classifier model")
	}

	for {
		var frame cv.Mat

		start_time := time.Now()

		if !vcap.Read(&frame) {
			return errors.New("failed to read frame")
		}

		cv.CvtColor(frame, &frame, cv.ColorBGRToGray)
		cars := cars_cascade.DetectMultiScaleWithParams(frame, 1.05, 4, 0, image.Point{}, image.Point{})

		for _, car := range cars {
			cv.Rectangle(&frame, car, red, 2)
			cv.Circle(&frame, car.Min.Add(car.Max).Div(2), 2, green, 2)

			window.IMShow(frame)

			// Press q to close the video windows before it ends if you want
			if window.WaitKey(22)&0xFF == 'q' {
				break
			}

			diff := time.Until(start_time.Add(wt))
			if diff > 0 {
				time.Sleep(diff)
			}
		}
	}
}
