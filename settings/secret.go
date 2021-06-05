package settings

import "DocumentSystem/untils/RSA"

func InitRSAKey()error{
	err := RSA.GenerateRSAKey(2048)
	return err
}