# Vlc Playlist Conversion Tool

This was designed to enable an easy way to add pause tracks between an existing Vlc playlist file for dubbing to MiniDisc.

```shell
Create new VLC playlist with pause tracks for recording

Usage:
  vlc-pl-convert [flags]

Flags:
  -h, --help                help for vlc-pl-convert
  -i, --input string        input VLC xspf playlist file
  -o, --output string       output VLC xspf playlist file (default "output.xspf")
  -p, --pause-seconds int   pause delay between tracks (default 2)
  -v, --verbose             enable verbose debugging output
```
