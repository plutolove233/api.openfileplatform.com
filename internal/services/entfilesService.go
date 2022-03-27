/*
@Coding : utf-8
@Time : 2022/2/14 14:21
@Author : 刘浩宇
@Software: GoLand
*/
package services

import (
	"api.openfileplatform.com/internal/dao"
	"time"
)

type EnterpriseFilesService struct {
	dao.EntFiles
}

type entFilesParser struct {
	FileId       string    `json:"FileId"`
	CategoryId   string    `json:"CategoryId"`
	CategoryName string	   `json:"CategoryName"`
	ProjectId    string    `json:"ProjectId"`
	ProjectName  string    `json:"ProjectName"`
	EnterpriseId string    `json:"EnterpriseId"`
	FileName     string    `json:"FileName"`
	FileUrl      string    `json:"FileUrl"`
	FileTypeId   string    `json:"FileTypeId"`
	Status       int8      `json:"Status"`
	UserId       string    `json:"UserId"`
	UserName	 string	   `json:"UserName"`
	FileCabinet  string    `json:"FileCabinet"`
	UpdateTime   time.Time `json:"UpdateTime"`
}

func (m *EnterpriseFilesService)GetFileInformation()([]entFilesParser,error){
	fileInfo, err := m.GetAll()
	if err != nil {
		return nil, err
	}
	data := []entFilesParser{}

	for _,item := range fileInfo{
		a := entFilesParser{
			FileId: 		item.FileID,
			CategoryId:		item.CategoryID,
			ProjectId:		item.ProjectID,
			EnterpriseId:	item.EnterpriseID,
			FileName:     	item.FileName,
			FileUrl:      	item.FileURL,
			FileTypeId:   	item.FileTypeID,
			Status:      	item.Status,
			UserId:       	item.UserID,
			FileCabinet:	item.FileCabinet,
			UpdateTime:		item.UpdateTime,
		}
		project := dao.EntProject{}
		project.ProjectID = item.ProjectID
		if err = project.Get(); err != nil{
			return nil, err
		}
		category := dao.EntFileCategory{}
		category.CategoryID = item.CategoryID
		if err = category.Get(); err != nil{
			return nil,err
		}
		entuser := dao.EntUsers{}
		entuser.UserID = item.UserID
		if err = entuser.Get(); err != nil {
			return nil,err
		}
		a.ProjectName  = project.ProjectName
		a.CategoryName = category.CategoryName
		a.UserName = entuser.UserName
		data = append(data, a)
	}
	return data,nil
}