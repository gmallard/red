//
// Copyright © 2020 Guy M. Allard
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

package red

import (
	"fmt"
	"unicode/utf8"
)

// New returns an instance that uses the defailt alphabet and shift values.
func New(p *NewParms) (ri *EDData) {
	ri = &EDData{pta: dfltAlpha,
		us: dfltShift,
		ds: false,
		ns: dfltShift,
		//
		l: nil,
	}
	//
	if p != nil {
		if p.Log != nil {
			ri.l = p.Log
		}
		if p.SwapDir {
			ri.ds = true
		}
	}
	//
	ri.commonInit()
	//
	return ri
}

// NewAlpha returns an instance that uses the default alphabet and shift
// values.
func NewAlpha(a string, s int, p *NewParms) (ri *EDData, e error) {
	//
	if !utf8.ValidString(a) {
		c := 0
		size := 0
		for i := 0; i < len(a); i += size {
			r, size := utf8.DecodeRuneInString(a)
			if r == utf8.RuneError {
				return nil, fmt.Errorf(
					"alphabet has invalid UTF8 at offset: %d, hex value: 0x%X",
					c, a[i:i+1])
			}
			a = a[size:]
			c++
		}

	}
	//
	ri = &EDData{pta: a,
		us: s,
		ds: false,
		//
		l: nil,
	}
	//
	if p != nil {
		if p.Log != nil {
			ri.l = p.Log
		}
		if p.SwapDir {
			ri.ds = true
		}
	}
	//
	ri.commonInit()
	//
	return ri, nil
}
