package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println(UserInfo("ardanlabs"))
}

func demo() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", resp.Status)
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type:", ctype)
	//io.Copy(os.Stdout, resp.Body)
	var reply struct {
		Login    string
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERRROR:", err)
		return
	}
	fmt.Println(reply.Login, reply.Name, reply.NumRepos)
}

func UserInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login

	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
	}

	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Login    string
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERRROR:", err)
		return "", 0, err
	}
	return reply.Name, reply.NumRepos, nil
}

/* JSON <-> Go
Types:
string <-> string ** string cannot be nil in Go
true/false <-> bool
number <-> float65, float32, int, int8 ... uint, uint8 ** defaults to float64 if not specified
array <-> []T, []any
object <-> map[string]any, struct

encoding/json API
JSON -> []byte -> Go: use Unmarshal
Go -> []byte -> JSON: use Marshall
JSON -> io.Reader -> GO: Decoder
Go -> io.Writer -> JSON: Encoder

*/
