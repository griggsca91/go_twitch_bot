package main

import (
	"fmt"
	"net"
	//	"os"
	"bufio"
	"net/textproto"
	"time"
)

type Bot struct {
	server  string
	port    string
	nick    string
	channel string
	conn    net.Conn
}

func (bot *Bot) ConsoleInput() {
	// _ := bufio.NewReader(os.Stdin)

}

func (bot *Bot) Connect() {

	var err error
	fmt.Printf("Attempting to connect to server...\n")
	bot.conn, err = net.Dial("tcp", bot.server+":"+bot.port)
	if err != nil {
		fmt.Printf("Unable to connect to Twitch IRC server! Reconnecting in 10 seconds...\n")
		time.Sleep(10 * time.Second)
		bot.Connect()
	}

	fmt.Printf("Connected to IRC server %s\n", bot.server)
}

func (bot *Bot) Chat(message string) {
	fmt.Fprintf(bot.conn, message+"\r\n")
	fmt.Printf(message + "\r\n")
}



func ParseMessage(message string) {




}

func main() {
	bot := &Bot{
		server:  "irc.twitch.tv",
		port:    "6667",
		nick:    "bad_hombres_bot",
		channel: "bad_hombres",
		conn:    nil,
	}

	pass := "oauth:30zdoyl8vntrr7p0g2ub3sx3ferb2r"

	bot.Connect()

	bot.Chat(fmt.Sprintf("USER %s 8 * :%s", bot.nick, bot.nick))
	bot.Chat(fmt.Sprintf("PASS %s", pass))
	bot.Chat(fmt.Sprintf("NICK %s", bot.nick))
	bot.Chat(fmt.Sprintf("JOIN #%s", bot.channel))

	reader := bufio.NewReader(bot.conn)
	tp := textproto.NewReader(reader)

	for {
		line, err := tp.ReadLine()

		if err != nil {
			fmt.Printf("ERROR")
			break
		}
		fmt.Println(line)

	}

}
