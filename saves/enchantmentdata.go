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
	"io"
)

// A EnchantmentData object represents enchantment data from Dominions 4.
type EnchantmentData struct {
	// 1x int32
	// 6x int16
	// 5x int32
	// 1x int16
}

// Read extracts EnchantmentData from the stream.
func (e *EnchantmentData) Read(r io.ReadSeeker) (err error) {
	// todo ...
	r.Seek(38, 1)
	return
}
