package main

import (
	"log"
	"video_web/bootstrap"
)

func main() {

	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}

}
