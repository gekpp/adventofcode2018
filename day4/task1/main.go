package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	var stdin io.Reader = bytes.NewReader([]byte(`[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up
	`))
	stdin = os.Stdin

	r := bufio.NewReader(stdin)

	var (
		lastGN   string
		lastTime time.Time
		sleepReg = make(map[string][60]int)
	)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.Trim(line, "\n")

		prts := strings.Split(line, " ")
		timePart := prts[0] + " " + prts[1]
		t, err := time.Parse("2006-01-02 15:04", timePart[1:len(timePart)-1])
		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(prts[2], "Guard") {
			lastTime = t
			lastGN = prts[3]
			continue
		}

		if strings.HasPrefix(prts[2], "falls") {
			lastTime = t
			continue
		}

		if strings.HasPrefix(prts[2], "wakes") {
			hour := sleepReg[lastGN]

			for m := lastTime.Minute(); m < t.Minute(); m++ {
				hour[m] += 1
			}
			sleepReg[lastGN] = hour
		}
	}

	var (
		maxTimes    int
		maxGN       string
		minute      int
		maxSleepDur int
	)
	for guard, hour := range sleepReg {
		var (
			localMaxTimes int
			localMinute   int
			localSleepDur int
		)

		for minute, sleepTimes := range hour {
			localSleepDur += sleepTimes
			if sleepTimes > localMaxTimes {
				localMaxTimes = sleepTimes
				localMinute = minute
			}
		}

		if localSleepDur > maxSleepDur {
			maxGN = guard
			maxSleepDur = localSleepDur
			maxTimes = localMaxTimes
			minute = localMinute
		}
	}
	fmt.Printf("%v %v on min %v\n", maxGN, maxTimes, minute)
}
