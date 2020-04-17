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
	"runtime"
)

func (r *EDData) normalizedShift() int {
	if r.us > 0 {
		ts := r.us
		if r.us > len(r.ptr) {
			ts = ts % len(r.ptr)
		}
		r.log("normalizdedShift returns", ts)
		return ts
	}
	s := r.us
	if s < -1*len(r.ptr) {
		s = s % len(r.ptr)
	}
	s = -1 * s
	ts := len(r.ptr) - s
	r.log("normalizdedShift returns", ts)
	return ts
}

func (r *EDData) makeEncData() {
	r.ctr = make([]rune, 0)
	if r.us == 0 {
		r.ctr = r.ptr
		r.log("makeEncData copyied ctr", fmt.Sprintf("%q", r.ctr))
		return
	}
	t := r.us
	if t < 0 {
		t = -1 * t
	}
	if t%len(r.pta) == 0 {
		r.ctr = r.ptr
		r.log("makeEncData copyied ctr", fmt.Sprintf("%q", r.ctr))
		return
	}
	//
	if r.ds {
		// shift right
		// abcdefghijklmnopqrstuvwxyz
		// vwxyzabcdefghijklmnopqrstu
		r.log("makeEncData swaps direction", r.ns)
		r.ctr = r.ptr[len(r.ptr)-r.ns:]
		r.ctr = append(r.ctr, r.ptr[0:len(r.ptr)-r.ns]...) // pr[0:len(pr)-s]...]...)
	} else {
		// shhift left
		r.ctr = r.ptr[r.ns:]
		r.ctr = append(r.ctr, r.ptr[0:r.ns]...)
	}
	r.log("makeEncData made ctr", fmt.Sprintf("%q", r.ctr))
	//
	return
}

// For debugging only
func (r *EDData) dumpMaps() {
	fmt.Printf("ptr: %q\n", r.ptr)
	fmt.Printf("ctr: %q\n", r.ctr)
	fmt.Println()

	fmt.Println("Map dump:")
	for _, k := range r.em {
		av := string(k)
		ev := string(r.em[k])
		bv := string(r.dm[r.em[k]])
		//
		fmt.Printf("%s -> %s -> %s : %t\n", av, ev, bv, av == bv)
	}

}

func (r *EDData) makeEDMaps() {
	//
	r.em = make(map[rune]rune)
	for k, v := range r.pta {
		r.em[v] = r.ctr[k]
	}
	//
	r.dm = make(map[rune]rune)
	for k, v := range r.ctr {
		r.dm[v] = r.ptr[k]
	}
	r.log("makeEncData made em", fmt.Sprintf("%q", r.em))
	r.log("makeEncData made dm", fmt.Sprintf("%q", r.dm))
	// r.dumpMaps()
}

func (r *EDData) log(v ...interface{}) {
	if r.l == nil {
		return
	}
	_, fn, ld, ok := runtime.Caller(1)

	if ok {
		r.l.Printf("%s %d %v\n", fn, ld, v)
	} else {
		r.l.Printf("%v\n", v)
	}
	return
}

func (r *EDData) commonInit() {
	r.ptr = []rune(r.pta)
	r.ns = r.normalizedShift()
	r.makeEncData()
	r.makeEDMaps()
}
