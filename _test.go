package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"

	"crypto/aes"
	"crypto/cipher"
	"time"
)

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func writeInt32(data int) []byte {
	buf := make([]byte, 4)
	buf[0] = byte(data)
	buf[1] = byte(data >> 8)
	buf[2] = byte(data >> 16)
	buf[3] = byte(data >> 24)
	return buf
}

func readInt32(data []byte) int32 {
	return int32(uint32(data[0]) + uint32(data[1])<<8 + uint32(data[2])<<16 + uint32(data[3])<<24)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Encrypt() {
	num := random(-2147483647, 2147483647)
	b := (byte)(num & 127)
	num = int((num & -65281) | (65280 & (19 << 8)))
	fmt.Println(num)
	fmt.Println(writeInt32(num))
	wbytes := writeInt32(num)
	h := fmt.Sprintf("%x", wbytes)
	fmt.Println(h)
	fmt.Println(readInt32(wbytes))

	key := []byte("7gyPmqc54dVNB3Te6pIpd2THj2y3hjOP")
	pass := []byte(RandomString(16))
	//fmt.Println(pass)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	Pack := []byte{0x84, 0xAE, 0x6D, 0x5F, 0x43, 0x6F, 0x6E, 0x66, 0x69, 0x72, 0x6D, 0x65, 0x64, 0x56, 0x65, 0x72, 0x00, 0xA6, 0x6D, 0x5F, 0x55, 0x55, 0x49, 0x44, 0xDA, 0x00, 0x24}
	Pack2 := []byte{0xAD, 0x6D, 0x5F, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6F, 0x6B, 0x65, 0x6E, 0xDA, 0x00, 0x24}
	Pack3 := []byte{0xA8, 0x6D, 0x5F, 0x4D, 0x79, 0x43, 0x6F, 0x64, 0x65, 0xAA}
	Pack4 := []byte{0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8}
	uuuid := []byte("59153c80-d295-4209-ad8f-840cfd0ebae4")
	acctoken := []byte("6130584a-0626-42a2-ba5c-78c45b178676")
	mycode := []byte("A2N94657PV")

	Pack = append(Pack, uuuid...)
	Pack = append(Pack, Pack2...)
	Pack = append(Pack, acctoken...)
	Pack = append(Pack, Pack3...)
	Pack = append(Pack, mycode...)
	Pack = append(Pack, Pack4...)

	mode := cipher.NewCBCEncrypter(block, pass)
	mode.CryptBlocks(Pack, Pack)

	for i := 0; i < len(pass); i++ {
		var pass2 [16]byte
		pass2[i] = pass[i]
		pass[i] = pass2[i] + (byte)(96+i)
	}

	sum := (byte)(len(pass) + int(b))
	wbytes = append(wbytes, sum)
	wbytes = append(wbytes, pass...)
	wbytes = append(wbytes, (byte)(len(Pack)))
	wbytes = append(wbytes, (byte)(0x00), (byte)(0x00), (byte)(0x00))
	wbytes = append(wbytes, Pack...)
	_ = ioutil.WriteFile("aaa.d", wbytes, 0644)
}

func main() {
	Encrypt()
}
