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

import "testing"

func TestSelector_Matches(t *testing.T) {
	t.Run("nil selector should match everything", func(t *testing.T) {
		var selector *NamespaceNameSelector
		podNamespace, podName := "podNamespace", "podName"
		if !selector.Matches(podNamespace, podName) {
			t.Fatalf("expected nil selector to match everything")
		}
	})

	t.Run("empty match should match everything", func(t *testing.T) {
		podName := "podName"
		podNamespace := "podNamespace"
		tt := map[string]struct {
			matchName      []string
			matchNamespace []string
		}{
			"EmptyName": {
				matchName:      []string{},
				matchNamespace: []string{podNamespace},
			},
			"EmptyNamespace": {
				matchName:      []string{podName},
				matchNamespace: []string{},
			},
			"EmptyNamespaceAndName": {
				matchName:      []string{},
				matchNamespace: []string{},
			},
		}

		for name, tc := range tt {
			t.Run(name, func(t *testing.T) {
				selector := NamespaceNameSelector{
					MatchName:      tc.matchName,
					MatchNamespace: tc.matchNamespace,
				}
				if !selector.Matches(podNamespace, podName) {
					t.Errorf("expected selector to match podNamespace and podName")
				}
			})
		}
	})

	t.Run("namespace name matching", func(t *testing.T) {
		selector := NamespaceNameSelector{
			MatchName:      []string{"podName"},
			MatchNamespace: []string{"podNamespace"},
		}

		tt := []struct {
			podName      string
			podNamespace string
			expect       bool
		}{
			{
				podName:      "",
				podNamespace: "",
				expect:       false,
			},
			{
				podName:      "",
				podNamespace: "podNamespace",
				expect:       false,
			},
			{
				podName:      "podName",
				podNamespace: "",
				expect:       false,
			},
			{
				podName:      "podName",
				podNamespace: "podNamespace",
				expect:       true,
			},
		}

		for _, tc := range tt {
			got := selector.Matches(tc.podNamespace, tc.podName)
			if got != tc.expect {
				t.Errorf("Matches(%q, %q)=%v, expect=%v", tc.podNamespace, tc.podName, got, tc.expect)
			}
		}
	})
}
