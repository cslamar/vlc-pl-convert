/*
Copyright © 2022 Chris Slamar <chris@slamar.com>
*/
package playlist

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Generate compile new playlist
func Generate(inputFile, outputFile string, pauseSeconds int, verbose bool) error {
	playlist, err := parse(inputFile)
	if err != nil {
		return err
	}

	playlist.TrackList = generateNewTrackList(playlist.TrackList, pauseSeconds)

	trackIds := make([]VlcItem, 0)
	for i := 0; i < len(playlist.TrackList); i++ {
		trackIds = append(trackIds, generateTrackId(i))
	}

	playlist.Extension.TrackIds = trackIds

	if err := output(playlist, outputFile, verbose); err != nil {
		return err
	}

	return nil
}

// parse input file into Playlist object
func parse(filename string) (Playlist, error) {
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

// output write new playlist to file
func output(input any, outFile string, verbose bool) error {
	output, err := xml.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	finalOut := []byte(xml.Header + string(output))

	if verbose {
		fmt.Println(string(finalOut))
	}

	if err := os.WriteFile(outFile, finalOut, 0644); err != nil {
		return err
	}

	return nil
}
