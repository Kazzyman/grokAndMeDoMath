package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

// @formatter:off

// fyneFunc(fmt.Sprintf("   ... via the Gauss–Legendre algorithm ... \n"))

var (
	Table_of_perfect_Products []int
	sortedResults             []Results
	perfectResult2            float64
	perfectResult3            float64
	diffOfLarger              int
	diffOfSmaller             int
	precisionOfRoot           int
)

type Results struct {
	result float64
	pdiff  float64
}

// SetupRootsDemo sets up the roots demo UI and returns the button for window2
func SetupRootsDemo(mgr *TrafficManager, radicalEntry, workEntry *widget.Entry, fyneFunc func(string)) *ColoredButton {
	rootsBtn := NewColoredButton(
		"Roots\n" +
			"2 or 3\n" +
			"any integer\n" +
			"                   -*-*- Rick's own-favorite method -*-*-     ",
		color.RGBA{255, 255, 100, 235},
		func() {
			if mgr.IsCalculating() {
				return
			}
			radical, err := strconv.Atoi(radicalEntry.Text)
			if err != nil || (radical != 2 && radical != 3) {
				updateOutput2("Invalid radical: enter 2 or 3\n")
				fyneFunc(fmt.Sprintf("Invalid radical: enter 2 or 3\n"))
				return
			}
			workPiece, err := strconv.Atoi(workEntry.Text)
			if err != nil || workPiece < 0 {
				updateOutput2("Invalid number: enter a non-negative integer\n")
				fyneFunc(fmt.Sprintf("Invalid number: enter a non-negative integer\n"))
				return
			}
			mgr.SetRadical(radical)
			mgr.SetWorkPiece(workPiece)
			mgr.SetCalculating(true)
			for _, btn := range buttons2 {
				btn.Disable()
			}
			for _, btn := range rootBut2 {
				btn.Enable()
			}
			go func() {
					defer func() {
						mgr.Reset()
						for _, btn := range buttons2 {
							btn.Enable()
						}
					}()
				xRootOfy(updateOutput2) // ::: formatted to highlight the meat
					mgr.SetCalculating(false)
			}()
		},
	)
	return rootsBtn
}

func xRootOfy(fyneFunc func(string)) {
	usingBigFloats = false
	var index = 0
	TimeOfStartFromTop := time.Now()

	radicalIndex := mgr.GetRadical()
	workPiece := mgr.GetWorkPiece()

	radicalIndex, workPiece = setStateOfSquareOrCubeRoot(mgr, radicalIndex, workPiece, updateOutput2)
	mgr.SetRadical(radicalIndex)
	mgr.SetWorkPiece(workPiece)

	// updateOutput2("\n\nBuilding table...\n")
	fyneFunc(fmt.Sprintf("\n\nBuilding table...\n"))
	buildTableOfPerfectProducts(radicalIndex)
	// updateOutput2("Table built, starting calculation...\n")
	fyneFunc(fmt.Sprintf("Table built, starting calculation...\n"))
	startBeforeCall := time.Now()
	for index < 400000 {
		if mgr.ShouldStop() {
			updateOutput2("Calculation aborted\n")
			fyneFunc(fmt.Sprintf("Calculation aborted\n"))
			return
		}
		readTheTableOfPP(index, startBeforeCall, radicalIndex, workPiece, updateOutput2)
		handlePerfectSquaresAndCubes(TimeOfStartFromTop, radicalIndex, workPiece, mgr)
		if diffOfLarger == 0 || diffOfSmaller == 0 {
			break
		}
		if index%80000 == 0 && index > 0 {
			updateOutput2(fmt.Sprintf("%d iterations completed...\n", index))
			updateOutput2(fmt.Sprintf("... still working ...\n")) // ok
			
			fmt.Printf("%d iterations completed...\n", index)
			fyneFunc(fmt.Sprintf("%d iterations completed...\n", index))
			
			fmt.Println(index, "... still working ...")
			fyneFunc(fmt.Sprintf("\n... still working ...\n"))
		}
		index += 2
	}
	fmt.Println("Loop completed at index:", index) // Debug
	
	// ::: Show the final result
	fmt.Println("Entering result block, perfectResult2:", perfectResult2, "perfectResult3:", perfectResult3) // Debug
	
	t_s2 := time.Now()
	elapsed_s2 := t_s2.Sub(TimeOfStartFromTop)
	if perfectResult2 == 0 && perfectResult3 == 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic in result block:", r)
				updateOutput2("\nError calculating result\n")
			}
		}()
		fileHandle, err31 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err31)
		defer fileHandle.Close()

		Hostname, _ := os.Hostname()
		fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of perfect Products -- on %s \n", radicalIndex, workPiece, Hostname)
		fmt.Fprint(fileHandle, "was run on: ", time.Now().Format(time.ANSIC), "\n")
		fmt.Fprintf(fileHandle, "%d was total Iterations \n", index)

		fmt.Println("Sorting results...") // Debug
		sort.Slice(sortedResults, func(i, j int) bool { return sortedResults[i].pdiff < sortedResults[j].pdiff })
		fmt.Println("Sorted results, length:", len(sortedResults)) // Debug

		if len(sortedResults) > 0 {
			if radicalIndex == 2 {
				result := fmt.Sprintf("%0.9f, it's the best approximation for the Square Root of %d", sortedResults[0].result, workPiece)
				fmt.Println("Updating GUI with:", result) // Debug
				updateOutput2(result)
				fmt.Println("GUI updated, printing to console...") // Debug
				fmt.Printf("%s\n", result)
				fyneFunc(fmt.Sprintf("\nSquare-root Result is: %s\n", result))
				fmt.Println("Writing to file...") // Debug
				fmt.Fprintf(fileHandle, "%s \n", result)
				fmt.Println("File written") // Debug
			}
			if radicalIndex == 3 {
				result := fmt.Sprintf("%0.9f, it's the best approximation for the Cube Root of %d", sortedResults[0].result, workPiece)
				fmt.Println("Updating GUI with:", result) // Debug
				updateOutput2(result)
				fmt.Println("GUI updated, printing to console...") // Debug
				fmt.Printf("%s\n", result)
				fyneFunc(fmt.Sprintf("\nCube-root Result is: %s\n", result))
				fmt.Println("Writing to file...") // Debug
				fmt.Fprintf(fileHandle, "%s \n", result)
				fmt.Println("File written") // Debug
			}
		} else {
			updateOutput2(fmt.Sprintf("No results found within precision %d after %d iterations", precisionOfRoot, index))
			fmt.Printf("No results found within precision %d after %d iterations\n", precisionOfRoot, index)
			fyneFunc(fmt.Sprint("\nNo results found within precision %d after %d iterations\n", precisionOfRoot, index))
		}

		TotalRun := elapsed_s2.String()
		fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
		fmt.Printf("Calculation completed in %s\n", elapsed_s2)
		fyneFunc(fmt.Sprintf("\nCalculation completed in %s\n", elapsed_s2))
	} else {
		fmt.Println("Skipped result block due to perfect result detection") // Debug
	}
}

func readTheTableOfPP(index int, startBeforeCall time.Time, radicalIndex, workPiece int, fyneFunc func(string)) {
	smallerPerfectProductOnce := Table_of_perfect_Products[index]
	RootOfsmallerPerfectProductOnce := Table_of_perfect_Products[index+1]

	iter := 0
	for iter < 410000 && index+2 < len(Table_of_perfect_Products) {
		iter++
		index += 2
		largerPerfectProduct := Table_of_perfect_Products[index]

		if largerPerfectProduct > smallerPerfectProductOnce*workPiece {
			ProspectiveHitOnLargeSide := largerPerfectProduct
			rootOfProspectiveHitOnLargeSide := Table_of_perfect_Products[index+1]
			ProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-2]
			rootOfProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-1]

			diffOfLarger = ProspectiveHitOnLargeSide - workPiece*smallerPerfectProductOnce
			diffOfSmaller = workPiece*smallerPerfectProductOnce - ProspectiveHitOnSmallerSide

			if diffOfLarger == 0 {
				fmt.Println(colorCyan, "\n The", radicalIndex, "root of", workPiece, "is", colorGreen,
					float64(rootOfProspectiveHitOnLargeSide)/float64(RootOfsmallerPerfectProductOnce), colorReset, "\n")
				fyneFunc(fmt.Sprintf("\n The %d root of %d is %0.33f\n\n", radicalIndex, workPiece, float64(rootOfProspectiveHitOnLargeSide)/float64(RootOfsmallerPerfectProductOnce)))
				perfectResult3 = math.Cbrt(float64(workPiece))
				break
			}
			if diffOfSmaller == 0 {
				fmt.Println(colorCyan, "\n The", radicalIndex, "root of", workPiece, "is", colorGreen,
					float64(rootOfProspectiveHitOnSmallerSide)/float64(RootOfsmallerPerfectProductOnce), colorReset, "\n")
				fyneFunc(fmt.Sprintf("\n The %d root of %d is %0.33f\n\n", radicalIndex, workPiece, float64(rootOfProspectiveHitOnSmallerSide)/float64(RootOfsmallerPerfectProductOnce)))
				perfectResult2 = math.Sqrt(float64(workPiece))
				perfectResult3 = math.Cbrt(float64(workPiece))
				break
			}

			if diffOfLarger < precisionOfRoot {
				result := float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce)
				pdiff := float64(diffOfLarger) / float64(ProspectiveHitOnLargeSide)
				sortedResults = append(sortedResults, Results{result: result, pdiff: pdiff})
				fmt.Printf("Found large prospect at index %d: result=%f, diff=%d\n", index, result, diffOfLarger) // Debug
				fyneFunc(fmt.Sprintf("Found large prospect at index %d: result=%f, diff=%d\n", index, result, diffOfLarger)) // Debug
			}
			if diffOfSmaller < precisionOfRoot {
				result := float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce)
				pdiff := float64(diffOfSmaller) / float64(ProspectiveHitOnSmallerSide)
				sortedResults = append(sortedResults, Results{result: result, pdiff: pdiff})
				fmt.Printf("Found small prospect at index %d: result=%f, diff=%d\n", index, result, diffOfSmaller) // Debug
				fyneFunc(fmt.Sprintf("Found small prospect at index %d: result=%f, diff=%d\n", index, result, diffOfSmaller)) // Debug
			}
			break // Exit after finding a prospect
		}
	}
	if iter >= 410000 || index+2 >= len(Table_of_perfect_Products) {
		fmt.Printf("No prospect found at index %d, iter %d\n", index, iter) // Debug
	}
}

// handlePerfectSquaresAndCubes reports perfect squares/cubes to file and UI
func handlePerfectSquaresAndCubes(TimeOfStartFromTop time.Time, radicalIndex, workPiece int, mgr *TrafficManager) {
	if diffOfLarger == 0 || diffOfSmaller == 0 {
		t_s1 := time.Now()
		elapsed_s1 := t_s1.Sub(TimeOfStartFromTop)

		fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		check(err1)
		defer fileHandle.Close()

		Hostname, _ := os.Hostname()
		fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of PerfectProducts -- selection #%d on %s \n", radicalIndex, workPiece, 1, Hostname)
		fmt.Fprint(fileHandle, "was run on: ", time.Now().Format(time.ANSIC), "\n")
		fmt.Fprintf(fileHandle, "Total run was %s \n ", elapsed_s1.String())

		if radicalIndex == 2 {
			result := fmt.Sprintf("Perfect square: %0.2f is the %d root of %d", perfectResult2, radicalIndex, workPiece)
			updateOutput2(result)
			fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radicalIndex, workPiece, perfectResult2)
		}
		if radicalIndex == 3 {
			result := fmt.Sprintf("Perfect cube: %0.2f is the %d root of %d", perfectResult3, radicalIndex, workPiece)
			updateOutput2(result)
			fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radicalIndex, workPiece, perfectResult3)
		}
	}
}


// setStateOfSquareOrCubeRoot adjusts precision based on radical and workPiece
func setStateOfSquareOrCubeRoot(mgr *TrafficManager, radicalIndex, workPiece int, fyneFunc func(string)) (int, int) {
	if radicalIndex == 3 {
		if workPiece > 4 {
			precisionOfRoot = 1700
			fmt.Println("\n Default precision is 1700 \n")
			fyneFunc(fmt.Sprintf("\n Default precision is 1700 \n"))
		}
		if workPiece == 2 || workPiece == 11 || workPiece == 17 {
			precisionOfRoot = 600
			fmt.Println("\n resetting precision to 600 \n")
			fyneFunc(fmt.Sprintf("\n resetting precision to 600 \n"))
		}
		if workPiece == 3 || workPiece == 4 || workPiece == 14 {
			precisionOfRoot = 900
			fmt.Println("\n resetting precision to 900 \n")
			fyneFunc(fmt.Sprintf("\n resetting precision to 900 \n"))
		}
	}
	if radicalIndex == 2 {
		precisionOfRoot = 4
	}
	return radicalIndex, workPiece
}

// buildTableOfPerfectProducts builds a table of perfect squares or cubes
func buildTableOfPerfectProducts(radicalIndex int) {
	var PerfectProduct int
	Table_of_perfect_Products = nil // this fixed my bug
	root := 10
	iter := 0
	for iter < 825000 {
		iter++
		root++
		if radicalIndex == 3 {
			PerfectProduct = root * root * root
		}
		if radicalIndex == 2 {
			PerfectProduct = root * root
		}
		Table_of_perfect_Products = append(Table_of_perfect_Products, PerfectProduct)
		Table_of_perfect_Products = append(Table_of_perfect_Products, root)
	}
}
