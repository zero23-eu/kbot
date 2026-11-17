# kbot

Simple Telegram bot in Go (cobra + telebot.v4).

Bot: t.me/pro_me_te_bot

## Install

```bash
git clone https://github.com/zero23-eu/kbot
cd kbot
go mod tidy
go build -o kbot .
```

## Run

```bash
export TELE_TOKEN=<token from BotFather>
./kbot start
```

## Commands

- `/start` – greeting
- `/hello` – greets you by name
- any text – bot echoes it back
