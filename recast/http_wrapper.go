package recast

import (
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type httpWrapper struct {
	inner *gorequest.SuperAgent
}

func newHttpWrapper() *httpWrapper {
	return &httpWrapper{gorequest.New()}
}

func (w *httpWrapper) Post(url string) *httpWrapper {
	w.inner = w.inner.Post(url)
	return w
}

func (w *httpWrapper) Get(url string) *httpWrapper {
	w.inner = w.inner.Get(url)
	return w
}

func (w *httpWrapper) Delete(url string) *httpWrapper {
	w.inner = w.inner.Delete(url)
	return w
}

func (w *httpWrapper) Put(url string) *httpWrapper {
	w.inner = w.inner.Put(url)
	return w
}

func (w *httpWrapper) Send(data interface{}) *httpWrapper {
	w.inner = w.inner.Send(data)
	return w
}

func (w *httpWrapper) Set(key, value string) *httpWrapper {
	w.inner = w.inner.Set(key, value)
	return w
}

func (w *httpWrapper) EndStruct(v interface{}) (*http.Response, []byte, []error) {
	return w.inner.EndStruct(v)
}

func (w *httpWrapper) Type(typeStr string) *httpWrapper {

	w.inner = w.inner.Type(typeStr)
	return w
}

func (w *httpWrapper) SendFile(file interface{}, args ...string) *httpWrapper {

	w.inner = w.inner.SendFile(file, args...)
	return w
}

func (w *httpWrapper) End(callback ...func(response gorequest.Response, body string, errs []error)) (*http.Response, string, []error) {
	return w.inner.End(callback...)
}
