package main

import (
	"log"

	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/pkg/config"
	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/pkg/db"
	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/pkg/logger"
)

func main() {
	log.Println("Basket service is getting started")

	// Setting environments
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("environment variables could not be set %v", err)
	}
	//

	// Set logger
	logger.NewLogger(cfg)
	defer logger.Close()
	//

	// Connect to DB
	db.NewPsqlDB(cfg)
	//

}
