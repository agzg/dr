package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func digiroot(n int) int {
	a := strconv.Itoa(n)

	for n > 9 {
		n = 0
		for _, char := range a {
			i, _ := strconv.Atoi(string(char))
			n += i
		}
		a = strconv.Itoa(n)
	}

	return n
}

const url = "https://api.github.com/users/%s"

type User struct {
	Id int `json:"id"`
}

func fetchGHId(name string) (int, error) {
	r, err := http.Get(fmt.Sprintf(url, name))
	if err != nil {
		return 0, err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 0, err
	}

	var u User
	if err := json.Unmarshal(b, &u); err != nil {
		return 0, err
	}

	return u.Id, nil
}

func main() {
	id, err := fetchGHId("Abdul-Muiz-Iqbal")
	if err != nil {
		panic(err)
	}
	fmt.Println(digiroot(id))
}
