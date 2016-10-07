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

// CalendarEntry represents an entry in the "new" Dom4 calendar format.
type CalendarEntry struct {
	Index int32
	Unk1  int32
	Unk2  int32
}

// ReadCalendar reads the array of calendar entries found in the
// Fatherland file.
func ReadCalendar(r io.ReadSeeker) (array []CalendarEntry, err error) {
	array = make([]CalendarEntry, 0)
	var ce *CalendarEntry
	for {
		ce = &CalendarEntry{}

		ce.Index, err = filepacking.ReadInt32(r)
		if err != nil {
			err = newReadError("calendar error: index failed to read", err)
			return
		}

		// Bad value
		if ce.Index <= -1 {
			break
		}

		ce.Unk1, err = filepacking.ReadInt32(r)
		if err != nil {
			err = newReadError("calendar error: unk1 failed to read", err)
			return
		}

		ce.Unk2, err = filepacking.ReadInt32(r)
		if err != nil {
			err = newReadError("calendar error: unk2 failed to read", err)
			return
		}

		// Valid range
		if ce.Index > 999 {
			continue
		}

		array = append(array, *ce)
	}

	var test int32
	test, err = filepacking.ReadInt32(r)
	if err != nil {
		err = newReadError("calendar error: test value not read", err)
	}
	if test != 8283 {
		err = newReadError("calendar error: test value invalid", nil)
	}
	return
}
