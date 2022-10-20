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

	err = test_renderer.Page(w, r, "not-exist", nil, nil)

	if err == nil {
		t.Error("Error rendering non-existent Go template", err)
	}

	test_renderer.Renderer = "jet"

	err = test_renderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error(err)
	}

	err = test_renderer.Page(w, r, "not-exist", nil, nil)

	if err == nil {
		t.Error("Error rendering non-existent Jet template", err)
	}

	test_renderer.Renderer = "some other value"

	err = test_renderer.Page(w, r, "home", nil, nil)

	if err == nil {
		t.Error(err)
	}

}

func TestRender_GoPage(t *testing.T) {

	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		t.Error(err)
	}

	test_renderer.Renderer = "go"
	test_renderer.RootPath = "./tests"

	err = test_renderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestRender_JetPage(t *testing.T) {

	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		t.Error(err)
	}

	test_renderer.Renderer = "jet"
	test_renderer.RootPath = "./tests"

	err = test_renderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error(err)
	}
}
