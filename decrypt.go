package main

import (
    "crypto/aes"
    "crypto/cipher"
    "io/ioutil"
    "strings"
)

func Decrypt() bool {
    dat, _ := ioutil.ReadFile(MyName + "/a.d")

    pass := dat[5:21]
    for i := 0; i < len(pass); i++ {
        var pass2 [16]byte
        pass2[i] = pass[i]
        pass[i] = pass2[i] - (byte)(96+i)
    }

    key := []byte("7gyPmqc54dVNB3Te6pIpd2THj2y3hjOP")
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    mode := cipher.NewCBCDecrypter(block, pass)
    mode.CryptBlocks(dat[25:], dat[25:])

    if strings.Contains(string(dat[25:]), UUID) {
        return true
    }

    return false
}
