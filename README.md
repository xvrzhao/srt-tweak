# srt-tweak

A srt subtitle file timeline adjustment tool.

## Installation
```
$ go install github.com/xvrzhao/srt-tweak@latest
```

## Usage

- Delay 5 seconds and 100 milliseconds:
```
$ srt-tweak -d 5s100ms -f subtitle.srt
```

- Advance 1 minute and 12 milliseconds:
```
$ srt-tweak -d -1m12ms -f subtitle.srt
```
