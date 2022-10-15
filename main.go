package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Playlist struct {
	XMLName   xml.Name          `xml:"playlist"`
	Ns        string            `xml:"xmlns,attr"`
	VlcNs     string            `xml:"xmlns:vlc,attr"`
	Version   int               `xml:"version,attr"`
	Title     string            `xml:"title"`
	TrackList []Track           `xml:"trackList>track"`
	Extension PlaylistExtension `xml:"extension"`
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
	Id          int    `xml:"http://www.videolan.org/vlc/playlist/ns/0/ id"`
}

type VlcItem struct {
	XMLName xml.Name `xml:"http://www.videolan.org/vlc/playlist/ns/0/ item"`
	TrackId int      `xml:"tid,attr"`
}

type PlaylistExtension struct {
	XMLName     xml.Name  `xml:"extension"`
	Application string    `xml:"application,attr"`
	TrackIds    []VlcItem `xml:",any"`
}

func main() {

	//testItOut()
	pl, err := parseFile("test.xspf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(pl)

	xmlPrint(pl)
}

func parseFile(filename string) (Playlist, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Playlist{}, err
	}

	var playlist = Playlist{}

	if err := xml.Unmarshal(data, &playlist); err != nil {
		return Playlist{}, err
	}

	return playlist, nil
}

//func generatePauseTrack(pauseSeconds, trackNumber int) Track {
//	return Track{
//		Location: fmt.Sprintf("vlc://pause:%d", pauseSeconds),
//		Duration: pauseSeconds,
//		TrackExtensions: VlcId{
//			Application: "http://www.videolan.org/vlc/playlist/0",
//			Id:          fmt.Sprint(trackNumber),
//		},
//	}
//}
//
//func generateTrackId(trackNumber int) VlcItem {
//	return VlcItem{
//		TrackId: fmt.Sprint(trackNumber),
//	}
//}

func xmlPrint(input any) {
	output, err := xml.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	finalOut := []byte(xml.Header + string(output))

	fmt.Println(string(finalOut))
}
