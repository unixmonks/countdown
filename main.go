package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: countdown duration\n")
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "-h" || arg == "--help" {
		printHelp()
		os.Exit(0)
	}

	duration, err := parseDuration(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countdown: %s\n", err)
		os.Exit(1)
	}

	runCountdown(duration)
}

func printHelp() {
	fmt.Print(`countdown - display a countdown timer

USAGE
    countdown duration

DESCRIPTION
    countdown displays a countdown timer for the specified duration.
    The duration can be specified using Go's duration format.

EXAMPLES
    countdown 30s        30 seconds
    countdown 5m         5 minutes
    countdown 1h30m      1 hour 30 minutes
    countdown 1m30s      1 minute 30 seconds

OPTIONS
    -h, --help    show this help message

`)
}

func parseDuration(input string) (time.Duration, error) {
	return time.ParseDuration(input)
}

func runCountdown(duration time.Duration) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	remaining := duration

	for {
		select {
		case <-c:
			fmt.Println("\nCountdown interrupted")
			return
		case <-ticker.C:
			if remaining <= 0 {
				fmt.Println("")
				return
			}

			fmt.Printf("\r%s", formatDuration(remaining))
			remaining -= time.Second
		}
	}
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
