# system-setup

> Effortless System Configuration

## Introduction

I like to clear my laptop frequently,
but configuring it can be hard and time-consuming.
As a **DevOps engineer**, I started this journey with **Ansible**.
However, it's not perfect.

Before starting the main playbook, you need to install dependencies.
I want a tool that has no dependencies and can be run with a single command.

Another approach is to use a simple **bash scripts**,
but I don't like bash a lot :smile.

After some reflection, I decided to use **Golang**.
It's simple to use: just download the binary, execute it, and wait.
That's why this project was created.

My approach is to use the Golang standard library as much as possible
to configure my system.
I used Ubuntu 24.04 when I created this project.

## What is included ?

1. Standard APT packages installation.
2. Add APT repositories and install Packages.
3. Install packages from github repositories (neovim, lazygit...)
4. Download Git repositories like dotfile repo and create all necessary symlinks. (In progress)
5. Gnome Configuration (In progress (In progress))

> Much details in the next sections

### Standard APT Packages

All installed packages are defined in the file `[cmd/systemSetup/templates/config.yaml.tmpl](./cmd/systemSetup/templates/config.yaml.tmpl)`

> Check the variable `aptStandardPackages`
> all lists can be overridden by creating a `config.yaml` file
> use the command `config -init`

This includes packages such as:

- ansible
- cargo
- fzf
- git
- stow
- tmux
- tree

Many of these packages are prerequisites for other tools installed later

### Additionnal APT repositories and Packages

Many tools I use are available in other APT repositories.
I must include them by default in my setup

The definition is in the config file

> Check the variable `aptRepositories` ans `additionnalPackages`

I include:

- brave
- google chrome
- docker
- gcloud
- golang
- helm
- hashicorp
- minikube
- vscode

### Tools from Github

I use some tools in my workflow that are not available in an APT repository,
but are available on GitHub.

By default, I install:

- [Neovim](https://github.com/neovim/neovim)
- [LazyGit](https://github.com/jesseduffield/lazygit)

I install also some fonts from **Nerd Fonts** such as:

- FiraCode
- Meslo

### Personal Github repository

To finish my system configuration, I need to get some stuff from my GitHub,
like my dotfiles and my personal CLI. All links are below:

[Dotfiles](https://github.com/q-sw/dotfiles)
[Personal CLI](https://github.com/q-sw/go-cli)

### Gnome Configuration

And the cherry on the cake is the Gnome configuration. For a long time, I used I3,
but I had some trouble with screen,
sound, and network management, especially at the office.

So, I decided to use Gnome extensions to replicate my I3 workflow.
I include:

- Just Perfection
- Space Bar
- Tactile
- TopHat

## How To use

### Build from the source

```shell
git clone https://github.com/q-sw/system-setup.git
cd system-setup
make build
mv ./bin/config /usr/local/bin/
```

### With custom configuration

```shell
config -init
config -path config.yaml
```

## Dev environment

I develop this project in **Golang 1.24**, I try the well exuction of the
binary with **Docker**
The command below build a Docker image with the Go binary include and run
the command `config`

```shell
make docker-run
```
