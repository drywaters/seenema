package layout

import (
	"net/url"
	"strings"
)

var assetsVersion string

// SetAssetsVersion sets the version string appended to static asset URLs.
func SetAssetsVersion(version string) {
	assetsVersion = version
}

func AssetURL(path string) string {
	if assetsVersion == "" {
		return path
	}

	separator := "?"
	if strings.Contains(path, "?") {
		separator = "&"
	}

	return path + separator + "v=" + url.QueryEscape(assetsVersion)
}
