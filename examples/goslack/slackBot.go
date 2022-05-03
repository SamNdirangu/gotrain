package goslack

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func SlackBot() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3438197530485-3438244034453-o4h2EZSpjrcIhUaB8rIddTKu")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03D1DPUFFE-3426552130615-1535ade7ab9bfa7382cbc09aeb36950d7151ca5b441494c231ecbed3d1f99a5d")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("Hello <year>", &slacker.CommandDefinition{
		Description: "Age Calculator",
		Example:     "Hello 1995",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			fmt.Println(year)
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
