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

package workflows_test

import (
	"testing"

	"github.com/sacloud/packages-go/testutil"
	"github.com/sacloud/saclient-go"
	"github.com/sacloud/workflows-api-go"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionAPI(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	ctx := t.Context()

	var theClient saclient.Client
	client, err := workflows.NewClient(&theClient)
	require.NoError(t, err)

	subscriptionAPI := workflows.NewSubscriptionOp(client)

	// ListPlans
	respListPlans, err := subscriptionAPI.ListPlans(ctx)
	require.NoError(t, err)
	require.NotNil(t, respListPlans)

	// Read
	respRead, err := subscriptionAPI.Read(ctx)
	require.NoError(t, err)
	require.NotNil(t, respRead) // TODO: must be nil at first?

	// TODO: add tests for Create and Delete
}
