// Copyright 2020 Layer5, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import "github.com/layer5io/meshery-adapter-library/meshes"

func (h *Adapter) StreamErr(e *meshes.EventsResponse, err error) {
	h.Log.Error(err)
	e.EventType = 2
	*h.Channel <- e
}

func (h *Adapter) StreamInfo(e *meshes.EventsResponse) {
	h.Log.Info("Sending event")
	e.EventType = 0
	*h.Channel <- e
}
