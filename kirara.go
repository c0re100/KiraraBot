package main

import (
    "fmt"
    "time"
)

var (
    UUID        string
    AccessToken string
    SessionId   string
    MyCode      string
    MyName      string
    BoxID       string
    Wishlist    = make([]float64, 0, 8)
    wishCount   int
)

func main() {
    fmt.Println("====================Kirara Fantasia 自動首抽模擬器 v3.0====================")
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
    time.Sleep(1 * time.Second)
    wishlistAsk()
    gachaGet()
    time.Sleep(3 * time.Second)
    FirstDraw()
    //time.Sleep(2 * time.Second)
    //CharSave()
    time.Sleep(3 * time.Second)
    SaveFile()
    time.Sleep(500 * time.Millisecond)
    ShowDrawn()
    time.Sleep(500 * time.Millisecond)
    fmt.Println("====================Kirara Fantasia 自動首抽模擬器 v3.0====================")
    time.Sleep(24 * time.Hour)
}
