package routes

import (
    "fmt"
    "net/http"
	"io/ioutil"
    "encoding/json"
    "gopkg.in/yaml.v2"
    "os"
    "log"
    "path/filepath"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
    UserAgent string `yaml:"user_agent"`
    ClientID string `yaml:"client_id"`
    ClientSecret string `yaml:"client_secret"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
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


func TimeSeriesIntraday(timestamp string, open string, high string, low string, close string, volume string) {
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


func GetHottestPosts(count string) {
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
