#!/usr/bin/env bash
# Copyright 2022 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

ROOT_DIR="$(dirname "${BASH_SOURCE[0]}")/.."

CODE_GENRTETOR_VERSION="v0.26.0"
CONTROLLER_TOOLS_VERSION="v0.11.1"

function deepcopy-gen() {
  go run k8s.io/code-generator/cmd/deepcopy-gen@${CODE_GENRTETOR_VERSION} "$@"
}

function defaulter-gen() {
  go run k8s.io/code-generator/cmd/defaulter-gen@${CODE_GENRTETOR_VERSION} "$@"
}

function conversion-gen() {
  go run k8s.io/code-generator/cmd/conversion-gen@${CODE_GENRTETOR_VERSION} "$@"
}

function client-gen() {
  go run k8s.io/code-generator/cmd/client-gen@${CODE_GENRTETOR_VERSION} "$@"
}

function lister-gen() {
  go run k8s.io/code-generator/cmd/lister-gen@${CODE_GENRTETOR_VERSION} "$@"
}

function informer-gen() {
  go run k8s.io/code-generator/cmd/informer-gen@${CODE_GENRTETOR_VERSION} "$@"
}

function controller-gen() {
  go run sigs.k8s.io/controller-tools/cmd/controller-gen@${CONTROLLER_TOOLS_VERSION} "$@"
}

function gen() {
  deepcopy-gen \
    --input-dirs ./pkg/apis/v1alpha1/ \
    --trim-path-prefix sigs.k8s.io/kwok/pkg/apis \
    --output-file-base zz_generated.deepcopy \
    --go-header-file ./hack/tools/boilerplate.go.txt
  defaulter-gen \
    --input-dirs ./pkg/apis/v1alpha1/ \
    --trim-path-prefix sigs.k8s.io/kwok/pkg/apis \
    --output-file-base zz_generated.defaults \
    --go-header-file ./hack/tools/boilerplate.go.txt
  deepcopy-gen \
    --input-dirs ./pkg/apis/internalversion/ \
    --trim-path-prefix sigs.k8s.io/kwok/pkg/apis \
    --output-file-base zz_generated.deepcopy \
    --go-header-file ./hack/tools/boilerplate.go.txt
  conversion-gen \
    --input-dirs ./pkg/apis/internalversion/ \
    --trim-path-prefix sigs.k8s.io/kwok/pkg/apis \
    --output-file-base zz_generated.conversion \
    --go-header-file ./hack/tools/boilerplate.go.txt

  client-gen \
    --clientset-name versioned \
    --input-base "" \
    --input sigs.k8s.io/kwok/pkg/apis/v1alpha1 \
    --output-package sigs.k8s.io/kwok/pkg/client/clientset \
    --go-header-file ./hack/tools/boilerplate.go.txt
  lister-gen \
    --input-dirs sigs.k8s.io/kwok/pkg/apis/v1alpha1 \
    --output-package sigs.k8s.io/kwok/pkg/client/listers \
    --go-header-file ./hack/tools/boilerplate.go.txt
  informer-gen \
    --input-dirs sigs.k8s.io/kwok/pkg/apis/v1alpha1 \
    --versioned-clientset-package sigs.k8s.io/kwok/pkg/client/clientset/versioned \
    --listers-package sigs.k8s.io/kwok/pkg/client/listers \
    --output-package sigs.k8s.io/kwok/pkg/client/informers \
    --go-header-file ./hack/tools/boilerplate.go.txt

  controller-gen rbac:roleName=manager-role crd \
    paths=./pkg/apis/v1alpha1/ \
    output:crd:artifacts:config=crd
}

cd "${ROOT_DIR}" && gen || exit 1
