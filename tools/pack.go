package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var s strings.Builder
	s.WriteString("package main\n\nvar files = map[string]string{\n")
	for _, n := range []string{"index.html", "manifest.webmanifest", "sw.js", "icon-192.png", "icon-512.png"} {
		a, err := os.ReadFile(filepath.Join("web", n))
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(&s, "\t%q: %q,\n", n, string(a))
	}
	s.WriteString("}\n")
	if err := os.WriteFile(filepath.Join("windows", "assets.go"), []byte(s.String()), 0644); err != nil {
		panic(err)
	}
}
