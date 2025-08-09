package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var showProgress bool
	flag.BoolVar(&showProgress, "v", false, "verbose output")
	flag.BoolVar(&showProgress, "verbose", false, "verbose output")
	
	flag.Usage = printUsage
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		printUsage()
		os.Exit(1)
	}

	duration, err := parseDuration(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "countdown: invalid time interval '%s'\n", args[0])
		os.Exit(1)
	}

	runCountdown(duration, showProgress)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: countdown [-v] NUMBER[SUFFIX]...\n")
	fmt.Fprintf(os.Stderr, "Pause for NUMBER seconds.  SUFFIX may be 's' for seconds (the default),\n")
	fmt.Fprintf(os.Stderr, "'m' for minutes, 'h' for hours.  Multiple numbers may be specified\n")
	fmt.Fprintf(os.Stderr, "by combining them: 1h30m means 1 hour 30 minutes.\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  -v, --verbose     display countdown progress\n")
	fmt.Fprintf(os.Stderr, "      --help        display this help and exit\n")
}

func parseDuration(input string) (time.Duration, error) {
	return time.ParseDuration(input)
}

func runCountdown(duration time.Duration, showProgress bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	remaining := duration

	for {
		select {
		case <-c:
			if showProgress {
				fmt.Println("\nCountdown interrupted")
			}
			return
		case <-ticker.C:
			if remaining <= 0 {
				if showProgress {
					fmt.Println("")
				}
				return
			}

			if showProgress {
				fmt.Printf("\r%s", formatDuration(remaining))
			}
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
