package main

import (
    "strings"
    "io/ioutil"
    "net/http"
)

func main() {
    http.HandleFunc("/", uptime)
    http.ListenAndServe(":8080", nil)
}

func uptime(w http.ResponseWriter, h *http.Request) {
    url := "https://api.uptimerobot.com/v2/getMonitors"

    r := strings.NewReader("api_key=xxxx")
    req, err := http.NewRequest("POST", url, r)
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    w.Write([]byte(body))
}
