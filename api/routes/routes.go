package routes

import (
        "encoding/json"
        "fmt"
        "gopkg.in/yaml.v2"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "path/filepath"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
        UserAgent          string `yaml:"user_agent"`
        ClientID           string `yaml:"client_id"`
        ClientSecret       string `yaml:"client_secret"`
        Username           string `yaml:"username"`
        Password           string `yaml:"password"`
        AlphaVantageApiKey string `yaml:"alpha_vantage_api_key"`
}

func Load_config() {
        absPath, _ := filepath.Abs("configs/config.yaml")
        yamlFile, error := ioutil.ReadFile(absPath) // just pass the file name
        if error != nil {
                fmt.Print(error)
        }

        t := T{}
        err := yaml.Unmarshal(yamlFile, &t)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- t:\n%v\n\n", t)

}

// TO DO: parse json
func getReddit() {
        type Name struct {
                Name string `json:"subreddit_url"`
        }
        var data []byte
        data, _ = ioutil.ReadFile("../../api/reddit/reddit.json")

        var str Name
        _ = json.Unmarshal(data, &str)

        if str.Name == "apple" {
                // Do Stuff
        }
}

func TimeSeriesIntraday(timestamp string, open string, high string, low string, close string, volume string, isCsv bool) {
        url := "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=MSFT&interval=5min&apikey=demo"
        response, err := http.Get(url)
        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        } else {
                defer response.Body.Close()
                contents, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        fmt.Printf("%s", err)
                        os.Exit(1)
                }
                fmt.Printf("%s\n", string(contents))
        }
}

func TimeSeriesDaily(date string, dailyOpen string, dailyHigh string, dailyLow string, dailyClose string, dailyVolume string, isCsv bool) {
        url := "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=MSFT&outputsize=full&apikey=demo"
        response, err := http.Get(url)
        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        } else {
                defer response.Body.Close()
                contents, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        fmt.Printf("%s", err)
                        os.Exit(1)
                }
                fmt.Printf("%s\n", string(contents))
        }
}


func TimeSeriesDailyAdjusted(date string, dailyOpen string, dailyHigh string, dailyLow string, dailyAdjustedClose string, splitEvents string, dailyVolume string, isCsv bool) {
        url := "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT&outputsize=full&apikey=demo"
        response, err := http.Get(url)
        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        } else {
                defer response.Body.Close()
                contents, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        fmt.Printf("%s", err)
                        os.Exit(1)
                }
                fmt.Printf("%s\n", string(contents))
        }
}



func TimeSeriesWeekly(lastWeeklyDate string, weeklyOpen string, weeklyHigh string, weeklyLow string, weeklyClose string, weeklyVolume string, isCsv bool) {
        url := "https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=MSFT&apikey=demo"
        response, err := http.Get(url)
        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        } else {
                defer response.Body.Close()
                contents, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        fmt.Printf("%s", err)
                        os.Exit(1)
                }
                fmt.Printf("%s\n", string(contents))
        }
}

func TimeSeriesWeeklyAdjusted(lastWeeklyDate string, weeklyOpen string, weeklyHigh string, weeklyLow string, weeklyClose string, weeklyVolume string, weeklyDividend string, isCsv bool) {
        url := "https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY_ADJUSTED&symbol=MSFT&apikey=demo"
        response, err := http.Get(url)
        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        } else {
                defer response.Body.Close()
                contents, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        fmt.Printf("%s", err)
                        os.Exit(1)
                }
                fmt.Printf("%s\n", string(contents))
        }
}


func GetHottestPosts(count string) {

        client := &http.Client{}

        url := "http://reddit.com/r/wallstreetbets/hot/.json?count=" + count

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                log.Fatalln(err)
        }

        req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

        response, err := client.Do(req)
        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        } else {
                defer response.Body.Close()
                contents, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        fmt.Printf("%s", err)
                        os.Exit(1)
                }
                fmt.Printf("%s\n", string(contents))
        }
}