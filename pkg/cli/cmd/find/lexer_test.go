/* Copyright (C) 2019 Monomax Software Pty Ltd
 *
 * This file is part of Dnote CLI.
 *
 * Dnote CLI is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Dnote CLI is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Dnote CLI.  If not, see <https://www.gnu.org/licenses/>.
 */

package find

import (
	"fmt"
	"testing"

	"github.com/dnote/dnote/pkg/assert"
)

func TestScanToken(t *testing.T) {
	testCases := []struct {
		input  string
		idx    int
		retTok token
		retIdx int
	}{
		{
			input:  "foo bar",
			idx:    1,
			retTok: token{Value: 'o', Kind: tokenKindChar},
			retIdx: 2,
		},
		{
			input:  "foo bar",
			idx:    6,
			retTok: token{Value: 'r', Kind: tokenKindChar},
			retIdx: -1,
		},
		{
			input:  "foo <bar>",
			idx:    4,
			retTok: token{Value: '<', Kind: tokenKindChar},
			retIdx: 5,
		},
		{
			input:  "foo <dnotehL>",
			idx:    4,
			retTok: token{Value: '<', Kind: tokenKindChar},
			retIdx: 5,
		},
		{
			input:  "foo <dnotehl>bar</dnotehl> foo bar",
			idx:    4,
			retTok: token{Kind: tokenKindHLBegin},
			retIdx: 13,
		},
		{
			input:  "foo <dnotehl>bar</dnotehl> <dnotehl>foo</dnotehl> bar",
			idx:    4,
			retTok: token{Kind: tokenKindHLBegin},
			retIdx: 13,
		},
		{
			input:  "foo <dnotehl>bar</dnotehl> <dnotehl>foo</dnotehl> bar",
			idx:    27,
			retTok: token{Kind: tokenKindHLBegin},
			retIdx: 36,
		},
		{
			input:  "foo <dnotehl>bar</dnotehl> foo bar",
			idx:    13,
			retTok: token{Value: 'b', Kind: tokenKindChar},
			retIdx: 14,
		},
		{
			input:  "foo <dnotehl>bar</dnotehl> foo bar",
			idx:    16,
			retTok: token{Kind: tokenKindHLEnd},
			retIdx: 26,
		},
		{
			input:  "<dno<dnotehl>tehl>",
			idx:    0,
			retTok: token{Value: '<', Kind: tokenKindChar},
			retIdx: 1,
		},
		{
			input:  "<dno<dnotehl>tehl>",
			idx:    4,
			retTok: token{Kind: tokenKindHLBegin},
			retIdx: 13,
		},
		{
			input:  "foo <dnotehl>bar</dnotehl>",
			idx:    16,
			retTok: token{Kind: tokenKindHLEnd},
			retIdx: -1,
		},
		// user writes reserved token
		{
			input:  "foo <dnotehl>",
			idx:    4,
			retTok: token{Kind: tokenKindHLBegin},
			retIdx: -1,
		},
	}

	for tcIdx, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", tcIdx), func(t *testing.T) {
			tok, nextIdx := scanToken(tc.idx, tc.input)

			assert.Equal(t, nextIdx, tc.retIdx, "retIdx mismatch")
			assert.DeepEqual(t, tok, tc.retTok, "retTok mismatch")
		})
	}
}

func TestTokenize(t *testing.T) {
	testCases := []struct {
		input  string
		tokens []token
	}{
		{
			input: "ab<dnotehl>c</dnotehl>",
			tokens: []token{
				{
					Kind:  tokenKindChar,
					Value: 'a',
				},
				{
					Kind:  tokenKindChar,
					Value: 'b',
				},
				{
					Kind: tokenKindHLBegin,
				},
				{
					Kind:  tokenKindChar,
					Value: 'c',
				},
				{
					Kind: tokenKindHLEnd,
				},
				{
					Kind: tokenKindEOL,
				},
			},
		},
		{
			input: "ab<dnotehl>c</dnotehl>d",
			tokens: []token{
				{
					Kind:  tokenKindChar,
					Value: 'a',
				},
				{
					Kind:  tokenKindChar,
					Value: 'b',
				},
				{
					Kind: tokenKindHLBegin,
				},
				{
					Kind:  tokenKindChar,
					Value: 'c',
				},
				{
					Kind: tokenKindHLEnd,
				},
				{
					Kind:  tokenKindChar,
					Value: 'd',
				},
				{
					Kind: tokenKindEOL,
				},
			},
		},
		// user writes a reserved token
		{
			input: "<dnotehl><dnotehl></dnotehl>",
			tokens: []token{
				{
					Kind: tokenKindHLBegin,
				},
				{
					Kind: tokenKindHLBegin,
				},
				{
					Kind: tokenKindHLEnd,
				},
				{
					Kind: tokenKindEOL,
				},
			},
		},
		{
			input: "<dnotehl></dnotehl></dnotehl>",
			tokens: []token{
				{
					Kind: tokenKindHLBegin,
				},
				{
					Kind: tokenKindHLEnd,
				},
				{
					Kind: tokenKindHLEnd,
				},
				{
					Kind: tokenKindEOL,
				},
			},
		},
	}

	for tcIdx, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", tcIdx), func(t *testing.T) {
			tokens := tokenize(tc.input)

			assert.DeepEqual(t, tokens, tc.tokens, "tokens mismatch")
		})
	}
}
