package execution

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	// "strings"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/homedir"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	
	internal "github.com/k3ai/internal"
	color "github.com/k3ai/pkg/color"
	db "github.com/k3ai/pkg/db"
	http "github.com/k3ai/pkg/http"
	gh "github.com/k3ai/pkg/io/git"
)

const (
	k3aiKube =".tools/kubectl"
	k3aiHelm = ".tools/helm"
	lnxApp = "/bin/bash"	
)
var (
	appPlugin = internal.AppPlugin{}
	kubeconfig *string
	// cmd *exec.Cmd
	rootPlugin = &internal.K3aiRootPlugin{}
	subRootPlugin = &internal.K3aiInternalPlugin{}
	subPlugin = &internal.AppPlugin{}
	subPluginResources = &internal.AppPluginResources{}
	restatus = false
	strQuiet = true
) 


func Deployment (actionType string,name string, ctype string) (status bool, err error) {
		
		if actionType == "cluster" {
			url := db.List(ctype)
			data,_ := http.Download(url)
			_ = yaml.Unmarshal([]byte(data), &rootPlugin)
			if len(rootPlugin.Resources) > 1 {
				for i:=0; i < len(rootPlugin.Resources);i++{
					_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[i]),url,"install",name)
				}
			} else {
				_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[0]),url,"install",name)
			}
		}else {
			url := db.List(name)
			data,_ := http.Download(url)
			_ = yaml.Unmarshal([]byte(data), &rootPlugin)
			if len(rootPlugin.Resources) > 1 {
				for i:=0; i < len(rootPlugin.Resources);i++{
					_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[i]),url,"install",ctype)
				}
			} else {
				_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[0]),url,"install",ctype)
			}
		}
	return true,nil
}

func Removal (actionType string,name string, ctype string) (status bool, err error) {

	if actionType == "cluster" {
		clusterResults := db.ListClustersByName()
		url := db.List(clusterResults[1])
		data,_ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &rootPlugin)
		if len(rootPlugin.Resources) > 1 {
			for i:=0; i < len(rootPlugin.Resources);i++{
				_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[i]),url,"remove",name)
			}
		} else {
			_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[0]),url,"remove",name)
		}
	}else {
		url := db.List(name)
		data,_ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &rootPlugin)
		if len(rootPlugin.Resources) > 1 {
			for i:=0; i < len(rootPlugin.Resources);i++{
				_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[i]),url,"remove",name)
			}
		} else {
			_ = innerPluginResource(rootPlugin.Metadata.Name,string(rootPlugin.Resources[0]),url,"remove",name)
		}
	}
return true,nil
}


func innerPluginResource (name string,base string,url string, action string,clusterName string) error {
	if strings.Contains(base,"../../") {
		name = strings.ToLower(name)
		base = strings.TrimLeft(base,"../..") + "/k3ai.yaml"
		// path := "/home/alefesta/.k3ai/"+name+"/"+strings.TrimRight(base,"/k3ai.yaml")
		url = strings.Replace(url,"/apps/" + name +"/k3ai.yaml","/"+base,-1)
		data,_ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &subPlugin)
		subUrl := strings.Replace(url,"/k3ai.yaml","/base/plugin.yaml",-1)
		subData,_ := http.Download(subUrl)
		_ = yaml.Unmarshal([]byte(subData), &subPlugin)
		if len(subPlugin.Resources) > 1 {
			for k:=0; k < len(subPlugin.Resources); k++{
				if subPlugin.Resources[k].PluginType == "kustomize" {
					gh.Clone(subPlugin.Resources[k].Path, name)
					kustomize(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args,name,action, strings.ToLower(clusterName))
				} else if subPlugin.Resources[k].PluginType == "kubectl" {
					kubectl(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args,name,action, strings.ToLower(clusterName))

				}else  if subPlugin.Resources[k].PluginType == "shell"{
					shell(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args,false,action)
					}
			}	
		}		
	} else {
		url = strings.Replace(url,"/k3ai.yaml","/base/plugin.yaml",-1)
		subData,_ := http.Download(url)
		_ = yaml.Unmarshal([]byte(subData), &subPlugin)
		if len(subPlugin.Resources) > 1 {
			for k:=0; k < len(subPlugin.Resources); k++{
				if subPlugin.Resources[k].PluginType == "kustomize" {
					gh.Clone(subPlugin.Resources[k].Path,name)
					kustomize(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args,name,action, strings.ToLower(clusterName))
				} else if subPlugin.Resources[k].PluginType == "kubectl" {
					kubectl(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args,name,action, strings.ToLower(clusterName))

				}else if subPlugin.Resources[k].PluginType == "shell"{
					path := strings.Replace(subPlugin.Resources[k].Path,"{{name}}",strings.ToLower(clusterName),-1)
					if action == "remove"{
						shell(path, subPlugin.Resources[k].Remove,false,action)
					} else {
						shell(path, subPlugin.Resources[k].Args,false,action)
					}
					
				}
			}
		} else {
			if subPlugin.Resources[0].PluginType == "kustomize" {
				gh.Clone(subPlugin.Resources[0].Path,name)
				kustomize(subPlugin.Resources[0].Path, subPlugin.Resources[0].Args,name,action,clusterName)
			} else if subPlugin.Resources[0].PluginType == "shell"{
				path := strings.Replace(subPlugin.Resources[0].Path,"{{name}}",strings.ToLower(clusterName),-1)
				if action == "remove"{
					removePath := strings.Replace(subPlugin.Resources[0].Remove,"{{name}}",strings.ToLower(clusterName),-1)
					shell(removePath, "",false,action)
					db.DeleteCluster(strings.ToLower(clusterName))
				} else {
					shell(path, subPlugin.Resources[0].Args,false,action)
				}
				
			}
		}
	}
	return nil
}


func shell(pluginEx string, pluginArgs string, outPrint bool, action string) error {
		home,_ := os.UserHomeDir()
		shellPath := home + "/.k3ai"
		
		if pluginEx == "post" {
			pluginEx = ""
			cmd := exec.Command("/bin/bash","-c",pluginArgs)
			cmd.Dir = shellPath
			cmd.Output()


		}
		if action == "install" {
		color.Done()
		fmt.Println(" 🚀 Starting installation...")
		fmt.Println(" ")
		} else if action == "remove" {
			color.Done()
			fmt.Println(" 🚀 Removing installation...")
			fmt.Println(" ")
		}
		cmd := exec.Command("/bin/bash","-c",pluginEx,pluginArgs)
		cmd.Dir = shellPath
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})

		scanner := bufio.NewScanner(r)
		go func() {
			// Read line by line and process it
			msg := "⏳	Working..."
			fmt.Printf("\r %v", msg)
			fmt.Println(" ")
			for scanner.Scan() {
				line := scanner.Text()
				color.Disable()
				if strQuiet {
					fmt.Println(" 🚀 " + line)
				}
				
			}
			done <- struct{}{}
		}()
		// Start the command and check for errors
		err := cmd.Start()
		if err != nil {
			log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")	
		}
		<-done
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
		return err
}

func kustomize(pluginEx string, pluginArgs string, pluginName string, action string, target string) error {
	_,clusterType := db.CheckClusterName(target)
	out := Client(target,clusterType)
	home,_ := os.UserHomeDir()
	shellPath := home + "/.k3ai"
	
	// kustomizeBinary := home + k3aiKube
	if pluginEx == "post" {
		pluginEx = ""
		cmd := exec.Command("/bin/bash","-c",pluginArgs)
		cmd.Dir = shellPath
		cmd.Output()


	}
	if action == "install" {
	if !restatus {
		restatus = true
		color.Done()
		fmt.Println(" 🚀 Starting installation...")
		fmt.Println(" ")
		color.Disable()
	}

	gh.Clone(pluginEx,pluginName)
	// path := w.Filesystem.Root()
	
    cmd:= exec.Command(k3aiKube,"apply","-k",shellPath +"/git/"+ pluginName + "/" + pluginArgs,"--kubeconfig="+ out, "--wait")
	cmd.Dir=shellPath
	r, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	done := make(chan struct{})

	scanner := bufio.NewScanner(r)
	go func() {
		// Read line by line and process it
		msg := "⏳	Working..."
		fmt.Printf("\r %v", msg)
		fmt.Println(" ")
		for scanner.Scan() {
			line := scanner.Text()
			color.Disable()
			if strQuiet {
				fmt.Println(" 🚀 " + line)
			}
			
		}
		done <- struct{}{}
	}()
	// Start the command and check for errors
	err := cmd.Start()
	if err != nil {
		log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")	
	}
	<-done
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	}
	return nil
}

func kubectl(pluginEx string, pluginArgs string, pluginName string, action string, target string) error {
	_,clusterType := db.CheckClusterName(target)
	out := Client(target,clusterType)
	home,_ := os.UserHomeDir()
	shellPath := home + "/.k3ai"

	// if action == "install" {
	// 	if !restatus {
	// 		restatus = true
	// 		color.Done()
	// 		fmt.Println(" 🚀 Starting installation...")
	// 		fmt.Println(" ")
	// 		color.Disable()
	// 	}
	
		// gh.Clone(pluginEx,pluginName)
		// path := w.Filesystem.Root()
		
		cmd:= exec.Command(k3aiKube,"apply",pluginArgs,"--kubeconfig="+ out, "--wait")
		cmd.Dir=shellPath
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})
	
		scanner := bufio.NewScanner(r)
		go func() {
			// Read line by line and process it
			msg := "⏳	Working..."
			fmt.Printf("\r %v", msg)
			fmt.Println(" ")
			for scanner.Scan() {
				line := scanner.Text()
				color.Disable()
				if strQuiet {
					fmt.Println(" 🚀 " + line)
				}
				
			}
			done <- struct{}{}
		}()
		// Start the command and check for errors
		err := cmd.Start()
		if err != nil {
			log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")	
		}
		<-done
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
		// }

	return err
}

func WaitForDeployment(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		
}

func Client (name string,ctype string) (kubeconfig string) { 
	var cPath string
	if ctype == "k3s" {
		cPath ="/etc/rancher/k3s/k3s.yaml"
	} else if ctype == "eks-a" {
		cPath = homedir.HomeDir() + "/.k3ai/"+ name +"/"+name+"-eks-a-cluster.kubeconfig"
	}else {
		cPath = homedir.HomeDir() + "/.kube/config"
	}
	if home := homedir.HomeDir(); home != "" {
		if ctype != "eks-a"{
			out,_ := os.Create(homedir.HomeDir() + "/.k3ai/" + name +".config")
			in,_ := os.Open(cPath)
		
			
			_, err := io.Copy(out,in)
			if err != nil {
				log.Print(err)
			}
			out.Close()
			kubeconfig = homedir.HomeDir() + "/.k3ai/" + name +".config"
		} else {
			kubeconfig = homedir.HomeDir() + "/.k3ai/"+ name +"/"+name+"-eks-a-cluster.kubeconfig"
		}

		
	}
	
	return kubeconfig
}
