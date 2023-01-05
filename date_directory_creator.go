package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// determine which OS program is being run
func date_directory_creatorOS() string {
	var OS string
	switch runtime.GOOS {
	case "windows":
		OS = "date_directory_creator.exe"
	case "linux":
		OS = "date_directory_creator.bin"
	case "darwin":
		OS = "date_directory_creator-darwin"
	default:
		OS = "date_directory_creator"
	}
	return OS
}

// version info
func versionFunc() {
	fmt.Println(" ------------------------")
	fmt.Println("| Date Directory Creator |")
	fmt.Println(" ------------------------")
	fmt.Println("version:\t0.3.1")
	fmt.Println("build date:\t2023-01-05-1600")
	fmt.Println("written by:\tcyclone")
}

// help menu
func helpFunc(OS string) {
	versionFunc()
	checkYear()
	checkDay()
	fmt.Println("\nUsage Examples:")
	fmt.Println("(create directories for all days of 2025)")
	fmt.Println(OS, "-year 2025\n")
	fmt.Println("(create directories for all days of 2020-2030)")
	fmt.Println(OS, "-year 2020-2030\n")
	fmt.Println("(create directories for all sunday's of 2025)")
	fmt.Println(OS, "-year 2025 -day sunday\n")
	fmt.Println("(create directories for monday-friday of 2020-2030)")
	fmt.Println(OS, "-year 2020-2030 -day monday-friday\n")
	fmt.Println("Program will attempt to not overwrite existing directories, but use with caution.\n")
}

// year flag info
func checkYear() {
	fmt.Println()
	fmt.Println("--> Year/s <--")
	fmt.Println("Example:")
	fmt.Println("-year 2023")
	fmt.Println("-year 2020-2030")
}

// day flag info
func checkDay() {
	fmt.Println()
	fmt.Println("--> Day/s <--")
	fmt.Println("Example:")
	fmt.Println("-day monday")
	fmt.Println("-day monday-friday")
	fmt.Println("(weekday starts on sunday)")
	fmt.Println("(if -day is not specified, defaults to full week)")
}

// press any key to continue
func anyKeyFunc() {
	// Print message
	fmt.Println("Press any key to continue...")
	// Create a new buffered reader
	reader := bufio.NewReader(os.Stdin)
	// Read a single byte from the input
	_, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// parses string representation of the day of the week into time.Weekday
func parseWeekday(day string) time.Weekday {
	switch strings.ToLower(day) {
	case "sunday", "Sunday":
		return time.Sunday
	case "monday", "Monday":
		return time.Monday
	case "tuesday", "Tuesday":
		return time.Tuesday
	case "wednesday", "Wednesday":
		return time.Wednesday
	case "thursday", "Thursday":
		return time.Thursday
	case "friday", "Friday":
		return time.Friday
	case "saturday", "Saturday":
		return time.Saturday
	default:
		checkDay()
		os.Exit(0)
		return time.Sunday
	}
}

// main function
func main() {
	// Parse the year(s) and day(s) of the week from the command line flags
	var yearFlag, dayFlag string
	flag.StringVar(&yearFlag, "year", "", "Year or range of years.\nExampe:\n-year 2020\n-year 2020-2022")
	flag.StringVar(&dayFlag, "day", "sunday-saturday", "Day or range of days of the week.\nExample:\n-day monday\n-day monday-friday\n")
	version := flag.Bool("version", false, "Prints program version.")
	help := flag.Bool("help", false, "Prints help menu.")
	flag.Parse()

	// run sanity checks for -version & -help
	if *version == true {
		versionFunc()
		os.Exit(0)
	} else if *help == true {
		helpFunc(date_directory_creatorOS())
		os.Exit(0)
	}

	// sanity check if no flags are specified
	if len(yearFlag) < 2 || len(dayFlag) < 2 {
		helpFunc(date_directory_creatorOS())
		anyKeyFunc()
		os.Exit(0)
	}

	// sanity check on yearFlag
	if len(yearFlag) < 2 {
		checkYear()
		os.Exit(0)
	}
	// sanity check on dayFlag
	if len(dayFlag) < 2 {
		checkDay()
		os.Exit(0)
	}
	// split yearFlag string into a start year and end year
	var startYear, endYear int
	years := strings.Split(yearFlag, "-")
	if len(years) == 1 {
		startYear, _ = strconv.Atoi(years[0])
		endYear = startYear
	} else if len(years) == 2 {
		startYear, _ = strconv.Atoi(years[0])
		endYear, _ = strconv.Atoi(years[1])
	} else if len(years) < 2 {
		checkYear()
		os.Exit(0)
	} else {
		checkYear()
		os.Exit(0)
	}

	// split dayFlag string into a start day and end day
	var startDay, endDay time.Weekday
	days := strings.Split(dayFlag, "-")
	if len(days) == 1 {
		startDay = parseWeekday(days[0])
		endDay = startDay
	} else if len(days) == 2 {
		startDay = parseWeekday(days[0])
		endDay = parseWeekday(days[1])
	} else {
		checkDay()
		os.Exit(0)
	}

	// create a directory for each year in range
	for year := startYear; year <= endYear; year++ {
		yearDir := fmt.Sprintf("%d", year)
		// check if directory already exists
		if _, err := os.Stat(yearDir); os.IsNotExist(err) {
			// create directory if does not exist
			os.Mkdir(yearDir, os.ModePerm)
		} else {
			// print out directory if it already exists
			fmt.Println("Directory", yearDir, "already exists.")
		}
		// create a directory for each month in year
		for month := time.January; month <= time.December; month++ {
			monthDir := fmt.Sprintf("%s/%02d-%s", yearDir, month, month.String())
			// check if directory already exists
			if _, err := os.Stat(monthDir); os.IsNotExist(err) {
				// create directory if does not exist
				os.Mkdir(monthDir, os.ModePerm)
			} else {
				// print out directory if it already exists
				fmt.Println("Subdirectory", monthDir, "already exists.")
			}

			// get number of days in each month
			daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

			// create directory for each day in the month
			for day := 1; day <= daysInMonth; day++ {
				// check if day is in specified range of days
				date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
				if date.Weekday() >= startDay && date.Weekday() <= endDay {
					dayDir := fmt.Sprintf("%s/%02d-%02d-%v", monthDir, month, day, year)
					// check if directory already exists
					if _, err := os.Stat(dayDir); os.IsNotExist(err) {
						// create directory if does not exist
						fmt.Println("Creating directory:", dayDir)
						os.Mkdir(dayDir, os.ModePerm)
					} else {
						// print out directory if it already exists
						fmt.Println("Subdirectory", dayDir, "already exists.")
					}
				}
			}
		}
	}
	fmt.Println()
	os.Exit(0)
}
// end of program
