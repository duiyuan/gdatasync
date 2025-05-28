package util

import (
	"crypto/md5"
	"encoding/hex"
	"path/filepath"
	"strings"
)

type Predicate[T comparable] func(T) bool

// FixWindowPath return Window-Sytle Path, eg: C://a/b/c instead of /C/a/b/c
func FixWindowsPath(p string) string {
	if strings.HasPrefix(p, "/") && len(p) >= 3 && p[2] == '/' {
		drive := strings.ToUpper(string(p[1]))
		rest := p[2:]
		return drive + ":/" + rest
	}

	return filepath.Clean(p)
}

func FindIndex[T ~[]E, E comparable](slice T, fn Predicate[E]) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

// CombinedMD5 combine md5 encoded string, output it with hex encoded
func CombinedMD5(parts ...string) string {
	hash := md5.New()

	for _, part := range parts {
		hash.Write([]byte(part))
	}
	return hex.EncodeToString(hash.Sum(nil))
}
