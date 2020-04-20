package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Fetching/config"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Fetching/controllers"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Fetching/controllers/auth"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Fetching/controllers/storage"
)

var (
	router = gin.Default()
	port   string
)

func StartApplication() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router.POST("/login", inDB.LoginUser)
	router.GET("/showuser", auth.Auth, inDB.ShowUser)
	router.GET("/showstorage", auth.Auth, storage.ShowStorage)
	router.POST("/adminpage", auth.IsAdmin, storage.StorageAdmin)

	log.Print("Starting the application\n")

	log.Print("Enter Port:\n")
	fmt.Scanln(&port)
	po := (":")
	po_rt := po + port
	router.Run(po_rt)
}
