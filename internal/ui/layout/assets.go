package layout

import (
	"net/url"
	"strings"
	"sync/atomic"
)

var assetsVersion atomic.Value

// SetAssetsVersion sets the version string appended to static asset URLs.
func SetAssetsVersion(version string) {
	assetsVersion.Store(version)
}

func AssetURL(path string) string {
	version, _ := assetsVersion.Load().(string)
	if version == "" {
		return path
	}

	separator := "?"
	if strings.Contains(path, "?") {
		separator = "&"
	}

	return path + separator + "v=" + url.QueryEscape(version)
}
