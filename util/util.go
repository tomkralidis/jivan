///////////////////////////////////////////////////////////////////////////////
//
// The MIT License (MIT)
// Copyright (c) 2018 Jivan Amara
// Copyright (c) 2018 Tom Kralidis
// Copyright (c) 2018 James Lucktaylor
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.
//
///////////////////////////////////////////////////////////////////////////////

package util

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
)

// DefaultGpkg is used when a data source isn't provided.
// Scans for files w/ '.gpkg' extension following logic described below, and chooses the first alphabetically.
// Returns the absolute path as a string.
//
// 1) scan the current working directory
// 2) if none found in the working directory and it exists, scan ./data/
// 3) if none found in either of the previous and it exists, scan ./test_data/
// 4) if none found, return an empty string
func DefaultGpkg() string {
	// Directories to check in decreasing order of priority
	dirs := []string{"", "data", "test_data"}
	gpkgPath := ""
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Unable to get working directory: %v", err)
		return ""
	}

	for _, dir := range dirs {
		searchGlob := path.Join(dir, "*.gpkg")
		gpkgFiles, err := filepath.Glob(searchGlob)
		if err != nil {
			panic("Invalid glob pattern hardcoded")
		}
		if len(gpkgFiles) == 0 {
			continue
		}
		sort.Strings(gpkgFiles)
		gpkgFilename := gpkgFiles[0]
		gpkgPath = path.Clean(path.Join(wd, gpkgFilename))
		break
	}

	if len(gpkgPath) == 0 {
		return ""
	} else {
		return filepath.Clean(gpkgPath)
	}
}

func RenderTemplate(templateString string, data map[string]interface{}) ([]byte, error) {
	var tpl bytes.Buffer
	t := template.New("template")
	t, _ = t.Parse(templateString)

	if err := t.Execute(&tpl, data); err != nil {
		return tpl.Bytes(), err
	}

	return tpl.Bytes(), nil
}
