# Psychotron - A Slack Bot 🤖

Psychotron is a slack-bot written using slack, wit.ai and wolfram alpha go apis.

## Motivation behind the name `Psychotron` 🤖

During one weekend, I was listening to Psychotron by Megadeath (https://www.youtube.com/watch?v=sjXI5JI_4nc). Just felt it would be a great bot name. In my defense, it was late in the night. 

## How to setup your repo:

Create `keys` file under psychotron directory. Contents of `psychotron/keys` look like this:

```
export WOLFRAM_APP_ID="<WOLFRAM_APP_ID>"
export WIT_AI_ACCESS_TOKEN=<WIT_AI_ACCESS_TOKEN_VALUE>
export SLACK_ACCESS_TOKEN="<SLACK_ACCESS_TOKEN_VALUE>"
```

## How to run psychotron:
```
cd $GOPATH/src/github.com/nlopes/slack && git checkout tags/v0.4.0

Reference: https://github.com/go-chat-bot/bot/issues/92

source psychotron/keys

go run psychotron/main.go
```

*Note*

1] Slack go-client is open-source (https://github.com/nlopes/slack)

2] wit.ai go client is open-source (https://github.com/christianrondeau/go-wit)

2] WolframAlpha allows only 2000 API calls per month for unpaid free account.
(wolfram go-client: https://github.com/Krognol/go-wolfram)