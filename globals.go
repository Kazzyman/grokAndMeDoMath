package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"time"
)

// @formatter:off

// var digits int 
var calculating bool

var archiBut []*ColoredButton
var walisBut []*ColoredButton
var chudBut []*ColoredButton
var chudBut2 []*ColoredButton
var spigotBut []*ColoredButton
var BPPMaxbut2 []*ColoredButton
var montBut []*ColoredButton
var gaussBut []*ColoredButton
var customBut []*ColoredButton
var gottfieBut []*ColoredButton
var nilaBut2 []*ColoredButton

var scoreBut2 []*ColoredButton
var rootBut2 []*ColoredButton

var buttons1 []*ColoredButton // Change to ColoredButton
var buttons2 []*ColoredButton // Change to ColoredButton
var buttons3 []*ColoredButton // Change to ColoredButton
var buttons4 []*ColoredButton // Change to ColoredButton

var copyOfLastPosition int

// convenience globals:

var usingBigFloats = false // a variable of type bool which is passed by many funcs to print Result Stats Long()

var iterationsForMonte16i int
var iterationsForMonte16j int
var iterationsForMonteTotal int
var four float64 // is initialized to 4 where needed
var π float64    // a var can be any character, as in this Pi symbol/character
var LinesPerSecond float64
var LinesPerIter float64
var iterInt64 int64     // to be used primarily in selections which require modulus calculations
var iterFloat64 float64 // to be used in selections which do not require modulus calculations
var t2 time.Time

// The following globals, are used in multiple funcs of case 18: calculate either square or cube root of any integer

var Tim_win float64             // Time Window

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"

// Theme
type myTheme struct { // ::: - -
	Theme fyne.Theme
}
func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{245, 245, 245, 255}
	case theme.ColorNameButton:
		return color.RGBA{200, 200, 200, 255}
	case theme.ColorNameForeground:
		return color.RGBA{0, 0, 0, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{255, 165, 0, 255}
	}
	return m.Theme.Color(name, variant)
}
func (m *myTheme) Font(style fyne.TextStyle) fyne.Resource    { return m.Theme.Font(style) }
func (m *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource { return m.Theme.Icon(name) }
func (m *myTheme) Size(name fyne.ThemeSizeName) float32       { return m.Theme.Size(name) }
