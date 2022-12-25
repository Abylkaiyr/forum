package main

import (
	"github.com/Abylkaiyr/forum/pkg/delivery"
	"github.com/Abylkaiyr/forum/pkg/server"
)

func init() {
	delivery.StorageInit()
}

func main() {
	server.Server()
}
