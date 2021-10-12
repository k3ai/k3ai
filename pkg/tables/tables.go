package tables

import (
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

)

func List(listStr string,appsResults []string,infraResults []string,bundlesResults []string,commsResults[]string) {
	if listStr == "plugin" {
		if len(appsResults) >0 {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.SetStyle(table.StyleLight)
			t.Style().Color.Header = text.Colors{text.Bold}
			t.SetTitle("APPLICATIONS")
			t.AppendHeader(table.Row{"Name","Description","Type","Tag","Version","Status"})
			t.SetColumnConfigs([]table.ColumnConfig{
				{Number: 1, Colors: text.Colors{text.Color(text.FgYellow),},},
				{Number: 2, WidthMax: 90, WidthMaxEnforcer: text.WrapText},
				{Number: 6, Colors: text.Colors{text.Color(text.FgHiGreen)}},
			})
			limit := 6
			for i:=0; i < len(appsResults)-1;i+= limit{
				batch := appsResults[i:min(i+limit, len(appsResults))]
				t.AppendRow(table.Row{strings.ToUpper(batch[0]),strings.Title(batch[1]),batch[2],batch[3],batch[4],batch[5]})
				t.AppendSeparator()
			}
			t.Render()
		}
	}

	if listStr =="infra"{
		if len(infraResults) > 0 {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.SetStyle(table.StyleLight)
			t.Style().Color.Header = text.Colors{text.Bold}
			t.SetTitle("INFRASTRUCTURE")
			t.AppendHeader(table.Row{"Name","Description","Type","Tag","Version","Status"})
			t.SetColumnConfigs([]table.ColumnConfig{
				{Number: 1, Colors: text.Colors{text.Color(text.FgYellow),},},
				{Number: 2, WidthMax: 90, WidthMaxEnforcer: text.WrapText},
				{Number: 6, Colors: text.Colors{text.Color(text.FgHiGreen)}},
			})
			limit := 6
			for i:=0; i < len(infraResults)-1;i+= limit{
				batch := infraResults[i:min(i+limit, len(infraResults))]
				t.AppendRow(table.Row{strings.ToUpper(batch[0]),strings.Title(batch[1]),batch[2],batch[3],batch[4],batch[5]})
				t.AppendSeparator()				
			}
			t.Render()
		}
	}

	if listStr == "plugin" {
		if len(bundlesResults) > 0 {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.SetStyle(table.StyleLight)
			t.Style().Color.Header = text.Colors{text.Bold}
			t.SetTitle("BUNDLES")
			t.AppendHeader(table.Row{"Name","Description","Type","Tag","Version","Status"})
			t.SetColumnConfigs([]table.ColumnConfig{
				{Number: 1, Colors: text.Colors{text.Color(text.FgYellow),},},
				{Number: 2, WidthMax: 90, WidthMaxEnforcer: text.WrapText},
				{Number: 6, Colors: text.Colors{text.Color(text.FgHiGreen)}},
			})
			limit := 6
			for i:=0; i < len(bundlesResults)-1;i+= limit{
				batch := bundlesResults[i:min(i+limit, len(bundlesResults))]

				t.AppendRow(table.Row{strings.ToUpper(batch[0]),strings.Title(batch[1]),batch[2],batch[3],batch[4],batch[5]})
				t.AppendSeparator()
			}
			t.Render()
		}
	}


	if len(commsResults)>0{
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleLight)
		t.Style().Color.Header = text.Colors{text.Bold}
		t.SetTitle("COMMUNITY")
		t.AppendHeader(table.Row{"Name","Description","Type","Tag","Version","Status"})
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, Colors: text.Colors{text.Color(text.FgYellow),},},
			{Number: 2, WidthMax: 90, WidthMaxEnforcer: text.WrapText},
			{Number: 6, Colors: text.Colors{text.Color(text.FgHiGreen)}},
		})
		limit := 5
		for i:=0; i < len(commsResults)-1;i+= limit{
			batch := commsResults[i:min(i+limit, len(commsResults))]
			t.AppendRow(table.Row{strings.ToUpper(batch[0]),strings.Title(batch[1]),batch[2],batch[3],batch[4],batch[5]})
			t.AppendSeparator()
		}
		t.Render()
	}

}

func ListByName(Results []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Style().Color.Header = text.Colors{text.Bold}
	t.SetTitle("COMMUNITY")
	t.AppendHeader(table.Row{"Name","Description","Type","Tag","Version","Status"})
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Colors: text.Colors{text.Color(text.FgYellow),},},
		{Number: 2, WidthMax: 90, WidthMaxEnforcer: text.WrapText},
		{Number: 6, Colors: text.Colors{text.Color(text.FgHiYellow)}},
	})
	limit := 5
	for i:=0; i < len(Results)-1;i+= limit{
		batch := Results[i:min(i+limit, len(Results))]
		t.AppendRow(table.Row{strings.ToUpper(batch[0]),batch[1],batch[2],batch[3],batch[4],batch[5]})
		t.AppendSeparator()
	}
	t.Render()
}

func ListClusters(Results []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Style().Color.Header = text.Colors{text.Bold}
	t.SetTitle("CLUSTERS DEPLOYED")
	t.AppendHeader(table.Row{"Name","Type","Status"})
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Colors: text.Colors{text.Color(text.FgYellow),},},
		{Number: 2, WidthMax: 90, WidthMaxEnforcer: text.WrapText},
		{Number: 6, Colors: text.Colors{text.Color(text.FgHiYellow)}},
	})
	limit := 3
	for i:=0; i < len(Results)-1;i+= limit{
		batch := Results[i:min(i+limit, len(Results))]
		t.AppendRow(table.Row{strings.ToUpper(batch[0]),batch[1],batch[2]})
		t.AppendSeparator()
	}
	t.Render()
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}