package enterprise

type UserModel struct {
	UserID 		string 	`json:"UserID"`
	Account 	string 	`json:"Account"`
	UserName 	string 	`json:"UserName"`
	FacePicUrl	string 	`json:"FacePicUrl"`
	IsEntUser	bool 	`json:"IsEntUser"`
	IsAdmin 	int 	`json:"IsAdmin"`
	Token 		string 	`json:"Token"`
}

func (u *UserModel)GetUserID() string{
	return u.UserID
}

func (u *UserModel)SetUserID(id string) error {
	u.UserID = id
	return nil
}

func (u *UserModel)VerifyAdminRole() bool{
	if u.IsEntUser {
		return true
	}
	if u.IsAdmin==1 {
		return true
	}
	return false
}