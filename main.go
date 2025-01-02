package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	// "regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix string = "!gobot"
const name string = "rahul"

type Answers struct {
	OriginChannelID string
	FavFood         string
	FavGame         string
}

var responses map[string]Answers = map[string]Answers{}

func main() {
	godotenv.Load()
	token := os.Getenv("DISCORD_BOT_TOKEN") // Ensure the token is read from the environment
	sess, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		result := printx()

		args := strings.Split(m.Content, " ")

		if args[0] == name && args[1] == "--help" {
			s.ChannelMessageSend(m.ChannelID, "lodu its simple for now, just properly type !gobot and then write the correct name")
			return
		}

		if args[0] == name && args[1] == "printx" {
			s.ChannelMessageSend(m.ChannelID, result)
			return
		}

		if args[0] != prefix {
			return
		}

		if args[1] == "sam" {
			useWordHandler(s, m)
		}

		if args[1] == "prompt" {
			usePromptHandler(s, m)
		}

	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	fmt.Println("the bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func useWordHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "gay!")
}

func usePromptHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	channel, err := s.UserChannelCreate(m.Author.ID)

	if err != nil {
		log.Panic(err)
	}

	if _, ok := responses[channel.ID]; !ok {
		responses[channel.ID] = Answers{
			OriginChannelID: m.ChannelID,
			FavFood:         "",
			FavGame:         "",
		}

		s.ChannelMessageSend(channel.ID, "Hey there ! here are some questions for u")
		s.ChannelMessageSend(channel.ID, "whats your favortie food? ")
	} else {
		s.ChannelMessageSend(channel.ID, "We're still waiting 😅")
	}
}
