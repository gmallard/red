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

// PlainAlpha returns the plaintext alphabet in use.
func (r *EDData) PlainAlpha() string {
	return r.pta
}

// CipherAlpha returns the ciphertext alphabet in use.
func (r *EDData) CipherAlpha() string {
	return string(r.ctr)
}

// UserShift returns the shift value the user specified.
func (r *EDData) UserShift() int {
	return r.us
}

// NormalizedShift returns the shift value the user specified.
func (r *EDData) NormalizedShift() int {
	return r.ns
}

// ShiftDirection returns the shift direction value in use.
func (r *EDData) ShiftDirection() string {
	d := "plus is left, negative is right"
	if r.ds {
		d = "plus is right, negative is left"
	}
	return d
}
