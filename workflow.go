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
	"context"
	"errors"
	"net/http"

	v1 "github.com/sacloud/workflows-api-go/apis/v1"
)

type WorkflowAPI interface {
	Create(ctx context.Context, request v1.CreateWorkflowReq) (*v1.CreateWorkflowCreatedWorkflow, error)
	Delete(ctx context.Context, id string) error
}

var _ WorkflowAPI = (*workflowOp)(nil)

type workflowOp struct {
	client *v1.Client
}

func NewWorkflowOp(client *v1.Client) WorkflowAPI {
	return &workflowOp{client: client}
}

func (op *workflowOp) Create(ctx context.Context, req v1.CreateWorkflowReq) (*v1.CreateWorkflowCreatedWorkflow, error) {
	const methodName = "Workflow.Create"

	// TODO: why opt req
	res, err := op.client.CreateWorkflow(ctx, v1.NewOptCreateWorkflowReq(req))
	if err != nil {
		return nil, NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.CreateWorkflowCreated:
		return &r.Workflow, nil
	case *v1.CreateWorkflowBadRequest:
		return nil, NewAPIError(methodName, http.StatusBadRequest, errors.New(r.Message))
	case *v1.CreateWorkflowUnauthorized:
		return nil, NewAPIError(methodName, http.StatusUnauthorized, errors.New(r.Message))
	case *v1.CreateWorkflowForbidden:
		return nil, NewAPIError(methodName, http.StatusForbidden, errors.New(r.Message))
	case *v1.CreateWorkflowNotFound:
		return nil, NewAPIError(methodName, http.StatusNotFound, errors.New(r.Message))
	case *v1.CreateWorkflowInternalServerError:
		return nil, NewAPIError(methodName, http.StatusInternalServerError, errors.New(r.Message))
	default:
		return nil, NewAPIError(methodName, 0, err)
	}
}

func (op *workflowOp) Delete(ctx context.Context, id string) error {
	const methodName = "Workflow.Delete"

	res, err := op.client.DeleteWorkflow(ctx, v1.DeleteWorkflowParams{
		ID: id,
	})
	if err != nil {
		return NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.DeleteWorkflowOK:
		return nil
	case *v1.DeleteWorkflowBadRequest:
		return NewAPIError(methodName, http.StatusBadRequest, errors.New(r.Message))
	case *v1.DeleteWorkflowUnauthorized:
		return NewAPIError(methodName, http.StatusUnauthorized, errors.New(r.Message))
	case *v1.DeleteWorkflowForbidden:
		return NewAPIError(methodName, http.StatusForbidden, errors.New(r.Message))
	case *v1.DeleteWorkflowNotFound:
		return NewAPIError(methodName, http.StatusNotFound, errors.New(r.Message))
	case *v1.DeleteWorkflowInternalServerError:
		return NewAPIError(methodName, http.StatusInternalServerError, errors.New(r.Message))
	default:
		return NewAPIError(methodName, 0, err)
	}
}
