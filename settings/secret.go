package settings

import "DocumentSystem/utils/RSA"

func InitRSAKey()error{
	err := RSA.GenerateRSAKey(2048)
	return err
}