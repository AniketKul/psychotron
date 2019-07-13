package main

import (
	"log"
	"os"

	wolfram "github.com/Krognol/go-wolfram"
	wit "github.com/christianrondeau/go-wit"
	"github.com/nlopes/slack"
)

var (
	slackClient   *slack.Client
	witaiClient   *wit.Client
	wolframClient *wolfram.Client
)

func handleMessage(event *slack.MessageEvent) {
	// fmt.Printf("%v\n", event)
	response, err := witaiClient.Message(event.Msg.Text)
	if err != nil {
		log.Fatalf("func handleMessage - unable to get wit.ai response: %v", err)
		return
	}

	var (
		confidenceThreshold = 0.5
		topEntity           wit.MessageEntity
		topEntityKey        string
	)

	for entityKey, entityList := range response.Entities {
		for _, entity := range entityList {
			if entity.Confidence > confidenceThreshold && entity.Confidence > topEntity.Confidence {
				topEntity = entity
				topEntityKey = entityKey
			}
		}
	}
	replyToUser(event, topEntityKey, topEntity)
}

func replyToUser(event *slack.MessageEvent, entityKey string, entity wit.MessageEntity) {
	// if you do not set AsUser: true then the bot will send a message to slackbot and it won't visible in the same thread to user.
	switch entityKey {
	case "greetings":
		slackClient.PostMessage(event.User, "Wassup fam! how is it going?", slack.PostMessageParameters{
			AsUser: true,
		})
		return

	case "wolfram_search_query":
		res, err := wolframClient.GetShortAnswerQuery(entity.Value.(string), wolfram.Metric, 1000)
		if err != nil {
			log.Fatalf("func replyToUser - unable to get wolfram result: %v", err)
		}
		slackClient.PostMessage(event.User, res, slack.PostMessageParameters{
			AsUser: true,
		})
		return
	}
	slackClient.PostMessage(event.User, "Apologies! I am not sure I follow!", slack.PostMessageParameters{
		AsUser: true,
	})
}

func main() {
	slackClient = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	witaiClient = wit.NewClient(os.Getenv("WIT_AI_ACCESS_TOKEN"))
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(event)
		}
	}
}
