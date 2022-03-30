package services

import (
	"api.openfileplatform.com/internal/dao"
	"errors"
	"time"
)

type EntUserService struct {
	dao.EntUsers
}

type borrowFileParser struct {
	FileId       	string    `json:"FileId"`
	FileName		string	  `json:"FileName"`
	FileURL	     	string    `json:"FileURL"`
	EnterpriseId 	string    `json:"EnterpriseId"`
	BorrowerId   	string    `json:"BorrowerId"`
	BorrowerName	string	  `json:"BorrowerName"`
	BorrowTime   	time.Time `json:"BorrowTime"`
	BorrowTerm   	int8      `json:"BorrowTerm"`
	ReturnTime   	time.Time `json:"ReturnTime"`
}

func (m *EntUserService)BorrowFile(file_id string,borrow_term int8)(borrowFileParser,error){
	msg := dao.EntFileLend{}
	msg.FileID = file_id
	msg.BorrowTime = time.Now()
	msg.BorrowTerm = borrow_term
	msg.BorrowerID = m.UserID
	msg.ReturnTime = time.Now().AddDate(0,int(borrow_term),0)
	msg.CreatTime = time.Now()
	msg.EnterpriseID = m.EnterpriseID
	file := dao.EntFiles{}
	file.FileID = file_id
	err := file.Get()
	if err != nil {
		return borrowFileParser{},err
	}
	data := borrowFileParser{
		FileId: file_id,
		FileName: file.FileName,
		FileURL: file.FileURL,
		EnterpriseId: m.EnterpriseID,
		BorrowTime: msg.BorrowTime,
		BorrowTerm: msg.BorrowTerm,
		BorrowerId: m.UserID,
		BorrowerName: m.UserName,
		ReturnTime: msg.ReturnTime,
	}
	if file.EnterpriseID != m.EnterpriseID{
		err = errors.New("Enterprise is not matched")
		return borrowFileParser{},err
	}
	if file.Status == 1 {
		err = errors.New("This book has been borrowed")
		return borrowFileParser{},err
	}
	err = file.Update(map[string]interface{}{
		"Status":1,
	})
	err = msg.Add()
	if err != nil {
		return borrowFileParser{},err
	}
	return data,err
}

func (m *EntUserService)ReturnFile(file_id string) error{
	fileLend := dao.EntFileLend{}
	fileLend.FileID = file_id
	fileLend.BorrowerID = m.UserID
	fileLend.EnterpriseID = m.EnterpriseID
	err := fileLend.Update(map[string]interface{}{
		"ReturnTime": time.Now(),
		"IsDeleted":1,
	})
	if err != nil {
		return err
	}

	file := dao.EntFiles{}
	file.FileID = file_id
	file.EnterpriseID = m.EnterpriseID
	err = file.Update(map[string]interface{}{
		"Status":0,
	})
	return err
}