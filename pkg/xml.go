/*
Copyright © 2022 Chris Slamar <chris@slamar.com>
*/
package playlist

import (
	"encoding/xml"
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
