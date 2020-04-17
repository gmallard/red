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

package main

import (
	"log"
	"os"

	//
	"gmallard.com/cryplay/red"
)

var (
	ll = log.New(os.Stdout, "SHSHFT ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	ss = "=============================================================="
	r  *red.EDData
)

func showData() {
	ll.Println(ss)
	ll.Printf("Plaintext_Alphabet:  %s\n", r.PlainAlpha())
	ll.Printf("Ciphertext_Alphabet: %s\n", r.CipherAlpha())
	ll.Printf("Shift Direction: %s\n", r.ShiftDirection())
	ll.Printf("User Shift Value: %d\n", r.UserShift())
	ll.Printf("Normalized Shift Value: %d\n", r.NormalizedShift())
}

func main() {
	//
	r, _ = red.NewAlpha(red.AlLc, 5, nil)
	showData()
	//
	r, _ = red.NewAlpha(red.AlLc, -5, nil)
	showData()
	//
	rpf := &red.NewParms{Log: nil, SwapDir: false}
	r, _ = red.NewAlpha(red.AlLc, 5, rpf)
	showData()
	//
	rpt := &red.NewParms{Log: nil, SwapDir: true}
	r, _ = red.NewAlpha(red.AlLc, 5, rpt)
	showData()
	//
	r, _ = red.NewAlpha(red.AlLc, 1776, nil)
	showData()
	//
	r, _ = red.NewAlpha(red.AlLc, -1776, nil)
	showData()
	//
	r, _ = red.NewAlpha(red.AlLc, 0, nil)
	showData()
}
