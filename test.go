package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/ergochat/ergo/irc"
	"github.com/ergochat/ergo/irc/logger"
)

func ergo() {
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
	log.Println("Starting ergo")
	server.Run()
}

func cliRecv(conn net.Conn, wg *sync.WaitGroup) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	wg.Done()
}

func join(nick string) (net.Conn, error) {
	conn, err := net.Dial("tcp", "localhost:6667")
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(conn, "NICK", nick)
	fmt.Fprintln(conn, "USER test test test :test")
	fmt.Fprintln(conn, "JOIN #foo")
	return conn, nil
}

func main() {
	go ergo()

	time.Sleep(750 * time.Millisecond)

	conn, err := join("test")
	if err != nil {
		log.Fatal(err)
	}
	conn2, err := join("test2")
	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go cliRecv(conn, wg)
	go cliRecv(conn2, wg)

	fmt.Fprintln(conn, "PRIVMSG #foo :channel message")
	fmt.Fprintln(conn2, "PRIVMSG test :direct message")
	time.Sleep(200 * time.Millisecond)
	fmt.Fprintln(conn, "QUIT :ta ta")
	fmt.Fprintln(conn2, "QUIT :ta ta")

	wg.Wait()
}
