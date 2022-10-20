package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Render(t *testing.T) {

	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	test_renderer.Renderer = "go"
	test_renderer.RootPath = "./tests"

	err = test_renderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error(err)
	}

	test_renderer.Renderer = "jet"

	err = test_renderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error(err)
	}

}
