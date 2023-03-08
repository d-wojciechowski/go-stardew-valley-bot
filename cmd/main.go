package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"github.com/kindlyfire/go-keylogger"
	"image"
	"time"
)

func setupKetLogger() chan rune {
	kl := keylogger.NewKeylogger()
	runes := make(chan rune, 1024)
	go func() {
		for {
			logger := kl.GetKey()
			if !logger.Empty {
				fmt.Println(logger.Rune)
				runes <- logger.Rune
			}
		}
	}()
	return runes
}

func main() {
	images := make(chan image.Image)
	go keyLoggerRoutine(images)
	a := app.New()
	w := a.NewWindow("Image")
	w.Resize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	go func() {
		imageRecreation(images, w)
	}()
	w.ShowAndRun()
}

func imageRecreation(images chan image.Image, w fyne.Window) {
	for src := range images {
		fmt.Println("I CAN SEE THE IMAGE")
		img := canvas.NewImageFromImage(src)
		img.FillMode = canvas.ImageFillOriginal
		w.SetContent(img)
	}
}

func keyLoggerRoutine(images chan image.Image) {
	keys := setupKetLogger()
	for key := range keys {
		if key == '(' {
			castRod()
			time.Sleep(2 * time.Second)

			previous, _ := screenshot.Capture(740, 387, 50, 30)
			for {
				time.Sleep(200 * time.Millisecond)
				current := takeScreenShot(740, 387, 50, 30)
				images <- current
				if !equals(previous, current) {
					break
				}
				println("Wait")
			}
		} else if key == ')' {
			for i := 0; i < 40; i++ {
				time.Sleep(200 * time.Millisecond)
				images <- takeScreenShot(740, 387, 200, 100)
			}
		}
	}
}

func takeScreenShot(x, y, width, height int) *image.RGBA {
	start := time.Now()
	current, _ := screenshot.Capture(x, y, width, height)
	fmt.Println("Screenshot took :" + time.Now().Sub(start).String())
	return current
}

func equals(previous *image.RGBA, current *image.RGBA) bool {
	for i, pix := range previous.Pix {
		if current.Pix[i] != pix {
			return false
		}
		return true
	}
	return false
}

func castRod() {
	fmt.Println("CASTING ROD")
	robotgo.Toggle("left")
	robotgo.MilliSleep(1049)
	robotgo.Toggle("left", "up")
	fmt.Println("CASTING ROD END")
}
