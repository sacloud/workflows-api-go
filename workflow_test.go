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

func TestWorkflowAPI(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	ctx := t.Context()

	var theClient saclient.Client
	client, err := workflows.NewClient(&theClient)
	require.NoError(t, err)

	workflowAPI := workflows.NewWorkflowOp(client)

	// CreateWorkflow
	createRes, err := workflowAPI.Create(ctx, v1.CreateWorkflowReq{
		Name: "test-workflow",
	})
	require.Error(t, err)
	require.NotNil(t, createRes)
	assert.Equal(t, createRes.Name, "test-workflow")
}
