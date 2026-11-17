package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var appVersion = "v1.0.0"

var kbotCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("TELE_TOKEN")
		if token == "" {
			log.Fatal("TELE_TOKEN environment variable is not set")
		}

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("cannot create bot: %s", err)
		}

		kbot.Handle("/start", func(m telebot.Context) error {
			return m.Send("Hello! I'm kbot " + appVersion + ". Send /hello or any text.")
		})

		kbot.Handle("/hello", func(m telebot.Context) error {
			return m.Send(fmt.Sprintf("Hello, %s!", m.Sender().FirstName))
		})

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf("message from %s: %s", m.Sender().Username, m.Text())
			return m.Send("You said: " + m.Text())
		})

		log.Printf("kbot %s started", appVersion)
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}
