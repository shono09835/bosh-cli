package cmd

import (
	"fmt"

	boshui "github.com/shono09835/bosh-cli/v7/ui"
)

type Diff struct {
	lines [][]interface{}
}

func NewDiff(lines [][]interface{}) Diff {
	return Diff{
		lines: lines,
	}
}

func (d Diff) Print(ui boshui.UI) {
	for _, line := range d.lines {
		lineMod, _ := line[1].(string)

		if lineMod == "added" {
			ui.BeginLinef("+ %s\n", line[0])
		} else if lineMod == "removed" {
			ui.BeginLinef("- %s\n", line[0])
		} else {
			ui.BeginLinef("  %s\n", line[0])
		}
	}
}

func (d Diff) String() string {
	var result string
	for _, line := range d.lines {
		lineMod, _ := line[1].(string)

		if lineMod == "added" {
			result += fmt.Sprintf("+ %s\n", line[0])
		} else if lineMod == "removed" {
			result += fmt.Sprintf("- %s\n", line[0])
		} else {
			result += fmt.Sprintf("  %s\n", line[0])
		}
	}
	return result
}
