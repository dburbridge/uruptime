package main

import (
    "strings"
    "encoding/json"
    "net/http"
    "time"
    "strconv"
    "os"
)

func main() {
    http.HandleFunc("/", hello)

    http.HandleFunc("/uptime/", func(w http.ResponseWriter, u *http.Request) {
        UtDates := strings.SplitN(u.URL.Path, "/", 3)[2]
                if len(UtDates) !=17 {
                    w.Write([]byte("please include dates in the url path"))
                    return
                }
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
                for _, monitor := range data.Monitors {
                    w.Write([]byte(monitor.Friendly_name + " : " + monitor.Uptime + "\n"))
                }
    })
    
        http.HandleFunc("/outages/", func(w http.ResponseWriter, u *http.Request) {
        UtDates := strings.SplitN(u.URL.Path, "/", 3)[2]
                if len(UtDates) !=17 {
                    w.Write([]byte("please include dates in the url path"))
                    return
                }
                UtFrom := strings.SplitN(UtDates, "_", 2)[0]
                UtTo := strings.SplitN(UtDates, "_", 2)[1]
                t1, _ := time.Parse("20060102", UtFrom)
                t2, _ := time.Parse("20060102", UtTo)
                r1 := strconv.FormatInt(t1.Unix(), 10)
                r2 := strconv.FormatInt(t2.Unix(), 10)
                Utout := ("&logs_start_date="+ r1 + "&logs_end_date=" + r2)

        data, err := outages(Utout)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
                for _, monitor := range data.Monitors {
                    w.Write([]byte(monitor.Friendly_name + " : " + monitor.Logs + "\n"))
                }
    })

    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Enter /uptime/yyyymmdd_YYYYMMDD where yyyymmdd is the start date and YYYYMMDD is the end date in the address bar.\n For example /uptime/20161001_20161101 to get the uptime values for October 2016"))
}

func query(Utrange string) (Utdata, error){
    url := "https://api.uptimerobot.com/v2/getMonitors"
    r := strings.NewReader("api_key=" + os.Getenv("API_KEY") + "&custom_uptime_ranges=" + Utrange)
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

func outages(Utout string) (Utdata, error){
    url := "https://api.uptimerobot.com/v2/getMonitors"
    r := strings.NewReader("api_key=" + os.Getenv("API_KEY") + "&logs=1" + Utout)
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
      Logs []struct {
                Type int `json:"type"`
                Duration int `json:"duration"
            } `json:"logs"`
        } `json:"monitors"`
}

