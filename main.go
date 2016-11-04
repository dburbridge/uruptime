package main

import (
    "strings"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    url := "https://api.uptimerobot.com/v2/getMonitors"
    fmt.Println("URL:>", url)

    r := strings.NewReader("api_key=xxxxx")
    req, err := http.NewRequest("POST", url, r)
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
