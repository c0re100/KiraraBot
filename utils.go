package main

import (
    "crypto/sha256"
    "encoding/hex"
    "io/ioutil"
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
