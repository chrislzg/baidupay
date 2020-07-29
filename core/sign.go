// copy from https://github.com/baidu-smart-app/auth-demo-backend

package core

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
)

var ErrBadPem = errors.New("decode pem failed")

type SignatureStruct interface {
	FieldMap() map[string]interface{}
	FieldForm() url.Values
}

// RSA 算法：
// 用公钥解密数据
// 用私钥加密数据

// Sign 签名
// 	1 用私钥对明文进行加密
// 	2 对密文进行sha1摘要
// https://dianshang.baidu.com/platform/doclist/index.html#!/doc/nuomiplus_2_base/anchor/sign.md
func Sign(v SignatureStruct, rsaPrivateKey []byte) (string, error) {
	fieldMap := v.FieldMap()
	plainText := BuildSignatureString(fieldMap)
	plainText, err := url.QueryUnescape(plainText)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(rsaPrivateKey)
	if block == nil {
		return "", ErrBadPem
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	hash := crypto.Hash.New(crypto.SHA1)
	hash.Write([]byte(plainText))

	sign, err := privKey.Sign(rand.Reader, hash.Sum(nil), crypto.SHA1)

	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

// CheckSign 验签：对采用sha1算法进行签名后转base64格式的数据进行验签
func CheckSign(originalData, signData string, publicKey []byte) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	// public, _ := base64.StdEncoding.DecodeString(pubKey)

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return ErrBadPem
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hash.Sum(nil), sign)
}

// 构建需要签名的字符串
func BuildSignatureString(fieldMap map[string]interface{}) string {
	values := url.Values{}
	for key := range fieldMap {
		values.Add(key, fmt.Sprint(fieldMap[key]))
	}
	return values.Encode()
}
