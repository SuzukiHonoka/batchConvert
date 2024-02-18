package converter

import "path/filepath"

type Extension []string

var VideoExtension Extension = []string{
	".mp4",
	".avi",
	".mov",
	".mkv",
	".wmv",
	".flv",
	".mpeg",
	".mpg",
	".webm",
	".m4v",
	".3gp",
	".ts",
	".vob",
	".rmvb",
	".divx",
}

// Match check if given file name matches the specific Extensions
func (e Extension) Match(path string) bool {
	for _, ext := range VideoExtension {
		if ext == filepath.Ext(path) {
			return true
		}
	}
	return false
}
