// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package backend

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

// Key is the unique identifier for an [Item].
type Key struct {
	components []string
	s          string
	exactKey   bool
	noEnd      bool
}

// Separator is used as a separator between key parts
const Separator = '/'

// NewKey joins parts into path separated by Separator,
// makes sure path always starts with Separator ("/")
func NewKey(components ...string) Key {
	k := internalKey("", components...)
	k.exactKey = k.s == string(Separator) || (len(k.s) > 0 && k.s[len(k.s)-1] == Separator)
	return k
}

// ExactKey is like Key, except a Separator is appended to the result
// path of Key. This is to ensure range matching of a path will only
// math child paths and not other paths that have the resulting path
// as a prefix.
func ExactKey(components ...string) Key {
	k := NewKey(append(components, "")...)
	k.exactKey = true
	return k
}

func KeyFromString(s string) Key {
	components := strings.Split(s, string(Separator))
	if components[0] == "" && len(components) > 1 {
		components = components[1:]
	}
	return NewKey(components...)
}

func (k Key) IsZero() bool {
	return len(k.components) == 0 && k.s == ""
}

func internalKey(internalPrefix string, components ...string) Key {
	return Key{components: components, s: strings.Join(append([]string{internalPrefix}, components...), string(Separator))}
}

func (k Key) ExactKey() Key {
	if k.exactKey {
		return k
	}

	return ExactKey(k.components...)
}

// String returns the textual representation of the key with
// each component concatenated together via the [Separator].
func (k Key) String() string {
	if k.noEnd {
		return string(noEnd)
	}

	return k.s
}

// HasPrefix reports whether the key begins with prefix.
func (k Key) HasPrefix(prefix Key) bool {
	return strings.HasPrefix(k.s, prefix.s)
}

// TrimPrefix returns the key without the provided leading prefix string.
// If the key doesn't start with prefix, it is returned unchanged.
func (k Key) TrimPrefix(prefix Key) Key {
	key := strings.TrimPrefix(k.s, prefix.s)
	if key == "" {
		return Key{}
	}

	return KeyFromString(key)
}

func (k Key) PrependKey(p Key) Key {
	return NewKey(append(slices.Clone(p.components), slices.Clone(k.components)...)...)
}

func (k Key) AppendKey(p Key) Key {
	return p.PrependKey(k)
}

// HasSuffix reports whether the key ends with suffix.
func (k Key) HasSuffix(suffix Key) bool {
	return strings.HasSuffix(k.s, suffix.s)
}

// TrimSuffix returns the key without the provided trailing suffix string.
// If the key doesn't end with suffix, it is returned unchanged.
func (k Key) TrimSuffix(suffix Key) Key {
	key := strings.TrimSuffix(k.s, suffix.s)
	if key == "" {
		return Key{}
	}

	return KeyFromString(key)
}

func (k Key) Components() []string {
	return slices.Clone(k.components)
}

func (k Key) Compare(o Key) int {
	return strings.Compare(k.s, o.s)
}

// Scan implement sql.Scanner, allowing a [Key] to
// be directly retrieved from sql backends without
// an intermediary object.
func (k *Key) Scan(scan any) error {
	switch key := scan.(type) {
	case []byte:
		if len(key) == 0 {
			return nil
		}
		*k = KeyFromString(string(bytes.Clone(key)))
	case string:
		if len(key) == 0 {
			return nil
		}

		*k = KeyFromString(strings.Clone(key))
	default:
		return fmt.Errorf("invalid Key type %T", scan)
	}

	return nil
}
