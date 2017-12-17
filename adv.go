package main

import (
    "io/ioutil"
    "log"
    "net/http"

    "github.com/Jeffail/gabs"
    "strings"
)

func adv() {
    url := "https://krr-prd.star-api.com/api/player/adv/add"

    advJson := gabs.New()
    advJson.Set("1000000", "advId")
    advJson.Set(3, "stepCode")

    req, _ := http.NewRequest("POST", url, strings.NewReader(advJson.String()))

    hash := SHA256withSid(SessionId, "/api/player/adv/add", advJson.String())

    req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
    req.Header.Add("x-star-requesthash", hash)
    req.Header.Add("x-unity-version", "5.5.4f1")
    req.Header.Add("X-STAR-AB", "3")
    req.Header.Add("X-STAR-SESSION-ID", SessionId)
    req.Header.Add("content-type", "application/json; charset=UTF-8")
    req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; Nexus 5X Build/N2G48C)")
    req.Header.Add("Host", "krr-prd.star-api.com")

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    jsonParsed, _ := gabs.ParseJSON(body)
    if jsonParsed.S("resultCode").Data().(float64) == 0 {
        log.Println("載入首抽畫面...")
    }
}
