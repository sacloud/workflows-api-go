// Copyright 2025- The sacloud/workflows-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workflows

import (
	"fmt"
	"runtime"

	"github.com/sacloud/saclient-go"
	v1 "github.com/sacloud/workflows-api-go/apis/v1"
)

// DefaultAPIRootURL デフォルトのAPIルートURL
const DefaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/zone/tk1b/api/workflow/1.0/"

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"workflows-api-go/%s (%s/%s; +https://github.com/sacloud/workflows-api-go)",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
)

// NewClient creates a new workflows API client with default settings
func NewClient(client saclient.ClientAPI) (*v1.Client, error) {
	return NewClientWithAPIRootURL(client, DefaultAPIRootURL)
}

// NewClientWithAPIRootURL creates a new workflows API client with a custom API root URL
func NewClientWithAPIRootURL(client saclient.ClientAPI, apiRootURL string) (*v1.Client, error) {
	return v1.NewClient(apiRootURL, v1.WithClient(client))
}
