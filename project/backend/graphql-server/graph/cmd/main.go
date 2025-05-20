package main

import (
	"mindmap-backend/graphql-server/graph/api"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	api.ServerHandler()
}
