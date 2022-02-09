// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 23:36
// @Software: GoLand

package ginModels

type UserModel struct {
	UserID     string `json:"user_id"`
	Account    string `json:"account"`
	IsPlatUser bool   `json:"is_plat_user"`
	IsAdmin    bool   `json:"is_admin"`
}

func (u UserModel) GetUserID() string {
	return u.UserID
}

func (u *UserModel) SetUserID(userID string) error {
	u.UserID = userID
	return nil
}

func (u UserModel) VerifyAdminRole() bool {
	if u.IsPlatUser {
		return true
	}

	if u.IsAdmin {
		return true
	}
	return false
}
