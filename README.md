# 📚 Mobitag

![GitHub release (latest by date)](https://img.shields.io/github/v/release/opt-nc/mobitag)
![GitHub Workflow Status](https://github.com/opt-nc/mobitag/actions/workflows/test-release.yml/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/opt-nc/mobitag)
![GitHub License](https://img.shields.io/github/license/opt-nc/mobitag)
![GitHub Repo stars](https://img.shields.io/github/stars/opt-nc/mobitag)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg)](https://github.com/goreleaser)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)
[![GoReportCard](https://goreportcard.com/badge/github.com//opt-nc/mobitag)](https://goreportcard.com/report/github.com/opt-nc/mobitag)
[![GoDoc](https://godoc.org/github.com/opt-nc/mobitag?status.svg)](https://pkg.go.dev/github.com/opt-nc/mobitag)


![Exemple d'utilisation](media/auth/mobitag.gif)

# 🔑 Prérequis

- ✅ La clé d'API `OPTNC_MOBITAGNC_API_KEY` chargée dans la session

# 🚀 QuickStart

Pour installer :

- **Manuellement** : Se rendre sur la page de [releases](https://github.com/opt-nc/mobitag-cli/releases) et télécharger la version correspondant à votre système d'exploitation
- **Automatique** : Installation via `brew` cf [homebrew-tap](https://github.com/opt-nc/homebrew-tap)

1. Installer `mobitag`

Avec `brew` : 

```sh
brew install opt-nc/homebrew-tap/mobitag
```

Ou tout simplement avec `go` :

```sh
go install github.com/opt-nc/mobitag@latest
```

2. Mettre à jour :

```sh
brew update && brew upgrade
mobitag version
```

# 🦥 Autocomplétion

Pour une UX optimale dans le terminal, il est possible d'activer l'autocomplétion :

```sh
# Pour avoir plus d'informations sur l'autocompletion :
mobitag completion <shell> --help

# <shell> peut être bash, zsh, fish, powershell.
```

Sous `zsh` :

```sh
# Pour activer l'autocompletion dans le shell courant
source <(mobitag completion zsh)
```

```sh
# Pour installer l'autocompletion de manière permanente
mobitag completion zsh > "${fpath[1]}/_mobitag"
```

# 🕹️ Utiliser

Afficher l'aide :

```sh
mobitag
```

Tester l'environnement afin de vérifier la présence de la clé API : 

```sh
mobitag dryRun
```

Envoyer un `mobit@g` : 

```sh
mobitag send --to xxxxxx --message "Hello World : a mobit@g from Go(lang) XD"
```

En indiquant également le numéro de l'expéditeur

```sh
mobitag send --to xxxxxx --message "Hello World : a mobit@g from Go(lang) XD" --from yyyyyy
```

# 🗑️ Désinstaller

## 🍺 `brew`

```sh
brew uninstall opt-nc/homebrew-tap/mobitag
brew list | grep mobitag
```

# 🤓 Cool oneliners

Depuis le terminal, les oneliners sont super cools : en une commande concise exécutée en une seule ligne dans un terminal ou un script
cela permet d’accomplir des tâches rapidement et efficacement, sans avoir à écrire un programme complet.

## Gestion du `pipe` avec la commande `pipe`

> "Hey I don't have to do anything here except glue together things that somebody else did 
for me already" - Brian Kernighan ([see short](https://youtube.com/clip/UgkxtOCaReaRRQCOu5Oo5rrOgCwb56JoX7Gw?si=cJ1TTdKZbArizMmt))


```sh
# Exemple avec la commande `whoami`
echo "Hello c'est $(whoami) : alors on se le fait ce café ?" |\
    mobitag pipe --to $DIDI_MOBILE
```

## ㊙️ Envoyer un fichier ou des secrets avec `privatebin`

[`privatebin`](https://privatebin.info/) est...

> a minimalist, open source online pastebin where the server has zero knowledge of pasted data.

On va ici l'utiliser pour envoyer des fichiers directement par `sms` depuis le terminal.

1. Disposer d'une instance à soi ou en choisir une sur [privatebin.info/directory/](https://privatebin.info/directory/)
2. Créer le [fichier de conf](https://github.com/gearnode/privatebin/blob/master/doc/privatebin.conf.5.md#examples) `~/.config/privatebin/config.json`
3. Télécharger et installer [`gearnode/privatebin`](https://github.com/gearnode/privatebin)
4. Profiter

### 🐮 Un petit coup de `cowsay`

Avec [`cowsay`](https://cowsay.diamonds/):

```sh
cowsay -f tux "Mobitag c'est VACHEMENT cool...surtout depuis le terminal et pipé avec privatebin"\
    | privatebin create\
    | mobitag pipe --to $MOBILIS_DEST
```

### 🔐 Communiquer un fichier de secrets

```sh
cat secrets.txt\
    | privatebin create\
    | mobitag pipe --to $MOBILIS_DEST
```

## 📤 Envoi de `sms` en masse avec `awk`

### A propos de `awk`

Comment développer un `cli` sans proposer une intégration avec le légendaire `awk` ?
Voyons donc comment envoyer des sms en masse en combinant `mobitag` et `awk` 🚀

🎤 D'abord, bien comprendre la phitlosophie de `awk` : [7' interview of Brian Kernighan for the legacy.](https://www.youtube.com/watch?v=W5kr7X7EG4o)

### Envoyer un `csv`

Supposons que l'on ait le `csv` suivant (par exemple en sortie d'un traitement précédent) : 

```
dest,msg,from
NUMERO_1,"NERD ALERT : Demo automatisation csv avec awk - exemple 1",""
NUMERO_2,"NERD ALERT : Demo automatisation csv avec awk - exemple 2",""
```
Alors on peut générer un "dry run" (ie. génération de commandes sans les éxécuter) : 

```sh
# Génération des commandes : 
awk -F',' 'NR > 1 && $1 != "" {print "mobitag send --to " $1 " --message " $2 " --from " $3 }' mobitags.csv
```

Puis exécuter :


```
awk -F',' 'NR > 1 && $1 != "" {print "mobitag send --to " $1 " --message " $2 " --from " $3 }' mobitags.csv |\
    bash
```


# 🧑‍🤝‍🧑 Equipe

Ce projet d'innovation frugale n'aurait pas vu le jour sans une équipe, par ordre d'entré sur le projet :

1. [👱‍♀️ Michèle BARRE](https://www.linkedin.com/in/michelebarre/), aka. [`@mbarre`](https://github.com/mbarre/) : backend dev et UX bêta-testeuse
2. [🤓 Adrien SALES](https://www.linkedin.com/in/adrien-sales/), aka. [`@adriens`](https://github.com/adriens/) : Premier proto Go, Story Teller, Product Owner et alpha testeur
3. [🥋 Vinh FAUCHER](https://www.linkedin.com/in/vinh-faucher/) aka. [`@supervinh`](https://github.com/supervinh/) : Core Go dev
4. [🧑🏾‍🦱 Romain PELIZZO](https://www.linkedin.com/in/romain-pelizzo/) aka. [`@Draks898`](https://github.com/Draks898) : Bêta-testeur
5. 📊 Liste exhaustive des [contributeurs](https://github.com/opt-nc/mobitag-cli/graphs/contributors)

# 📖 Histoire de ce `cli`

Cette repo était à l'origine une **première expérimentation dont le but était de découvrir le language [`Go`](https://go.dev/)**,
sur un cas concret car... c'est plus amusant et beaucoup plus motivant 🤓.

Cette expérimentation avait donc pour but de :

> créer un `cli` permettant d'envoyer des mobitags depuis le terminal.


# 🔖 Ressources

- 🔖 Site web officiel [`mobitag.nc`](http://www.mobitag.nc)
- [🥳 Mobitag.nc... 25 ans plus tard, des sms en SaaS via API{GEE}](https://dev.to/optnc/mobitagnc-25-ans-plus-tard-des-sms-en-saas-via-apigee-2h9e)
- [📲 Mobitag.nc for dummies](https://www.kaggle.com/code/optnouvellecaldonie/mobitag-nc-for-dummies)
- [⏱️ Mobitag Go Hackathon 2024-06-22 week-end](https://dev.to/adriens/mobitag-go-hackathon-2024-06-22-week-end-2n16)
- [⏱️ Mobitag Hackathon week-end du 2024-06-22](https://youtu.be/yVoMg7CXgaM)
