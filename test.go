package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ergochat/ergo/irc"
	"github.com/ergochat/ergo/irc/logger"
)

type Client struct {
	net.Conn
	name  string
	lines chan string
}

func NewClient(network, addr, name string) (*Client, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	cli := Client{
		Conn:  conn,
		name:  name,
		lines: make(chan string, 100),
	}
	go cli.readLines()
	return &cli, nil
}

func (c *Client) readLines() {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%s: %s", c.name, line)
		c.lines <- line
	}
}

func (c *Client) ReadUntil(match string) {
	for line := range c.lines {
		if strings.Contains(line, match) {
			return
		}
	}
}

func (c *Client) Join(channel string) {
	fmt.Fprintln(c, "JOIN", channel)
	c.ReadUntil("JOIN")
}

func (c *Client) Login(nick, password string, firstRun bool) {
	authStr := fmt.Sprintf("%s\000%s\000%s", nick, nick, password)
	auth64 := base64.StdEncoding.EncodeToString([]byte(authStr))
	if !firstRun {
		fmt.Fprintln(c, "CAP REQ :sasl")
	}
	fmt.Fprintln(c, "NICK", nick)
	fmt.Fprintln(c, "USER test test test :test")
	if firstRun {
		fmt.Fprintln(c, "PRIVMSG nickserv :register", password)
	} else {
		fmt.Fprintln(c, "AUTHENTICATE PLAIN")
		fmt.Fprintln(c, "AUTHENTICATE ", auth64)
	}
	c.ReadUntil("903 * :Authentication successful")
	fmt.Fprintln(c, "CAP END")
}

func main() {
	config, err := irc.LoadConfig("ircd.yaml")
	if err != nil {
		log.Fatal("Config file did not load successfully: ", err.Error())
	}

	logman, err := logger.NewManager(config.Logging)
	if err != nil {
		log.Fatal("Logger did not load successfully:", err.Error())
	}

	server, err := irc.NewServer(config, logman)
	if err != nil {
		log.Fatal("Could not load server:", err)
	}
	go server.Run()

	first := false
	cli1, err := NewClient("tcp", "localhost:6667", "1")
	if err != nil {
		log.Fatal(err)
	}
	cli2, err := NewClient("tcp", "localhost:6667", "2")
	if err != nil {
		log.Fatal(err)
	}

	cli1.Login("test", "password123", first)
	cli2.Login("test2", "password456", first)

	cli1.Join("#foo")
	cli2.Join("#foo")

	fmt.Fprintln(cli1, "PRIVMSG #foo :channel message")
	fmt.Fprintln(cli2, "PRIVMSG test :direct message")

	cli2.ReadUntil("channel message")
	cli1.ReadUntil("direct message")

	fmt.Fprintln(cli1, "QUIT :ta ta")
	fmt.Fprintln(cli2, "QUIT :ta ta")
	cli1.ReadUntil("ERROR")
	cli2.ReadUntil("ERROR")
}
