package data

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"model"
)

type FileTable struct {
	lock *sync.Mutex
}

func NewFileTable() *FileTable {
	return &FileTable{
		lock: &sync.Mutex{},
	}
}

func (a *FileTable) Insert(o orm.Ormer, File *model.File) error {

	_, err := o.Insert(File)
	if err != nil {
		log.Errorf("FileTable.Insert the err is [%+v]", err)
		return err
	}
	return nil
}

func (a *FileTable) Update(o orm.Ormer, File *model.File) error {

	_, err := o.Insert(File)
	if err != nil {
		log.Errorf("FileTable.Update the err is [%+v]", err)
		return err
	}
	return nil
}

func (a *FileTable) Get(o orm.Ormer, File *model.File,colume string) error {

	err := o.Read(File,colume)
	if err != nil {
		log.Infof("FileTable.Get the err is [%+v]", err)
		return err
	}
	return nil
}


func (a *FileTable) Delete(o orm.Ormer, File *model.File) error {

	_, err := o.Delete(File)
	if err != nil {
		log.Errorf("FileTable.Delete the err is [%+v]", err)
		return err
	}
	return nil
}


