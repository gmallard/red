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

// Common alphabets generally useful.
const (
	AlLc  = "abcdefghijklmnopqrstuvwxyz"
	AlUc  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlLU  = AlLc + AlUc
	AlNuM = "0123456789"
	AlLUN = AlLc + AlUc + AlNuM
)

// Error messages
const (
	ErrU8Alpha = "Invalid UTF8 in Alphabet"
	ErrU8PT    = "Invalid UTF8 in plaintext"
	ErrU8CT    = "Invalid UTF8 in ciphertext"
)

// Unexported constants.
const (
	dfltAlpha = AlLc
	dfltShift = 5
)
