package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Auth/config"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Auth/controllers"
)

var (
	router = gin.Default()
	port   string
)

func StartApplication() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router.POST("/login", inDB.LoginUser)
	router.POST("/user", inDB.CreateUser)
	log.Print("Starting the application\n")
	log.Print("Enter Port:\n")
	fmt.Scanln(&port)
	po := (":")
	po_rt := po + port
	router.Run(po_rt)
}
