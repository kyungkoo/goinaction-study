package search

// Result 구조체는 검색 결과를 저장한다.
type Result struct {
	Field   string
	Content string
}

// Matcher 인터페이스는 새로운 검색 타입을 구현할 때 필요한 동작을 정의한다.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 함수는 고루틴으로서 호출되며 개별 피드 타입에 대한 검색을 동시에 수행한다.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {

}

// Display 함수는 개별 고루틴이 전달한 검색 결과를 콘솔 창에 출력한다.
func Display(results chan *Result) {

}

// Register 는 프로그램에서 사용할 검색기를 등록할 함수를 정의한다.
func Register(feedType string, matcher Matcher) {

}
