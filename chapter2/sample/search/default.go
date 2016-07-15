package search

// defaultMatcher 는 기본 검색기를 구현하는데 사용된다.
type defaultMatcher struct{}

// Search 함수는 기본 검색기의 동작을 구현한다.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
