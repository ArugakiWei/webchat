package service

import (
	log "github.com/Sirupsen/logrus"
	"model"
	"github.com/astaxie/beego/orm"
	"errors"
)

var(
	DB		string
)

type BusinessService struct {
	MessageServiceInstance     	*MessageService
	UserServiceInstance			*UserService
	FileServiceInstance			*FileService
}

func GetOrmer() orm.Ormer {
	ormer := orm.NewOrm()
	return ormer
}

func Init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", DB)
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 100)
	orm.RegisterModel(new(model.User))
	orm.RegisterModel(new(model.Message))
	orm.RegisterModel(new(model.File))
}

func NewBusinessService() *BusinessService {
	return &BusinessService{
		MessageServiceInstance:    	NewMessageService(),
		UserServiceInstance:		NewUserService(),
		FileServiceInstance:		NewFileService(),
	}
}


func (a *BusinessService) NewUser(user *model.User) error {

	u, err  := a.UserServiceInstance.GetByUserName(user.UserName)
	if err != nil{
		log.Errorf("UserServiceInstance.GetByUserName, Get the User is failed:the err is [%+v]", err)
		return err
	}
	if u != nil{
		return errors.New("exist")
	}
	user.IsOnline = "false"
	err = a.UserServiceInstance.Insert(user)
	if err != nil{
		log.Errorf("UserServiceInstance.Insert, Insert into the msyql is failed, the err is [%+v]",err)
		return err
	}
	return nil
}


func (a *BusinessService) UserLogin(user *model.User) error {

	u, err  := a.UserServiceInstance.GetByUserName(user.UserName)
	if err != nil{
		log.Errorf("UserLogin.GetByUserName, Get the User is failed:the err is [%+v]", err)
		return err
	}
	if u != nil{
		if u.PassWord == user.PassWord{
			if u.IsOnline == "true"{
				log.Error("The user have been logined.")
				return errors.New("The user have been logined ")
			}else{
				u.IsOnline = "true"
				err = a.UserServiceInstance.Update(u)
				if err != nil{
					log.Errorf("Login is failed, the err is [%+v]", err)
					return err
				}
				return nil
			}
		}else{
			return errors.New("The user name or password is incorrect. ")
		}
	}
	return errors.New("The user is no existed. ")
}

func (a *BusinessService) UserLogout(userName string) error {

	u, err  := a.UserServiceInstance.GetByUserName(userName)
	if err != nil{
		log.Errorf("UserLogout.GetByUserName, Get the User is failed:the err is [%+v]", err)
		return err
	}
	if u != nil{
		u.IsOnline = "false"
		err = a.UserServiceInstance.Update(u)
		if err != nil{
			log.Errorf("Login is failed, the err is [%+v]", err)
			return err
		}
		return nil
	}
	return errors.New("The user is no existed. ")
}

func (a *BusinessService) NewFile(file *model.File) error {

	f, err  := a.FileServiceInstance.GetByFileName(file.FileName)
	if err != nil{
		log.Errorf("FileService.InsertToDB.GetByFileName, Get the File is failed:the err is [%+v]", err)
		return err
	}
	if f != nil{
		return errors.New("exist")
	}
	err = a.FileServiceInstance.Insert(file)
	if err != nil{
		log.Errorf("FileService.InsertToDB.Insert, Insert into the msyql is failed, the err is [%+v]",err)
		return err
	}
	return nil
}

func (a *BusinessService) GetAllMessage() *model.MessageResponse {

	allMessage := a.MessageServiceInstance.GetAllMessage()
	response := model.NewMessageResponse(allMessage)
	return response
}

func (a *BusinessService) GetAllUser() *model.UserResponse {

	allUser := a.UserServiceInstance.GetAllUser()
	var res []model.UserRes
	var temp model.UserRes
	for _, user := range allUser{
		temp.UserName = user.UserName
		temp.IsOnline = user.IsOnline
		res = append(res, temp)
	}
	response := model.NewUserResponse(res)
	return response
}


func (a *BusinessService) GetAllFile() *model.FileResponse {

	allFile := a.FileServiceInstance.GetAllFile()
	response := model.NewFileResponse(allFile)
	return response
}

func (a *BusinessService) CreateMsgAndCheck(message string) (*model.Message, bool){

	msg := a.MessageServiceInstance.StrToMsg(string(message))
	if msg == nil{
		log.Error("BusinessService.CreateMsgAndCheck is failed")
		return nil, false
	}
	user, err := a.UserServiceInstance.GetByUserName(msg.UserName)
	if err != nil{
		log.Errorf("BusinessService.CreateMsgAndCheck.GetByUserName is failed, the err is [%+v]", err)
		return nil, false
	}
	if user != nil && user.IsOnline == "true" {
		err := a.MessageServiceInstance.Insert(msg)
		if err != nil{
			log.Errorf("BusinessService.CreateMsgAndCheck.Insert is failed, the err is [%+v]", err)
			return nil, false
		}
		return msg, true
	}
	return nil, false
}
