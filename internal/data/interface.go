/*
 * Copyright 2024 Carsten Lenz
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package data

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/carslen/altpapier-alb-donau-kreis-ics/internal/globals"
)

type Data interface {
	// CheckAvailability returns a bool and tests if next year's collection dates are available.
	// CheckAvailability expects the municipal name as string as input.
	CheckAvailability(municipal string) bool

	// Get downloads the PDF with the collection dates for a municipal. Get expects fullURL, FilePath and fileName as string.
	// Default for filePath is `test/data` if filePath is empty sting.
	Get(fullURL string, filePath string, fileName string)

	// Metadata returns download URL and file name used for Get. Metadata expects municipal as string and year as int
	// (use CurrentYear() or NextYear()).
	Metadata(municipal string, year int) (string, string)

	// Parse reads the PDF and returns is as string
	Parse() string
}

// PDF is a representation of a PDF file.
type PDF struct {
}

// CheckAvailability returns a bool and tests if next year's collection dates are available.
// CheckAvailability expects the municipal name as string as input.
func (P PDF) CheckAvailability(municipal string) bool {
	nextYear := time.Now().Year() + 1

	resp, err := http.Head(globals.BaseUrl + globals.BaseName + strconv.Itoa(nextYear) + municipal + globals.FileExtension)
	if err != nil {
		return false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode == 200 {
		return true
	}
	return false
}

// Get downloads the PDF with the collection dates for a municipal. Get expects fullURL, FilePath and fileName as string.
// Default for filePath is `test/data` if filePath is empty sting.
func (P PDF) Get(fullURL string, filePath string, fileName string) {
	if filePath == "" {
		filePath = globals.BasePath
	}

	resp, _ := http.Get(fullURL)
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode == 200 {
		file, err := os.Create(path.Join(filePath, fileName))

		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalf("HTTP Status:\t%d\nURL:\t%s", resp.StatusCode, fullURL)
	}
}

// Metadata returns download URL and file name used for Get. Metadata expects municipal as string and year as int
// (use CurrentYear() or NextYear()).
func (P PDF) Metadata(municipal string, year int) (string, string) {
	fullURL := globals.BaseUrl + globals.BaseName + strconv.Itoa(year) + "-" + municipal + globals.FileExtension
	fileName, _ := strings.CutPrefix(fullURL, globals.BaseUrl)
	return fullURL, fileName
}

// Parse reads the PDF and returns is as string
func (P PDF) Parse() string {
	//TODO implement me
	panic("implement me")
	return ""
}

// New returns a new PDF struct of type Data
func New() Data {
	return &PDF{}
}
