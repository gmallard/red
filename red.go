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
	"log"
)

// EDData is the primary data structure used for the red application.
type EDData struct {
	// Available to clients via getters
	pta string // Plaintext Alphabet in use
	ptr []rune // Plaintext runes of Alphabet in use
	ctr []rune // Ciphertext runes of the ciphertext Alphabet in use
	us  int    // User Rotation / shift count
	ns  int    // Normalized shift count
	ds  bool   // Shift direction control, false=+ is left, true=+ is right

	// Not available to clients
	em map[rune]rune // Encryption map
	dm map[rune]rune // Decryption map
	l  *log.Logger   // Logger, if any
}
