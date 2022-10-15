package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Playlist struct {
	XMLName   xml.Name          `xml:"playlist"`
	Ns        string            `xml:"xmlns,attr"`
	VlcNs     string            `xml:"xmlns vlc,attr"`
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
	pl, err := parseFile("test.xspf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pl.TrackList = generateNewTrackList(pl.TrackList)

	tids := make([]VlcItem, 0)
	for i := 0; i < len(pl.TrackList); i++ {
		tids = append(tids, generateTrackId(i))
	}

	pl.Extension.TrackIds = tids

	xmlPrint(pl)
}

func generateNewTrackList(trackList []Track) []Track {
	newTracks := make([]Track, 0)

	ticks := 0

	for _, track := range trackList {
		t := track
		t.TrackExtensions.Id = ticks

		newTracks = append(newTracks, t)
		newTracks = append(newTracks, generatePauseTrack(2, ticks+1))
		ticks += 2
	}

	return newTracks
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

func generatePauseTrack(pauseSeconds, trackNumber int) Track {
	return Track{
		Location: fmt.Sprintf("vlc://pause:%d", pauseSeconds),
		Duration: pauseSeconds,
		TrackExtensions: VlcId{
			Application: "http://www.videolan.org/vlc/playlist/0",
			Id:          trackNumber,
		},
	}
}

func generateTrackId(trackNumber int) VlcItem {
	return VlcItem{
		TrackId: trackNumber,
	}
}

func xmlPrint(input any) {
	output, err := xml.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	finalOut := []byte(xml.Header + string(output))

	fmt.Println(string(finalOut))
}
