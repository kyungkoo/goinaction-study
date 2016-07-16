package search

import "fmt"

// defaultMatcher 는 기본 검색기를 구현하는데 사용된다.
type defaultMatcher struct{}

// init 함수에서는 기본 검색기를 프로그램에 등록한다.
func init() {
	fmt.Println("defaultMatcher init() called!")
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 함수는 기본 검색기의 동작을 구현한다.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
