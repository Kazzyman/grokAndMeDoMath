package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// @formatter:off

var (
			bgsc = canvas.NewRectangle(color.NRGBA{R: 150, G: 180, B: 160, A: 240}) // Light green
			bgwc = canvas.NewRectangle(color.NRGBA{R: 110, G: 160, B: 255, A: 150}) // Light blue, lower number for A: means less opaque, or more transparent
		
		pie float64
		
			outputLabel1 = widget.NewLabel("\nSelect one of the brightly-colored panels to estimate π via featured method...\n\n")
			scrollContainer1 = container.NewVScroll(outputLabel1)
		
			myApp = app.New()
			window1 = myApp.NewWindow("Rick's Pi calculation Demo, set #1")
			currentDone    chan bool
)

func main() {
	countAndLogSLOC() 
	calculating = false
	myApp.Settings().SetTheme(theme.LightTheme()) 
	window1.Resize(fyne.NewSize(1900, 1600))

	scrollContainer1 = container.NewVScroll(outputLabel1)
	
	scrollContainer1.SetMinSize(fyne.NewSize(1900, 930)) // was 1900, 1050 before adding the scoreBoard, which required this adjustment 
	
	outputLabel1.Wrapping = fyne.TextWrapWord 

		coloredScroll := container.NewMax(bgsc, scrollContainer1) 
	
		windowContent := container.NewMax(bgwc, coloredScroll) 

/*
.
.
 */
	terminalDisplay := widget.NewTextGrid()
	terminalDisplay.SetText("Terminal Output:\n\nWaiting for calculation...")

	// Button only being used as a title-label for nifty_scoreBoard
	calcButton := widget.NewButton("Calculate Pi on a ScoreBoard", func() {
		updateOutput1("\n- * - * - that button does nothing - * - * -\n\n")
	})

	contentForScoreBoard := container.NewVBox(
		calcButton,
		terminalDisplay,
	)
/*
.
.
 */
	// Custom colored ::: Buttons1 - - - - - - - - - follow - - - - - - - - - - - v v v v v v v v v - - - - - - 
	/*
	.
	.
	 */
	archimedesBtn1 := NewColoredButton(
	"Archimedes method for finding π, modified by Richard Woolley\n" +
		"easy to understand geometric method using big.Float variables\n" +
		"produces 3,012 digits of delicious Pi in under a minute, 230BCE\n" +
		"             -*-*-*- Rick's personal favorite -*-*-*-          ",
		color.RGBA{255, 110, 110, 215},
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			// We want to cause the button that corresponds to the currently executing method to remain bright, while the other buttons remain dimmed during calculating ...
			for _, btn := range archiBut { 
				calculating = true 
				btn.Enable() 
			} //  ... we do it like this ^ ^ ^ because, we are inside the creation of archimedesBtn1 [ it was the simplest way to deal with a timing and scoping issue ]
			currentDone = make(chan bool) 
			updateOutput1("\nRunning ArchimedesBig...\n\n")
			go func(done chan bool) {
					defer func() { 
						calculating = false
						updateOutput1("Calculation definitely finished; possibly aborted\n")
					}()
				ArchimedesBig(updateOutput1, done) // ::: func < - - - - - - - - - - - - - < -
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
			}(currentDone)
		},
	)
	/*
	.
	.
	 */
	
	JohnWallisBtn1 := NewColoredButton(
	"John Wallis infinite series -- 40 billion iterations -- runs 5m30s\n" +
		"π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ...\n" +
		"only manages to do 10 digits of Pi in well-over five minutes\n" +
		"an infinite series circa 1655    --- served here by Rick Woolley ---",
		color.RGBA{110, 110, 255, 185}, 
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range walisBut {
				btn.Enable()
			}
			currentDone = make(chan bool) 
			updateOutput1("\nRunning John Wallis...\n\n")
			go func(done chan bool) { 
					defer func() { 
						calculating = false 
						updateOutput1("Calculation definitely finished; possibly aborted\n")
					}()
				JohnWallis(updateOutput1, done) 
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
			}(currentDone)
			fmt.Printf("here at the end of JohnWallisBtn1 calculating is %t\n", calculating)
		},
	)
/*
.
.
 */

	SpigotBtn1 := NewColoredButton(
	"The Spigot Algorithm, a Leibniz series. Served hot, bite by byte\n" +
		"spits out a nearly-unlimited, continuous stream of Pi goodness\n" +
		"This trick made possible by a bit of code mooched off of GitHub\n" +
		"bakes π without using any floating-point arithmetic",
		color.RGBA{255, 255, 100, 235},
		
		func() {
			var spigotDigits int = 1460 // to resolve a scoping issue 
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range spigotBut {
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool)
			updateOutput1("\nRunning The Spigot...\n\n")
			
			// In the case of the spigot, retain this dialog please. 
			showCustomEntryDialog(
				"Input Desired number of digits",
				"Any number less than 1461",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input) 
						val, err := strconv.Atoi(input)
						if err != nil { // we may force val to become 460, or leave it alone ...
							fmt.Println("Error converting input:", err)
							updateOutput1("\nInvalid input, using default 1460 digits\n")
							val = 1460
						} else if val <= 0 {
							updateOutput1("\nInput must be positive, using default 1460 digits\n")
							val = 1460
						} else if val > 1460 {
							updateOutput1("\nInput must be less than 1461 -- using default of 1460 digits\n")
							val = 1460 
						} else {
							spigotDigits = val 
						}
						
						go func(done chan bool) { 
								defer func() { 
									calculating = false 
									updateOutput1("\nCalculation definitely finished; possibly aborted\n")
								}()
							TheSpigot(updateOutput1, spigotDigits, done) 
								calculating = false
								for _, btn := range buttons1 {
									btn.Enable()
								}
						}(currentDone)
					} else {
						// dialog canceled 
						updateOutput1("\nspigot calculation canceled, make another selection\n")
						for _, btn := range buttons1 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					}
				},
			)
		}, 
	)
	/*
	.
	.
	 */ 
	
	ChudnovskyBtn1 := NewColoredButton(
	"Chudnovsky -- by David & Gregory Chudnovsky -- late 1980s\n" +
		"extremely efficient, quickly bakes world-record quantities of Pi\n" +
		"this algorithm is a rapidly converging infinite series which\n" +
		"leverages properties of j-invariant from elliptic function theory",
		color.RGBA{100, 255, 100, 215}, 
		
		func() {
			// 
			var chudDigits int
				if calculating {
					return
				}
				calculating = true
				for _, btn := range buttons1 {
					btn.Disable()
				}
				for _, btn := range chudBut { 
					calculating = true 
					btn.Enable() 
				}
			currentDone = make(chan bool) // ::: New channel per run
			updateOutput1("\nRunning Chudnovsky...\n\n")
	
			// unsure about retaining dialog for chud
			showCustomEntryDialog(
				"Input Desired number of digits",
				"Any number less than 49,999",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input)
						val, err := strconv.Atoi(input)
						if err != nil {
							fmt.Println("Error converting input:", err)
							updateOutput1("Invalid input, using default 49,000 digits")
						} else if val <= 0 {
							updateOutput1("Input must be positive, using default 49000 digits")
						} else if val > 50000 {
							updateOutput1("Input must be less than 50,000 -- using default of 49,000 digits")
						} else {
							chudDigits = val
						}
						go func(done chan bool) { 
								defer func() { 
									calculating = false
									updateOutput1("Calculation definitely finished; possibly aborted\n")
								}()
							chudnovskyBig(updateOutput1, chudDigits, done) 
								calculating = false
								for _, btn := range buttons1 {
									btn.Enable()
								}
						}(currentDone)
					} else {
						// dialog canceled 
							updateOutput1("chudnovsky calculation canceled, make another selection")
							for _, btn := range buttons1 {
								btn.Enable()
							}
							calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					}
				},
			)
		},
	)
	/*
	.
	.
	 */

	MontyBtn1 := NewColoredButton(
		"Monte Carlo method for converging on π  --  big floats, & float64\n" +
			"Flavor: no fancy equations are used, only Go's pure randomness\n" +
			"4 digits of pi in 21s; 7 digits possible in 1h30m with a 119k grid\n" +
			"                   -*-*- Rick's second-favorite method -*-*-     ",
		color.RGBA{255, 255, 100, 235},

		func() {
			var MontDigits string
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range montBut {
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool) 
			updateOutput1("\nRunning Monte Carlo ...\n\n")

			showCustomEntryDialog(
				"Input Desired number of grid elements",
				"max 120,000; 10,000 will produce 4 pi digits, 110,00 may get you 5 digits",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input)
						val, err := strconv.Atoi(input) 
						if err != nil {
							fmt.Println("Error converting input:", err)
							updateOutput1("Invalid input, using default 10,000 digits")
						} else if val <= 1 {
							updateOutput1("Input must be greater than 1, using default 10,000 digits")
						} else if val > 120000 {
							updateOutput1("Input must be less than 120,001 -- using default of 10,000 digits")
						} else {
							MontDigits = strconv.Itoa(val) 
						}
						go func(done chan bool) { 
								defer func() { 
									calculating = false 
									updateOutput1("Calculation definitely finished; possibly aborted\n")
								}()
							Monty(updateOutput1, MontDigits, done)
								calculating = false
								for _, btn := range buttons1 {
									btn.Enable()
								}
						}(currentDone)
					} else {
						// dialog canceled 
						updateOutput1("Monte Carlo calculation canceled, make another selection")
						for _, btn := range buttons1 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					}
				},
			)
		},
	)
	/*
		.
		.
	*/
	
	GaussBtn1 := NewColoredButton(
	"Gauss-Legendre -- C F Gauss, refined by Adrien-Marie Legendre\n" +
		"π ≈ (aₙ + bₙ)² / (4 tₙ)\n" +
		"only manages to do 10 digits of Pi in well-over five minutes\n" +
		"an infinite series circa 1655    --- served here by Rick Woolley ---",
		color.RGBA{100, 255, 100, 215},
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range gaussBut {
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool)
			updateOutput1("\nRunning Gauss...\n\n")
			go func(done chan bool) { 
					defer func() {  
						calculating = false 
						updateOutput1("Calculation definitely finished; possibly aborted\n")
					}()
				Gauss_Legendre(updateOutput1, done) // ::: func < - - - - - - - - - - - - - < -
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
			}(currentDone)
		},
	)
	/*
	.
	.
	 */
	
	CustomSeriesBtn1 := NewColoredButton(
	"Custom series -- I don't remember where it's from ... \n" +
		"but it is very quick -- 4s gets us 9 digits of Pi\n" +
		"π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...",
		color.RGBA{255, 120, 120, 215}, // Greenish for variety
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range customBut {
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool) 
			updateOutput1("\nRunning Custom Series ...\n\n")
			go func(done chan bool) { 
					defer func() { 
						calculating = false
						updateOutput1("Calculation definitely finished; possibly aborted\n")
					}()
				CustomSeries(updateOutput1, done)
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
			}(currentDone)
		},
	)
	/*
	.
	.
	 */
	
	GregoryLeibnizBtn1 := NewColoredButton(
	"Gregory-Leibniz -- runs 20sec -- gives 10 digits of Pi\n" +
		"James Gregory 1638–1675  Gottfried Wilhelm Leibniz 1646-1716\n" +
		"π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ...) ",
		color.RGBA{110, 110, 255, 185},
		
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range gottfieBut { 
				calculating = true
				btn.Enable()
			}
			currentDone = make(chan bool)
			updateOutput1("\nRunning Gregory-Leibniz...\n\n")
			go func(done chan bool) { 
				defer func() { 
					calculating = false
					updateOutput1("Calculation definitely finished; possibly aborted\n")
				}()
				GregoryLeibniz(updateOutput1, done) 
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}(currentDone)
		},
	)
	/*
	.
	.
	 */
	
	archiBut = []*ColoredButton{archimedesBtn1} // All of these 8 *But items are a trick/kluge used as bug preventions -- to keep methods from being started or restarted in parallel (over-lapping) 
	walisBut = []*ColoredButton{JohnWallisBtn1} 
	spigotBut = []*ColoredButton{SpigotBtn1} 
	chudBut = []*ColoredButton{ChudnovskyBtn1} 
	montBut = []*ColoredButton{MontyBtn1} 
	gaussBut = []*ColoredButton{GaussBtn1}
	customBut = []*ColoredButton{CustomSeriesBtn1}
	gottfieBut = []*ColoredButton{GregoryLeibnizBtn1}
	
	buttons1 = []*ColoredButton{archimedesBtn1, JohnWallisBtn1, SpigotBtn1, ChudnovskyBtn1, MontyBtn1, GaussBtn1, CustomSeriesBtn1, GregoryLeibnizBtn1,} // used only for range btn.Enable()

		content1 := container.NewVBox(widget.NewLabel("\nSelect a method to estimate π:\n"),
			container.NewGridWithColumns(4, archimedesBtn1, JohnWallisBtn1, SpigotBtn1,
				ChudnovskyBtn1, MontyBtn1, GaussBtn1, CustomSeriesBtn1, GregoryLeibnizBtn1, contentForScoreBoard),
			windowContent,
		)
/*
.
.
 */
	// ::: drop-down menus -- same for all windows  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  -  -  --  
	logFilesMenu := fyne.NewMenu("Log-Files",
		fyne.NewMenuItem("View Log 1", func() { dialog.ShowInformation("Log Files", "Viewing Log 1", window1) }),
		fyne.NewMenuItem("View Log 2", func() { dialog.ShowInformation("Log Files", "Viewing Log 2", window1) }),
	)
	additionalMethodsMenu := fyne.NewMenu("Other-Methods",
		fyne.NewMenuItem("Home-Page (Pi methods)", func() { window1.Show() }),
		fyne.NewMenuItem("Second-page of Pi methods", func() { createWindow2(myApp).Show() }), 
		fyne.NewMenuItem("Odd Pi calculators", func() { createWindow3(myApp).Show() }),
		fyne.NewMenuItem("Misc Maths", func() { createWindow4(myApp).Show() }), // maybe our roots demo will live here some day
	)
	optionsMenu := fyne.NewMenu("Options",
		fyne.NewMenuItem("Begin the ScoreBoard of Pi", func() {
			
			// dialog.ShowInformation("ScoreBoard", "Use Abort in Menu\nPrior to dismissing with OK", window1)
			if calculating {
				fmt.Println("Calculation already in progress")
				return
			}
			calculating = true
			currentDone = make(chan bool)
			termsCount = 0

			go func(done chan bool) {
				defer func() {
					calculating = false
					terminalDisplay.SetText(fmt.Sprintf("Terminal Output:\n\nCalculation stopped.\nFinal Pi: %.11f\nTerms: %d", <-pichan, termsCount))
				}()

				pie := nifty_scoreBoardG(func(text string) {
					terminalDisplay.SetText(text)
				}, done)

				if pie != 0.0 {
					terminalDisplay.SetText(fmt.Sprintf("Terminal Output:\n\nComputed Value of Pi:\t\t%.11f\n# of Nilakantha Terms:\t\t%d", pie, termsCount))
				}
			}(currentDone)
		}),
		fyne.NewMenuItem("Abort any currently executing method", func() {
			if currentDone == nil {
				updateOutput1("\nNo active calculation to abort, no such currentDone channel exists\n")
				fmt.Println("No active calculation to abort, no such currentDone channel exists")
				return
			}
			select {
			case <-currentDone:
				updateOutput1("\nMenu select determined that currentDone-chan had already been closed; all Goroutines were PREVIOUSLY notified to terminate\n")
				fmt.Println("Menu select determined that currentDone-chan had already been closed; all Goroutines were PREVIOUSLY notified to terminate")
			default:
				close(currentDone)
				updateOutput1("\nTermination signals were sent to all current processes that may be listening\n")
				fmt.Println("Termination signals were sent to all current processes that may be listening")
			}
		}),
		fyne.NewMenuItem("Show the terminal -- Cmd+Tab to return", func() {
			err := openTerminal()
			if err != nil {
				fmt.Println(err)
				return
			}		}),
	)

	mainMenu := fyne.NewMainMenu(logFilesMenu, additionalMethodsMenu, optionsMenu)
	window1.SetMainMenu(mainMenu)
	
	windowWithBackground := container.NewMax(bgwc, content1)
	
	window1.SetContent(windowWithBackground)
	
	window1.ShowAndRun() 
}
