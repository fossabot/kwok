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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ClusterPortForwardKind is the kind of the ClusterPortForward.
	ClusterPortForwardKind = "ClusterPortForward"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPortForward provides cluster-wide port forward configuration.
type ClusterPortForward struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterPortForwardSpec `json:"spec"`
}

// ClusterPortForwardSpec holds spec for cluster port forward.
type ClusterPortForwardSpec struct {
	Selector *ClusterPortForwardSelector `json:"selector,omitempty"`
	Forwards []Forward                   `json:"forwards,omitempty"`
}

// ClusterPortForwardSelector holds information how to match based on namespace and name.
type ClusterPortForwardSelector struct {
	MatchNamespace []string `json:"matchNamespace,omitempty"`
	MatchName      []string `json:"matchName,omitempty"`
}
