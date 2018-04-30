package service

import (
	"data"
	log "github.com/Sirupsen/logrus"
	"model"
	"strings"
	"time"
)

type MessageService struct {
	Message		*data.MessageTable
}


func NewMessageService() *MessageService {
	return &MessageService{
		Message: 	data.NewMessageTable(),
	}
}


func (a *MessageService) Insert(Message *model.Message) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.Message.Insert(o, Message)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *MessageService) Update(Message *model.Message) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.Message.Update(o, Message)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *MessageService) Delete(Message *model.Message) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.Message.Delete(o, Message)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *MessageService) GetAllMessage() []model.Message {

	var allMessage []model.Message
	o := GetOrmer()
	o.QueryTable("Message").All(&allMessage)
	if len(allMessage) == 0{
		log.Infof("MessageService.GetAllMessage, there is no Message.")
		return nil
	}
	return allMessage
}

func (a *MessageService) StrToMsg(message string) *model.Message {

	m := strings.Split(message, ",")
	if len(m) != 2{
		log.Errorf("The message received is illegel.")
		return nil
	}
	userName := m[0]
	body := m[1]
	time := time.Now().Format("2006-01-02 15:04:05")
	msg := model.NewMessage(userName, body, time)
	return msg
}


