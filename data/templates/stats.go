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

package templates

import "github.com/verizonlabs/northstar/pkg/stats"

var (
	s = stats.New("templatesdata")

	CreateTemplate        = s.NewCounter("CreateTemplate")
	UserErrCreateTemplate = s.NewCounter("UserErrCreateTemplate")
	SvcErrCreateTemplate  = s.NewCounter("SvcErrCreateTemplate")
	ExtErrCreateTemplate  = s.NewCounter("ExtErrCreateTemplate")

	UpdateTemplate       = s.NewCounter("UpdateTemplate")
	SvcErrUpdateTemplate = s.NewCounter("SvcErrUpdateTemplate")
	ExtErrUpdateTemplate = s.NewCounter("ExtErrUpdateTemplate")

	GetTemplate       = s.NewCounter("GetTemplate")
	SvcErrGetTemplate = s.NewCounter("SvcErrGetTemplate")
	ExtErrGetTemplate = s.NewCounter("ExtErrGetTemplate")

	QueryTemplate        = s.NewCounter("QueryTemplate")
	UserErrQueryTemplate = s.NewCounter("UserErrQueryTemplate")
	SvcErrQueryTemplate  = s.NewCounter("SvcErrQueryTemplate")
	ExtErrQueryTemplate  = s.NewCounter("ExtErrQueryTemplate")

	DelTemplate          = s.NewCounter("DelTemplate")
	SvcErrDeleteTemplate = s.NewCounter("SvcErrDeleteTemplate")
	ExtErrDeleteTemplate = s.NewCounter("ExtErrDeleteTemplate")
)
