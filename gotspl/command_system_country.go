/*
 * Copyright 2020 Anton Globa
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
 */
package gotspl

import "bytes"

const (
	COUNTRY_NAME                                  = "COUNTRY"
	COUNTRY_USA                    CountryCommand = "001"
	COUNTRY_CANADIAN_FRENCH        CountryCommand = "002"
	COUNTRY_SPANISH_LATHIN_AMERICA CountryCommand = "003"
	COUNTRY_DUTCH                  CountryCommand = "031"
	COUNTRY_BELGIAN                CountryCommand = "032"
	COUNTRY_FRENCH_FRANCE          CountryCommand = "033"
	COUNTRY_SPANISH_SPAIN          CountryCommand = "034"
	COUNTRY_HUNGARIAN              CountryCommand = "036"
	COUNTRY_YUGOSLAVIAN            CountryCommand = "038"
	COUNTRY_ITALIAN                CountryCommand = "039"
	COUNTRY_SWITZERLAND            CountryCommand = "041"
	COUNTRY_SLOVAK                 CountryCommand = "042"
	COUNTRY_UNITED_KINGDOM         CountryCommand = "044"
	COUNTRY_DANISH                 CountryCommand = "045"
	COUNTRY_SWEDISH                CountryCommand = "046"
	COUNTRY_NORWEGIAN              CountryCommand = "047"
	COUNTRY_POLISH                 CountryCommand = "048"
	COUNTRY_GERMAN                 CountryCommand = "049"
	COUNTRY_BRAZIL                 CountryCommand = "055"
	COUNTRY_ENGLISH_INTERNATIONAL  CountryCommand = "061"
	COUNTRY_PORTUGUESE             CountryCommand = "351"
	COUNTRY_FINNISH                CountryCommand = "358"
)

type CountryCommand string

func (cc CountryCommand) GetMessage() ([]byte, error) {

	buf := bytes.NewBufferString(COUNTRY_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(string(cc))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}
