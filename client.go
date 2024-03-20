package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		log.Println("command received: ", cmd)

		switch cmd {
		case "/name":
			c.commands <- command{
				id:     CMD_NAME,
				client: c,
				args:   args,
			}

		case "/join":
			c.commands <- command{
				id:     CMD_NAME,
				client: c,
				args:   args,
			}

		case "/rooms":
			c.commands <- command{
				id:     CMD_NAME,
				client: c,
				args:   args,
			}

		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}

		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}

		case "/help":
			c.commands <- command{
				id:     CMD_HELP,
				client: c,
				args:   args,
			}

		case "/version":
			c.commands <- command{
				id:     CMD_VERSION,
				client: c,
				args:   args,
			}

		default:
			c.err(errors.New("unknown command"))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
