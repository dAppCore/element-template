// SPDX-License-Identifier: EUPL-1.2

package main

import (
	"io/fs"
	"net/http"
)

// noCache wraps a filesystem with cache-busting headers for development.
func noCache(fsys fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(fsys))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		fileServer.ServeHTTP(w, r)
	})
}
