<h1 align="center">
  <img src="https://raw.githubusercontent.com/k3ai/k3ai.github.io/main/static/img/logo-banner.jpg"/><br/>
  Welcome to K3ai Project
</h1>
<p align="center"><b>K3ai</b> is a lightweight tool to get an AI Infrastructure Stack up in minutes not days.</p>

<p align="center">
<img src="https://img.shields.io/badge/version-v1.0-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;
<img src="https://img.shields.io/badge/Go-1.14+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;
<a href="https://goreportcard.com/report/github.com/k3ai/k3ai" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;
<img src="https://img.shields.io/github/license/k3ai/k3ai?style=for-the-badge" alt="license" /></p>

---
**NOTE on the K3ai origins**

Original K3ai Project has been developed at the end of October 2020 in 2 weeks by:

- **Alessandro Festa** [https://github.com/alefesta](https://github.com/alefesta)
- **Gabriele Santomaggio** [https://github.com/GSantomaggio](https://github.com/GSantomaggio)

K3ai v1.0 has been entirely re-written by **Alessandro Festa** during the month of October 2021 to
offer a better User Experience.
---

Thanks to the amazing and incredible people and projects that have been instrumental to create K3ai project repositories,website,etc...


- [Docusaurs](https://docusaurus.io/) - How simple and amazing is to use for your own website (https://k3ai.in)
- [https://undraw.co/](https://undraw.co/) - The amazing work created by Katerina Limpitsouni (https://twitter.com/ninaLimpi) is a real piece of art
- [https://getemoji.com/](https://getemoji.com/) - We all need some emoji in our life isn't it?
- [https://clig.dev/](https://clig.dev/) -- K3ai is completly inspired by the Command Line Guidelines manifesto.



## ⚡️ Quick start

Let's discover **K3ai in three simple steps**.

## 🌘 Getting Started

Get started by **download k3ai** from the release page [here](https://github.com/k3ai/releases).

Or **try K3ai companion script** using this command:

```bash
curl -LO https://get.k3ai.in | sh -
```

## 🌗 Load K3ai configuration

Let's start loading the configuration:

```shell
k3ai up
```

---

First time k3ai run will ask for a **Github PAT (Personal Access Token)** that we will use to avoid API calls limitations.  Check [`Github Documentation`](#) to learn how to create one. Your personal GH PAT only need `read repository permission`.

---

## 🌖 Configure the base infrastructure

Choose your favourite `Kubernetes` flavor and run it:

To know which K8s flavors are available

```shell
k3ai cluster list --all
```

it should print something like:

```markdown
┌─────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ INFRASTRUCTURE                                                                                          │
├───────┬─────────────────────────────────────────────────────┬───────┬────────┬─────────┬────────────────┤
│ TYPE  │ DESCRIPTION                                         │ KIND  │ TAG    │ VERSION │ STATUS         │
├───────┼─────────────────────────────────────────────────────┼───────┼────────┼─────────┼────────────────┤
│ CIVO  │ The First Cloud Native Service Provider Power...    │ infra │ cloud  │ latest  │ Available      │
├───────┼─────────────────────────────────────────────────────┼───────┼────────┼─────────┼────────────────┤
│ EKS-A │ Amazon Eks Anywhere Is A New Deployment Option...   │ infra │ hybrid │ v0.5.0  │ Available      │
│       │ ate And Operate Kubernetes Clusters On Custome...   │       │        │         │                │
├───────┼─────────────────────────────────────────────────────┼───────┼────────┼─────────┼────────────────┤
│ K3S   │ K3s Is A Highly Available, Certified Kubernetes...  │ infra │ local  │ latest  │ Available      │
│       │ oads In Unattended, Resource-Constrained...         │       │        │         │                │
├───────┼─────────────────────────────────────────────────────┼───────┼────────┼─────────┼────────────────┤
│ KIND  │ Kind Is A Tool For Running Local Kubernetes...      │ infra │ local  │ v0.11.2 │ Available      │
│       │ as Primarily Designed For Testing Kubernetes...     │       │        │         │                │
│       │  Or Ci.                                             │       │        │         │                │
├───────┼─────────────────────────────────────────────────────┼───────┼────────┼─────────┼────────────────┤
│ TANZU │ Tanzu Community Edition Is A Fully-Featured...      │ infra │ hybrid │ latest  │ In Development │
│       │ ers And Users. It Is A Freely Available...          │       │        │         │                │
│       │  Of Vmware Tanzu.                                   │       │        │         │                │
└───────┴─────────────────────────────────────────────────────┴───────┴────────┴─────────┴────────────────┘
```

Now let start with something super fast and super simple:

<!-- ```bash
k3ai [COMMAND] [ACTION] [CHOICE] [SUB-CHOICE]
```
where:

- **K3AI [COMMAND]** : I want to do something with a `cluster` or a `plugin`
- **[ACTION]** : I want to deploy a `cluster` or a `plugin`
- **[CHOICE]** I want a specific type of a `cluster`. This could be shortened into `-t`
- **[SUB- CHOICE]** I want to identify later the `cluster` with this name. This could be shortened into `-n`
So in our case will be: -->

```bash
k3ai cluster deploy --type k3s --n mycluster
```

## 🌝 Install a plugin to do your AI experimentations

Now that the server is up and running let's type:

```bash
k3ai plugin deploy -n mlflow -t mycluster
```

K3ai will print the url where you may access to the MLFLow tracking server at the end of the installation.
That's all now just start having fun with K3ai!

## 🌈 Push a piece of code to the AI tools and focus on your goals

Let's push some code to the AI tool (i.e.: MLFlow)

```bash

k3ai run --source https://github.com/k3ai/quickstart --target mycluster --backend mlflow

```

wait the run to complete and login the backend AI tolls (i.e.: on the MLFlow UI `http://<YOUR IP>:30500`)

## Current Implementation support

### Operating Systems

| Operating System | K3ai v1.0.0 |
|:---|:---:|
|    Linux     |    Yes          |
|    Windows   |    In Progress  |
|    MacOs     |    In Progress  |
|    Arm       |    In Progress  |

### Clusters

| K8s Clusters | K3ai v1.0.0 |
|:---|:---:|
| Rancher K3s|Yes|
|Vmware Tanzu Community Ed.|Yes|
|Amazon EKS Anywhere|Yes|
|KinD|Yes|

### Plugins

| Plugins | K3ai v1.0.0 |
|:---|:---:|
|Kuebflow Components| Yes|
| MLFlow| Yes|
|Apache Airflow |Yes|
|Argo Workflows| Yes|


## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `K3ai Project`:

- Add a [GitHub Star](https://github.com/k3ai/k3ai) to the project.
- Tweet about project [on your Twitter](https://twitter.com/intent/tweet?text=%E2%9C%A8%20An%20AI%20stack%20including%20%23kubernetes%20and%20popular%20tools%20like%20%23kubeflow%20%23mlflow%20%23airflow.%20%20Deploy%20your%20AI%20projects%20in%20seconds%20in%20one%20command.%20Focus%20on%20writing%20code%20and%20thinking%20of%20business%20logic.K3ai%20will%20take%20care%20of%20the%20rest.%0A%0Ahttps%3A%2F%2Fgithub.com%2Fk3ai%2Fk3ai).
- Write interesting articles about K3ai project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.

Together, we can make this project **better** every day! 😘

## ⚠️ License


`K3ai` is free and open-source software licensed under the [BSD 3-Clause](https://github.com/k3ai/k3ai/blob/master/LICENSE). Official [logo](https://raw.githubusercontent.com/k3ai/k3ai.github.io/main/static/img/logo.jpg) was created by [Alessandro Festa](https://github.com/alefesta/).

