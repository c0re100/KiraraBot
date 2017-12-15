package main

import (
    "fmt"
    "time"
)

var (
    reguuid     string
    accessToken string
    sessionId   string
    BoxID       string
)

func main() {
    fmt.Println("===========Kirara Fantasia 自動首抽模擬器 v1.0===========")
    SimVersion()
    SignUp()
    Login()
    Getall()
    questGet()
    missionGet()
    questchapterGet()
    time.Sleep(3 * time.Second)
    presentGet()
    time.Sleep(2 * time.Second)
    getPresent()
    time.Sleep(10 * time.Second)
    adv()
    time.Sleep(3 * time.Second)
    gachaGet()
    time.Sleep(3 * time.Second)
    FirstDraw()
    time.Sleep(2 * time.Second)
    CharSave()
    time.Sleep(3 * time.Second)
    SaveData()
    fmt.Println("===========Kirara Fantasia 自動首抽模擬器 v1.0===========")
}
