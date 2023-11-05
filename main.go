package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const appName string = "GoLife"

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow(appName)

	appLogo := widget.NewLabel("LOGO HERE")

	startButton := widget.NewButton("Go life!", func() {
		startNewGame(myWindow)
	})

	settingsButton := widget.NewButton("Settings", func() {
		Settings(myWindow)
	})

	quitButton := widget.NewButton("Quit", func () {
		myApp.Quit()
	})

	menuContainer := container.NewVBox(appLogo, startButton, settingsButton, quitButton)

	myWindow.SetContent(menuContainer)

	myWindow.ShowAndRun()
}

func startNewGame(window fyne.Window) {
	dialog.ShowInformation("Go Life", "Generating Life...", window)
}

func Settings(window fyne.Window) {
	dialog.ShowInformation("Setting X", "ON", window)
}
