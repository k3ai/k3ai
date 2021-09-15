package internals

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
)
func List(result string)  {
    if result == "infra" {
        listTable("name","Description", "Version", "Type")
    }
    if result == "apps"{
        listTable("name","Description","Version","Group")
    }
    if result == "bundles"{
        listTable("name","Description","Version","")
    }
}

func listTable( column1 string, column2 string, column3 string, column4 string)  {
    
    t := table.NewWriter()
    t.SetTitle("List Plugins")
    t.SetOutputMirror(os.Stdout)
    t.SetStyle(table.StyleLight)
    t.Style().Options.SeparateRows = true
    t.AppendHeader(table.Row{"Infrastructure Plugins"})
    // t.AppendHeader(table.Row{column1, column2, column3, column4})
    t.SetColumnConfigs([]table.ColumnConfig{
        {Number: 2,WidthMaxEnforcer: text.Trim},
    })
    str := "Amazon EKS Anywhere is an open-source deployment option for Amazon EKS that allows customers to create and operate Kubernetes clusters on-premises,with optional support offered by AWS"
    wrapped := text.WrapSoft(str,100)
    t.AppendRows([]table.Row{
        {"k3s", "The certified Kubernetes distribution built for IoT & Edge computing", "v1.21", "Local"},
        {"EKS Anywhere", wrapped, "v1.21", "Hybrid"},
        {"Civo", "The first pure play cloud native service powered only by Kubernetes", "v1.21", "Cloud"},
        {"Kind", "kind is a tool for running local Kubernetes clusters using Docker container “nodes”.\n kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.", "v1.21", "Local"},

    })
    t.AppendRow(table.Row{"APPLICATION PLUGINS"})
    t.AppendRows([]table.Row{
        {"KFP", "[Argo Backend ] Kubeflow Pipelines is a platform for building and deploying portable,\n scalable machine learning (ML) workflows based on Docker containers", "v1.3", "Kubeflow"},
        {"KFPT", "[Tekton Backend] Kubeflow Pipelines is a platform for building and deploying portable,\n scalable machine learning (ML) workflows based on Docker containers", "v1.3", "Kubeflow"},
        {"Katib", "Katib is a Kubernetes-native project for automated machine learning (AutoML).\n Katib supports hyperparameter tuning, early stopping and neural architecture search (NAS)", "v1.3", "Kubeflow"},
        {"Notebooks", "Jupyter Notebooks integrated in Kubeflow to interact from any notebook with other KF components", "v1.3", "Kubeflow"},
    })
    t.AppendRow(table.Row{"BUNDLES PLUGINS"})
    t.SetColumnConfigs([]table.ColumnConfig{
        {Number: 2, WidthMaxEnforcer: text.Trim},
    })
    t.AppendRows([]table.Row{
        {"Kubeflow", wrapped, "v1.3", "Platform"},
        {"MLFlow", "An open source platform for the machine learning lifecycle", "v1.20.2", "Platform"},
    })

    t.Render()
}