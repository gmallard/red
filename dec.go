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

// Decrypt returns the plaintext of the input cipher.  An error is
// returned if the ciphertext is not valid UTF8.
func (r *EDData) Decrypt(ct string) (cs string, e error) {
	if !utf8.ValidString(ct) {
		c := 0
		size := 0
		for i := 0; i < len(ct); i += size {
			r, size := utf8.DecodeRuneInString(ct)
			if r == utf8.RuneError {
				return "", fmt.Errorf(
					"ciphertext has invalid UTF8 at offset: %d, hex value: 0x%X",
					c, ct[i:i+1])
			}
			ct = ct[size:]
			c++
		}
	}
	rr := make([]rune, 0)
	ter := []rune(ct)
	for _, v := range ter {
		uv, ok := r.dm[v]
		if !ok {
			continue
		}
		rr = append(rr, uv)
	}
	return string(rr), nil
}
