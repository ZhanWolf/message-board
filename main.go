package main

import (
	"github.com/gin-gonic/gin"
	"message-board/cmd"
)

func main() {
	r := gin.Default()
	cmd.Userroute(r)
	cmd.Messageroute(r)
	r.Run(":6060")
}
