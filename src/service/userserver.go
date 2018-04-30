package service

import (
	"data"
	log "github.com/Sirupsen/logrus"
	"model"
	"github.com/astaxie/beego/orm"
)

type UserService struct {
	User		*data.UserTable
}


func NewUserService() *UserService {
	return &UserService{
		User: 	data.NewUserTable(),
	}
}


func (a *UserService) Insert(User *model.User) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.User.Insert(o, User)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *UserService) Update(User *model.User) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.User.Update(o, User)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *UserService) Delete(User *model.User) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.User.Delete(o, User)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *UserService) GetByUserName(userName string) (*model.User, error) {

	o := GetOrmer()
	User := new(model.User)
	User.UserName = userName
	err := a.User.Get(o, User, "username")
	if err != nil {
		if err != orm.ErrNoRows {
			log.Infof("UserService.GetByUserName error,Username is: [%s]", userName)
			return nil, err
		} else {
			return nil, nil
		}
	}
	return User, nil
}

func (a *UserService) GetAllUser() []model.User {

	var allUser []model.User
	o := GetOrmer()
	o.QueryTable("User").All(&allUser)
	if len(allUser) == 0{
		log.Infof("UserService.GetAllUser, there is no user.")
		return nil
	}
	return allUser
}

func (a *UserService) GetOnline() ([]model.User, error) {

	var onlineUser []model.User
	o := GetOrmer()
	cmd := "select * from user where isonline='true'"
	_, err := o.Raw(cmd).QueryRows(&onlineUser)
	if err != nil{
		log.Errorf("Get the online user is err, the err is [%+v]", err)
		return nil, err
	}
	return onlineUser, nil
}