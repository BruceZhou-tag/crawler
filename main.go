package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Fetch(url string) ([]byte, error) {
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

	return body,nil
	// return Decodeutf8(body),nil
}

// func Decodeutf8(b []byte) []rune {
// 	u8 := []rune(string(b))
// 	return u8
// }

func main() {
	url := "https://www.thepaper.cn/"
	
	
	body, err := Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("read content failed: %v", err)
	}

	doc.Find("div.ant-card div.ant-card-body a[target=_blank] h2").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

/* Result:css选择器可能有问题，只匹配出了4条内容
Review 0: 驻英国使馆提醒我公民注意英方出入境检查，敦促英方停止任意扣留盘查
Review 1: 众声丨AI“苦命男孩小海”收割老人，平台审核不能总是被动响应
Review 2: 四川警方披露新型智能化非法狩猎：涉案者私改无人机，高空坠箭猎杀野生动物
Review 3: 女性是否应服兵役？韩国兵源缩减，在野党提“性别平等征兵”
*/