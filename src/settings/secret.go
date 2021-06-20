package settings

import "api.openfileplatform.com/utils/RSA"

func InitRSAKey()error{
	err := RSA.GenerateRSAKey(2048)
	return err
}