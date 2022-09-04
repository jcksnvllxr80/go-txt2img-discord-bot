package bot

import (
	"bufio"
	"io"
	"fmt"
	"strings"
	"os/exec"
	"io/ioutil"
	"os"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/jcksnvllxr80/go-txt2img-discord-bot/config"
)

var WORK_DIR string = "/home/aaron/stable-diffusion/"
var OUT_DIR string = WORK_DIR + "samples"
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

func clearOutputDir() {
	fmt.Println("Clearing previous files...")
	cmd := exec.Command("rm", "-rf", OUT_DIR)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)

	cmd.Wait()
}

func createImageFromText(imageText string) {
	fmt.Println("Creating images...")
	script := WORK_DIR + "scripts/txt2img.py"
	cmd := exec.Command("conda", "run", "-n", "ldm", "--cwd", WORK_DIR, "python3", script, "--prompt", "\"" + imageText + "\"", "--plms", "--ckpt", "sd-v1-4.ckpt", "--skip_grid", "--n_samples", "1", "--outdir", "./")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)

	cmd.Wait()
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	botCallStr := "<@1015747298719973459>"
	// fmt.Println("Message content is: " + m.Content)  // temp debug logging
	if strings.Index(m.Content, botCallStr) == 0 {
		imageToCreateStr := strings.TrimLeft(m.Content[len(botCallStr):]," ")

		clearOutputDir()
		_, _ = s.ChannelMessageSend(m.ChannelID, "Creating two images... please wait about 15 seconds.")
		createImageFromText(imageToCreateStr)
		
		_, _ = s.ChannelMessageSend(m.ChannelID, "Uploading mages... ")
		files, err := ioutil.ReadDir(OUT_DIR)
		if err != nil {
			log.Fatal(err)
		}

		for i, file := range files {
			message := "Image #" + strconv.Itoa(i + 1) + " for description, \"" + imageToCreateStr + "\"."
			// read the file
			f, err := os.Open(OUT_DIR + "/" + file.Name())
			check(err)
			// send the file to discord
			_, _ = s.ChannelFileSendWithMessage(m.ChannelID, message, file.Name(), f)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
