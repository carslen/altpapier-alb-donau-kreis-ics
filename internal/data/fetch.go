/*
 * Copyright 2023 Carsten Lenz
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
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/carslen/altpapier-alb-donau-kreis-ics/internal/globals"
)

var (
	year     = time.Now().Year()
	nextYear = year + 1
)

// Fetch uses http.Get to retrieve PDFs with collection dates for each municipal defined in globals.Municipals and stores
// it in `test/data` location.
func Fetch(filePath string) {
	stringYear := strconv.Itoa(year)
	stringNextYear := strconv.Itoa(nextYear)

	for _, municipal := range globals.Municipals() {

		if nextYearAvailable() {
			fullURL := globals.BaseUrl + globals.BaseName + stringNextYear + "-" + municipal + globals.FileExtension
			fileName, _ := strings.CutPrefix(fullURL, globals.BaseUrl)
			fmt.Printf("URL:\t%s\nFile:\t%s\n", fullURL, fileName)
			getFile(fullURL, filePath+fileName)
		}

		fullURL := globals.BaseUrl + globals.BaseName + stringYear + "-" + municipal + globals.FileExtension
		fileName, _ := strings.CutPrefix(fullURL, globals.BaseUrl)
		fmt.Printf("URL:\t%s\nFile:\t%s\n", fullURL, fileName)
		getFile(fullURL, filePath+fileName)
	}
}

func getFile(url string, fileName string) {
	resp, _ := http.Get(url)
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode == 200 {
		file, err := os.Create(fileName)

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
		log.Fatalf("HTTP Status:\t%d\nURL:\t%s", resp.StatusCode, url)
	}

}

func nextYearAvailable() bool {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	municipals := globals.Municipals()
	municipal := municipals[r.Intn(len(municipals))]

	fullURL := globals.BaseUrl + globals.BaseName + strconv.Itoa(nextYear) + "-" + municipal + globals.FileExtension
	fmt.Printf("\navailable:\t%s\n", fullURL)

	resp, _ := http.Get(fullURL)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return false
	}
	return true
}
