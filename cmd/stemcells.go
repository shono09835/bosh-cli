package cmd

import (
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
	boshtbl "github.com/shono09835/bosh-cli/v7/ui/table"
)

type StemcellsCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewStemcellsCmd(ui boshui.UI, director boshdir.Director) StemcellsCmd {
	return StemcellsCmd{ui: ui, director: director}
}

func (c StemcellsCmd) Run() error {
	stemcells, err := c.director.Stemcells()
	if err != nil {
		return err
	}

	table := boshtbl.Table{
		Content: "stemcells",

		Header: []boshtbl.Header{
			boshtbl.NewHeader("Name"),
			boshtbl.NewHeader("Version"),
			boshtbl.NewHeader("OS"),
			boshtbl.NewHeader("CPI"),
			boshtbl.NewHeader("CID"),
		},

		SortBy: []boshtbl.ColumnSort{
			{Column: 0, Asc: true},
			{Column: 1, Asc: false},
		},

		Notes: []string{"(*) Currently deployed"},
	}

	for _, stem := range stemcells {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString(stem.Name()),
			boshtbl.NewValueSuffix(
				boshtbl.NewValueVersion(stem.Version()),
				stem.VersionMark("*"),
			),
			boshtbl.NewValueString(stem.OSName()),
			boshtbl.NewValueString(stem.CPI()),
			boshtbl.NewValueString(stem.CID()),
		})
	}

	c.ui.PrintTable(table)

	return nil
}
