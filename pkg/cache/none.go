/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import "errors"

// NoneStore is a dummy, nop store for use
// with read-only storages (e.g. github).
type NoneStore struct{}

var _ Store = NoneStore{}

// NewNoneStore creates a dummy, nop store.
func NewNoneStore() NoneStore {
	return NoneStore{}
}

// ErrNoneStore is returned by NoneStore.Get to force rebuilding
// index each time it was requested.
var ErrNoneStore = errors.New(`"none" cache does not store any entry`)

// Get is a nop.
func (NoneStore) Get(string) ([]byte, error) { return nil, ErrNoneStore }

// Set is a nop.
func (NoneStore) Set(string, []byte) error { return nil }

// Delete is a nop.
func (NoneStore) Delete(string) error { return nil }
