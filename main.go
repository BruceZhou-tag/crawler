package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Fetch(url string) ([]rune, error) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("read content failed:%v\n", err)
		err = errors.New("read content failed!")
		return nil,err
	}

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links!\n", numLinks)

	return Decodeutf8(body),nil
}

func Decodeutf8(b []byte) []rune {
	u8 := []rune(string(b))
	return u8
}

func main() {
	url := "https://www.thepaper.cn/"
	content, err := Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(content))
}