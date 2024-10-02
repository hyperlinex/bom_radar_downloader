package bom_radar_downloader

import (
	"fmt"
	"testing"

	"github.com/jlaffaye/ftp"
)

func TestHello(t *testing.T) {

	Hello()
}

func TestGetFileNames(t *testing.T) {
	var prodID = "IDR66A"
	var ftpUrl = "ftp.bom.gov.au:21"
	var user = "anonymous"
	var pass = "anonymous"
	var ftpDir = "/anon/gen/radar"

	c, err := ftp.Connect(ftpUrl)
	if err != nil {
		panic(err)
	}
	defer c.Quit()

	err = c.Login(user, pass)
	if err != nil {
		panic(err)
	}

	if err := c.ChangeDir(ftpDir); err != nil {
		panic(err)
	}

	files, err := GetFileNames(c, prodID, 10)
	if err != nil {
		panic(err)

	}

	for _, file := range files {
		fmt.Println(file)
	}
}
