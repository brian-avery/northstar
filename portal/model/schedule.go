/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

// Defines the supported event categories. E.g., Device
// events represents an event type generated by a device.
const (
	DeviceEvent string = "Device"
	TimerEvent  string = "Timer"
	NoneEvent   string = "None"
)

const (
	DeviceEventDescription string = "Enables execution of transformations on device event (e.g., location, battery, etc.)."
	TimerEventDescription  string = "Enables execution of transformations on a time-base event (e.g., event 5 minutes, etc.)."
	NoneEventDescription   string = "This transformation is not scheduled."
)

// Defines the type that describes the event and time for triggering a specified transformation.
type Schedule struct {
	Id          string        `json:"id,omitempty"`
	CreatedOn   string        `json:"createdOn,omitempty"`
	LastUpdated string        `json:"lastUpdated,omitempty"`
	Event       ScheduleEvent `json:"event,omitempty"`
}

// Defines the type that represents an event on transformation.
type ScheduleEvent struct {
	Category string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
}

// Defines the type that contains the description of the
// supported schedule events.
type ScheduleEventSchema struct {
	Category    string                     `json:"category,omitempty"`
	DeviceKind  string                     `json:"deviceKind,omitempty"`
	Name        string                     `json:"name"`
	Description string                     `json:"description,omitempty"`
	Type        string                     `json:"type"`
	Semantic    string                     `json:"semantic,omitempty"`
	Fields      []ScheduleEventFieldSchema `json:"fields,omitempty"`
}

// Defines the type that contains the description of a schedule event field.
//	- Type     - Specifies the field primitive data type. E.g., string, number, etc.
//	- Constant - What this is for?
type ScheduleEventFieldSchema struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Constant bool   `json:"constant"`
	Value    string `json:"value"`
}