package internals

import (
    "github.com/k3ai/log"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
    callDB "github.com/k3ai/shared"
)
func List(result string)  {
    if result == "infra" {
        listTable(result,"name","Description", "Version", "Tag","Status")
    }
    if result == "apps"{
        listTable(result,"name","Description","Version","Tag","Status")
    }
    if result == "bundles"{
        listTable(result,"name","Description","Version","Tag","Status")
    }
}
var results []string
func listTable( title string,column1 string, column2 string, column3 string, column4 string, column5 string)  {
    switch title {
    case "apps":
        results = callDB.AppsDisplaySQL()
    case "infra":
        results = callDB.InfraDisplaySQL()
    case "bundles":
        results = callDB.BundleDisplaySQL()
    }

    if len(results) < 1 {
        log.Error("Sorry the table is empty")
    } else {
        t := table.NewWriter()
        t.SetTitle("List Plugins")
        t.SetOutputMirror(os.Stdout)
        t.SetStyle(table.StyleLight)
        t.Style().Options.SeparateRows = true
        switch title {
        case "apps":
            t.AppendHeader(table.Row{"Application Plugins"})
        case "infra":
            t.AppendHeader(table.Row{"Infrastructure Plugins"})
        case "bundles":
            t.AppendHeader(table.Row{"Bundles Plugins"})
        }
        t.AppendSeparator()
        t.AppendHeader(table.Row{column1, column2, column3, column4, column5})
        t.SetColumnConfigs([]table.ColumnConfig{
            {Number: 2,WidthMaxEnforcer: text.Trim},
        })
        var i int
        limit := 5
        for i = 0; i < len(results); i+= limit{
            batch := results[i:min(i+limit, len(results))]
            if batch[4] == "" {
                batch[4] = "available"
            }
            t.AppendRow(table.Row{batch[0],text.WrapSoft(string(batch[1]),100),batch[2],batch[3], batch[4]})
        }
         
    
        t.Render()
    }

}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}