package mysql

import (
	"log"

	"github.com/ergochat/ergo/irc/history"
	"github.com/ergochat/ergo/irc/logger"
)

// XXX: Remove this stub and tie in the existing code
type MySQLHistory struct {
	*history.LogHistory
}

func NewMySQLHistory(logger *logger.Manager, config Config) (MySQLHistory, error) {
	log.Println("XXX: implement this")
	// 		server.historyDB.Initialize(server.logger, config.Datastore.MySQL)
	// server.historyDB.Open()
	hi := MySQLHistory{
		LogHistory: new(history.LogHistory),
	}

	return hi, nil
}
