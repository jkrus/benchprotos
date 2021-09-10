// Copyright © 2020-2021 The EVEN Solutions Developers Team

package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	AppDefaultHost     = "localhost"
	AppDefaultPortGRPC = "4317"
	AppDefaultPortHTTP = "8080"
	DstDefaultPathName = "storage"

	appRootPathName = "swcapital"
)

// OsAppRootPath fetches a path for the root of the app based on the system.
func OsAppRootPath() string {
	path := appRootPathName
	dir, _ := os.UserConfigDir()
	switch runtime.GOOS {
	case "windows", "darwin":
		path = strings.ToUpper(path[:1]) + path[1:]
	}

	return filepath.Join(dir, path)
}
