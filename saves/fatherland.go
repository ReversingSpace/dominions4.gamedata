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
	"encoding/binary"
	"github.com/ReversingSpace/dominions4.gamedata/filepacking"
	"io"
)

// Fatherland data file, stored in data order.
type Fatherland struct {
	// File Header
	Header FileHeader

	// Game Info/Settings fields
	Settings Settings

	// CalendarA
	CalendarA map[int32]int32

	// CalenderB
	CalendarB map[int32]int32 // lunar?

	// This is the default zoom value; it's from Dom3.
	// Fairly sure it isn't remotely useful in Dom4.
	Zoom float32

	// Land information
	Lands map[int32]*Land

	// Kingdom information
	Kingdoms map[int32]*Kingdom

	// Unit information
	Units map[int32]*Unit
}

// Read processes a fatherland file from the stream.
func (f *Fatherland) Read(r io.ReadSeeker) (err error) {
	err = f.Header.Read(r)
	if err != nil {
		return
	}

	err = f.Settings.Read(r)
	if err != nil {
		return
	}

	{
		if f.CalendarA == nil {
			f.CalendarA = make(map[int32]int32)
		}
		if f.CalendarB == nil {
			f.CalendarB = make(map[int32]int32)
		}

		var a int32
		var b int32
		var k int32
		for {
			k, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("fatherland: failed to read calendar value", err)
			}
			if k < 0 {
				break
			}
			a, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("fatherland: failed to read calendar a value", err)
			}
			b, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("fatherland: failed to read calendar b value", err)
			}
			if k > 999 {
				break
			}
			f.CalendarA[k] = a
			f.CalendarB[k] = b
		}
		a, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: failed to read calendar sentry value", err)
		}
		if a != 8283 {
			return newReadError("fatherland: bad calendar sentry value", err)
		}
	}

	err = binary.Read(r, binary.LittleEndian, &f.Zoom)
	if err != nil {
		return newReadError("fatherland: failed to read default zoom", err)
	}

	f.Lands = make(map[int32]*Land)
	{
		var index int32

		for {
			index, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("fatherland: failed to read land index", err)
			}

			if index < 0 {
				break
			}

			if index > 0x5E0 {
				return newReadError("fatherland: invalid index (exceeds 0x5e0)", err)
			}

			var land Land
			land.TreatAsFatherland = true

			err = land.Read(r)
			if err != nil {
				return newReadError("fatherland: land read failure", err)
			}

			f.Lands[index] = &land
		}
	}

	f.Kingdoms = make(map[int32]*Kingdom)
	{
		var index int32
		for {
			index, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("fatherland: kingdom index", err)
			}
			if index < 0 {
				break
			}
			if index > 0xC7 {
				return newReadError("fatherland: kingdom index exceeds 0xc7", nil)
			}

			var kingdom Kingdom
			err = kingdom.Read(r)
			if err != nil {
				return err
			}
			f.Kingdoms[index] = &kingdom
		}
	}

	// Units.
	f.Units = make(map[int32]*Unit)

	// todo
	// units
	// commanders
	// dominions
	// spells?
	// mercs
	// enchantments
	// scores
	// items
	// war info
	// highest heroes
	// fatherland passwords (?) - 200 of them?
	// end stats
	// events
	// 12346 sentinel

	return
}
