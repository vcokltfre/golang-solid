package frontend

import (
	"embed"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed data
var data embed.FS

var types = map[string]string{
	".js":  "application/javascript",
	".css": "text/css",
	".png": "image/png",
}

func endsAny(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}

	return false
}

func WebMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if strings.HasPrefix(path, "/api") {
			return next(c)
		}

		if endsAny(path, ".js", ".css", ".png") {
			data, err := data.ReadFile("data/dist" + path)
			if err != nil {
				if os.IsNotExist(err) {
					return c.NoContent(404)
				} else {
					return err
				}
			}

			ext := filepath.Ext(path)

			return c.Blob(200, types[ext], data)
		}

		data, err := data.ReadFile("data/dist/index.html")
		if err != nil {
			return err
		}

		return c.HTMLBlob(200, data)
	}
}
