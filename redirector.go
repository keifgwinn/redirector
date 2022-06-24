package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("unable to parse " + key))
	}
	return val
}

func getIntEnv(key string) int {
	val := getStrEnv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to convert " + key))
	}
	return ret
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	url := getStrEnv("REDIRECT_URL")
	code := getStrEnv("REDIRECT_CODE")
	urlpattern := `[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)` // must be recognisable url
	codepattern := `3[0-9][0-9]$`                                                                                  // 300 to 399 valid redirect code
	urlmatched, urlerr := regexp.Match(urlpattern, []byte(url))
	check(urlerr)
	if urlmatched {
		fmt.Printf("√ '%s' is a valid url\n", url)
	} else {
		fmt.Printf("X '%s' is not a valid url\n", url)
		log.Fatal(fmt.Sprintf("Invalid configuration."))
	}

	codematched, codeerr := regexp.Match(codepattern, []byte(code))
	check(codeerr)
	if codematched {
		fmt.Printf("√ '%s' is a valid code\n", code)
	} else {
		fmt.Printf("X '%s' is not a valid code\n", code)
		log.Fatal(fmt.Sprintf("Invalid configuration."))
	}

	ret, codestrerr := strconv.Atoi(code)
	if codestrerr != nil {
		log.Fatal("Could not convert code", codestrerr)
	}

	http.Handle("/", http.RedirectHandler(url, ret))
	httperr := http.ListenAndServe(":8080", nil)
	if httperr != nil {
		log.Fatal("ListenAndServe: ", httperr)
	}
}
