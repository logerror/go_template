package main

import "welights.net/go_template/pkg/log"

func main() {
	logger := log.InitLogger("error")
	defer logger.Sync()
}
