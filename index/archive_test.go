package index

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	blob1 string = `{
	"name" : "song one",
	"artist" : "artist one"
}`
	blob2 string = `{
	"name" : "song two",
	"artist" : "artist two"
}`
	blob3 string = `{
	"name" : "song three",
	"artist" : "artist three"
}`
)

func TestMain(t *testing.T) {
	Convey("it adds json objects to the index", t, func() {
		var key string
		key, err := AddMetadata(blob1)
		if err != nil {
			panic(err)
		}
		So(key, ShouldResemble, "I don't know the key")
	})
}
