package uiutil

import (
	"fmt"
	"strings"

	"github.com/rodaine/table"
	"krzywicki.pro/freqr/internal/radio/models"
)

func PrepareChannelsTable(channels []*models.Channel) table.Table {
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}
	tbl := table.New("ID", "Name", "Freq (MHz)", "Mode", "Delay", "Lockout", "Priority")

	for _, ch := range channels {
		tbl.AddRow(ch.Index, ch.Name, ch.Freq, ch.Mod, ch.Delay, ch.Lockout, ch.Priority)
	}

	return tbl
}
