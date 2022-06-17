package main

import (
	db "ecommerce/db/sqlc"
	_ "ecommerce/docs"
	api "ecommerce/handlers"
	"ecommerce/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @title           Asarfi Acqurier API
// @version         1.0
// @description     Asarfi Acqurier API is an api for card management and payment processing
// @termsOfService  http://asarfi.live./terms/

// @contact.name   Asarfi Pay
// @contact.url    http://www.asarfi.live/contact
// @contact.email  asarfi.pay@gmaiil.com

// @host     localhost:8080
// @schemes  http
// @BasePath
func main() {
	log.Info("Starting ASARFI API ...")
	log.Info("Initiating connection to core_asarfi")
	err := models.InitCBS()
	if err != nil {
		log.Panic(err)
		return
	}
	config := models.GetConfiguration()

	connection_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.CBS.Host,
		config.CBS.Port,
		config.CBS.User,
		config.CBS.Password,
		config.CBS.Dbname)

	conn, err := sql.Open("postgres", connection_string)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
}

