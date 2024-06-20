// +build linux
package webview

import (
	webview "github.com/webview/webview_go"
)

func New(debug bool) webview.WebView {
	return webview.New(debug)
}
