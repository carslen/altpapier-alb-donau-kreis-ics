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

package globals

// Municipals returns a slice of strings with Municipals in which Gebr. Braig Ehingen collects waste paper. The strings
// are used to concatenate the download URLs in data.Fetch.
func Municipals() []string {
	return []string{
		"allmendingen--altheim",
		"berghuelen",
		"blaubeuren",
		"blaustein",
		"dornstadt",
		"ehingen",
		"erbach",
		"heroldstatt--merklingen",
		"laichingen",
		"lauterach-untermarchtal-hausen-unterwachingen",
		"munderkingen",
		"nellingen",
		"griesingen--oberdischingen",
		"obermarchtal-rechtenstein-emeringen",
		"oberstadion-unterstadion-grundsheim",
		"oepfingen",
		"rottenacker--emerkingen",
		"schelklingen",
		"ulm",
		"westerheim",
		"westerstetten",
	}
}
