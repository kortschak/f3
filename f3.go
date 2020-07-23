// Copyright ©2020 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package f3 provides a universal F3 error type.
package f3

// Error is an F3 error. To conform to the standard F3 specification
// it is not comparable with anything except by F3 lookup.
type Error struct{ _ []byte }

func (Error) Error() string { return "F3" }
