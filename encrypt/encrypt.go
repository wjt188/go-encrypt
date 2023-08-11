package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

const saltSize = 10

// 计算sha256的hash值
func hsha256(bytes []byte) string {
	hash := sha256.New()
	hash.Write(bytes)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// 口令加密
func EncryPassWord(bytes []byte) string {
	salt := make([]byte, saltSize)
	rand.Read(salt)
	return fmt.Sprintf("%s$$%s", hex.EncodeToString(salt), hsha256(append(salt, bytes...)))
}

// 口令认证
func VerfiyPassWord(password, input []byte) bool {
	nodes := strings.SplitN(string(password), "$$", 2)
	bytes, _ := hex.DecodeString(nodes[0])
	return nodes[1] == hsha256(append(bytes, input...))
}

// aes加密
func EncryptAES(cxt []byte, key []byte) ([]byte, error) {
	//创建加密算法块对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("创建加密对象失败")
		fmt.Println(err)
		return nil, err
	}
	blockSize := block.BlockSize()
	//按照PKCS7填充字节
	padding := blockSize - len(cxt)%blockSize
	cxt = append(cxt, bytes.Repeat([]byte{byte(padding)}, padding)...)
	//设置初始化向量
	crypted := make([]byte, len(cxt)+aes.BlockSize)
	iv := crypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("初始化向量失败")
		return nil, err
	}
	//获取CBC加密对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//执行加密逻辑
	blockMode.CryptBlocks(crypted[aes.BlockSize:], cxt)
	return crypted, nil
}

// aes解密
func DecryptAES(crypted, key []byte) ([]byte, error) {
	//创建加密算法块对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取CBC解密对象
	blockMode := cipher.NewCBCDecrypter(block, crypted[:aes.BlockSize])
	//执行解密逻辑
	cxt := make([]byte, len(crypted)-aes.BlockSize)
	blockMode.CryptBlocks(cxt, crypted[aes.BlockSize:])
	//按照PKCS7规则移除填充的字节
	return cxt[:len(cxt)-int(cxt[len(cxt)-1])], nil
}
