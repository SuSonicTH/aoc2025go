package main

import (
	"aoc2025/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const baseUrl string = "https://adventofcode.com/2025/day/%d/input"
const baseFilePath string = "./input/input%d.txt"

const baseDirectoryPath string = "./day%d"
const baseSourcePath string = "%s/day%d.go"
const baseTestPath string = "%s/day%d_test.go"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: aocget [DAY]")
		os.Exit(1)
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("error %s: expecting day as integer argument got '%s'", err, os.Args[1]))
	}

	downloadInput(day)
	dir := createDayDirectory(day)
	createDaySource(dir, day)
	createDayTestSource(dir, day)
}

func downloadInput(day int) {
	filepath := fmt.Sprintf(baseFilePath, day)
	if _, err := os.Stat(filepath); err == nil {
		fmt.Printf("input file %s for day %d already exists\n", filepath, day)
		return
	}

	session := util.ReadFile(".session")
	fmt.Printf("Getting input data for day %d\n", day)

	url := fmt.Sprintf(baseUrl, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		panic(fmt.Sprintf("http response %d", resp.StatusCode))
	}

	out, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

func createDayDirectory(day int) (filepath string) {
	filepath = fmt.Sprintf(baseDirectoryPath, day)
	if _, err := os.Stat(filepath); err == nil {
		fmt.Printf("directory %s for day %d already exists\n", filepath, day)
		return
	}
	os.Mkdir(filepath, 0755)
	fmt.Printf("created directory %s for day %d\n", filepath, day)
	return
}

func createDaySource(dir string, day int) {
	filepath := fmt.Sprintf(baseSourcePath, dir, day)
	if _, err := os.Stat(filepath); err == nil {
		fmt.Printf("source file %s for day %d already exists\n", filepath, day)
		return
	}
}

func createDayTestSource(dir string, day int) {
	filepath := fmt.Sprintf(baseTestPath, dir, day)
	if _, err := os.Stat(filepath); err == nil {
		fmt.Printf("test source file %s for day %d already exists\n", filepath, day)
		return
	}
}
