package main

import (
	"bufio"
	"fmt"
	"go-encrypt/encrypt"
	"os"
	"strings"
	"time"
)

const (
	password string = ".password"
	contacts string = ".contacts"
)

type record struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birthDay"`
	Tel      string    `json:"tel"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
}

func (red *record) String() string {
	format := "ID:%d\nName:%s\nBirthDay:%s\nTel:%s\nAddr:%s\nDesc:%s\n"
	return fmt.Sprintf(format, red.ID, red.Name, red.BirthDay, red.Tel, red.Addr, red.Desc)
}

func readString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt)
	if line, err := reader.ReadString('\n'); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(line), nil
	}
}

////修改口令
//func modfiyPassword()bool{
//	fmt.Print("请输入初始化口令")
//	input,_ :=
//}
////口令鉴权
//func auth()bool{
//	if cxt,err := ioutil.ReadFile(password);os.IsNotExist(err){
//		return
//	}
//}

func main() {

	key := "qwertyuiopasdfgh"
	test_text := "The Curry who is from Gold Warriors is changing the development of NBA"
	fmt.Println("原始数据：", test_text)
	encry_result, err := encrypt.EncryptAES([]byte(test_text), []byte(key))
	if err != nil {
		fmt.Println("加密失败")
	}
	fmt.Println("加密后的结果", encry_result)
	decrypt_result, err := encrypt.DecryptAES(encry_result, []byte(key))
	if err != nil {
		fmt.Println("解密失败")
	}
	fmt.Println("解密后的结果：", string(decrypt_result))

}
