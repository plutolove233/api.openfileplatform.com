// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 20:29
// @Software: GoLand

package services

import (
	"api.openfileplatform.com/internal/dao"
)

type PlatUsersService struct {
	dao.PlatUsers
}

func (*PlatUsersService)SetEntUserAdmin(entprise_id string, user_id string) (msg string, err error){
	var entUsers dao.EntUsers
	var platEnterprise dao.PlatEnterprise
	entUsers.UserID = user_id
	entUsers.EnterpriseID = entprise_id
	platEnterprise.EnterpriseID = entprise_id

	err = entUsers.Get()
	if err != nil {
		msg = "企业目标管理员用户信息获取失败"
		return msg, err
	}

	err = platEnterprise.Get()
	if err != nil {
		msg = "获取平台企业表失败"
		return msg, err
	}

	err = entUsers.Update(map[string]interface{}{
		"IsAdmin":true,
	})
	if err != nil {
		msg = "设置企业管理员失败"
		return msg, err
	}

	err = platEnterprise.Update(map[string]interface{}{
		"AdminID":user_id,
	})
	if err != nil {
		msg = "企业平台管理员信息更新失败"
		return msg, err
	}

	return "企业管理员设置成功",nil
}

func (*PlatUsersService) RemoveEntUserAdmin(enterpriseID string,userID string) (string, error) {
	entUser := dao.EntUsers{}
	entUser.UserID = userID
	entUser.IsAdmin = true
	entUser.EnterpriseID = enterpriseID
	if err := entUser.Get(); err != nil {
		return "无法找到该用户",err
	}
	platEnterprise := dao.PlatEnterprise{}
	platEnterprise.EnterpriseID = enterpriseID
	platEnterprise.AdminID = userID
	if err := platEnterprise.Get(); err != nil {
		return "无法找到相关企业",err
	}
	if err:=entUser.Update(map[string]interface{}{
		"IsAdmin":false,
	});err!=nil {
		return "移除管理员身份失败",err
	}
	if err:=platEnterprise.Update(map[string]interface{}{
		"AdminID":"",
	});err!=nil {
		return "移除管理员身份失败",err
	}
	return "移除管理员身份成功",nil
}