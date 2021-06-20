package RSA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateRSAKey(bits int)error{
	//私钥生成
	privateKey, err:= rsa.GenerateKey(rand.Reader,bits)
	if err!=nil{
		panic(err)
	}
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateFile,err:= os.Create("utils/RSA/private.pem")
	if err!=nil{
		panic(err)
	}
	defer privateFile.Close()

	privateBlock := pem.Block{Type:"RSA Private Key",Bytes:X509PrivateKey}

	err = pem.Encode(privateFile,&privateBlock)
	if err!=nil{
		return err
	}
	//公钥生成
	publicKey := privateKey.PublicKey
	X509PublicKey,err := x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		panic(err)
	}

	publicFile,err:= os.Create("utils/RSA/public.pem")
	if err!=nil{
		panic(err)
	}
	defer publicFile.Close()

	publicBlock := pem.Block{Type:"RSA Public Key",Bytes:X509PublicKey}

	err = pem.Encode(publicFile,&publicBlock)
	if err!=nil{
		return err
	}
	return nil
}

func Encrypt(plaintext []byte) ([]byte,error){//加密
	file,err := os.Open("utils/RSA/public.pem")
	if err!=nil {
		panic(err)
	}
	defer file.Close()

	info,_ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)
	block,_ := pem.Decode(buf)
	publicKeyInterface,err:= x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		return nil,err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	cipherText,err:=rsa.EncryptPKCS1v15(rand.Reader,publicKey,plaintext)
	if err!=nil{
		return nil,err
	}
	return cipherText,nil
}

func Decrypt(cipherText []byte)([]byte,error){//解密
	file,err:=os.Open("utils/RSA/private.pem")
	if err!=nil{
		return nil,err
	}
	defer file.Close()

	info,_ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)

	block,_ := pem.Decode(buf)

	privateKey,err:= x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		return nil,err
	}

	plainText,err:=rsa.DecryptPKCS1v15(rand.Reader,privateKey,cipherText)
	return plainText,nil
}