package bom_radar_downloader

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

// Eg. IDR66A == QLD Mt Stapylton 6m rain radar, see http://www.bom.gov.au/catalogue/anon-ftp.shtml for full list

func Hello() {
	fmt.Println("First time doing this :3, lets see if it works")
}

// UNFINSHED: Encode a given date and product ID into BoM style naming convention
func Encode(t time.Time) string {
	var str = "time is, " + t.String()
	return str
}

// Return string slice of paths to files, error, of specified number of files,
// in the current directory of an FTP server, sorted by last modified first.
func GetFileNames(c *ftp.ServerConn, productID string, numFiles int) ([]string, error) {
	entries, err := c.List(".")
	if err != nil {
		panic(err)
		//return nil, err
	}

	var filteredEntries []*ftp.Entry
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name, productID+".T.") {
			filteredEntries = append(filteredEntries, entry)
			// fmt.Println(entry)
		}
	}

	// Sort by descending order
	sort.Slice(filteredEntries, func(i, j int) bool {
		return filteredEntries[i].Time.After(filteredEntries[j].Time)
	})

	var latestFiles []string
	for i, entry := range filteredEntries {
		if i >= numFiles {
			break
		}
		latestFiles = append(latestFiles, entry.Name)
	}

	// Then return the slice
	return latestFiles, nil
}
