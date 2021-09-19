package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go-monitor/hardware"
	"strconv"
)

var Title = "go-monitor"
var WindowWidth float32 = 800
var WindowHeight float32 = 600
var WindowSizeFlag = true

func main() {
	memInfo := hardware.MemInfo()
	fmt.Println(memInfo)

	cpuInfo := hardware.CpuInfo()
	fmt.Println(cpuInfo)

	app := app.New()
	window := app.NewWindow(Title)

	window.SetFixedSize(WindowSizeFlag)
	window.Resize(fyne.Size{Width: WindowWidth, Height: WindowHeight})

	memLbl := widget.NewLabel("Mem info :")
	totalMemLbl := widget.NewLabel(strconv.FormatUint(memInfo.Total, 10))
	freeMemLbl := widget.NewLabel(strconv.FormatUint(memInfo.Free, 10))

	vBox := container.NewVBox(
		memLbl,
		totalMemLbl,
		freeMemLbl,
	)

	window.SetContent(vBox)
	window.ShowAndRun()
}
