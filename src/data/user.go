package data

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"model"
)

type UserTable struct {
	lock *sync.Mutex
}

func NewUserTable() *UserTable {
	return &UserTable{
		lock: &sync.Mutex{},
	}
}

func (a *UserTable) Insert(o orm.Ormer, User *model.User) error {

	_, err := o.Insert(User)
	if err != nil {
		log.Errorf("UserTable.Insert the err is [%+v]", err)
		return err
	}
	return nil
}

func (a *UserTable) Update(o orm.Ormer, User *model.User) error {

	_, err := o.Update(User, "isonline")
	if err != nil {
		log.Errorf("UserTable.Update the err is [%+v]", err)
		return err
	}
	return nil
}

func (a *UserTable) Get(o orm.Ormer, User *model.User,colume string) error {

	err := o.Read(User, colume)
	if err != nil {
		log.Infof("UserTable.Get the err is [%+v]", err)
		return err
	}
	return nil
}


func (a *UserTable) Delete(o orm.Ormer, User *model.User) error {

	_, err := o.Delete(User)
	if err != nil {
		log.Errorf("UserTable.Delete the err is [%+v]", err)
		return err
	}
	return nil
}



