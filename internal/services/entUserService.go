package services

import (
	"api.openfileplatform.com/internal/dao"
	"errors"
	"time"
)

type EntUserService struct {
	dao.EntUsers
}

func (m *EntUserService)BorrowFile(file_id string,borrow_term int8)(dao.EntFileLend,error){
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
		return msg,err
	}
	if file.EnterpriseID != m.EnterpriseID{
		err = errors.New("Enterprise is not matched")
		return msg,err
	}
	if file.Status == 1 {
		err = errors.New("This book has been borrowed")
		return msg,err
	}
	err = file.Update(map[string]interface{}{
		"Status":1,
	})
	err = msg.Add()
	if err != nil {
		return msg,err
	}
	return msg,err
}