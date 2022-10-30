package main

import (
	"fmt"

	"github.com/mattn/go-shellwords"
)

const (
	TERMINAL = "terminal"
	PIPE     = "pipe"
)

func main() {
	var inputText string

	fmt.Scan(&inputText)

	args, err := shellwords.Parse(inputText)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", args)

	// app := tview.NewApplication()
	// flex := tview.NewFlex().
	// 	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
	// 		AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlue).SetTitle("Keyword").SetTitleAlign(tview.AlignLeft), 3, 1, false).
	// 		AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlue).SetTitle(inputText).SetTitleAlign(tview.AlignLeft), 0, 1, false), 0, 2, false)
	// if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
	// 	panic(err)
	// }
}
