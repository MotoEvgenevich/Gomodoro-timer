package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	cfg, err := parseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Welcome to Pomodoro timer with the following settings:\n- Work Time: %d min\n- Break Time: %d min\n- Cycles: %d\n",
		cfg.WorkTime, cfg.BreakTime, cfg.Cycles)

	runPomodoro(cfg)
}

type Config struct {
	WorkTime  int
	BreakTime int
	Cycles    int
}

func parseFlags() (*Config, error) {
	work := flag.String("work", "25m", "minutes/hours of work")
	pause := flag.String("break", "5m", "time of coffee break")
	cycles := flag.Int("cycles", 4, "number of cycles")

	flag.Parse()

	workTime, err := convertTime(*work)
	if err != nil {
		return nil, err
	}
	breakTime, err := convertTime(*pause)
	if err != nil {
		return nil, err
	}

	if *cycles <= 0 {
		return nil, fmt.Errorf("invalid value for cycles: %d", *cycles)
	}

	return &Config{WorkTime: workTime, BreakTime: breakTime, Cycles: *cycles}, nil
}

func convertTime(input string) (int, error) {
	if len(input) < 2 {
		return 0, fmt.Errorf("invalid time format: %s", input)
	}

	unit := input[len(input)-1]
	time, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		return 0, fmt.Errorf("invalid number in time: %s", input)
	}

	if unit == 'h' || unit == 'H' {
		time *= 60
	} else if unit != 'm' && unit != 'M' {
		return 0, fmt.Errorf("unknown time unit in: %s", input)
	}

	return time, nil
}

func runPomodoro(cfg *Config) {
	fmt.Println("Timer Started!")
	for i := 0; i < cfg.Cycles; i++ {
		if startTimer(cfg.WorkTime) {
			beep()
			fmt.Println("Stop work! Time for a break.")
		}
		if i < cfg.Cycles-1 { // Проверяем, нужно ли начинать новый цикл работы
			if startTimer(cfg.BreakTime) {
				beep()
				fmt.Println("Break over! Start working.")
			}
		}
	}
	fmt.Println("Pomodoro session completed!")
}

func startTimer(duration int) bool {
	fmt.Printf("Starting %d minutes timer...\n", duration)
	timer := time.NewTimer(time.Duration(duration) * time.Minute)
	<-timer.C
	return true
}

func beep() {
	f, err := os.Open("beep.wav")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open beep sound: %v\n", err)
		return
	}
	defer f.Close()

	streamer, format, err := wav.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode beep sound: %v\n", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool) // Создаем канал для уведомления о завершении воспроизведения

	// Воспроизведение стримера и ожидание его завершения
	speaker.Play(streamer)
	go func() {
		speaker.Lock()
		defer speaker.Unlock()
		for streamer.Position() < streamer.Len() {
			// Ожидаем, пока позиция стримера не достигнет его конца
			time.Sleep(100 * time.Millisecond)
		}
		done <- true
	}()

	<-done                             // Ожидание завершения воспроизведения
	time.Sleep(500 * time.Millisecond) // Краткая пауза после завершения воспроизведения
}
