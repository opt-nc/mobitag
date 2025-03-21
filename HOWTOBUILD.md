# ✅ Prérequis

- ✅ Tooling `Go` ([installer `Go`](https://go.dev/doc/install))
- ✅ Une clé d'API, chargée dans l'environnement `OPTNC_MOBITAGNC_API_KEY`
- ✅ [Task](https://taskfile.dev/#/installation) pour builder le projet
- ✅ [Goreleaser](https://goreleaser.com/install/) pour builder le projet sur plusieurs plateformes

# 🚀 Getting started

## ⚙️ Builder avec task

```shell
task build
# Ensuite l'ajouter dans le PATH
echo "export PATH=$PATH:$(pwd)/bin" >> ~/.zshrc # ou ~/.bashrc selon votre shell
source ~/.zshrc # ou ~/.bashrc selon votre shell
```

## Compilation avec goreleaser

```shell
task release
```

## Sans compilation

Si vous ne souhaitez pas utiliser le binaire, vous pouvez utiliser directement le code source :

```shell
go run main.go [command]
```

# 🕹️ Essayer

```sh
mobitag -h
```

```sh
# Tester l'environnement
mobitag dryRun
```

# 🥳 Envoyer un `mobit@g`

```sh
mobitag send --to xxxxxx --message "Hello World : a mobit@g from Go(lang) XD"
```

# 🎯 Autocompletion

Pour avoir plus d'informations sur l'autocompletion :

```sh
mobitag completion <shell> --help
```

Exemple pour un shell zsh :

```sh
source <(mobitag completion zsh)  # pour activer l'autocompletion dans le shell courant
mobitag completion zsh > "${fpath[1]}/_mobitag" # pour installer l'autocompletion de manière permanente
```

# 📼 Buidler la demo video

La video de demo est buildée avec [`charmbracelet/vhs`](https://github.com/charmbracelet/vhs):

```sh
task auth
```
