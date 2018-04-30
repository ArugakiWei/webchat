package main

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
	"time"
	"handler"
	"service"
	"chatserver"
)


func main() {

	app := cli.NewApp()
	app.Name = "Chat Server"
	app.Usage = "Chat Server"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "",
			Usage: "config file used to run programmer",
		},
	}
	app.Action = func(c *cli.Context) {
		router := gin.Default()
		router.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		DB := os.Getenv("db")
		DB = "root:18434361185@tcp(127.0.0.1:3306)/webchat?charset=utf8"
		service.DB = DB
		StoreFilePath := os.Getenv("storefilepath")
		service.Init()
		//StoreFilePath = "/Users/arugaki/test/"
		handler.StoreFilePath = StoreFilePath
		businessService := service.NewBusinessService()
		handler.BusinessService = businessService
		chatserver.BussinessService = businessService
		hub := chatserver.NewHub()
		handler.Hub = hub
		go hub.Run()

		router.GET("/user/all", handler.GetAllUser)				//获得所有用户  (用户列表)
		router.POST("/user/registered", handler.NewUser)			//用户注册
		router.POST("/user/login", handler.UserLogin)			//用户登录
		router.POST("/user/logout/:username", handler.UserLogout)	//用户注销

		router.GET("/chat", handler.Chat)						//进入聊天室

		router.GET("/message/all", handler.GetAllMessage)		//获得所有消息 （消息记录）
		router.GET("/file/all", handler.GetAllFile)				//获得所有文件  (文件共享)
		router.POST("/file/upload", handler.NewFile)				//上传文件
		router.StaticFS("/download", http.Dir(StoreFilePath))	//下载文件

		server := &http.Server{
			Addr:           ":8989",
			Handler:        router,
			ReadTimeout:    300 * time.Second,
			WriteTimeout:   300 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		server.ListenAndServe()
	}
	app.Run(os.Args)
}