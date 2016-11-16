package main

import (
    "strings"
    "encoding/json"
    "net/http"
    "time"
    "strconv"
)

func main() {
    http.HandleFunc("/hello", hello)

    http.HandleFunc("/uptime/", func(w http.ResponseWriter, u *http.Request) {
        UtDates := strings.SplitN(u.URL.Path, "/", 3)[2]
                UtFrom := strings.SplitN(UtDates, "_", 2)[0]
                UtTo := strings.SplitN(UtDates, "_", 2)[1]
                t1, _ := time.Parse("20060102", UtFrom)
                t2, _ := time.Parse("20060102", UtTo)
                r1 := strconv.FormatInt(t1.Unix(), 10)
                r2 := strconv.FormatInt(t2.Unix(), 10)
                Utrange := (r1 + "_" + r2)

        data, err := query(Utrange)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
                w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(data)
        w.Write([]byte(Utrange))
    })

    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

func query(Utrange string) (Utdata, error){
    url := "https://api.uptimerobot.com/v2/getMonitors"
    r := strings.NewReader("api_key=xxxx5&custom_uptime_ranges=" + Utrange)
    req, err := http.NewRequest("POST", url, r)
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
       return Utdata{}, err
    }
    defer resp.Body.Close()

    var d Utdata

    if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
        return Utdata{}, err
    }

    return d, nil
}


type Utdata struct {
  Monitors []struct {
    Friendly_name string `json:"friendly_name"`
    Uptime string `json:"custom_uptime_ranges"`
        } `json:"monitors"`
}

