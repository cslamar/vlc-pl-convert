package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Playlist struct {
	XMLName   xml.Name `xml:"playlist"`
	Ns        string   `xml:"xmlns,attr"`
	VlcNs     string   `xml:"xmlns:vlc,attr"`
	Version   int      `xml:"version,attr"`
	Title     string   `xml:"title"`
	TrackList []Track  `xml:"trackList>track"`
	Extension Ext      `xml:"extension"`
}

type Track struct {
	XMLName         xml.Name `xml:"track"`
	Location        string   `xml:"location"`
	Duration        int      `xml:"duration"`
	Title           string   `xml:"title,omitempty"`
	Creator         string   `xml:"creator,omitempty"`
	TrackExtensions VlcId    `xml:"extension"`
}

type VlcId struct {
	Application string `xml:"application,attr"`
	Id          string `xml:",any"`
}

type VlcItem struct {
	TrackId string `xml:",attr,any"`
}

type Ext struct {
	XMLName     xml.Name  `xml:"extension"`
	Application string    `xml:"application,attr"`
	TrackIds    []VlcItem `xml:",any"`
}

func main() {

	//testItOut()
	parseFile("test.xspf")
}

func parseFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var playlist = Playlist{}

	if err := xml.Unmarshal(data, &playlist); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(playlist)
}

//func testItOut() {
//	t := Track{
//		Location: "vlc://pause:2",
//		Duration: 2,
//		TrackExtensions: VlcId{
//			Application: "http://www.videolan.org/vlc/playlist/0",
//			Id:          1,
//		},
//	}
//
//	item := VlcItem{
//		Tid: 1,
//	}
//
//	e := Ext{
//		Application: "http://www.videolan.org/vlc/playlist/0",
//		Tids:        []VlcItem{item, item},
//	}
//
//	p := Playlist{
//		Ns:        "http://xspf.org/ns/0/",
//		VlcNs:     "http://www.videolan.org/vlc/playlist/ns/0/",
//		Version:   1,
//		Title:     "Something",
//		TrackList: []Track{t, t},
//		Extension: e,
//	}
//
//	xmlPrint(p)
//}

func xmlPrint(input any) {
	output, err := xml.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	finalOut := []byte(xml.Header + string(output))

	fmt.Println(string(finalOut))
}
