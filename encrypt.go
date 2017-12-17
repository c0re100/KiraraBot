package main

import (
    "io/ioutil"

    "crypto/aes"
    "crypto/cipher"
    "os"
)

func Encrypt() {
    num := random(-2147483647, 2147483647)
    b := (byte)(num & 127)
    num = int((num & -65281) | (65280 & (19 << 8)))
    //fmt.Println(num)
    //fmt.Println(writeInt32(num))
    wbytes := writeInt32(num)
    //h := fmt.Sprintf("%x", wbytes)
    //fmt.Println(h)
    //fmt.Println(readInt32(wbytes))

    key := []byte("7gyPmqc54dVNB3Te6pIpd2THj2y3hjOP")
    pass := []byte(RandomString(16))

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    Pack := []byte{0x84, 0xAE, 0x6D, 0x5F, 0x43, 0x6F, 0x6E, 0x66, 0x69, 0x72, 0x6D, 0x65, 0x64, 0x56, 0x65, 0x72, 0x00, 0xA6, 0x6D, 0x5F, 0x55, 0x55, 0x49, 0x44, 0xDA, 0x00, 0x24}
    Pack2 := []byte{0xAD, 0x6D, 0x5F, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6F, 0x6B, 0x65, 0x6E, 0xDA, 0x00, 0x24}
    Pack3 := []byte{0xA8, 0x6D, 0x5F, 0x4D, 0x79, 0x43, 0x6F, 0x64, 0x65, 0xAA}
    Pack4 := []byte{0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8}
    uuid := []byte(UUID)
    token := []byte(AccessToken)
    code := []byte(MyCode)

    Pack = append(Pack, uuid...)
    Pack = append(Pack, Pack2...)
    Pack = append(Pack, token...)
    Pack = append(Pack, Pack3...)
    Pack = append(Pack, code...)
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
    os.Mkdir(MyName, 0644)
    _ = ioutil.WriteFile(MyName+"/a.d", wbytes, 0644)
    _ = ioutil.WriteFile(MyName+"/a.d2", wbytes, 0644)
}
