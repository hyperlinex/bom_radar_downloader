package bom_radar_downloader

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

// Eg. IDR66A == QLD Mt Stapylton 6m rain radar, see http://www.bom.gov.au/catalogue/anon-ftp.shtml for full list

func Hello() {
	fmt.Println("First time doing this :3, lets see if it works")
}

// Encode a given date and product ID into BoM style naming convention
func Encode(productID string, t time.Time) string {
	var year = strconv.Itoa(t.Year())
	var month = fmt.Sprintf("%02d", int(t.Minute()))
	var day = fmt.Sprintf("%02d", t.Day())
	var hour = fmt.Sprintf("%02d", t.Hour())
	var minute = fmt.Sprintf("%02d", t.Minute())

	var str = productID + ".T." + year + month + day + hour + minute

	return str
}

// Decode input file name string, return productID string and Date object
func Decode(fileName string) (time.Time, string, error) {
	// First, split between .T.
	parts := strings.Split(fileName, ".T.")
	if len(parts) != 2 {
		return time.Time{}, "", fmt.Errorf("Name is invalid %s", fileName)
	}

	// Decode time
	year, err := strconv.Atoi(parts[1][0:4])
	if err != nil {
		return time.Time{}, "", fmt.Errorf("Invalid year %s", parts[1][0:4])
	}

	month, err := strconv.Atoi(parts[1][4:6])
	if err != nil {
		return time.Time{}, "", fmt.Errorf("Invalid month %s", parts[1][4:6])
	}

	day, err := strconv.Atoi(parts[1][6:8])
	if err != nil {
		return time.Time{}, "", fmt.Errorf("Invalid day %s", parts[1][6:8])
	}

	hour, err := strconv.Atoi(parts[1][8:10])
	if err != nil {
		return time.Time{}, "", fmt.Errorf("Invalid hour %s", parts[1][8:10])
	}

	minute, err := strconv.Atoi(parts[1][10:12])
	if err != nil {
		return time.Time{}, "", fmt.Errorf("Invalid minute %s", parts[1][10:12])
	}

	// Create date type with the data
	var date = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	// Return the date
	return date, parts[0], nil
}

// Return string slice of paths to files, error, of specified number of files,
// in the current directory of an FTP server, sorted by last modified first.
func GetFileNames(c *ftp.ServerConn, productID string, numFiles int) ([]string, error) {
	entries, err := c.List(".")
	if err != nil {
		return nil, err
	}

	var filteredEntries []*ftp.Entry
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name, productID+".T.") {
			filteredEntries = append(filteredEntries, entry)
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
