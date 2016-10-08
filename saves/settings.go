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
	"math"
)

// Settings represents game settings (found in the fatherland)
type Settings struct {
	// Game Mode
	// 1 = text
	// 2 = GUI/visual
	Mode byte

	// Unknown 01
	unk01 byte

	// 0 = Waiting for Missing Turns
	// 1 = Skip missing turns (host without)
	Autoplay byte

	// 0-9
	StrengthOfIndependents byte

	// #features
	MagicSiteFrequency byte

	// 1 = common
	// 2 = Rare
	RandomEventCommonality byte

	// Number of scores in the hall of fame section
	HallOfFameCount byte

	// -1 = Easy
	//  1 = Standard
	//  2 = Difficult
	//  3 = Very Difficult
	DifficultyOfMagicalResearch byte

	// Map Wrapping Mode
	// 0 = None
	// 1 = Horizontal
	// 2 = Vertical
	// 3 = Horizontal and Vertical
	MapWrap byte

	// 0xFF is undecided;
	// Non-zero is a realm ID.
	WinnerID uint8

	// Unknown
	unk10 byte

	// Unknown counter; this is sort of annoying.
	counter int16

	// Victory Mode
	// 1 = Victory Points
	// 2 = Dominion
	// 3 = Provinces
	// 4 = Research
	// 5 = Cumulative
	// 6 = Thrones
	//
	// Default seems to be 6 (thrones)
	VictoryMode int16

	// Required amount for the given Victory Mode
	VictoryRequirement int16

	// Money Multiplier
	// Default value is 100;
	// Range is 50-300
	MultiplierMoney int16

	// default value 100
	// valid in-game range of 50-300
	MultiplierResource int16

	// default value 100
	// valid in-game range of 50-300
	MultiplierSupply int16

	// unknown short array
	unknownShorts []int16

	// Game Era
	GameEra int16

	// Internal Game Flags
	GameFlags int32

	//  unknown int32 array
	unknownLongs []int32

	// set of unknown strings
	unknownStrings []string

	// unknown int32, not game turn
	unknownInt32 int32

	// Sail distance
	SailDistance int32

	// Mods
	Mods []GameMod

	// Map Picture (filename)
	MapPicture string

	// Colours (four float32s, I believe)
	Colours [4]float32
}

// Read extracts the Settings block from the stream.
func (s *Settings) Read(r io.ReadSeeker) (err error) {
	s.Mode, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: mode", err)
	}

	s.unk01, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: unk01", err)
	}

	s.Autoplay, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: autoplay", err)
	}

	s.StrengthOfIndependents, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: strength of independents", err)
	}

	// #features
	s.MagicSiteFrequency, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: magical site frequency", err)
	}

	s.RandomEventCommonality, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: random event rarity/commonality", err)
	}

	s.HallOfFameCount, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: hall of fame count", err)
	}

	s.DifficultyOfMagicalResearch, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: difficulty of magical research", err)
	}

	s.MapWrap, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: map wrapping", err)
	}

	s.WinnerID, err = filepacking.ReadUInt8(r)
	if err != nil {
		return newReadError("failed to read settings: winner id", err)
	}

	s.unk10, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("failed to read settings: unk10", err)
	}

	s.counter, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("failed to read settings: victory mode", err)
	}

	counter := int(s.counter*2) + 8

	if counter > 5 {
		s.VictoryMode, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("failed to read settings: victory mode", err)
		}

		s.VictoryRequirement, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("failed to read settings: victory requirement", err)
		}

		s.MultiplierMoney, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("failed to read settings: multiplier money", err)
		}

		s.MultiplierResource, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("failed to read settings: multiplier resource", err)
		}

		s.MultiplierSupply, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("failed to read settings: multiplier supply", err)
		}
	}

	if counter > 6 {
		s.unknownShorts = make([]int16, counter-6)
		for i := 0; i < counter-6; i++ {
			s.unknownShorts[i], err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("failed to read settings: unknown short loop", err)
			}
		}
	}

	s.GameEra, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("failed to read settings: era", err)
	}

	counter = int(s.counter*4) + 5
	if counter > 0 {
		s.GameFlags, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("failed to read settings: flags", err)
		}

		s.unknownLongs = make([]int32, counter-1)
		for i := 0; i < counter-1; i++ {
			s.unknownLongs[i], err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("failed to read settings: unknown int32s", err)
			}
		}
	}

	// When the counter is not -1 there's an array of strings; this is
	// almost never in modern Dom4.  I can't remember the last time I
	// actually saw the string array, but I keep the code here and stash
	// the strings for the eventuality that they are there ...
	if s.counter > -1 {
		s.unknownStrings = make([]string, s.counter)
		for i := int16(0); i < s.counter; i++ {
			s.unknownStrings[i], err = filepacking.ReadFileString(r)
			if err != nil {
				return newReadError("failed to read settings: unknown strings", err)
			}
		}
	}

	s.unknownInt32, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read settings: unknown i32", err)
	}

	s.SailDistance, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read settings: sail distance", err)
	}

	var modCount int16
	modCount, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("failed to read settings: mod count", err)
	}

	if modCount > 0 {
		modCount++
		s.Mods = make([]GameMod, int(modCount))
		for i := int16(0); i < modCount; i++ {
			s.Mods[i].VersionMajor, err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("failed to read settings: mod - version major", err)
			}

			s.Mods[i].VersionMinor, err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("failed to read settings: mod - version minor", err)
			}

			s.Mods[i].Name, err = filepacking.ReadFileString(r)
			if err != nil {
				return newReadError("failed to read settings: mod - name", err)
			}

			s.Mods[i].Unk3, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("failed to read settings: mod - unk3", err)
			}

			s.Mods[i].Unk4, err = filepacking.ReadInt32(r)
			if err != nil {
				return newReadError("failed to read settings: mod - unk4", err)
			}
		}
	}

	s.MapPicture, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("failed to read settings: map picture name", err)
	}

	var col int32
	for i := 0; i < 4; i++ {
		col, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("failed to read settings: colour", err)
		}
		s.Colours[i] = math.Float32frombits(uint32(col))
	}
	return
}

// GameMod is a game modification.
type GameMod struct {
	VersionMajor int16
	VersionMinor int16
	Name         string
	Unk3         int32
	Unk4         int32
}
