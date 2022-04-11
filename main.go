package main

import (
	"log"

	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/pkg/config"
)

func main() {
	log.Println("Basket service is getting started")

	// Setting environments
	_, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("environment variables could not be set %v", err)
	}
	//

}
