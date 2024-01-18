package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora/v4"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		now := time.Now()
		printInfo(fmt.Sprintf("Now: %d, %s, %s", aurora.Green(now.Unix()), aurora.Yellow(now.Format(time.RFC3339)), aurora.Cyan(now.UTC().Format(time.RFC3339))))
		return
	}
	if tt, err := strconv.ParseInt(args[1], 10, 64); err == nil {
		t := time.Unix(tt, 0)
		printInfo(fmt.Sprintf("Parse: %d, %s, %s", aurora.Green(tt), aurora.Yellow(t.Format(time.RFC3339)), aurora.Cyan(t.UTC().Format(time.RFC3339))))
		return
	}
	if t, err := time.Parse("2006-01-02 15:04:05", args[1]); err == nil {
		printInfo(fmt.Sprintf("Parse: %d, %s, %s", aurora.Green(t.Unix()), aurora.Yellow(t.Format(time.RFC3339)), aurora.Cyan(t.UTC().Format(time.RFC3339))))
		return
	} else {
		printInfo(fmt.Sprintf("Error: %s\n", aurora.Red(err)))
	}
}

func printInfo(info string) {
	f, err := os.OpenFile("./print.ans", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	txt := fmt.Sprintf("%s %s", aurora.Gray(0, time.Now().Format(time.RFC3339)), info)
	fmt.Println(txt)
	fmt.Fprintln(f, txt)
}
