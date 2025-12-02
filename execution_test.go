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
	v1 "github.com/sacloud/workflows-api-go/apis/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecutionAPI(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	ctx := t.Context()

	var theClient saclient.Client
	client, err := workflows.NewClient(&theClient)
	require.NoError(t, err)

	// setup
	workflowAPI := workflows.NewWorkflowOp(client)
	workflow, err := workflowAPI.Create(ctx, v1.CreateWorkflowReq{
		Name:    "test-workflow",
		Runbook: sampleRunbook,
		Publish: false,
		Logging: false,
	})
	require.NoError(t, err)
	require.NotNil(t, workflow)
	defer func() {
		err := workflowAPI.Delete(ctx, workflow.ID)
		require.NoError(t, err)
	}()

	executionAPI := workflows.NewExecutionOp(client)

	// Create
	respCreate, err := executionAPI.Create(ctx, workflow.ID, v1.CreateExecutionReq{})
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	assert.Equal(t, workflow.ID, respCreate.Workflow.ID)
}
