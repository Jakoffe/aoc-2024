package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const YEAR = 2024
const SESSION_ID_FILE = "session.cookie"

var SESSION = string(ReadFile(SESSION_ID_FILE))

var HEADERS = map[string]string{
	"User-Agent": "github.com/tomfran/advent-of-code-setup reddit:u/fran-sch, discord:@tomfran#5786",
}
var COOKIES = map[string]string{
	"session": SESSION,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	check(err)
	return content
}

func getUrl(year int, day int) string {
	return fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", strconv.Itoa(year), strconv.Itoa(day))
}

func GetInput(day int) string {
	path := fmt.Sprintf("inputs/%d", day)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {

		client := &http.Client{}
		url := getUrl(YEAR, day)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Cookie", fmt.Sprintf("session=%s", SESSION))
		res, err := client.Do(req)

		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}

		if res.StatusCode != 200 {
			fmt.Printf("error making http request.. Got statuscode %d\n", res.StatusCode)
			os.Exit(1)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		os.WriteFile(path, body, 0644)
	}

	dat, err := os.ReadFile(path)
	check(err)
	return string(dat)
}
