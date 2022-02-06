package platform

type UserList struct {
	UserName string `json:"UserName" binding:""`
	UserID   string `json:"UserID" binding:"required"`
	Phone    string `json:"Phone" binding:""`
	Email    string `json:"Email" binding:""`
}

func (u UserList) GetUserName()  string{
	return u.UserName
}

func (u UserList) GetUserAccount()	string {
	return u.UserID
}

func (u *UserList) SetPhone(p string)error{
	u.Phone = p
	return nil
}

func (u *UserList) SetEmail(e string )error {
	u.Email = e
	return nil
}
