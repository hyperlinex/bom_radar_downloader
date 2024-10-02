package bom_radar_downloader

import (
	"fmt"
	"testing"
	"time"

	"github.com/jlaffaye/ftp"
)

var prodID = "IDR46A"
var ftpUrl = "ftp.bom.gov.au:21"
var user = "anonymous"
var pass = "anonymous"
var ftpDir = "/anon/gen/radar"

func TestHello(t *testing.T) {

	Hello()
}

func TestGetFileNames(t *testing.T) {

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

func TestEncode(t *testing.T) {
	var prodID = "IDR66A"
	var time = time.Now()
	fileName := Encode(prodID, time.UTC())
	fmt.Println(fileName)
}

func TestDecode(t *testing.T) {
	// Grab latest file and decode its properties

	// Connect to FTP server
	c, err := ftp.Dial(ftpUrl)
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

	fileNames, err := GetFileNames(c, prodID, 1)
	if err != nil {
		panic(err)
	}

	date, prodId, err := Decode(fileNames[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("Input:", fileNames[0])
	fmt.Println("Product ID:", prodId, ", date:", date)
}
