package main

import (
	"errors"
	"regexp"
)

type URLNode struct {
	next  *URLNode
	value string
}

type URLList struct {
	head *URLNode
	end  *URLNode
}

func (URLList *URLList) Add(url string) {
	newNode := &URLNode{value: url}
	if URLList.head == nil {
		URLList.head = newNode
		URLList.end = newNode
	} else {
		URLList.end.next = newNode
		URLList.end = newNode
	}
}

func (URLList *URLList) AddToURLListFromHtml(html string) {
	re := regexp.MustCompile(`(?:href|src)\s*=\s*["'](https?://[^"']+)["']`)
	matches := re.FindAllStringSubmatch(html, -1)

	for _, value := range matches {
		URLList.Add(value[1])
	}
}

func (URLList *URLList) PopFront() string {
	if URLList.head == nil {
		return ""
	}
	currentValue := URLList.head.value
	URLList.head = URLList.head.next
	if URLList.head == nil {
		URLList.end = nil
	}
	return currentValue
}

func newURLList(URL string) *URLList {
	node := &URLNode{
		next:  nil,
		value: URL,
	}
	return &URLList{
		head: node,
		end:  node,
	}
}

func URLListFromArray(data []string) (*URLList, error) {
	if len(data) == 0 {
		// Return error
		return nil, errors.New("empty input array")
	}

	head := newURLList(data[0])

	for _, v := range data[1:] {
		head.Add(v)
	}

	return head, nil
}
