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

## psychotron 🤖 in action:


<img width="1097" alt="ss_1" src="https://user-images.githubusercontent.com/21697139/61167225-a93d8b80-a4f0-11e9-99a1-3e7d339397bf.png">

<img width="944" alt="ss_2" src="https://user-images.githubusercontent.com/21697139/61167228-b22e5d00-a4f0-11e9-82eb-da4c7b1f455a.png">

<img width="1104" alt="ss_3" src="https://user-images.githubusercontent.com/21697139/61167229-b8243e00-a4f0-11e9-8f3b-4926b57c97f4.png">

<img width="1246" alt="ss_4" src="https://user-images.githubusercontent.com/21697139/61167230-bf4b4c00-a4f0-11e9-8324-ce3c6deac08f.png">

<img width="1243" alt="ss_5" src="https://user-images.githubusercontent.com/21697139/61167231-c3776980-a4f0-11e9-92a4-d9b11436c650.png">

<img width="961" alt="ss_6" src="https://user-images.githubusercontent.com/21697139/61167235-c83c1d80-a4f0-11e9-84fe-6c45a5de7355.png">

<img width="839" alt="ss_7" src="https://user-images.githubusercontent.com/21697139/61167238-cb370e00-a4f0-11e9-8f17-4af72d3f805f.png">


**Notes:**

1] Slack go-client is open-source (https://github.com/nlopes/slack)

2] wit.ai go client is open-source (https://github.com/christianrondeau/go-wit)

2] WolframAlpha allows only 2000 API calls per month for unpaid free account.
(wolfram go-client: https://github.com/Krognol/go-wolfram)