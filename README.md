# ❔ A propos

Cette repo est une **première expérimentation dont le but est de découvrir le
language [`Go`](https://go.dev/)**, sur un cas concret car... c'est plus amusant
et beaucoup plus motivant 🤓.

Cette expérimentation a donc pur but de créer un cli permettant d'envoyer des mobitags
depuis le terminal.

![](media/mobitag-cli.gif)


# 🔖 Ressources

- Site web officiel http://www.mobitag.nc
- [🥳 Mobitag.nc... 25 ans plus tard, des sms en SaaS via API{GEE}](https://dev.to/optnc/mobitagnc-25-ans-plus-tard-des-sms-en-saas-via-apigee-2h9e)
- [📲 Mobitag.nc for dummies](https://www.kaggle.com/code/optnouvellecaldonie/mobitag-nc-for-dummies)
- [⏱️ Mobitag Go Hackathon 2024-06-22 week-end 🤓](https://dev.to/adriens/mobitag-go-hackathon-2024-06-22-week-end-2n16)
- [⏱️ Mobitag Hackathon week-end du 2024-06-22 🤓](https://youtu.be/yVoMg7CXgaM)

# ✅ Prérequis

- [x] Tooling `Go` ([installer `Go`](https://go.dev/doc/install))
- [x] Une clé d'API, chargée dans l'environnement `OPTNC_MOBITAGNC_API_KEY`

# 🚀 Getting started

<!-- ## 🤓 `go install`

```shell
go install github.com/opt-nc/mobitag@latest
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bashrc

```

Puis : 

```sh
mobitag -h

``` -->

## ⚙️ Builder

```shell
task build
# Ensuite l'ajouter dans le PATH
echo "export PATH=$PATH:$(pwd)/bin" >> ~/.zshrc # ou ~/.bashrc selon votre shell
```

# 🕹️ Essayer

```sh
mobitag-cli -h
```

```sh
# Tester l'environnement
mobitag-cli dryRun
```

# 🥳 Envoyer un `mobit@g`

```sh
mobitag-cli send --to xxxxxx --message "Hello World : a mobit@g from Go(lang) XD"
```

# 🎯 Autocompletion

Pour avoir plus d'informations sur l'autocompletion :

```sh
mobitag-cli completion <shell> --help
```

Exemple pour un shell zsh :

```sh
source <(mobitag-cli completion zsh)  # pour activer l'autocompletion dans le shell courant
mobitag-cli completion zsh > "${fpath[1]}/_mobitag-cli" # pour installer l'autocompletion de manière permanente
```


# 📼 Buidler la demo video

La video de demo est buildée avec [`charmbracelet/vhs`](https://github.com/charmbracelet/vhs):

```sh
vhs mobitag-cli.tape
```
