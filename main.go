package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/sensor"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Camera")

	cameraDevice, isCameraDevice := fyne.CurrentDevice().(sensor.CameraDevice)

	resultImage := &canvas.Image{
		FillMode: canvas.ImageFillContain,
	}

	takePhoto := widget.NewButton("Take a photo!", func() {
		img, err := cameraDevice.CapturePhoto()
		if err == nil {
			resultImage.Image = img
			resultImage.Refresh()
		} else {
			log.Println("error returned from CapturePhoto: " + err.Error())
		}
	})

	startLoop := widget.NewButton("Start Viewfinder", func() {
		preview := cameraDevice.Preview()
		cameraDevice.StartPreview()
		go func() {
			for img := range preview {
				resultImage.Image = img
				resultImage.Refresh()
			}
		}()
	})

	stopLoop := widget.NewButton("Stop Viewfinder", func() {
		cameraDevice.StopPreview()
	})

	buttons := container.NewHBox(
		takePhoto,
		startLoop,
		stopLoop,
	)
	if !isCameraDevice {
		takePhoto.Disable()
		startLoop.Disable()
		stopLoop.Disable()
	}
	w.SetContent(container.New(
		layout.NewBorderLayout(nil, buttons, nil, nil),
		buttons,
		container.NewStack(resultImage),
	))

	w.ShowAndRun()
}
