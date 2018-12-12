package routes

import (
    "fmt"
    "net/http"
	"io/ioutil"
    "encoding/json"
    // "gopkg.in/yaml.v2"
    "os"
)


var alphavantage_api_key = ""

// TO DO: parse json
func get_reddit() {
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


func TIME_SERIES_INTRADAY(timestamp string, open string, high string, low string, close string, volume string) {
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


func Get_hottest_posts(count string) {
	url := "http://reddit.com/r/wallstreetbets/hot/.json?count=" + count
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
