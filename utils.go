package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "time"
)

func SimVersion() {
    url := "https://krr-prd.star-api.com/api/app/version/get?platform=2&version=1.0.2"

    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
    req.Header.Add("x-star-requesthash", RandomString(64))
    req.Header.Add("x-unity-version", "5.5.4f1")
    req.Header.Add("content-type", "application/json; charset=UTF-8")
    req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; Nexus 5X Build/N2G48C)")
    req.Header.Add("Host", "krr-prd.star-api.com")

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    _, _ = ioutil.ReadAll(res.Body)
}

func random(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
}

func RandomString(strlen int) string {
    rand.Seed(time.Now().UTC().UnixNano())
    const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
    result := make([]byte, strlen)
    for i := 0; i < strlen; i++ {
        result[i] = chars[rand.Intn(len(chars))]
    }
    return string(result)
}

func SHA256(apiEndpoint string, json string) string {
    hash := sha256.New()
    hash.Write([]byte(apiEndpoint + " " + json + " " + "85af4a94ce7a280f69844743212a8b867206ab28946e1e30e6c1a10196609a11"))
    sha256hash := hash.Sum(nil)
    return hex.EncodeToString(sha256hash)
}

func SHA256withSid(sessionId string, apiEndpoint string, json string) string {
    hash := sha256.New()
    if json != "" {
        hash.Write([]byte(sessionId + " " + apiEndpoint + " " + json + " " + "85af4a94ce7a280f69844743212a8b867206ab28946e1e30e6c1a10196609a11"))
    } else {
        hash.Write([]byte(sessionId + " " + apiEndpoint + " " + "85af4a94ce7a280f69844743212a8b867206ab28946e1e30e6c1a10196609a11"))
    }
    sha256hash := hash.Sum(nil)
    return hex.EncodeToString(sha256hash)
}

func ListOfChar() string {
    var msg string
    for _, value := range wishlist {
        if value == 10002000 {
            msg += "悠乃,"
        } else if value == 11002000 {
            msg += "野野原柚子,"
        } else if value == 12002000 {
            msg += "丈槍由紀,"
        } else if value == 13002000 {
            msg += "一井透,"
        } else if value == 14002000 {
            msg += "九條可憐,"
        } else if value == 15002000 {
            msg += "涼風青葉,"
        } else if value == 16002000 {
            msg += "本田珠輝,"
        } else if value == 17002000 {
            msg += "千矢,"
        }
    }
    return msg[:len(msg)-1]
}

func wishlistAsk() {
    var wishId string
    for wishId == "" {
        fmt.Println("1=悠乃 | 2=野野原柚子 | 3=丈槍由紀 | 4=一井透")
        fmt.Println("5=九條可憐 | 6=涼風青葉 | 7=本田珠輝 | 8=千矢")
        fmt.Println("你想要邊幾隻5星做首抽? (eg: 12345678=要晒8隻5星)")
        fmt.Printf("Pattern: ")
        fmt.Scanln(&wishId)

        wishDrawn = len(wishId)
        if wishDrawn >= 1 && wishDrawn <= 10 {
            for i := 0; i < len(wishId); i++ {
                if string(wishId[i]) == "1" {
                    wishlist = append(wishlist, 10002000)
                } else if string(wishId[i]) == "2" {
                    wishlist = append(wishlist, 11002000)
                } else if string(wishId[i]) == "3" {
                    wishlist = append(wishlist, 12002000)
                } else if string(wishId[i]) == "4" {
                    wishlist = append(wishlist, 13002000)
                } else if string(wishId[i]) == "5" {
                    wishlist = append(wishlist, 14002000)
                } else if string(wishId[i]) == "6" {
                    wishlist = append(wishlist, 15002000)
                } else if string(wishId[i]) == "7" {
                    wishlist = append(wishlist, 16002000)
                } else if string(wishId[i]) == "8" {
                    wishlist = append(wishlist, 17002000)
                }
            }
        } else {
            wishId = ""
            log.Println("數值輸入錯誤...請重新輸入!!!")
            continue
        }
    }
    log.Println("依家開始抽呢幾隻5星", ListOfChar())
}

func idContains(drawnID float64) bool {
    for _, value := range wishlist {
        if value == drawnID {
            return true
        }
    }
    return false
}

func ShowDrawn() {
    log.Println("恭喜晒, 你已經抽到5星" + ListOfChar())
    log.Println("你仲唔快D去引繼!!!!!!!!!!!!!!!!!!!!!")
}
