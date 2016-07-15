package search

// Feed 를 처리할 정보를 표현하는 구조체
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 함수는 피드 데이터 파일을 읽어 구조체로 변환한다.
func RetrieveFeeds() ([]*Feed, error) {
	return nil, nil
}
