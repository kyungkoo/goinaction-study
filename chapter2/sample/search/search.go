package search

import (
	"fmt"
	"log"
	"sync"
)

// 검색을 처리할 검색기으 매핑 정보를 저장할 맵
var matchers = make(map[string]Matcher)

// Run 함수는 검색 로직을 수행한다.
func Run(searchTerm string) {
	// 검색할 피드의 목록을 조회한다.
	feeds, err := RetrieveFeeds()
	for err != nil {
		log.Fatal(err)
	}

	// 버퍼가 없는 채널을 생성하여 화면에 표시할 검색 결과를 전달받는다.
	results := make(chan *Result)

	// 모든 피드를 처리할 때까지 기다릴 대기 그룹을 설정한다.
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대기해야 할 고루틴의 개수를 설정한다.
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["defaults"]
		}

		// 고루틴 생성
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 모니터링 하는 고루틴
	go func() {
		fmt.Println("모니터링 하는 고루틴 실행")
		waitGroup.Wait()
		fmt.Println("WaitGroup.Wait() 끝!")
		close(results)
	}()

	Display(results)
}

// Register 는 프로그램에서 사용할 검색기를 등록할 함수를 정의한다.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다.")
	}

	log.Println("등록완료:", feedType, " 검색기")
	matchers[feedType] = matcher
}
