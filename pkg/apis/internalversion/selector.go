/*
Copyright 2023 The Kubernetes Authors.

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

package internalversion

import (
	"golang.org/x/exp/slices"
)

// NamespaceNameSelector holds information how to match based on namespace and name.
type NamespaceNameSelector struct {
	MatchNamespace []string
	MatchName      []string
}

// Matches returns true if name and namespace is specified within the selector
//
// If the match field is empty, the match on that field is considered to be true.
func (s *NamespaceNameSelector) Matches(namespace, name string) bool {
	if s == nil {
		return true
	}
	if len(s.MatchNamespace) > 0 && !slices.Contains(s.MatchNamespace, namespace) {
		return false
	}
	if len(s.MatchName) > 0 && !slices.Contains(s.MatchName, name) {
		return false
	}
	return true
}
