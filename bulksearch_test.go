package bulksearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSearchFor(t *testing.T) {
	Convey("when the term is only once in the specific file", t, func() {
		got := SearchFor("search-term", "./test-dir/test-file.js")

		expected := 1

		So(got, ShouldResemble, expected)
	})

	Convey("when the term is multiple in the dir", t, func() {
		got := SearchFor("search-term", "./test-dir/")

		expected := 2

		So(got, ShouldResemble, expected)
	})

	Convey("when the term is multiple in the dir and file", t, func() {
		got := SearchFor("ok", "./test-dir/")

		//  should only list the amount of files the term appears in
		// not the number of appearances
		expected := 2

		So(got, ShouldResemble, expected)
	})

	Convey("when the term is in only one of the files in the dir", t, func() {
		got := SearchFor("other-search-term", "./test-dir/")

		expected := 1

		So(got, ShouldResemble, expected)
	})

	Convey("when the term is nowhere in the dir", t, func() {
		got := SearchFor("asdf", "./test-dir/")

		expected := 0

		So(got, ShouldResemble, expected)
	})
}

// func TestSearch(t *testing.T) {
// 	Convey("when the search term is in a file in the dir", t, func() {
// 		got := Search([]string{`search-term`, "not-found"}, "./testdir/test-file.js")

// 		expected := map[string]int{}
// 		expected["search-term"] = 1

// 		// doesn't get the rule or the value
// 		So(got, ShouldResemble, expected)
// 	})
// }

