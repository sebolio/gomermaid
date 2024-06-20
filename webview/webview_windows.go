// +build windows
package webview

import (
	webview "github.com/jchv/go-webview2"
)

func New(debug bool) webview.WebView {
	return webview.New(debug)
}

var HintNone = webview.HintNone