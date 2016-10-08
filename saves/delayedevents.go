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

// DelayedEvents represents delayed events.
type DelayedEvents struct {
	Base []int32
	Turn []int32

	// I *assume* this is right.
	Lunar []int32
}

// Read extracts DelayedEvents from the stream.
func (e *DelayedEvents) Read(r io.Reader) (err error) {

	var maxIndex int32
	maxIndex, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("delayed events: failed to read max index", err)
	}

	if maxIndex < 0 {
		return
	}

	maxIndex++

	var value int32

	e.Base = make([]int32, maxIndex)
	e.Turn = make([]int32, maxIndex)
	e.Lunar = make([]int32, maxIndex)

	for index := int32(0); index < maxIndex; index++ {
		value, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("delayed events: failed to read base event value", err)
		}
		e.Base[index] = value

		value, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("delayed events: failed to read turn event value", err)
		}
		e.Turn[index] = value

		value, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("delayed events: failed to read lunar event value", err)
		}
		e.Lunar[index] = value

	}

	return
}
