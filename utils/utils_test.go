package utils

import (
	"log"
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	parseT := StringToTime("2020-10-20 06:35:00")
	now := time.Now()
	log.Println(parseT)
	log.Println(now)
	log.Println(now.After(parseT))
}

func TestMilliToTime(t *testing.T) {
	log.Println(MilliToTime(1578964135000))
	log.Println(MilliToTime(NowMillisecond()))
}
