package services

import (
	"api.openfileplatform.com/internal/dao"
	"api.openfileplatform.com/internal/globals/snowflake"
	"os"
	"time"
)

type PlatEnterpriseService struct {
	dao.PlatEnterprise
}

func (m *PlatEnterpriseService)CreatePartition() error{
	cate := dao.EntFileCategory{}
	cate.CategoryName = m.EnterpriseName
	cate.CategoryID = snowflake.GetSnowflakeID()
	cate.CategoryParentID = ""
	cate.ProjectID = ""
	cate.EnterpriseID = m.EnterpriseID
	cate.CreatTime = time.Now()
	err := cate.Add()
	if err != nil {
		return err
	}
	err = os.Mkdir("./save/"+cate.CategoryName,os.ModePerm)
	return err
}