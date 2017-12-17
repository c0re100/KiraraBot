package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"

    "fmt"
    "time"

    "github.com/Jeffail/gabs"
)

func CharSave() {
    url := "https://krr-prd.star-api.com/api/player/tutorial/party/set"

    charJson := gabs.New()
    charJson.Set(5, "stepCode")

    req, _ := http.NewRequest("POST", url, strings.NewReader(charJson.String()))

    hash := SHA256withSid(SessionId, "/api/player/tutorial/party/set", charJson.String())

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
        log.Println("成功儲存首抽角色...")
    }
}

func SaveFile() {
    Encrypt()
    if Decrypt() {
        log.Println("成功儲存帳戶存檔...存檔位置: /" + MyName + "/")
        log.Println("如要使用，請複製a.d及a.d2到以下位置")
        log.Println("位置: /Android/data/com.aniplex.kirarafantasia/files")
    } else {
        log.Println("儲存失敗...")
        log.Println("UUID:", UUID)
        log.Println("AccessToken:", AccessToken)
        log.Println("MyCode:", MyCode)
        log.Println("請保留以上")
    }
}

func SaveData() {
    log.Println("正在進行帳戶引繼...")
    time.Sleep(1 * time.Second)
    upass := ""
    for upass == "" {
        fmt.Printf("請輸入引繼密碼: ")
        fmt.Scanln(&upass)
    }

    url := "https://krr-prd.star-api.com/api/player/move/get"

    moveJson := gabs.New()
    moveJson.Set(upass, "password")

    req, _ := http.NewRequest("POST", url, strings.NewReader(moveJson.String()))

    hash := SHA256withSid(SessionId, "/api/player/move/get", moveJson.String())

    req.Header.Add("unity-user-agent", "app/0.0.0; Android OS 7.1.2 / API-25 N2G48C/4104010; LGE Nexus 5X")
    req.Header.Add("x-star-requesthash", hash)
    req.Header.Add("x-unity-version", "5.5.4f1")
    req.Header.Add("X-STAR-AB", "3")
    req.Header.Add("X-STAR-SESSION-ID", SessionId)
    req.Header.Add("content-type", "application/json; charset=UTF-8")
    req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 7.1.2; Nexus 5X Build/N2G48C)")

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    jsonParsed, _ := gabs.ParseJSON(body)
    if jsonParsed.S("resultCode").Data().(float64) == 0 {
        log.Println("成功設定引繼...")
        moveCode := jsonParsed.S("moveCode").Data().(string)
        moveDeadline := jsonParsed.S("moveDeadline").Data().(string)
        log.Println("引繼ID:", moveCode)
        log.Println("引繼密碼:", upass)
        log.Println("引繼限期:", moveDeadline)
    }
}
