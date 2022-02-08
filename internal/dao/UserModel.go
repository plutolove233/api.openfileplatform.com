package dao

type UserInterface interface {
	Get()error
	GetUserID()string
	SetUserID(p string)
	GetIsAdmin()bool
	SetAccount(p string)
	GetPassword()string
}