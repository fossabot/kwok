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

package client

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestDecodeObjects(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []runtime.Object
		wantErr bool
	}{
		{
			args: args{
				data: bytes.NewBufferString(`
apiVersion: v1
kind: Pod
metadata:
  name: test
  namespace: test
spec:
  containers: []
---
apiVersion: v1
kind: Pod
metadata:
  name: test-2
  namespace: test
spec:
  containers: []
`),
			},
			want: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "Pod",
						"metadata": map[string]interface{}{
							"name":      "test",
							"namespace": "test",
						},
						"spec": map[string]interface{}{
							"containers": []interface{}{},
						},
					},
				},
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "Pod",
						"metadata": map[string]interface{}{
							"name":      "test-2",
							"namespace": "test",
						},
						"spec": map[string]interface{}{
							"containers": []interface{}{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []runtime.Object{}
			if err := DecodeObjects(tt.args.data, func(obj runtime.Object) error {
				got = append(got, obj)
				return nil
			}); (err != nil) != tt.wantErr {
				t.Errorf("DecodeObjects() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(got) != len(tt.want) {
				t.Errorf("DecodeObjects() got = %v, want %v", got, tt.want)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("expected match (-want +got):\n%s", diff)
			}
		})
	}
}
