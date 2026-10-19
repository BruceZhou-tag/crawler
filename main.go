package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Printf("fetch url error:%v\n", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("read content failed:%v\n", err)
		return
	}

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links!\n", numLinks)
}