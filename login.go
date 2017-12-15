package main

import (
    "log"
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/Jeffail/gabs"
)

func Login() {
    url := "https://krr-prd.star-api.com/api/player/login"

    loginJson := gabs.New()
    loginJson.Set(reguuid, "uuid")
    loginJson.Set(accessToken, "accessToken")
    loginJson.Set(2, "platform")
    loginJson.Set("1.0.2", "appVersion")

    payload := strings.NewReader(loginJson.String())

    req, _ := http.NewRequest("POST", url, payload)

    hash := SHA256("/api/player/login", loginJson.String())

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
        sessionId = jsonParsed.S("sessionId").Data().(string)
        log.Println("登入成功")
    }
}
