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
	"errors"
	"fmt"
)

var (
	// ErrorBadVersion is returned when an unsupported version is found.
	ErrorBadVersion = errors.New("bad or unsupported version encountered")

	// ErrorBadSignature is returned when a bad signature is encountered
	ErrorBadSignature = errors.New("file signature or magic not supported")
)

// A ReadError represents an error which occurs during the reading of a file.
// It wraps the internal error and displays additional information.
type ReadError struct {
	// Display error
	Display string

	// Internal error value
	Err error
}

// Error implements the error interface; it returns the string value of an
// error message.  In this case it performs the wrapping.
func (r ReadError) Error() string {
	return fmt.Sprintf("readerror: %s (inner: %s)", r.Display, r.Err)
}

// newReadError creates a read error
func newReadError(str string, err error) error {
	return ReadError{
		Display: str,
		Err:     err,
	}
}
