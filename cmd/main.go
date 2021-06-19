package main

import (
	"fmt"

	"github.com/shoppehub/fastcms/server"
)

func main() {
	r := server.New()
	r.Run("0.0.0.0:" + fmt.Sprint(server.Port))
}
