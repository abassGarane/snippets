package main

import (
	"html"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/abassGarane/snippet/pkg/models/mock"
	"github.com/golangcollege/sessions"
)

func NewTestApplication(t *testing.T) *application {
	templateCache, err := newTemplateCache("./../../ui")
	if err != nil {
		t.Fatal(err)
	}
	session := sessions.New([]byte("xddvasuJAaxusafvyZXH2385328"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	return &application{
		errorLog:      log.New(io.Discard, "", 0),
		infoLog:       log.New(io.Discard, "", 0),
		session:       session,
		snippets:      &mock.SnippetModel{},
		users:         &mock.UserModel{},
		templateCache: templateCache,
	}
}

type TestServer struct {
	*httptest.Server
}

func NewTestServer(t *testing.T, h http.Handler) *TestServer {
	ts := httptest.NewTLSServer(h)
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	//automatically store cookies
	ts.Client().Jar = jar

	//Disable redirects
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &TestServer{ts}
}

func (ts *TestServer) Get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	return rs.StatusCode, rs.Header, body
}

var CSRFTokenRX = regexp.MustCompile(`<input type="hidden" name="csrf_token" value="{{.CSRFToken}}">`)

func ExtractCSRFToken(t *testing.T, body []byte) string {
	matches := CSRFTokenRX.FindSubmatch(body)
	if len(matches) < 2 {
		t.Fatal("no csrf token in body")
	}
	return html.UnescapeString(string(matches[1]))
}
