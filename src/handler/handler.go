package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service"
	"model"
	log "github.com/Sirupsen/logrus"
	"chatserver"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"common"
	"errors"
)

var(
	BusinessService  *service.BusinessService
	Hub				 *chatserver.Hub
	StoreFilePath 	 string
)

func Chat(c *gin.Context){

	chatserver.Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := chatserver.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Errorf("Handler.chatserver.Upgrader.Upgrade is failed, the err is [%+v]", err)
		return
	}
	client := &chatserver.Client{Hub: Hub, Conn: conn, Send: make(chan []byte, 1024)}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}

func NewUser(c *gin.Context){

	var	user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Errorf("Bind Json is err, the err is [%+v]", err)
		return
	}
	log.Println("Bind the user is succeed.")
	err = BusinessService.NewUser(&user)
	if err != nil {
		HandleError(c, "Handler.NewUser is err, the err is [%+v]",  err)
		return
	}
	response := model.NewInfoResponse("success")
	HandleSuccess(c, response)
}

func NewFile(c *gin.Context){

	newfile := new(model.File)
	c.Request.ParseForm()
	file, handle, err := c.Request.FormFile("file")
	if err != nil {
		log.Errorf("Received the file is failed, the err is [%+v]", err)
		return
	}
	fileext := filepath.Ext(handle.Filename)
	if common.CheckFileFormat(fileext) == false {
		log.Error("the file is illegel. ")
		HandleError(c, "the file is illegel. ",  errors.New("the file is illegel. "))
		return
	}else{
		f, err := os.OpenFile( StoreFilePath + handle.Filename, os.O_WRONLY|os.O_CREATE, 0660)
		if err != nil{
			log.Errorf("Open the file is failed, the err is [%+v]", err)
			HandleError(c, "Upload is failed, please wait a moment and try again",  errors.New("Upload is failed, please wait a moment and try again. "))
			return
		}
		io.Copy(f, file)
		defer f.Close()
		defer file.Close()
		newfile.FileName = handle.Filename
		err = BusinessService.NewFile(newfile)
		if err != nil{
			log.Errorf("Create the new file is failed, the err is [%v]", err)
			return
		}
		go func() {
			Hub.Broadcast <- []byte("upload")
		}()
		response := model.NewInfoResponse("success")
		HandleSuccess(c, response)
	}
}

func UserLogin(c *gin.Context){

	var	user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Errorf("Bind Json is err, the err is [%+v]", err)
		return
	}
	log.Println("Bind the user is succeed.")
	err = BusinessService.UserLogin(&user)
	if err != nil {
		HandleError(c, "Handler.UserLogin is err",  err)
		return
	}
	response := model.NewInfoResponse("success")
	HandleSuccess(c, response)
}

func UserLogout(c *gin.Context){

	userName := c.Param("username")
	err := BusinessService.UserLogout(userName)
	if err != nil {
		HandleError(c, "Handler.UserLogout is err",  err)
		return
	}
	response := model.NewInfoResponse("success")
	HandleSuccess(c, response)
}

func GetAllUser(c *gin.Context){

	response := BusinessService.GetAllUser()
	HandleSuccess(c, response)
}


func GetAllMessage(c *gin.Context){

	response := BusinessService.GetAllMessage()
	HandleSuccess(c, response)
}

func GetAllFile(c *gin.Context){

	response := BusinessService.GetAllFile()
	HandleSuccess(c, response)
}


func HandleSuccess(c *gin.Context, result interface{}) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, result)
}

func HandleError(c *gin.Context, request string, err error) {
	c.Header("Access-Control-Allow-Origin", "*")
	response := make(map[string]interface{})
	if err.Error() == "exist" {
		c.JSON(409, response)
	} else if err.Error() == "noexist" {
		c.JSON(410, response)
	} else {
		response["error"] = request
		response["description"] = fmt.Sprint(err)
		c.JSON(500, response)
	}
}
