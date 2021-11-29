package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

//global stuff for shortcuts
var p = fmt.Println

func main() {

	// Setting up the token
	btok, _ := ioutil.ReadFile("token")
	token := string(btok)
	token = strings.Replace(token, "\n", "", -1)

	// Setting up Bot connection
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is up")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	} else {
		fmt.Println(m)
	}

	if strings.Contains(strings.ToLower(m.Content), "!test") {
		s.ChannelMessageSend(m.ChannelID, "testSuccesfull")
	}

	// run command on all nodes
	// firstCharacter := m.Content[0:2]

	if strings.HasPrefix(m.Content, "! ") {
		command_string := m.Content[2:len(m.Content)] // get everything after the '! '
		fmt.Println(command_string)

		out, errorMessage := runCommand(command_string)

		fmt.Println(out)
		fmt.Println(errorMessage)

		if errorMessage != "" {
			fmt.Println(errorMessage)

			s.ChannelMessageSend(m.ChannelID, errorMessage)

			return
		}

		s.ChannelMessageSend(m.ChannelID, string(out))

	}

	fmt.Println(m.Author.Username, ": ", m.Content)

}

func runCommand(command string) (outString string, errorMessage string) {

	var shell string
	errorMessage = ""

	if runtime.GOOS == "windows" {
		shell = "PS"
	} else {
		shell = "sh"
	}

	p("Command: " + command)
	p("Shell: " + shell)

	out, err := exec.Command(shell, "-c", command).Output()
	if err != nil {
		fmt.Println(err.Error())
		errorMessage = err.Error()

		return
	}

	outString = string(out)

	return

}
