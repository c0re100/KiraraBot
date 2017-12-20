package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/Jeffail/gabs"
)

func FirstDraw() {
    log.Println("正在進行首抽...")

    isFirst := true
    fDrawJson := gabs.New()
    count := 0

    for {
        url := "https://krr-prd.star-api.com/api/player/gacha/draw"

        if isFirst {
            fDrawJson.Set(1, "gachaId")
            fDrawJson.Set(3, "drawType")
            fDrawJson.Set(4, "stepCode")
            fDrawJson.Set(false, "reDraw")
            isFirst = false
        } else {
            fDrawJson.Set(1, "gachaId")
            fDrawJson.Set(3, "drawType")
            fDrawJson.Set(0, "stepCode")
            fDrawJson.Set(true, "reDraw")
        }

        payload := strings.NewReader(fDrawJson.String())

        req, _ := http.NewRequest("POST", url, payload)

        hash := SHA256withSid("/api/player/gacha/draw", fDrawJson.String())

        req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
        req.Header.Add("x-star-requesthash", hash)
        req.Header.Add("x-unity-version", "5.5.4f1")
        req.Header.Add("x-star-ab", "3")
        req.Header.Add("x-star-session-id", SessionId)
        req.Header.Add("content-type", "application/json; charset=UTF-8")
        req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; Nexus 5X Build/N2G48C)")

        res, _ := http.DefaultClient.Do(req)
        //defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)

        jsonParsed, _ := gabs.ParseJSON(body)
        Char := gabs.New()
        Char.Array("Gold")

        children, _ := jsonParsed.S("managedCharacters").Children()
        for _, child := range children {
            CharDrawn := child.Search("characterId").Data().(float64)
            if idContains(CharDrawn) {
                count++
            }
            if CharDrawn == 10002000 {
                Char.ArrayAppend("悠乃", "Gold")
            } else if CharDrawn == 11002000 {
                Char.ArrayAppend("野野原柚子", "Gold")
            } else if CharDrawn == 12002000 {
                Char.ArrayAppend("丈槍由紀", "Gold")
            } else if CharDrawn == 13002000 {
                Char.ArrayAppend("一井透", "Gold")
            } else if CharDrawn == 14002000 {
                Char.ArrayAppend("九條可憐", "Gold")
            } else if CharDrawn == 15002000 {
                Char.ArrayAppend("涼風青葉", "Gold")
            } else if CharDrawn == 16002000 {
                Char.ArrayAppend("本田珠輝", "Gold")
            } else if CharDrawn == 17002000 {
                Char.ArrayAppend("千矢", "Gold")
            }
        }

        log.Println(Char)
        if count >= wishCount {
            break
        }
        time.Sleep(time.Duration(random(17, 23)) * time.Second)
    }
}
