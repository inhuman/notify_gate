# Notify gate
[![Build Status](https://travis-ci.org/inhuman/notify_gate.svg?branch=master)](https://travis-ci.org/inhuman/notify_gate)

Service to send notifications to telegram or slack.

## Deploy

1. Binary
```
$ go get -v -t -d ./src/...
$ go build -i -v --ldflags '-extldflags "-static"' -o bin/notify-gate src/main.go
$ ./notify-gate
```

2. Doker image
```
$ docker run idconstruct/notify_gate
```

Postrges env variables are required
```
POSTGRES_HOST
POSTGRES_PORT 
POSTGRES_DB 
POSTGRES_USER
POSTGRES_PASSWORD

TELEGRAM_BOT_TOKEN 
SLACK_AUTH_TOKEN
NG_DEBUG
NG_UI_PORT
NG_INSTANCE_TITLE

```

## Usage

1. First of all, when image deployed or binary running (for example on localhost:8080), need to generate service token. 
```
            curl --header "Content-Type: application/json" /
            --request POST                                 /  
            --data '{"name": "YOUR_SERVICE_NAME"}' /
            http://localhost:8080/service/register 
```

2. Send notifications

Replace CHANNEL_ID with actual channel id

Telegram:
```
            curl --header "Content-Type: application/json" /
            --header "X-AUTH-TOKEN: service_token"         /
            --request POST                                 /  
            --data '{"type": "TelegramChannel", "message": "test for telegram",  "uids" : ["CHANNEL_ID", "CHANNEL2_ID"]}' /
            http://localhost:8080/notify 
```
Slack:
```
            curl --header "Content-Type: application/json" /
            --header "X-AUTH-TOKEN: service_token"         /
            --request POST                                 /
            --data '{"type": "SlackChannel", "message": "test for slack",  "uids" : ["CHANNEL_ID", "CHANNEL2_ID"]}' /
            http://localhost:8085/notify
```
3. Getting telegram bot token

Use [BotFatcher](https://www.siteguarding.com/en/how-to-get-telegram-bot-api-token) to create bot 

4. Getting telegram chat id

Forward message from chat to this bot 'getidsbot'

5. Getting slack auth token

Get slack [auth token](https://get.slack.help/hc/en-us/articles/215770388-Create-and-regenerate-API-tokens)

6. Getting slack chat id

Start conversation with bot, chat id in url https://slack.com/messages/CHAT_ID/


