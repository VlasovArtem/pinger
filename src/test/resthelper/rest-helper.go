package resthelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestRequest struct {
	Method     string
	URL        string
	Body       any
	Vars       map[string]string
	Parameters map[string]string
	Handler    http.HandlerFunc
	Request    *http.Request
	Recorder   *httptest.ResponseRecorder
	build      bool
}

func NewTestRequest() *TestRequest {
	return &TestRequest{
		Recorder:   httptest.NewRecorder(),
		Vars:       make(map[string]string),
		Parameters: make(map[string]string),
	}
}

type TestRequestBuilder interface {
	WithMethod(method string) *TestRequest
	WithURL(target string) *TestRequest
	WithBody(body any) *TestRequest
	WithHandler(handler http.HandlerFunc) *TestRequest
	WithVar(key string, value string) *TestRequest
	WithParameter(key string, value string) *TestRequest
	Build() *TestRequest
}

func (t *TestRequest) WithVar(key string, value string) *TestRequest {
	t.Vars[key] = value

	return t
}

func (t *TestRequest) WithMethod(method string) *TestRequest {
	t.Method = method
	return t
}
func (t *TestRequest) WithURL(URL string) *TestRequest {
	t.URL = URL
	return t
}

func (t *TestRequest) WithBody(body any) *TestRequest {
	t.Body = body
	return t
}

func (t *TestRequest) WithHandler(handler http.HandlerFunc) *TestRequest {
	t.Handler = handler
	return t
}

func (t *TestRequest) WithParameter(key string, value string) *TestRequest {
	t.Parameters[key] = value
	return t
}

func (t *TestRequest) Build() *TestRequest {
	body, _ := json.Marshal(t.Body)

	buffer := bytes.Buffer{}

	buffer.Write(body)

	t.Request = httptest.NewRequest(t.Method, t.mapURL(), &buffer)

	if len(t.Vars) != 0 {
		t.Request = mux.SetURLVars(t.Request, t.Vars)
	}

	t.build = true

	return t
}

func (t *TestRequest) execute() {
	t.Handler(t.Recorder, t.Request)
}

func (t *TestRequest) Verify(test *testing.T, expectedStatusCode int) []byte {
	if !t.build {
		t.Build()
	}

	t.execute()

	response := t.Recorder.Result()

	assert.Equal(test, expectedStatusCode, response.StatusCode, fmt.Sprintf("Response status code should be %d but was %d", expectedStatusCode, response.StatusCode))

	return ReadBytes(response)
}

func (t *TestRequest) mapURL() string {
	var newURL = t.URL
	for name, value := range t.Parameters {
		newURL = strings.Replace(newURL, "{"+name+"}", value, -1)
	}

	return newURL
}

func ReadBytes(response *http.Response) []byte {
	buffer := bytes.Buffer{}

	buffer.ReadFrom(response.Body)

	return buffer.Bytes()
}
