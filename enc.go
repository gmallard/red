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

// Encrypt returns the ciphertext of the input plaintext.  An error is
// returned if the plaintext is not valid UTF8.
func (r *EDData) Encrypt(pt string) (cs string, e error) {
	if !utf8.ValidString(pt) {
		return "", fmt.Errorf("%s", ErrU8PT)
	}
	//
	if !utf8.ValidString(pt) {
		return "", fmt.Errorf(ErrU8PT)
	}

	rr := make([]rune, 0)
	ter := []rune(pt)
	for _, v := range ter {
		uv, ok := r.em[v]
		if !ok {
			continue
		}
		rr = append(rr, uv)
	}
	return string(rr), nil
}
