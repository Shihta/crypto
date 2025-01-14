// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package siv

import (
	"log"

	"github.com/Shihta/crypto/common"
)

// The xorend operator of RFC 5297.
//
// Given strings A and B with len(A) >= len(B), let D be len(A) - len(B). Write
// A[:D] followed by xor(A[D:], B) into dst. In other words, xor B over the
// rightmost end of A and write the result into dst.
func xorend(dst, a, b []byte) {
	aLen := len(a)
	bLen := len(b)
	dstLen := len(dst)

	if dstLen < aLen || aLen < bLen {
		log.Panicf("Bad buffer lengths: %d, %d, %d", dstLen, aLen, bLen)
	}

	// Copy the left part.
	difference := aLen - bLen
	copy(dst, a[:difference])

	// XOR in the right part.
	common.Xor(dst[difference:difference+bLen], a[difference:], b)
}
