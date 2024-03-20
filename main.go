package main

import (
	// "fmt"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"
)

type LiveSplit struct {
	XMLName      xml.Name `xml:"Run"`
	Text         string   `xml:",chardata"`
	Version      string   `xml:"version,attr"`
	GameIcon     string   `xml:"GameIcon"`
	GameName     string   `xml:"GameName"`
	CategoryName string   `xml:"CategoryName"`
	LayoutPath   string   `xml:"LayoutPath"`
	Metadata     struct {
		Text string `xml:",chardata"`
		Run  struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"Run"`
		Platform struct {
			Text         string `xml:",chardata"`
			UsesEmulator string `xml:"usesEmulator,attr"`
		} `xml:"Platform"`
		Region    string `xml:"Region"`
		Variables struct {
			Text     string `xml:",chardata"`
			Variable []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"Variable"`
		} `xml:"Variables"`
	} `xml:"Metadata"`
	Offset         string `xml:"Offset"`
	AttemptCount   string `xml:"AttemptCount"`
	AttemptHistory struct {
		Text    string `xml:",chardata"`
		Attempt []struct {
			Text            string `xml:",chardata"`
			ID              string `xml:"id,attr"`
			Started         string `xml:"started,attr"`
			IsStartedSynced string `xml:"isStartedSynced,attr"`
			Ended           string `xml:"ended,attr"`
			IsEndedSynced   string `xml:"isEndedSynced,attr"`
			RealTime        string `xml:"RealTime"`
			GameTime        string `xml:"GameTime"`
			PauseTime       string `xml:"PauseTime"`
		} `xml:"Attempt"`
	} `xml:"AttemptHistory"`
	Segments struct {
		Text    string `xml:",chardata"`
		Segment []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"Name"`
			Icon       string `xml:"Icon"`
			SplitTimes struct {
				Text      string `xml:",chardata"`
				SplitTime []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"name,attr"`
					RealTime string `xml:"RealTime"`
					GameTime string `xml:"GameTime"`
				} `xml:"SplitTime"`
			} `xml:"SplitTimes"`
			BestSegmentTime struct {
				Text     string `xml:",chardata"`
				RealTime string `xml:"RealTime"`
				GameTime string `xml:"GameTime"`
			} `xml:"BestSegmentTime"`
			SegmentHistory struct {
				Text string `xml:",chardata"`
				Time []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					RealTime string `xml:"RealTime"`
					GameTime string `xml:"GameTime"`
				} `xml:"Time"`
			} `xml:"SegmentHistory"`
		} `xml:"Segment"`
	} `xml:"Segments"`
	AutoSplitterSettings string `xml:"AutoSplitterSettings"`
}

func getData(filename string) []byte {
	lsData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	return lsData
}

func DecodeRun(data []byte) LiveSplit {

	var run LiveSplit

	if err := xml.Unmarshal(data, &run); err != nil {
		log.Fatalf("Could not Decode file: %s", err)
		return LiveSplit{}
	}

	return run
}

func findPB(splits *LiveSplit) string {
	
	attempts := []time.Time
	
	for _, attempt := range splits.AttemptHistory.Attempt {
		if attempt.RealTime != "" {
			timeFormat := "15:04:05.0000000"
			t, err := time.Parse(timeFormat, attempt.RealTime)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%v, %v\n", attempt.RealTime, t)
		}
	}
	return "WIP"
}

func main() {
	lsData := getData("test_data/sm.lss")

	run := DecodeRun(lsData)

	pb := findPB(&run)
	println(pb)
}
