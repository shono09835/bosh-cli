package cmd

import (
	boshdir "github.com/shono09835/bosh-cli/v7/director"
	boshui "github.com/shono09835/bosh-cli/v7/ui"
	boshtbl "github.com/shono09835/bosh-cli/v7/ui/table"
)

type ConfigTable struct {
	Config boshdir.Config
	UI     boshui.UI
}

func (t ConfigTable) Print() {
	table := boshtbl.Table{
		Content: "config",

		Header: []boshtbl.Header{
			boshtbl.NewHeader("ID"),
			boshtbl.NewHeader("Type"),
			boshtbl.NewHeader("Name"),
			boshtbl.NewHeader("Created At"),
			boshtbl.NewHeader("Content"),
		},

		Notes: []string{},

		FillFirstColumn: true,

		Transpose: true,
	}

	table.Rows = append(table.Rows, []boshtbl.Value{
		boshtbl.NewValueString(t.Config.ID),
		boshtbl.NewValueString(t.Config.Type),
		boshtbl.NewValueString(t.Config.Name),
		boshtbl.NewValueString(t.Config.CreatedAt),
		boshtbl.NewValueString(t.Config.Content),
	})

	t.UI.PrintTable(table)
}
