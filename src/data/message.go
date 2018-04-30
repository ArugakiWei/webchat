package data

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"model"
)

type MessageTable struct {
	lock *sync.Mutex
}

func NewMessageTable() *MessageTable {
	return &MessageTable{
		lock: &sync.Mutex{},
	}
}

func (a *MessageTable) Insert(o orm.Ormer, Message *model.Message) error {

	_, err := o.Insert(Message)
	if err != nil {
		log.Errorf("MessageTable.Insert the err is [%+v]", err)
		return err
	}
	return nil
}

func (a *MessageTable) Update(o orm.Ormer, Message *model.Message) error {

	_, err := o.Insert(Message)
	if err != nil {
		log.Errorf("MessageTable.Update the err is [%+v]", err)
		return err
	}
	return nil
}

func (a *MessageTable) Get(o orm.Ormer, Message *model.Message,colume string) error {

	err := o.Read(Message,colume)
	if err != nil {
		log.Infof("MessageTable.Get the err is [%+v]", err)
		return err
	}
	return nil
}


func (a *MessageTable) Delete(o orm.Ormer, Message *model.Message) error {

	_, err := o.Delete(Message)
	if err != nil {
		log.Errorf("MessageTable.Delete the err is [%+v]", err)
		return err
	}
	return nil
}


