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
	respCreate, err := workflowAPI.Create(ctx, v1.CreateWorkflowReq{
		Name:        "test-workflow",
		Description: v1.NewOptString("test workflow"), // TODO: somehow it's required on the server side
		Runbook:     sampleRunbook,
		Publish:     false,
		Logging:     false,
	})
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	assert.Equal(t, "test-workflow", respCreate.Name)

	defer func() {
		// DeleteWorkflow
		err = workflowAPI.Delete(ctx, respCreate.ID)
		require.NoError(t, err)
	}()

	// ReadWorkflow
	respRead, err := workflowAPI.Read(ctx, respCreate.ID)
	require.NoError(t, err)
	require.NotNil(t, respRead)
	assert.Equal(t, "test-workflow", respRead.Name)
	assert.Equal(t, v1.NewOptString("test workflow"), respRead.Description)
	assert.Equal(t, false, respRead.Publish)
	assert.Equal(t, false, respRead.Logging)

	// ListWorkflows
	respList, err := workflowAPI.List(ctx, v1.ListWorkflowParams{})
	require.NoError(t, err)
	found := false
	for _, workflow := range respList.Workflows {
		if workflow.ID == respCreate.ID {
			found = true
			assert.Equal(t, "test-workflow", respRead.Name)
			assert.Equal(t, v1.NewOptString("test workflow"), respRead.Description)
			assert.Equal(t, false, respRead.Publish)
			assert.Equal(t, false, respRead.Logging)
		}
	}
	assert.True(t, found, "Created Workflow not found in list")

	// UpdateWorkflow
	respUpdate, err := workflowAPI.Update(ctx, respCreate.ID, v1.UpdateWorkflowReq{
		Name:    v1.NewOptString("test-workflow-updated"),
		Publish: v1.NewOptBool(true),
		Logging: v1.NewOptBool(true),
	})
	require.NoError(t, err)
	require.NotNil(t, respUpdate)
	assert.Equal(t, "test-workflow-updated", respUpdate.Name)
	assert.Equal(t, true, respUpdate.Publish)
	assert.Equal(t, true, respUpdate.Logging)
}

const sampleRunbook = `
meta:
  description: エラトステネスの篩
args:
  maxNumber:
    type: number
    description: 素数を求める最大の数
steps:
  setup:
    assign:
      sieve: ${array.fill(array.range(args.maxNumber), true)}
      primes: []
  initial:
    assign:
      _a: ${array.set(sieve, 0, false)}
      _b: ${array.set(sieve, 1, false)}
  loop:
    for:
      in: ${array.range(2, math.ceil(math.sqrt(args.maxNumber)))}
      as: index
      steps:
        if:
          switch:
            # falseだったら飛ばす
            - condition: ${sieve[index] == false}
              next: continue
            # trueだったら素数
            - condition: ${sieve[index] != false}
              steps:
                # 素数の倍数を篩にかける
                updateSieve:
                  for:
                    in: ${array.range(index * 2, args.maxNumber, index)}
                    as: n
                    steps:
                      set:
                        assign:
                          _a: ${array.set(sieve, n, false)}
        continue:
  printPrimes:
    for:
      in: ${array.range(2, args.maxNumber)}
      as: index
      steps:
        if:
          switch:
            - condition: ${sieve[index] == true}
              steps:
                push:
                  assign:
                    _a: ${array.push(primes, index)}
                log:
                  assign:
                    log: '${"素数: " + index}'
  done:
    return: ${primes}
`
