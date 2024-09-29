package bom_radar_downloader

import (
	"fmt"
	"time"
)

// Eg. IDR66A == QLD Mt Stapylton 6m rain radar, see http://www.bom.gov.au/catalogue/anon-ftp.shtml for full list

func Hello() {
	fmt.Println("First time doing this :3, lets see if it works")
}

// Encode
func Encode(t time.Time) string {
	var str = "time is, " + t
	return str
}
