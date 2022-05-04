package main

import (
	"log"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter"
)

func main() {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("TOKEN_KEY"), os.Getenv("TOKEN_SECRETET_KEY"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)


	tweet, _, err := client.Statuses.Update("uma mensagem puta", nil);

	if err := nil{
		log.Fatal(err)
	}

	tweet.Println(tweet.Text)
}