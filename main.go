package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var (
	flagD string
	flagF string

	duration time.Duration
	srtFile  *os.File
)

func handleLine(dstFile *os.File, line string) {
	if !strings.Contains(line, "-->") {
		dstFile.WriteString(line)
		return
	}
	strSli := strings.Split(line, "-->")

	start := strings.TrimSpace(strSli[0])
	end := strings.TrimSpace(strSli[1])

	startSli := strings.Split(start, ":")
	endSli := strings.Split(end, ":")

	startSli = append(startSli[:2], strings.Split(startSli[2], ",")...)
	endSli = append(endSli[:2], strings.Split(endSli[2], ",")...)

	startDuration, _ := time.ParseDuration(fmt.Sprintf("%sh%sm%ss%sms", startSli[0], startSli[1], startSli[2], startSli[3]))
	endDuration, _ := time.ParseDuration(fmt.Sprintf("%sh%sm%ss%sms", endSli[0], endSli[1], endSli[2], endSli[3]))
	startDuration += duration
	endDuration += duration

	line = fmt.Sprintf("%s --> %s\n", duration2srtTime(startDuration), duration2srtTime(endDuration))
	dstFile.WriteString(line)
}

func duration2srtTime(d time.Duration) string {
	h := int64(math.Floor(d.Hours()))
	m := int64(math.Floor(d.Minutes())) % 60
	s := int64(math.Floor(d.Seconds())) % 60
	ms := d.Milliseconds() % 1000

	return fmt.Sprintf("%02d:%02d:%02d,%d", h, m, s, ms)
}

func init() {
	flag.StringVar(&flagD, "d", "", "the time duration to adjust")
	flag.StringVar(&flagF, "f", "", "the srt file")
	flag.Parse()

	var err error
	duration, err = time.ParseDuration(flagD)
	if err != nil {
		log.Fatal(err)
	}

	srtFile, err = os.Open(flagF)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer srtFile.Close()

	bufReader := bufio.NewReader(srtFile)
	dstFile, err := os.Create("dst.srt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	for {
		line, err := bufReader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		// work
		handleLine(dstFile, line)
		// break
		if err == io.EOF {
			fmt.Println("Complete!")
			break
		}
	}
}
