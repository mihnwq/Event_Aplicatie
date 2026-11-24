package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func findBrowser() string {
	for _, b := range []string{"Edge", "Chrome"} {
		var a []string
		if b == "Edge" {
			a = []string{"ProgramFiles(x86)", "ProgramFiles", "LOCALAPPDATA"}
		} else {
			a = []string{"ProgramFiles", "ProgramFiles(x86)", "LOCALAPPDATA"}
		}
		for _, key := range a {
			var p string
			if b == "Edge" {
				p = filepath.Join(os.Getenv(key), "Microsoft", "Edge", "Application", "msedge.exe")
			} else {
				p = filepath.Join(os.Getenv(key), "Google", "Chrome", "Application", "chrome.exe")
			}
			if _, err := os.Stat(p); err == nil {
				return p
			}
		}
	}
	return ""
}

func main() {
	dir := os.Getenv("LOCALAPPDATA")
	if dir == "" {
		dir = os.TempDir()
	}
	dir = filepath.Join(dir, "SideChallenge3000")
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot create application folder:", err)
		return
	}
	for _, n := range []string{"index.html", "manifest.webmanifest", "sw.js", "icon-192.png", "icon-512.png"} {
		if err := os.WriteFile(filepath.Join(dir, n), []byte(files[n]), 0644); err != nil {
			return
		}
	}
	p := filepath.ToSlash(filepath.Join(dir, "index.html"))
	u := url.URL{Scheme: "file", Path: "/" + p}
	s := u.String()
	if len(os.Args) > 1 && strings.EqualFold(os.Args[1], "pop") {
		s += "?pop=1"
	}
	b := findBrowser()
	if b != "" {
		exec.Command(b, "--app="+s).Start()
		return
	}
	exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", s).Start()
}
