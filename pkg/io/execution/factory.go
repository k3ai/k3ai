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
	k3aiKube = "/.tools/kubectl"
	k3aiHelm = ".tools/helm" //nolint
	lnxApp   = "/bin/bash"   //nolint
	civoCli  = ".tools/civo" //nolint
	clusterTest = true
)

var (
	appPlugin  = internal.AppPlugin{} //nolint
	kubeconfig *string                //nolint
	// cmd *exec.Cmd
	rootPlugin         = &internal.K3aiRootPlugin{}
	subRootPlugin      = &internal.K3aiInternalPlugin{} //nolint
	subPlugin          = &internal.AppPlugin{}
	subPluginResources = &internal.AppPluginResources{} //nolint
	restatus           = false
	strQuiet           = true
)

func Deployment(actionType string, name string, ctype string, extras string) (status bool, err error) {

	if actionType == "cluster" {
		url := db.List(ctype)
		data, _ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &rootPlugin)
		if strings.ToLower(rootPlugin.Metadata.PluginStatus) != "available" && !clusterTest {
			color.Alert()
			fmt.Println("🥺 We are sorry, currently the plugin is unavailable")
			os.Exit(1)
		}
		if strings.ToLower(rootPlugin.Metadata.Name) == "tanzu" {
			_, err := exec.LookPath("kubectl")
			if err != nil {
				home, _ := os.UserHomeDir()
				shellPath := home + "/.k3ai/.tools/kubectl"
				_,err = exec.Command("/bin/bash", "-c", "sudo cp " + shellPath + " /usr/local/bin/").Output()
				if err != nil {
					log.Println(err)
				}
			}
		}

		if len(rootPlugin.Resources) > 1 {
			for i := 0; i < len(rootPlugin.Resources); i++ {
				_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[i]), url, "install", name, extras)
			}
		} else {
			_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[0]), url, "install", name, extras)
		}
	} else {
		url := db.List(name)
		data, _ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &rootPlugin)
		if strings.ToLower(rootPlugin.Metadata.PluginStatus) != "available" {
			color.Alert()
			fmt.Println("🥺 We are sorry, currently the plugin is unavailable")
			os.Exit(1)
		}
		if len(rootPlugin.Resources) > 1 {
			for i := 0; i < len(rootPlugin.Resources); i++ {
				_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[i]), url, "install", ctype, extras)
			}
		} else {
			_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[0]), url, "install", ctype, extras)
		}
	}
	return true, nil
}

func Removal(actionType string, name string, ctype string) (status bool, err error) {
	var extras string
	if actionType == "cluster" {
		clusterResults := db.ListClusterByName(name)
		url := db.List(clusterResults[1])
		data, _ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &rootPlugin)
		if len(rootPlugin.Resources) > 1 {
			for i := 0; i < len(rootPlugin.Resources); i++ {
				_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[i]), url, "remove", name, extras)
			}
		} else {
			_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[0]), url, "remove", name, extras)
		}
	} else {
		url := db.List(name)
		data, _ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &rootPlugin)
		if len(rootPlugin.Resources) > 1 {
			for i := 0; i < len(rootPlugin.Resources); i++ {
				_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[i]), url, "remove", name, extras)
			}
		} else {
			_ = innerPluginResource(rootPlugin.Metadata.Name, string(rootPlugin.Resources[0]), url, "remove", name, extras)
		}
	}
	return true, nil
}

func innerPluginResource(name string, base string, url string, action string, clusterName string, extras string) error {
	if strings.Contains(base, "../../") {
		name = strings.ToLower(name)
		base = strings.TrimLeft(base, "../..") + "/k3ai.yaml" //nolint:staticcheck
		url = strings.Replace(url, "/apps/"+name+"/k3ai.yaml", "/"+base, -1)
		data, _ := http.Download(url)
		_ = yaml.Unmarshal([]byte(data), &subPlugin)
		subUrl := strings.Replace(url, "/k3ai.yaml", "/base/plugin.yaml", -1)
		subData, _ := http.Download(subUrl)
		_ = yaml.Unmarshal([]byte(subData), &subPlugin)
		if len(subPlugin.Resources) > 1 {
			for k := 0; k < len(subPlugin.Resources); k++ {
				if subPlugin.Resources[k].PluginType == "kustomize" {
					err := gh.Clone(subPlugin.Resources[k].Path, name)
					if err != nil {
						log.Fatal(err)
					}
					err = kustomize(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args, name, action, strings.ToLower(clusterName))
					if err != nil {
						log.Fatal(err)
					}
				} else if subPlugin.Resources[k].PluginType == "kubectl" {
					err := kubectl(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args, name, action, strings.ToLower(clusterName))
					if err != nil {
						log.Fatal(err)
					}

				} else if subPlugin.Resources[k].PluginType == "shell" {
					err := shell(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args, false, action, extras)
					if err != nil {
						log.Fatal(err)
					}
				} else if subPlugin.Resources[k].PluginType == "helm" {
					err := helm(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args, name, action, strings.ToLower(clusterName))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	} else {
		url = strings.Replace(url, "/k3ai.yaml", "/base/plugin.yaml", -1)
		subData, _ := http.Download(url)
		_ = yaml.Unmarshal([]byte(subData), &subPlugin)
		if len(subPlugin.Resources) > 1 {
			for k := 0; k < len(subPlugin.Resources); k++ {
				if subPlugin.Resources[k].PluginType == "kustomize" {
					err := gh.Clone(subPlugin.Resources[k].Path, name)
					if err != nil {
						log.Fatal(err)
					}
					err = kustomize(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args, name, action, strings.ToLower(clusterName))
					if err != nil {
						log.Fatal(err)
					}
				} else if subPlugin.Resources[k].PluginType == "kubectl" {
					err := kubectl(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args, name, action, strings.ToLower(clusterName))
					if err != nil {
						log.Fatal(err)
					}

				} else if subPlugin.Resources[k].PluginType == "shell" {
					path := strings.Replace(subPlugin.Resources[k].Path, "{{name}}", strings.ToLower(clusterName), -1)
					if action == "remove" {
						if subPlugin.Resources[k].Remove != "" {
							pathRemove := strings.Replace(subPlugin.Resources[k].Remove, "{{name}}", strings.ToLower(clusterName), -1)
							err := shell(path, pathRemove, false, action, extras)
							if err != nil {
								log.Fatal(err)
							}
						}
					} else {
						err := shell(path, subPlugin.Resources[k].Args, false, action, extras)
						if err != nil {
							log.Fatal(err)
						}
					}

				} else if subPlugin.Resources[k].PluginType == "helm" {
					err := helm(subPlugin.Resources[k].Path, subPlugin.Resources[k].Args, name, action, strings.ToLower(clusterName))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		} else {
			if subPlugin.Resources[0].PluginType == "kustomize" {
				err := gh.Clone(subPlugin.Resources[0].Path, name)
				if err != nil {
					log.Fatal(err)
				}
				err = kustomize(subPlugin.Resources[0].Path, subPlugin.Resources[0].Args, name, action, clusterName)
				if err != nil {
					log.Fatal(err)
				}
			} else if subPlugin.Resources[0].PluginType == "shell" {
				path := strings.Replace(subPlugin.Resources[0].Path, "{{name}}", strings.ToLower(clusterName), -1)
				if action == "remove" {
					removePath := strings.Replace(subPlugin.Resources[0].Remove, "{{name}}", strings.ToLower(clusterName), -1)
					err := shell(removePath, "", false, action, extras)
					if err != nil {
						log.Fatal(err)
					}
					err = db.DeleteCluster(strings.ToLower(clusterName))
					if err != nil {
						log.Fatal(err)
					}
				} else {
					err := shell(path, subPlugin.Resources[0].Args, false, action, extras)
					if err != nil {
						log.Fatal(err)
					}
				}

			} else if subPlugin.Resources[0].PluginType == "helm" {
				err := helm(subPlugin.Resources[0].Path, subPlugin.Resources[0].Args, name, action, strings.ToLower(clusterName))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return nil
}

func shell(pluginEx string, pluginArgs string, outPrint bool, action string, extras string) error {
	home, _ := os.UserHomeDir()
	shellPath := home + "/.k3ai"

	if pluginEx == "post" {
		pluginEx = "" //nolint:ineffassign
		cmd := exec.Command("/bin/bash", "-c", pluginArgs)
		cmd.Dir = shellPath
		_, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}

	}
	if action == "install" {
		color.Done()
		fmt.Println(" 🚀 Starting installation...")
		fmt.Println(" ")
		if extras != "" {
			pluginArgs = pluginArgs + " " + extras
		}

		pluginEx = strings.Replace(pluginEx, "{{extras}}", pluginArgs, -1)
		cmd := exec.Command("/bin/bash", "-c", pluginEx)
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
	} else if action == "remove" {
		color.Done()
		fmt.Println(" 🚀 Removing installation...")
		fmt.Println(" ")
		cmd := exec.Command("/bin/bash", "-c", pluginArgs)
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
	return nil
}

func kustomize(pluginEx string, pluginArgs string, pluginName string, action string, target string) error {
	_, clusterType := db.CheckClusterName(target)
	out := Client(target, clusterType)
	home, _ := os.UserHomeDir()
	shellPath := home + "/.k3ai"

	// kustomizeBinary := home + k3aiKube
	if pluginEx == "post" {
		pluginEx = "" //nolint:ineffassign,staticcheck
		cmd := exec.Command("/bin/bash", "-c", pluginArgs)
		cmd.Dir = shellPath
		_, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}

	}
	if action == "install" {
		if !restatus {
			restatus = true
			color.Done()
			fmt.Println(" 🚀 Starting installation...")
			fmt.Println(" ")
			color.Disable()
		}

		cmd := exec.Command("/bin/bash", "-c", shellPath+k3aiKube+" "+"apply -k "+shellPath+"/git/"+pluginArgs+" --kubeconfig="+out+" --wait")
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
		_ = cmd.Start()
		// if err != nil {
		// 	log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")
		// }
		<-done
		_, _ = exec.Command("/bin/bash", "-c", shellPath+k3aiKube+" wait --for=condition=Ready pods --all --all-namespaces  --kubeconfig="+out).Output()
		// fmt.Println(string(outcome))
		_ = cmd.Wait()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
	os.RemoveAll(shellPath + "/git/")
	return nil
}

func kubectl(pluginEx string, pluginArgs string, pluginName string, action string, target string) error {
	_, clusterType := db.CheckClusterName(target)
	out := Client(target, clusterType)
	home, _ := os.UserHomeDir()
	shellPath := home + "/.k3ai"

	if action == "install" {
		if !restatus {
			restatus = true
			color.Done()
			fmt.Println(" 🚀 Starting installation...")
			fmt.Println(" ")
			color.Disable()
		}
		color.InProgress()
		fmt.Println(" 🚀 Working on the installation...")
		outcome, _ := exec.Command("/bin/bash", "-c", shellPath+k3aiKube+" "+pluginEx+" --kubeconfig="+out).Output()
		fmt.Println(string(outcome))
		_, _ = exec.Command("/bin/bash", "-c", shellPath+k3aiKube+" wait --for=condition=Ready pods --all --all-namespaces  --kubeconfig="+out).Output()
	}
	return nil
}

func helm(pluginEx string, pluginArgs string, pluginName string, action string, target string) error {
	// var cmd *exec.Cmd
	_, clusterType := db.CheckClusterName(target)
	out := Client(target, clusterType)
	home, _ := os.UserHomeDir()
	shellPath := home + "/.k3ai"

	if action == "install" {
		if !restatus {
			restatus = true
			color.Done()
			fmt.Println(" 🚀 Starting installation...")
			fmt.Println(" ")
			color.Disable()
		}

		if strings.Contains(pluginEx, "repo add") {
			// pluginEx = strings.Replace(pluginEx,"repo add","",-1)
			color.InProgress()
			fmt.Println(" 🚀 Adding Helm repo...")
			outcome, err := exec.Command("/bin/bash", "-c", shellPath+"/.tools/helm "+pluginEx).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(outcome))
		} else {
			color.InProgress()
			fmt.Println(" 🚀 Working on the installation...")
			outcome, err := exec.Command("/bin/bash", "-c", shellPath+"/.tools/helm "+pluginEx+" --kubeconfig="+out).Output()
			if err != nil {
				log.Fatal(err)
			}
			color.Disable()
			fmt.Println(string(outcome))

		}
	}

	return nil
}

func WaitForDeployment(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

}

func Client(name string, ctype string) (kubeconfig string) {
	var cPath string
	if ctype == "k3s" {
		cPath = "/etc/rancher/k3s/k3s.yaml"
	} else if ctype == "eks-a" {
		cPath = homedir.HomeDir() + "/.k3ai/" + name + "/" + name + "-eks-a-cluster.kubeconfig"
	} else if ctype == "tanzu" {
		cPath = homedir.HomeDir() + "/.kube-tkg/config"

	} else {
		cPath = homedir.HomeDir() + "/.kube/config"
	}
	if home := homedir.HomeDir(); home != "" {
		if ctype != "eks-a" {
			out, _ := os.Create(homedir.HomeDir() + "/.k3ai/" + name + ".config")
			in, _ := os.Open(cPath)

			_, err := io.Copy(out, in)
			if err != nil {
				log.Print(err)
			}
			out.Close()
			kubeconfig = homedir.HomeDir() + "/.k3ai/" + name + ".config"
		} else {
			kubeconfig = homedir.HomeDir() + "/.k3ai/" + name + "/" + name + "-eks-a-cluster.kubeconfig"
		}

	}

	return kubeconfig
}
