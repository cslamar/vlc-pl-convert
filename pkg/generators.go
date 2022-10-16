/*
Copyright © 2022 Chris Slamar <chris@slamar.com>
*/
package playlist

import "fmt"

// generatePauseTrack returns a Track object with new data
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

// generateTrackId returns new TrackId object
func generateTrackId(trackNumber int) VlcItem {
	return VlcItem{
		TrackId: trackNumber,
	}
}

// generateNewTrackList compiles new track list with pause tracks
func generateNewTrackList(trackList []Track, pauseSeconds int) []Track {
	newTracks := make([]Track, 0)

	ticks := 0

	for _, track := range trackList {
		t := track
		t.TrackExtensions.Id = ticks

		newTracks = append(newTracks, t)
		newTracks = append(newTracks, generatePauseTrack(pauseSeconds, ticks+1))
		ticks += 2
	}

	return newTracks
}
