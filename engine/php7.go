// Copyright 2016 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// +build php7

package engine

// #cgo CFLAGS: -Iinclude/php7 -Isrc/php7
// #cgo LDFLAGS: -lphp7
import "C"
