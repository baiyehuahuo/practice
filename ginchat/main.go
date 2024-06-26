package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	AutoMigrates()

	r := router.Router()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func AutoMigrates() {
	var err error
	if err = models.AutoMigrateUserBasic(); err != nil {
		log.Fatal(err)
	}
	if err = models.AutoMigrateMessage(); err != nil {
		log.Fatal(err)
	}
	if err = models.AutoMigrateContact(); err != nil {
		log.Fatal(err)
	}
	if err = models.AutoMigrateGroupBasic(); err != nil {
		log.Fatal(err)
	}
	if err = models.AutoMigrateCommunity(); err != nil {
		log.Fatal(err)
	}
}
