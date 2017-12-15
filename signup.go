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
    uname := ""
    for uname == "" {
        fmt.Printf("請輸入玩家名稱: ")
        fmt.Scanln(&uname)
    }

    url := "https://krr-prd.star-api.com/api/player/signup"

    reguuid = uuid.New().String()
    signJson := gabs.New()
    signJson.Set(reguuid, "uuid")
    signJson.Set(2, "platform")
    signJson.Set(uname, "name")
    signJson.Set(1, "stepCode")

    payload := strings.NewReader(signJson.String())
    req, _ := http.NewRequest("POST", url, payload)

    req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
    req.Header.Add("x-star-requesthash", RandomString(64))
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
        accessToken = jsonParsed.S("accessToken").Data().(string)
    }
}
