package service

import (
	"data"
	log "github.com/Sirupsen/logrus"
	"model"
	"github.com/astaxie/beego/orm"
)

type FileService struct {
	File		*data.FileTable
}


func NewFileService() *FileService {
	return &FileService{
		File: 	data.NewFileTable(),
	}
}


func (a *FileService) Insert(File *model.File) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.File.Insert(o, File)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *FileService) Update(File *model.File) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.File.Update(o, File)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *FileService) Delete(File *model.File) error {

	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		log.Error("Begin Transaction Fail")
		return err
	}
	err = a.File.Delete(o, File)
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
	}
	return nil
}

func (a *FileService) GetByFileName(FileName string) (*model.File, error) {

	o := GetOrmer()
	File := new(model.File)
	File.FileName = FileName
	err := a.File.Get(o, File, "Filename")
	if err != nil {
		if err != orm.ErrNoRows {
			log.Infof("FileService.GetByFileName error,Filename is: [%s]", FileName)
			return nil, err
		} else {
			return nil, nil
		}
	}
	return File, nil
}

func (a *FileService) GetAllFile() []model.File {

	var allFile []model.File
	o := GetOrmer()
	o.QueryTable("File").All(&allFile)
	if len(allFile) == 0{
		log.Infof("FileService.GetAllFile, there is no File.")
		return nil
	}
	return allFile
}
