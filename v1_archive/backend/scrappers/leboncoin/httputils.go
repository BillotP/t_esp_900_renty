package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Proxy Setup
// var (
// 	proxyHost     = goscrappy.MustGetSecret("proxy_url")
// 	proxyUsername = goscrappy.MustGetSecret("proxy_user")
// 	proxyPassword = goscrappy.MustGetSecret("proxy_password")
// )

// proxyURL is the scrapping proxy URL
// var proxyURL = &url.URL{
// 	Scheme: "http",
// 	User: url.UserPassword(
// 		proxyUsername,
// 		proxyPassword,
// 	),
// 	Host: proxyHost,
// }

// client is the GetphoneNumber http client
var client = http.Client{
	Timeout:   30 * time.Second,
	Jar:       http.DefaultClient.Jar,
	Transport: &http.Transport{
		// Proxy: http.ProxyURL(proxyURL),
	},
}

// userAgents is a list of the most popular user agent
var userAgents = [...]string{
	"Mozilla/5.0 (X11; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:53.0) Gecko/20100101 Firefox/53.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0;  Trident/5.0)",
	"Mozilla/5.0 (iPad; CPU OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H321 Safari/600.1.4",
}

// getrandomUA return a random user agent string from the below list
func getrandomUA() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(userAgents) - 1)
	return userAgents[n]
}

// getrandomSleepDuration return a random int between min and max
func getrandomSleepDuration(min, max int) (dur int) {
	rand.Seed(time.Now().UnixNano())
	if dur = rand.Intn(max); dur < min {
		return getrandomSleepDuration(min, max)
	}
	return dur
}

// GetMyIP return current external ip
func GetMyIP() (ip string, err error) {
	const MyIPAPIURL = "https://api.ipify.org?format=json"
	// MyIP is the response from ipify API
	type MyIP struct {
		IP string `json:"ip"`
	}
	resp, err := client.Get(MyIPAPIURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data MyIP
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	return data.IP, nil
}

// SetHeaders set the header before a leboncoin api request
func SetHeaders(req *http.Request, annonceID, annonceType string) {
	ua := getrandomUA()
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Host", "api.leboncoin.fr")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "fr-FR;fr;q=0.5")
	req.Header.Set("Origin", "https://www.leboncoin.fr")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", fmt.Sprintf("https://www.leboncoin.fr/%s/%s.htm/", annonceType, annonceID))
}
