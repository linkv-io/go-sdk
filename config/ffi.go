package config

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"plugin"
	"runtime"
)

const DownloadURL = "http://dl.linkv.fun/static/server"

func platformFile(name string) string {
	return fmt.Sprintf("lib%s%s", name, ext)
}

func dlopenPlatformSpecific(name, path string) (*plugin.Plugin, error) {
	if path == "" {
		path = os.TempDir()
	}
	return plugin.Open(path + platformFile(name))
}

func download(name, path, version string) (bool, error) {
	if path == "" {
		path = os.TempDir()
	}
	ok, err := fileExists(path + platformFile(name))
	if err != nil {
		return false, err
	}
	if ok {
		return true, nil
	}

	resp, err := http.Get(DownloadURL + "/" + version + "/" + runtime.Version() + "/" + platformFile(name))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	out, err := os.Create(path + platformFile(name))
	if err != nil {
		return false, err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return false, err
	}

	return true, nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
