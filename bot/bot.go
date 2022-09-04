package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jcksnvllxr80/go-txt2img-discord-bot/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	botCallStr := "<@1015747298719973459>"
	fmt.Println("Message content is: " + m.Content)  // temp debug logging

	if strings.Index(m.Content, botCallStr) == 0 {
		imageToCreateStr := strings.TrimLeft(m.Content[len(botCallStr):]," ")
		// TODO:
		// call the process to create the image with imageToCreateStr as an argument
		// get path to process output files
		// send the created images back in the messageSend
		_, _ = s.ChannelMessageSend(m.ChannelID, "This is a file upload of an image whose description is, \"" + imageToCreateStr + "\".")
	}
}
