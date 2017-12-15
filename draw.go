package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/Jeffail/gabs"
    "fmt"
)

func FirstDraw() {
    log.Println("正在進行首抽...")
    url := "https://krr-prd.star-api.com/api/player/gacha/draw"

    fDrawJson := gabs.New()
    fDrawJson.Set(1, "gachaId")
    fDrawJson.Set(3, "drawType")
    fDrawJson.Set(4, "stepCode")
    fDrawJson.Set(false, "reDraw")

    payload := strings.NewReader(fDrawJson.String())

    req, _ := http.NewRequest("POST", url, payload)

    hash := SHA256withSid(sessionId, "/api/player/gacha/draw", fDrawJson.String())

    req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
    req.Header.Add("x-star-requesthash", hash)
    req.Header.Add("x-unity-version", "5.5.4f1")
    req.Header.Add("x-star-ab", "3")
    req.Header.Add("x-star-session-id", sessionId)
    req.Header.Add("content-type", "application/json")
    req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; Nexus 5X Build/N2G48C)")
    req.Header.Add("Host", "krr-prd.star-api.com")

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    jsonParsed, _ := gabs.ParseJSON(body)
    Char := gabs.New()
    Char.Array("Gold")
    count := 0

    children, _ := jsonParsed.S("managedCharacters").Children()
    for _, child := range children {
        if child.Search("characterId").Data().(float64) == 10002000 {
            Char.ArrayAppend("悠乃", "Gold")
            count++
        } else if child.Search("characterId").Data().(float64) == 11002000 {
            Char.ArrayAppend("野野原柚子", "Gold")
            count++
        } else if child.Search("characterId").Data().(float64) == 12002000 {
            Char.ArrayAppend("丈槍由紀", "Gold")
            count++
        } else if child.Search("characterId").Data().(float64) == 13002000 {
            Char.ArrayAppend("一井透", "Gold")
            count++
        } else if child.Search("characterId").Data().(float64) == 14002000 {
            Char.ArrayAppend("九條可憐", "Gold")
            count++
        } else if child.Search("characterId").Data().(float64) == 15002000 {
            Char.ArrayAppend("涼風青葉", "Gold")
            count++
        } else if child.Search("characterId").Data().(float64) == 16002000 {
            Char.ArrayAppend("本田珠輝", "Gold")
            count++
        } else if child.Search("characterId").Data().(float64) == 17002000 {
            Char.ArrayAppend("千矢", "Gold")
            count++
        }
    }

    log.Println(Char)
    if count >= 2 {
        next := ""
        for next == "" {
            fmt.Println("雙5星 繼續刷定點? 1/0")
            fmt.Scanln(&next)
        }
        if next == "1" {
            reDraw()
        } else {
            return
        }
    }
    time.Sleep(time.Duration(random(17, 23)) * time.Second)
    reDraw()
}

func reDraw() {
    for {
        url := "https://krr-prd.star-api.com/api/player/gacha/draw"

        fDrawJson := gabs.New()
        fDrawJson.Set(1, "gachaId")
        fDrawJson.Set(3, "drawType")
        fDrawJson.Set(0, "stepCode")
        fDrawJson.Set(true, "reDraw")

        payload := strings.NewReader(fDrawJson.String())

        req, _ := http.NewRequest("POST", url, payload)

        hash := SHA256withSid(sessionId, "/api/player/gacha/draw", fDrawJson.String())

        req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
        req.Header.Add("x-star-requesthash", hash)
        req.Header.Add("x-unity-version", "5.5.4f1")
        req.Header.Add("x-star-ab", "3")
        req.Header.Add("x-star-session-id", sessionId)
        req.Header.Add("content-type", "application/json")
        req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; Nexus 5X Build/N2G48C)")
        req.Header.Add("Host", "krr-prd.star-api.com")

        res, _ := http.DefaultClient.Do(req)
        defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)

        jsonParsed, _ := gabs.ParseJSON(body)
        Char := gabs.New()
        Char.Array("Gold")
        count := 0

        children, _ := jsonParsed.S("managedCharacters").Children()
        for _, child := range children {
            if child.Search("characterId").Data().(float64) == 10002000 {
                Char.ArrayAppend("悠乃", "Gold")
                count++
            } else if child.Search("characterId").Data().(float64) == 11002000 {
                Char.ArrayAppend("野野原柚子", "Gold")
                count++
            } else if child.Search("characterId").Data().(float64) == 12002000 {
                Char.ArrayAppend("丈槍由紀", "Gold")
                count++
            } else if child.Search("characterId").Data().(float64) == 13002000 {
                Char.ArrayAppend("一井透", "Gold")
                count++
            } else if child.Search("characterId").Data().(float64) == 14002000 {
                Char.ArrayAppend("九條可憐", "Gold")
                count++
            } else if child.Search("characterId").Data().(float64) == 15002000 {
                Char.ArrayAppend("涼風青葉", "Gold")
                count++
            } else if child.Search("characterId").Data().(float64) == 16002000 {
                Char.ArrayAppend("本田珠輝", "Gold")
                count++
            } else if child.Search("characterId").Data().(float64) == 17002000 {
                Char.ArrayAppend("千矢", "Gold")
                count++
            }
        }

        log.Println(Char)
        if count >= 2 {
            next := ""
            for next == "" {
                fmt.Println("雙5星 繼續刷定點? 1/0")
                fmt.Scanln(&next)
            }
            if next == "1" {
                continue
            } else {
                return
            }
        }
        time.Sleep(time.Duration(random(17, 23)) * time.Second)
    }
}
