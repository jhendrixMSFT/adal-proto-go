package runtime

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"fmt"

	"github.com/Azure/azure-pipeline-go/pipeline"
)

type Responder func(resp pipeline.Response) (result pipeline.Response, err error)

func NewResponderPolicyFactory(r Responder) pipeline.FactoryFunc {
	return pipeline.FactoryFunc(func(next pipeline.Policy, po *pipeline.PolicyOptions) pipeline.PolicyFunc {
		return func(ctx context.Context, req pipeline.Request) (pipeline.Response, error) {
			resp, err := next.Do(ctx, req)
			if err != nil {
				return resp, err
			}
			return r(resp)
		}
	})
}

// ValidateResponse checks an HTTP response's status code against a legal set of codes.
// If the response code is not legal, then validateResponse reads all of the response's body
// (containing error information) and returns a response error.
func ValidateResponse(resp pipeline.Response, successStatusCodes ...int) error {
	if resp == nil {
		return pipeline.NewError(nil, "nil response")
	}
	responseCode := resp.Response().StatusCode
	for _, i := range successStatusCodes {
		if i == responseCode {
			return nil
		}
	}
	// TODO: better error
	return pipeline.NewError(nil, fmt.Sprintf("bad status code: %d", responseCode))
}
