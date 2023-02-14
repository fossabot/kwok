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
	"context"
	"io"

	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/pager"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/client-go/dynamic"

	oyaml "gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

type Client struct {
	client *dynamic.DynamicClient
}

func NewClient(inConfig *rest.Config) (*Client, error) {
	client, err := dynamic.NewForConfig(inConfig)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: client,
	}, nil
}

func DecodeObjects(data io.Reader, fn func(obj runtime.Object) error) error {
	tmp := map[string]interface{}{}
	decoder := oyaml.NewDecoder(data)
	for {
		err := decoder.Decode(&tmp)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		data, err := oyaml.Marshal(tmp)
		if err != nil {
			return err
		}
		data, err = yaml.YAMLToJSON(data)
		if err != nil {
			return err
		}
		obj := &unstructured.Unstructured{}
		err = obj.UnmarshalJSON(data)
		if err != nil {
			return err
		}

		if obj.IsList() {
			err = obj.EachListItem(fn)
			if err != nil {
				return err
			}
		} else {
			err = fn(obj)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetGVR(obj runtime.Object) schema.GroupVersionResource {
	gvk := obj.GetObjectKind().GroupVersionKind()
	gvr, _ := meta.UnsafeGuessKindToResource(gvk)
	return gvr
}

func (c *Client) Create(ctx context.Context, gvr schema.GroupVersionResource, obj *unstructured.Unstructured, opts metav1.CreateOptions) error {
	_, err := c.client.Resource(gvr).Create(ctx, obj, opts)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) EachListItem(ctx context.Context, gvr schema.GroupVersionResource, opts metav1.ListOptions, fn func(obj runtime.Object) error) error {
	cli := c.client.Resource(gvr)
	listPager := pager.New(func(ctx context.Context, opts metav1.ListOptions) (runtime.Object, error) {
		return cli.List(ctx, opts)
	})
	return listPager.EachListItem(ctx, opts, fn)
}
