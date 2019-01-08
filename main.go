package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

var start, end, apiKey, linkToVisit string

func main() {
	flag.StringVar(&start, "start", "", "Coordinate of the starting point (like 50.138197,2.273150)")
	flag.StringVar(&end, "end", "", "Coordinate of the ending point (like 49.238197,2.343180)")
	flag.StringVar(&apiKey, "apiKey", "", "Developer API key")
	flag.StringVar(&linkToVisit, "linkToVisit", "", "Link to visit via 'Open MyDrive map' when clicked")
	flag.Parse()

	if start == "" {
		fmt.Println("Argument start is not set")
		flag.PrintDefaults()
		return
	}
	if end == "" {
		fmt.Println("Argument end is not set")
		flag.PrintDefaults()
		return
	}
	if apiKey == "" {
		fmt.Println("Argument apiKey is not set")
		flag.PrintDefaults()
		return
	}
	if linkToVisit == "" {
		fmt.Println("Argument linkToVisit is not set")
		flag.PrintDefaults()
		return
	}

	go check()
	systray.Run(onReady, nil)
}

func check() {
	tick := time.Tick(30 * time.Second)
	process()
	for {
		select {
		case <-tick:
			process()
		}
	}
}

func process() {
	timeToHome, err := getTimeToHome()
	if err != nil {
		systray.SetTitle(fmt.Sprintf("Error: %v", err))
	} else {
		msg := fmt.Sprintf("%d min", timeToHome/60)
		log.Println(msg)
		systray.SetTitle(msg)
	}
}

type Response struct {
	Routes []struct {
		Summary struct {
			TravelTimeInSeconds int `json:"travelTimeInSeconds"`
		} `json:"summary"`
	} `json:"routes"`
}

func getTimeToHome() (seconds int, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.tomtom.com/routing/1/calculateRoute/%s:%s/json?key=%s", start, end, apiKey))
	if err != nil {
		log.Printf("Failed to reach remote server: %v", err)
		return -1, err
	}
	defer resp.Body.Close()

	var response Response

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		log.Printf("Failed to decode response: %v", err)
		return -1, err
	}
	if len(response.Routes) == 0 {
		log.Printf("No routes received!")
		return -1, err
	}

	return response.Routes[0].Summary.TravelTimeInSeconds, nil
}

func onReady() {
	go func() {
		systray.SetIcon(Data)
		systray.SetTitle("...")
		systray.SetTooltip("Estimated time to reach home")
		mVisit := systray.AddMenuItem("Open MyDrive map", "Visit the MyDrive map to see the route in details")
		systray.AddSeparator()
		mVisitHomePage := systray.AddMenuItem("Homepage", "Visit the ap home page")
		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
		for {
			select {
			case <-mVisit.ClickedCh:
				open.Run(linkToVisit)
			case <-mVisitHomePage.ClickedCh:
				open.Run("https://github.com/milanaleksic/timetohome")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}
