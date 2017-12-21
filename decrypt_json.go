package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"

	"os"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/vmihailenco/msgpack"
)

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func ToJson(data []byte) {
	var AccData struct {
		ConfirmedVer int64  `msgpack:"m_ConfirmedVer"`
		UUID         string `msgpack:"m_UUID"`
		AccessToken  string `msgpack:"m_AccessToken"`
		MyCode       string `msgpack:"m_MyCode"`
	}

	err := msgpack.Unmarshal(data, &AccData)
	if err != nil {
		panic(err)
	}

	acJson := gabs.New()
	acJson.Set(AccData.ConfirmedVer, "m_ConfirmedVer")
	acJson.Set(AccData.UUID, "m_UUID")
	acJson.Set(AccData.AccessToken, "m_AccessToken")
	acJson.Set(AccData.MyCode, "m_MyCode")

	_ = ioutil.WriteFile("user.json", []byte(acJson.StringIndent("", "  ")), 0644)

	fmt.Println("User json saved")
}

func readInt32(data []byte) int32 {
	return int32(uint32(data[0]) + uint32(data[1])<<8 + uint32(data[2])<<16 + uint32(data[3])<<24)
}

func Decrypt() []byte {
	dat, _ := ioutil.ReadFile("a.d")

	sum := readInt32(dat[0:4])
	b := (byte)(sum & 127)
	count := int(dat[4] - b)

	if count == 16 {
		fmt.Println("Checksum correct!")
	} else {
		fmt.Println("Checksum failed!")
		os.Exit(0)
	}

	pass := dat[5:21]
	for i := 0; i < len(pass); i++ {
		var pass2 [16]byte
		pass2[i] = pass[i]
		pass[i] = pass2[i] - (byte)(96+i)
	}

	fmt.Println("a.d password:", string(pass))

	key := []byte("7gyPmqc54dVNB3Te6pIpd2THj2y3hjOP")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, pass)
	mode.CryptBlocks(dat[25:], dat[25:])

	data := PKCS7UnPadding(dat[25:])

	data = bytes.Replace(data, []byte{0xDA, 0x00}, []byte{0xD9}, -1)

	return data
}

func main() {
	ToJson(Decrypt())
	time.Sleep(24 * time.Hour)
}
