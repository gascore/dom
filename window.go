package dom

import (
	"strings"

	"github.com/gascore/dom/js"
	sjs "syscall/js"
)

func GetWindow() *Window {
	win := js.Get("window")
	if !win.Valid() {
		return nil
	}
	return &Window{v: win}
}

var _ EventTarget = (*Window)(nil)

type Window struct {
	v js.Value
}

func (w *Window) JSValue() js.Ref {
	return w.v.JSValue()
}

func (w *Window) AddEventListener(typ string, h EventHandler) {
	w.v.Call("addEventListener", typ, js.NewEventCallback(func(v js.Value) {
		h(ConvertEvent(v))
	}))
}

func (w *Window) AddPassiveEventListener(typ string, h EventHandler) {
	obj := js.NewObject()
	obj.Set("passive", true)
	obj.Set("capture", false)

	w.v.Call("addEventListener", typ, js.NewEventCallback(func(v js.Value) {
		h(ConvertEvent(v))
	}), obj)
}

func (w *Window) RemoveEventListener(typ string, h EventHandler) {
	w.v.Call("removeEventListener", typ, js.NewEventCallback(func(v js.Value) {
		h(ConvertEvent(v))
	}))
}

func (w *Window) RemovePassiveEventListener(typ string, h EventHandler) {
	obj := js.NewObject()
	obj.Set("passive", true)
	obj.Set("capture", false)

	w.v.Call("removeEventListener", typ, js.NewEventCallback(func(v js.Value) {
		h(ConvertEvent(v))
	}), obj)
}

func (w *Window) DispatchEvent(value js.Value) {
	w.JSValue().Call("dispatchEvent", value.JSValue())
}

func (w *Window) Open(url, windowName string, windowFeatures map[string]string) {
	w.v.Call("open", url, windowName, joinKeyValuePairs(windowFeatures, ","))
}

func (w *Window) SetLocation(url string) {
	w.v.Set("location", url)
}

func (w *Window) OnResize(fnc func(e Event)) {
	w.AddEventListener("resize", fnc)
}

func (w *Window) LocalStorage() sjs.Value {
	return w.JSValue().Get("localStorage")
}

func (w *Window) RequestAnimationFrame(h func(timeStep js.Value)) {
	w.v.Call("requestAnimationFrame", js.NewEventCallback(h))
}

func (w *Window) GetHistory() sjs.Value {
	return w.JSValue().Get("history")
}

func (w *Window) GetLocation() sjs.Value {
	return w.JSValue().Get("location")
}

func (w *Window) GetLocationHash() string {
	return w.GetLocation().Get("hash").String()
}

func (w *Window) GetLocationPath() string {
	return w.GetLocation().Get("pathname").String()
}

func joinKeyValuePairs(kvpair map[string]string, joiner string) string {
	if kvpair == nil {
		return ""
	}

	pairs := make([]string, 0, len(kvpair))
	for k, v := range kvpair {
		pairs = append(pairs, k+"="+v)
	}
	return strings.Join(pairs, joiner)
}
