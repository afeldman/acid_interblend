package main

import (
	"acid_interblend/acid"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Parse()

	//output_volume_s1, output_volume_s2 := acid.AcidCrossCalculation(*solution1, *solution2, *result_solution, *volume)

	a := app.New()
	win := a.NewWindow("Acid Interbend")

	solution1 := widget.NewLabel("Acid:")
	input_s1 := widget.NewEntry()
	input_s1.SetText("85")

	solution2 := widget.NewLabel("Medium (Water=0%):")
	input_s2 := widget.NewEntry()
	input_s2.SetText("0")

	result_solution := widget.NewLabel("Resulting Solution")
	input_result := widget.NewEntry()
	input_result.SetText("0")

	volume := widget.NewLabel("Resulting Volume")
	input_volume := widget.NewEntry()
	input_volume.SetText("1000")

	grid := container.NewGridWithColumns(3,
		solution1, input_s1, widget.NewLabel("%"),
		solution2, input_s2, widget.NewLabel("%"),
		result_solution, input_result, widget.NewLabel("%"),
		volume, input_volume, widget.NewLabel("ml"))

	result_solution_1 := widget.NewLabel("ml")
	result_solution_2 := widget.NewLabel("ml")

	reslut_row := container.NewGridWithColumns(2,
		widget.NewLabel("Volume S1"), result_solution_1,
		widget.NewLabel("Volume S2"), result_solution_2)

	content := container.NewVBox(grid, reslut_row, widget.NewButton("Calc", func() {
		s1, err := strconv.ParseFloat(input_s1.Text, 64)
		log.Println(s1, err, reflect.TypeOf(s1))

		s2, err := strconv.ParseFloat(input_s2.Text, 64)
		log.Println(s2, err, reflect.TypeOf(s2))

		rs, err := strconv.ParseFloat(input_result.Text, 64)
		log.Println(rs, err, reflect.TypeOf(rs))

		vol, err := strconv.ParseFloat(input_volume.Text, 64)
		log.Println(vol, err, reflect.TypeOf(vol))

		output_volume_s1, output_volume_s2 := acid.AcidCrossCalculation(s1, s2, rs, vol)

		result_solution_1.SetText(fmt.Sprintf("%.2f ml", output_volume_s1))
		result_solution_2.SetText(fmt.Sprintf("%.2f ml", output_volume_s2))

	}))

	win.SetContent(content)

	win.ShowAndRun()

}
