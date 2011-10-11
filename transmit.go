//
// Copyright © 2011 Guy M. Allard
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
//

package stomp

import (
	"os"
)

// Common send for transaction verbs, some others
func (c *Connection) transmitCommon(v string, h Headers) (e os.Error) {
	ch := h.Clone()
	f := Frame{v, ch, make([]uint8, 0)}
	r := make(chan os.Error)
	c.output <- wiredata{f, r}
	e = <-r
	return e
}
