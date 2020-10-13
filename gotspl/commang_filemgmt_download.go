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

import (
	"bytes"
	"errors"
)

const (
	DOWNLOAD_NAME                                     = "DOWNLOAD"
	DOWNLOAD_STORGAE_DRAM             DownloadStorage = ""
	DOWNLOAD_STORAGE_FLASH            DownloadStorage = "F"
	DOWNLOAD_STORAGE_EXPANSION_MODULE DownloadStorage = "E"
)

type DownloadStorage string

type DownloadImpl struct {
	storage *string
	name    *string
}

type DownloadBuilder interface {
	TSPLCommand
	Storage(storage DownloadStorage) DownloadBuilder
	Name(name string) DownloadBuilder
}

func DownloadCmd() DownloadBuilder {
	return DownloadImpl{}
}

func (d DownloadImpl) GetMessage() ([]byte, error) {

	if d.name == nil || len(*d.name) == 0 {
		return nil, errors.New("ParseError DOWNLOAD Command: name should be specified")
	}

	buf := bytes.NewBufferString(DOWNLOAD_NAME)
	buf.WriteString(EMPTY_SPACE)
	if d.storage != nil {
		buf.WriteString(*d.storage)
		buf.WriteString(VALUE_SEPARATOR)
	}
	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(*d.name)
	buf.WriteString(DOUBLE_QUOTE)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (d DownloadImpl) Storage(storage DownloadStorage) DownloadBuilder {
	if d.storage == nil {
		d.storage = new(string)
	}
	*d.storage = string(storage)
	return d
}

func (d DownloadImpl) Name(name string) DownloadBuilder {
	if d.name == nil {
		d.name = new(string)
	}
	*d.name = name
	return d
}
