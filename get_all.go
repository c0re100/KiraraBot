package main

import (
    "github.com/Jeffail/gabs"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
)

func questGet() {
    url := "https://krr-prd.star-api.com/api/player/quest/get_all"

    req, _ := http.NewRequest("GET", url, nil)

    hash := SHA256withSid("/api/player/quest/get_all", "")

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
        log.Println("讀取關卡列表...")
    }
}

func missionGet() {
    url := "https://krr-prd.star-api.com/api/player/mission/get_all"

    req, _ := http.NewRequest("GET", url, nil)

    hash := SHA256withSid("/api/player/mission/get_all", "")

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
        log.Println("讀取任務列表...")
    }
}

func presentGet() {
    url := "https://krr-prd.star-api.com/api/player/present/get_all"

    req, _ := http.NewRequest("GET", url, nil)

    hash := SHA256withSid("/api/player/present/get_all", "")

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
        log.Println("讀取禮物盒...")

        presents, _ := jsonParsed.S("presents").Children()
        for _, box := range presents {
            BoxID += strconv.FormatFloat(box.Search("managedPresentId").Data().(float64), 'f', 0, 64) + ","
        }
        BoxID = BoxID[:len(BoxID)-1]
    }
}

func questchapterGet() {
    url := "https://krr-prd.star-api.com/api/quest_chapter/get_all"

    req, _ := http.NewRequest("GET", url, nil)

    hash := SHA256withSid("/api/quest_chapter/get_all", "")

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
        log.Println("讀取關卡章節...")
    }
}

func Getall() {
    url := "https://krr-prd.star-api.com/api/player/get_all"

    req, _ := http.NewRequest("GET", url, nil)

    hash := SHA256withSid("/api/player/get_all", "")

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
        log.Println("讀取玩家資料...")
        playerName := jsonParsed.S("player", "name").Data().(string)
        MyCode = jsonParsed.S("player", "myCode").Data().(string)
        log.Println("==========玩家資料==========")
        log.Println("玩家名稱:", playerName)
        log.Println("玩家ID:", MyCode)
    }
}

func getPresent() {
    url := "https://krr-prd.star-api.com/api/player/present/get"

    managedPresentId := "?managedPresentId=" + BoxID + "&stepCode=2"
    req, _ := http.NewRequest("GET", url+managedPresentId, nil)

    hash := SHA256withSid("/api/player/present/get"+managedPresentId, "")

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
        log.Println("領取禮物盒...")
    }
}

func gachaGet() {
    url := "https://krr-prd.star-api.com/api/player/gacha/get_all?gachaIds=1"

    req, _ := http.NewRequest("GET", url, nil)

    hash := SHA256withSid("/api/player/gacha/get_all?gachaIds=1", "")

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
        log.Println("載入首抽資料...")
    }
}
