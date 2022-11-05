package render

import (
	"os"
	"testing"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./tests/views"),
	jet.InDevelopmentMode(),
)

var test_renderer = Render{
	Renderer: "",
	RootPath: "",
	JetViews: views,
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
