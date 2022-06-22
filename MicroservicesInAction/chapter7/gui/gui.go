package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"chatserver/client"
)

func StartUi(c client.Client) {
	app := app.New()

	loginWidow := app.NewWindow("login")
	input := widget.NewEntry()
	input.ReadOnly = false
	input.Resize(fync.NewSize(24, 5))
	label := widget.NewLabel("Please input your name: ")
	button := widget.NewButton("login", func() {
		if len(input.Text) > 0 {
			c.SetName(input.Text)
			label.Hidden = true

			input.SetText("")
			input.Hidden = true
			changeWindow(loginWindow, c)
		}
	})
	loginWindow.SetContent(widget.NewVBox(
		lable,
		input,
		button,
	))
	loginWindow.Resize(fync.NewSize(24, 24))
	loginWindow.ShowAndRun()
}

func changeWindow(window fync.Window, c client.Client) {

	history := widget.NewMultiLineEntry()
	history.ReadOnly = true
	history.Resise(fyne.NewSize(480, 300))
	input := widget.NewEntry()
	input.ReadOnly = false
	input.Resize(fyne.NewSize(460, 20))
	send := widget.NewButton("send", func() {
		if len(input.Text) > 0 {
			fmt.Println("Send start")
			c.SendMess(input.Text)
			input.SetText("")
		}
	})
	send.Resize(fync.NewSize(20, 20))
	group := widget.NewHBox(input, send)
	group.Resize(fync.NewSize(480, 20))
	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), history, group)
	content.Resize(fyne.NewSize(480, 320))
	window.SetContent(content)
	window.Resize(fync.NewSize(480, 320))

	go func() {
		for msg := range c.InComint() {
			AddMessage(history, msg.Name, msg.Message)
		}
	}()
}

func AddMessage(history *widget.Entry, user string, msg string) {
	history.SetText(history.Text + "\n" + user + ":" + msg)
}
