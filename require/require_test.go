// +build !js

package require

import (
	"github.com/gascore/dom/js/jstest"
	"net/http"
	"strings"
	"testing"
	"time"
)

var (
	modtime = time.Now()
	files   = map[string]string{
		"env.js": `Val = 'ok'`,
		"err.js": `= 'ok'`,
	}
)

func TestRequire(t *testing.T) {
	jstest.RunTestChrome(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := strings.Trim(r.URL.Path, "/")
		if !strings.HasSuffix(name, ".js") {
			name += ".js"
		}
		data, ok := files[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.ServeContent(w, r, name, modtime, strings.NewReader(data))
	}))
}
