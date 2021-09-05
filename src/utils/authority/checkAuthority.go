package authority

import (
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
)

func CheckAuthority(userID int64,match int64) bool{
	var userRole []models.EntRole
	err :=dao.DB.Model(&models.EntRole{}).Where("UserID = ?",userID).Find(&userRole).Error
	if err != nil {
		return false
	}

	var userAuthority []models.EntRoleAuthority
	for _,aRole:=range userRole{
		err = dao.DB.Model(&models.EntRoleAuthority{}).Where("RoleID = ?",aRole.RoleID).Find(&userAuthority).Error
		if err != nil {
			return false
		}
		for _,oneAuthority := range userAuthority{
			if oneAuthority.AuthorityID == match {
				return true
			}
		}
	}
	return false
}