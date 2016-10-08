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

// A SpellData object represents spell data from the fatherland
type SpellData struct {
	// array of shorts; probably ID information
	shortArray []int16

	// array of shorts 2
	shortArray2 []int16
}

// Read extracts SpellData from the stream.
func (s *SpellData) Read(r io.Reader) (err error) {
	s.shortArray = make([]int16, 1000)
	for i := 0; i < 1000; i++ {
		s.shortArray[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("spelldata: unable to read short array: bad entry", err)
		}
	}

	s.shortArray2 = make([]int16, 24)
	for i := 0; i < 24; i++ {
		s.shortArray2[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("spelldata: unable to read short array 2: bad entry", err)
		}
	}

	return
}
