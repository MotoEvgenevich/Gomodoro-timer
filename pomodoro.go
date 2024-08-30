package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	work := flag.String("work", "25m", "minutes/hours of work")
	pause := flag.String("break", "5m", "time of coffee break:)")
	cycles := flag.Int("cycles", 4, "quantity of cycles")
	flag.Parse()

	if err := validateFlags(*work, *pause, *cycles); err != nil {
		fmt.Println("Error, reason:", err)
		os.Exit(1)
	}
	w, p := Convert(*work, *pause)

	fmt.Printf("Welcome to Pomodoro timer with next conditions: \n Time to work: %v min. \n Time to rest: %v min. \n", w, p)
	Beep()
	fmt.Println("Timer Started!")
	for i := 0; i < *cycles; i++ {
		if Timer(w) {
			Beep()
			fmt.Println("Stop work!")
		}
		if Timer(p) {
			Beep()
			if i+1 < *cycles {
				fmt.Println("Start work!")
			}

		}
	}

	fmt.Println("Timer stoped!")
}

func Timer(minutes int) bool {
	timer := time.NewTimer(time.Duration(minutes) * time.Minute)
	<-timer.C
	return true
}

func Beep() {

	f, err := os.Open("beep.wav")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	streamer, format, err := wav.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(beep.Loop(1, streamer))

	time.Sleep(2 * time.Second)
}

func Convert(work, pause string) (w, p int) {
	hourOrMinuteWork := work[len(work)-1:]
	hourOrMinutePause := pause[len(pause)-1:]

	w, err := strconv.Atoi(work[:len(work)-1])
	if err != nil {
		fmt.Println("Error converting work:", err)
	}

	if hourOrMinuteWork == "H" || hourOrMinuteWork == "h" {
		w = w * 60
	}

	p, err = strconv.Atoi(pause[:len(pause)-1])
	if err != nil {
		fmt.Println("Error converting pause:", err)
	}

	if hourOrMinutePause == "H" || hourOrMinutePause == "h" {
		p = p * 60
	}
	return w, p
}

func validateFlags(work, pause string, cycles int) error {

	if !isValidValue(work) {
		return fmt.Errorf("incorect value for flag --work: %v", work)
	}

	if !isValidValue(pause) {
		return fmt.Errorf("incorect value for flag --break: %v", pause)
	}

	if cycles <= 0 {
		return fmt.Errorf("incorect value for flag --cycles: %v", cycles)
	}

	return nil
}

func isValidValue(value string) bool {
	if len(value) < 2 || !(strings.HasSuffix(value, "m") || strings.HasSuffix(value, "h")) {
		return false
	}

	numberPart := value[:len(value)-1]
	if _, err := strconv.Atoi(numberPart); err != nil || numberPart[0] == '-' {
		return false
	}

	return true
}
