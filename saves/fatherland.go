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
	"fmt"
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

	// Commander information
	Commanders map[int32]*Commander

	// Dominion information
	Dominions map[int32]*Dominion

	// Spell information
	Spells map[int32]*SpellData

	// Mercenary Data
	Mercenaries map[int32]*MercenaryData

	// Mercenary related
	mercDataUnknown []byte

	// Enchantment Data
	Enchantments map[int32]*EnchantmentData

	// Scores
	Scores []int16

	// Items
	Items []byte

	// War Data
	WarData []byte

	// Heroes
	Heroes []int16

	// unknown strings
	unkRXN []string

	EndStats EndStats

	EventOccurrences []int16

	DelayedEvents DelayedEvents
}

// Read processes a fatherland file from the stream.
func (f *Fatherland) Read(r io.ReadSeeker) (err error) {
	var index int32
	var n int

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
		for {
			index, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("fatherland: failed to read calendar index", err)
			}
			if index < 0 {
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
			if index > 999 {
				break
			}
			f.CalendarA[index] = a
			f.CalendarB[index] = b
		}
		a, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: failed to read calendar sentry value", err)
		}
		if a != 0x205B {
			return fmt.Errorf("fatherland: bad calendar sentry value (is %d, should be 8283)", a)
		}
	}

	err = binary.Read(r, binary.LittleEndian, &f.Zoom)
	if err != nil {
		return newReadError("fatherland: failed to read default zoom", err)
	}

	f.Lands = make(map[int32]*Land)
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

	f.Kingdoms = make(map[int32]*Kingdom)
	for {
		index, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: kingdom index", err)
		}
		if index < 0 {
			break
		}
		if index > 0xF9 {
			return newReadError("fatherland: kingdom index exceeds 249", nil)
		}

		var kingdom Kingdom
		err = kingdom.Read(r)
		if err != nil {
			return err
		}
		f.Kingdoms[index] = &kingdom
	}

	// Units.
	f.Units = make(map[int32]*Unit)
	for {
		index, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: unit index", err)
		}
		if index < 0 {
			break
		}
		u := &Unit{}
		err = u.Read(r)
		if err != nil {
			return newReadError("fatherland: failed to read unit", err)
		}
		f.Units[index] = u
	}

	f.Commanders = make(map[int32]*Commander)
	for {
		index, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: commander index", err)
		}
		if index < 0 {
			break
		}
		c := &Commander{}
		err = c.Read(r)
		if err != nil {
			return newReadError("fatherland: failed to read commander", err)
		}
		f.Commanders[index] = c
	}

	f.Dominions = make(map[int32]*Dominion)
	for {
		index, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: dominion index", err)
		}
		if index < 0 {
			break
		}
		d := &Dominion{}
		err = d.Read(r)
		if err != nil {
			return newReadError("fatherland: failed to read dominion", err)
		}
		f.Dominions[index] = d
	}

	f.Spells = make(map[int32]*SpellData)
	for {
		index, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: spelldata index", err)
		}
		if index < 0 {
			break
		}
		d := &SpellData{}
		err = d.Read(r)
		if err != nil {
			return newReadError("fatherland: failed to read spelldata", err)
		}
		f.Spells[index] = d
	}

	f.Mercenaries = make(map[int32]*MercenaryData)
	for {
		index, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: mercenary index", err)
		}
		if index < 0 {
			break
		}

		// this does something weird, but it isn't read related so
		// we'll ignore it for now.

		d := &MercenaryData{}
		err = d.Read(r)
		if err != nil {
			return newReadError("fatherland: failed to read mercenarydata", err)
		}
		f.Mercenaries[index] = d
	}

	// unk
	f.mercDataUnknown = make([]byte, 100)
	n, err = r.Read(f.mercDataUnknown)
	if err != nil {
		return newReadError("fatherland: failed to read merc data (unknown)", err)
	}

	if n != 100 {
		return fmt.Errorf("fatherland: failed to read merc data (unknown)")
	}

	f.Enchantments = make(map[int32]*EnchantmentData)
	for {
		index, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: mercenary index", err)
		}
		if index < 0 {
			break
		}
		d := &EnchantmentData{}
		err = d.Read(r)
		if err != nil {
			return newReadError("fatherland: failed to read mercenarydata", err)
		}
		f.Enchantments[index] = d
	}

	if f.Settings.HallOfFameCount > 0 {
		f.Scores = make([]int16, f.Settings.HallOfFameCount)
		for i := 0; i < int(f.Settings.HallOfFameCount); i++ {
			f.Scores[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("fatherland: score failed to read", err)
			}
		}
	}

	f.Items = make([]byte, 1000)
	n, err = r.Read(f.Items)
	if err != nil {
		return newReadError("fatherland: failed to read items", err)
	}

	if n != 1000 {
		return fmt.Errorf("fatherland: failed to read items (needed 1000, not %d)", n)
	}

	f.WarData = make([]byte, 40000)
	n, err = r.Read(f.WarData)
	if err != nil {
		return newReadError("fatherland: failed to read war matrix", err)
	}

	if n != 40000 {
		return fmt.Errorf("fatherland: failed to read war matrix (needed 40000 bytes, got %d)", n)
	}

	var heroCount int32
	heroCount, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("fatherland: failed to read hero count", err)
	}

	f.Heroes = make([]int16, heroCount+1)
	for i := 0; i < len(f.Heroes); i++ {
		f.Heroes[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("fatherland: failed to read hero index", err)
		}
	}

	f.unkRXN = make([]string, 200)
	for i := 0; i < 200; i++ {
		f.unkRXN[i], err = filepacking.ReadFileStringRXN(r, 50)
		if err != nil {
			return newReadError("fatherland: failed to read rxn string", err)
		}
	}

	err = f.EndStats.Read(r)
	if err != nil {
		return newReadError("fatherland: failed to read end stats", err)
	}

	var unk4474 int32
	unk4474, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("fatherland: failed to read sentinel-like value (4474)", err)
	}

	if unk4474 < 4474 {
		return fmt.Errorf("fatherland: sentinel-like value must exceed 4474 and does not (is %d)", unk4474)
	}

	eventCount := int32(1000)

	if unk4474 == 4475 {
		eventCount, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("fatherland: failed to read event count", err)
		}
	}

	if eventCount > 4000 {
		return fmt.Errorf("fatherland: invalid event count (%d exceeds 4000)", eventCount)
	}

	f.EventOccurrences = make([]int16, eventCount)
	for i := 0; i < int(eventCount); i++ {
		f.EventOccurrences[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("fatherland: failed to read event occurrence", err)
		}
	}

	// re-use 4474
	unk4474, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("fatherland: failed to read sentinel-like value (4480)", err)
	}

	if unk4474 != 4480 {
		return fmt.Errorf("fatherland: invalid 4480 sentinel in delayed events (value is %d)", unk4474)
	}

	err = f.DelayedEvents.Read(r)
	if err != nil {
		return newReadError("fatherland: failed to read delayed events", err)
	}

	// 12346 sentinel; reuse unk4474
	unk4474, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("fatherland: failed to read sentinel", err)
	}

	if unk4474 != 12346 {
		return fmt.Errorf("fatherland: closing sentinel is not 12346 (is %d)", unk4474)
	}

	return
}
