package services

type UserInterface interface {
	Get()error
	Add()error
	GetUserID()string
	GetIsAdmin()bool
	GetPassword()string
	SetAccount(p string)
	SetUserID(p string)
	SetUserName(p string)
	SetPassword(p string)
	SetPhone(p string)
	SetEmail(p string)
}