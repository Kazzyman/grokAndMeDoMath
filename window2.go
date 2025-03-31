package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// @formatter:off

var (
	bgsc2           = canvas.NewRectangle(color.NRGBA{R: 130, G: 160, B: 250, A: 140})
	bgwc2           = canvas.NewRectangle(color.NRGBA{R: 110, G: 255, B: 160, A: 150})
	outputLabel2    = widget.NewLabel("Classic Pi calculators, make a selection")
	scrollContainer2 = container.NewScroll(outputLabel2)
	window2         = myApp.NewWindow("Rick's Pi calculation Demo, set #2")
	mgr             = NewTrafficManager(outputLabel2) // Move mgr to global scope
)

func createWindow2(myApp fyne.App) fyne.Window {
	window2.Resize(fyne.NewSize(1900, 1600))
	outputLabel2.Wrapping = fyne.TextWrapWord
	scrollContainer2.SetMinSize(fyne.NewSize(1900, 1000))

	coloredScroll2 := container.NewMax(bgsc2, scrollContainer2)

	// Add entry widgets here
	radicalEntry := widget.NewEntry()
	radicalEntry.SetPlaceHolder("Enter radical index (e.g., 2 or 3)")
	workEntry := widget.NewEntry()
	workEntry.SetPlaceHolder("Enter number to find the root of")
	scroll := widget.NewRichText() // For dual output

	RootsBtn2 := SetupRootsDemo(mgr, radicalEntry, workEntry, scroll) // Pass entries and scroll
	rootBut2 = []*ColoredButton{RootsBtn2}
	buttons2 = rootBut2 // Update global buttons2 for consistency

	content2 := container.NewVBox(
		widget.NewLabel("\nSelect a method to estimate π:\n"),
		outputLabel2,           // Show the output label
		container.NewHScroll(scroll), // Scrollable output
		radicalEntry,
		workEntry,
		container.NewGridWithColumns(4, RootsBtn2),
		coloredScroll2,
	)
	windowContent2 := container.NewMax(bgwc2, content2)
	window2.SetContent(windowContent2)
	return window2
}


// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow3(myApp fyne.App) fyne.Window {
	// Planning to have similar structure to createWindow2
	window3 := myApp.NewWindow("Odd Pi calculators")
	window3.Resize(fyne.NewSize(1900, 1600))
	outputLabel3 := widget.NewLabel("Odd Pi calculators, make a selection")
	outputLabel3.Wrapping = fyne.TextWrapWord
	scrollContainer3 := container.NewScroll(outputLabel3)
	scrollContainer3.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer3 := container.NewGridWithColumns(4,
		widget.NewButton("Button 9", func() {}),
		widget.NewButton("Button 10", func() {}),
		widget.NewButton("Button 11", func() {}),
		widget.NewButton("Button 12", func() {}),
		widget.NewButton("Button 13", func() {}),
		widget.NewButton("Button 14", func() {}),
		widget.NewButton("Button 15", func() {}),
		widget.NewButton("Button 16", func() {}),
	)
	content3 := container.NewVBox(buttonContainer3, scrollContainer3)
	window3.SetContent(content3)
	return window3
}

// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow4(myApp fyne.App) fyne.Window {
	// Planning to have similar structure to createWindow2
	window4 := myApp.NewWindow("Misc Maths")
	window4.Resize(fyne.NewSize(1900, 1600))
	outputLabel4 := widget.NewLabel("Misc Maths, make a selection")
	outputLabel4.Wrapping = fyne.TextWrapWord
	scrollContainer4 := container.NewScroll(outputLabel4)
	scrollContainer4.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer4 := container.NewGridWithColumns(4,
		widget.NewButton("Button 17", func() {}), widget.NewButton("Button 18", func() {}), widget.NewButton("Button 19", func() {}), widget.NewButton("Button 20", func() {}),
		widget.NewButton("Button 21", func() {}), widget.NewButton("Button 22", func() {}), widget.NewButton("Button 23", func() {}), widget.NewButton("Button 24", func() {}),
	)
	content4 := container.NewVBox(buttonContainer4, scrollContainer4)
	window4.SetContent(content4)
	return window4
}
