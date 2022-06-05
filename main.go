package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	a := app.New()
	w := a.NewWindow("Qr code Generator")
	w.Resize(fyne.NewSize(400, 400))
	url := widget.NewEntry()
	url.SetPlaceHolder("Enter url ...")
	file_name := widget.NewEntry()
	file_name.SetPlaceHolder("Enter file name ...")
	img := canvas.NewImageFromFile("./assets.jpeg")
	img.FillMode = canvas.ImageFillOriginal
	btn := widget.NewButton("Create", func() {
		fmt.Printf("%s.png", file_name.Text)
		err := qrcodes(url.Text, file_name.Text)
		img.File = file_name.Text + ".jpeg"
		img.Refresh()
		if err != nil {
			fmt.Printf("could not generate QRCode: %v", err)
		}
	})
	w.SetContent(container.NewVBox(
		url,
		file_name,
		btn,
		img,
	))
	w.ShowAndRun()
}

func qrcodes(text string, filname string) error {
	qrcs, err := qrcode.New(text)
	if err != nil {
		return err
	}

	w, err := standard.New("./" + filname + ".jpeg")
	if err != nil {
		return err
	}

	// save file
	if err = qrcs.Save(w); err != nil {
		return err
	}

	return nil
}
