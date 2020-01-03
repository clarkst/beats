// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
