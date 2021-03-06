// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

const unlinkatTrap = uintptr(472)
const openatTrap = uintptr(463)
const fstatatTrap = uintptr(469)

const AT_REMOVEDIR = 0x80
const AT_SYMLINK_NOFOLLOW = 0x0020
