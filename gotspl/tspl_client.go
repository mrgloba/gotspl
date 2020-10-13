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

type TSPLClient interface {
	Connect() error
	Disconnect() error
	SendData(data []byte) error
	ReadData(data []byte) (int,error)
	SendCommandSequence(commandSequence TSPLCommandSequence) error
	SendCommand(command TSPLCommand) error
	IsConnected() bool
	AddResponseListener(listener chan *RawResponseEvent)
}
