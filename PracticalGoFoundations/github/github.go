package main

import (
	"encoding/json"
	"fmt"

	// "io"
	"log"
	"net/http"
	"net/url"
	// "os"
)

func main() {
	fmt.Println(githubInfo("sarthakrawat-1"))
}

func demo() {
	resp, err := http.Get("https://api.github.com/users/sarthakrawat-1")

	if err != nil {
		log.Fatalf("error: %s", err)
		/*
			Equivalent to
			log.Printf("error: %s", err)
			os.Exit(1)
		*/
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("Content-Length: %d\n", resp.ContentLength)
	fmt.Printf("Protocol: %s\n", resp.Proto)

	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: %s", err)
	// }

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", r)
}

// Returns name and no of public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)

	if err != nil {
		return "", 0, nil
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	var r Reply

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.Public_Repos, nil
}

type Reply struct {
	Name string
	Public_Repos int
}

/*
	JSON <-> Go
	true/false <-> true/false
	string <-> string
	null <-> nil
	number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...const
	array <-> []any ([]interface{})
	object <-> map[string]any, struct

	encoding/json API
	JSON -> io.Reader -> Go : json.Decoder
	JSON -> []byte -> Go : json.Unmarshal
	Go -> io.Writer -> JSON : json.Encoder
	Go -> []byte -> JSON : json.Marshal
*/