package main

import "testing"

// Unit Test
// 모든 articles를 fetch하는 Test Function
// makes sure that the article list fetched by this function and the article list present in the global variable articleList are identical
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// fetch된 articles의 개수가 global variable의 articles의 개수와 동일한지 확인
	if len(alist) != len(articleList) {
		t.Fail()
	}
	// 모든 요소가 동일한지 확인
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}

}
