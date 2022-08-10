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

ARG GO_VERSION=1.18.5
ARG DOCKER_VERSION=20.10.17
ARG ALPINE_VERSION=3.16

FROM docker.io/library/golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS golang-builder

FROM golang-builder AS kind-builder
ARG KIND_VERSION=0.14.0
RUN wget -O ./kind https://kind.sigs.k8s.io/dl/v${KIND_VERSION}/kind-linux-$(go env GOARCH) && \
    chmod +x ./kind && \
    mv ./kind /usr/local/bin/kind

FROM golang-builder AS kubectl-builder
ARG KUBECTL_VERSION=1.24.2
RUN wget -O ./kubectl https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/$(go env GOARCH)/kubectl && \
    chmod +x ./kubectl && \
    mv ./kubectl /usr/local/bin/kubectl


FROM docker.io/library/docker:${DOCKER_VERSION}-dind-alpine${ALPINE_VERSION} AS docker-builder

FROM docker-builder

RUN apk add -U --no-cache bash curl ca-certificates jq

COPY --from=kind-builder /usr/local/bin/kind /usr/local/bin/
COPY --from=kubectl-builder /usr/local/bin/kubectl /usr/local/bin/

COPY --from=golang-builder /usr/local/go/ /usr/local/go/

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH
WORKDIR /workdir

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT [ "/entrypoint.sh" ]
