# ❔ A propos

Cette repo est une **première expérimentation dont le but est de découvrir le
language [`Go`](https://go.dev/)**, sur un cas concret car... c'est plus amusant
et beaucoup plus motivant 🤓.

Cette expérimentation a donc pur but de créer un cli permettant d'envoyer des mobitags
depuis le terminal.

# 🔖 Ressources

- Site web officiel http://www.mobitag.nc
- [🥳 Mobitag.nc... 25 ans plus tard, des sms en SaaS via API{GEE}](https://dev.to/optnc/mobitagnc-25-ans-plus-tard-des-sms-en-saas-via-apigee-2h9e)
- [📲 Mobitag.nc for dummies](https://www.kaggle.com/code/optnouvellecaldonie/mobitag-nc-for-dummies)

# ✅ Prérequis

- [x] Tooling Go
- [x] Une clé d'API, chargée dans l'environnement `OPTNC_MOBITAGNC_API_KEY`

# 🚀 Getting started

## ⚙️ Builder

```shell
go build mobitag.go

```

# 🕹️ Essayer

```sh
./mobitag -h
```

```sh
# Tester l'environnement
./mobitag -dry-run

```
# 🥳 Envoyer un `mobit@g`

```sh
/mobitag -to 842984 -message "Hello World : a mobit@g from Go(lang) XD"

```
