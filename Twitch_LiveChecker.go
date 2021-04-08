package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type data struct {
	Type        string `json:"type"`
	UserName    string `json:"user_name"`
	ViewerCount int    `json:"viewer_count"`
}
type chanel struct {
	Data []data `json:"data"`
}

var linksToRemove []string
var clientId = "change_this_to_your_clientId"
var authToken = "change_this_to_your_authorization_token"

func main() {

	allLinks, err := ioutil.ReadFile("input.dat")
	if err != nil {
		panic(err)
	}

	allLinksString := string(allLinks)
	allLinksSlice := strings.Split(allLinksString, ",")

	var links []string
	var linksCopy []string

	//Making api link for received Streamer
	for i := 0; i < len(allLinksSlice); i++ {
		makeLink := "https://api.twitch.tv/helix/streams?user_login=" + allLinksSlice[i]
		links = append(links, makeLink)
	}

	// Create a copy of all the links
	for _, item := range links {
		linksCopy = append(linksCopy, item)
	}

	
	c := time.Tick(30 * time.Minute)
	if len(links) >= 1 {
		for _, link := range links {
			checkLink(link)
		}
		fmt.Println("Finished checking...")
		fmt.Println("Sleeping, will check again in 15 mins..")
	}
	for range c {
		if len(linksToRemove) >= 1 {

			for _, lr := range linksToRemove {
				for i, link := range links {
					if lr == link {
						copy(links[i:], links[i+1:])
						links[len(links)-1] = ""
						links = links[:len(links)-1]
					}
				}
			}
		}

		// Run if links[] slice contains any links
		if len(links) >= 1 {
			for _, link := range links {
				checkLink(link)
			}
		}

	}			

}

func checkLink(link string) {

	fmt.Println("Checking User Status....")
	fmt.Println("---------------------->")
	fmt.Println("")
	checkerClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, link, nil)
	req.Header.Set("User-Agent", "Twitch_LiveChecker")
	req.Header.Add("Client-ID", clientId)
	req.Header.Add("Authorization", "Bearer "+authToken)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Twitch_LiveChecker")

	res, getErr := checkerClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	chnl := &chanel{}

	jsonErr := json.Unmarshal(body, &chnl)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if len(chnl.Data) < 1 {
		fmt.Println("")
		fmt.Println("User status: offline!")
		fmt.Println("###############")
		fmt.Println("")

	} else {
		fmt.Println("")
		fmt.Println("User status: Live")
		fmt.Println("###############")
		//fmt.Printf("%+v\n", chnl.Data)
		fmt.Println("Username: " + chnl.Data[0].UserName + " Viewer Count: " + strconv.Itoa(chnl.Data[0].ViewerCount))
		fmt.Println("Opening channel...")
		userName := string(chnl.Data[0].UserName)
		TwitchLink := "https://www.twitch.tv/" + userName
		fmt.Println(TwitchLink)
		openbrowser(TwitchLink)
		linksToRemove = append(linksToRemove, link)
	}

}
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
