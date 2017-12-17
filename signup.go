package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"

    "github.com/Jeffail/gabs"
    "github.com/google/uuid"
)

func SignUp() {
    for MyName == "" {
        fmt.Printf("請輸入玩家名稱: ")
        fmt.Scanln(&MyName)
    }

    url := "https://krr-prd.star-api.com/api/player/signup"

    UUID = uuid.New().String()
    signJson := gabs.New()
    signJson.Set(UUID, "uuid")
    signJson.Set(2, "platform")
    signJson.Set(MyName, "name")
    signJson.Set(1, "stepCode")

    payload := strings.NewReader(signJson.String())
    req, _ := http.NewRequest("POST", url, payload)

    hash := SHA256withSid(SessionId, "/api/player/signup", signJson.String())

    req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
    req.Header.Add("x-star-requesthash", hash)
    req.Header.Add("x-unity-version", "5.5.4f1")
    req.Header.Add("content-type", "application/json; charset=UTF-8")
    req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; Nexus 5X Build/N2G48C)")
    req.Header.Add("Host", "krr-prd.star-api.com")

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    jsonParsed, _ := gabs.ParseJSON(body)
    if jsonParsed.S("resultCode").Data().(float64) == 0 {
        log.Println("註冊成功")
        AccessToken = jsonParsed.S("accessToken").Data().(string)
    }
}
