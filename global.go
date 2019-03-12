 

package dom

import (
	"image"

	"github.com/gascore/dom/js"
)

var (
	Doc  = GetDocument()
	Body = Doc.GetElementsByTagName("body")[0]
	Head = Doc.GetElementsByTagName("head")[0]
)

// Value is an alias for js.Wrapper.
//
// Derprecated: use js.Wrapper
type Value = js.Wrapper

func ConsoleLog(args ...interface{}) {
	js.Get("console").Call("log", args...)
}

func ConsoleError(args ...interface{}) {
	js.Get("console").Call("error", args...)
}

func ConsoleDir(args ...interface{}) {
	js.Get("console").Call("dir", args...)
}

func ConsoleTime(args ...interface{}) {
	js.Get("console").Call("time", args...)
}

func ConsoleTimeEnd(args ...interface{}) {
	js.Get("console").Call("timeEnd", args...)
}

func Loop() {
	select {}
}

type Point = image.Point
type Rect = image.Rectangle
