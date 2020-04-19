package app

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Auth/config"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Auth/controllers"
)

var (
	router = gin.Default()
)

func StartApplication() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router.POST("/user", inDB.CreateUser)
	log.Print("Starting the application\n")
	router.Run(":8080")
}
