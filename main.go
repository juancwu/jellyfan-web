package main

import (
	"fmt"

	"github.com/juancwu/jellyfan/web/server"
)

func main() {
	fmt.Println("jellyfan")

	server.Serve()
}
