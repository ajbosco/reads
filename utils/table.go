package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// NewTable creates new tablewriter.Table
func NewTable(headers ...string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetHeader(headers)

	return table
}
