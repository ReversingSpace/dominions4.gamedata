/**
 * Reversing Space: Dominons 4 Game Data Files
 * Copyright (c) 2015-2016 A.W. Stanley.
 *
 * This software is provided 'as-is', without any express or implied warranty.
 * In no event will the authors be held liable for any damages arising from
 * the use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 *   1. The origin of this software must not be misrepresented; you must
 *      not claim that you wrote the original software. If you use this
 *      software in a product, an acknowledgment in the product
 *      documentation would be appreciated but is not required.
 *
 *   2. Altered source versions must be plainly marked as such, and
 *      must not be misrepresented as being the original software.
 *
 *   3. This notice may not be removed or altered from any
 *      source distribution.
 *
**/

package saves

import (
	"github.com/ReversingSpace/dominions4.gamedata/filepacking"
	"io"
)

// An EndStats object stores the end-game statistical information.
type EndStats struct {
	arr00 []int16
	arr01 []int16
	arr02 []int16
	arr03 []int16
	arr04 []int16
	arr05 []int16
	arr06 []int16
	arr07 []int16
}

// Read extracts EndStats from the stream.
func (e *EndStats) Read(r io.Reader) (err error) {

	var counter int32
	counter, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("endstats: failed to read counter", err)
	}

	if counter > 0 {
		sz := int(counter * 200)
		e.arr00 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr00[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 00 value", err)
			}
		}

		e.arr01 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr01[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 01 value", err)
			}
		}

		e.arr02 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr02[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 02 value", err)
			}
		}
		e.arr03 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr03[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 03 value", err)
			}
		}
		e.arr04 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr04[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 04 value", err)
			}
		}
		e.arr05 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr05[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 05 value", err)
			}
		}
		e.arr06 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr06[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 06 value", err)
			}
		}
		e.arr07 = make([]int16, sz)
		for i := 0; i < sz; i++ {
			e.arr07[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("endstats: unable to read short array 07 value", err)
			}
		}
	}

	return
}
