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
	"strconv"
)

const (
	SOUND_NAME         = "SOUND"
	SOUND_LEVEL_MIN    = 0
	SOUND_LEVEL_MAX    = 9
	SOUND_INTERVAL_MIN = 1
	SOUND_INTERVAL_MAX = 4095
)

type SoundImpl struct {
	volumeLevel    *int
	timingInterval *int
}

type SoundBuilder interface {
	TSPLCommand
	VolumeLevel(volumeLevel int) SoundBuilder
	TimingInterval(timingInterval int) SoundBuilder
}

func SoundCmd() SoundBuilder {
	return SoundImpl{}
}

func (s SoundImpl) GetMessage() ([]byte, error) {
	if s.timingInterval == nil || s.volumeLevel == nil {
		return nil, errors.New("ParseError SOUND Command: timingInterval and volumeLevel should be specified")
	}

	if !(*s.volumeLevel > SOUND_LEVEL_MIN && *s.volumeLevel < SOUND_LEVEL_MAX) {
		return nil, errors.New("ParseError SOUND Command: volumeLevel the parameter value must be between " +
			strconv.Itoa(SOUND_LEVEL_MIN) + " and " + strconv.Itoa(SOUND_LEVEL_MAX))
	}

	if !(*s.timingInterval > SOUND_INTERVAL_MIN && *s.timingInterval < SOUND_INTERVAL_MAX) {
		return nil, errors.New("ParseError SOUND Command: timingInterval the parameter value must be between " +
			strconv.Itoa(SOUND_INTERVAL_MIN) + " and " + strconv.Itoa(SOUND_INTERVAL_MAX))
	}

	buf := bytes.NewBufferString(SOUND_NAME)

	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*s.volumeLevel))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*s.timingInterval))
	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}

func (s SoundImpl) VolumeLevel(volumeLevel int) SoundBuilder {
	if s.volumeLevel == nil {
		s.volumeLevel = new(int)
	}
	*s.volumeLevel = volumeLevel
	return s
}

func (s SoundImpl) TimingInterval(timingInterval int) SoundBuilder {
	if s.timingInterval == nil {
		s.timingInterval = new(int)
	}
	*s.timingInterval = timingInterval
	return s
}
