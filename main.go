package main

import (
	// "fmt"
	"encoding/xml"
	"log"
	"os"
)

// type Run struct {
// 	startDate time.Time
// 	finishDate time.Time
// 	realTime time.Time
// 	gameTime time.Time
// 	pauseTime time.Time
// }
// type Results struct {
// 	personalBest Run
// }

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

// func findPB(splits *LiveSplit) string {

// 	// attempts := []time.Time

// 	for _, attempt := range splits.AttemptHistory.Attempt {
// 		if attempt.RealTime != "" {

// 			t, err := time.Parse(timerFormat, attempt.RealTime)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			log.Printf("%v, %v\n", attempt.RealTime, t)
// 		}
// 	}
// 	return "WIP"
// }

func main() {
	lsData := getData("test_data/sm.lss")

	// var result Results

	gameSplits := DecodeRun(lsData)
	log.Println(gameSplits.GameName)

	// pb := findPB(&gameSplits)
	// println(pb)
}
