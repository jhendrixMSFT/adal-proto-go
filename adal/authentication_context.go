package adal

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

import (
	"context"

	"github.com/Azure/azure-pipeline-go/pipeline"
)

// AuthenticationContext is used to retrieve authentication tokens from AAD and ADFS.
type AuthenticationContext interface {
	AcquireTokenFromClientCredentials(ctx context.Context, resource, clientID, secret string) (Token, error)
}

///////////////////////////////////////////////////////////////////////////////

type authCtx struct {
	a Authority
	p pipeline.Pipeline
}

func (ta authCtx) AcquireTokenFromClientCredentials(ctx context.Context, resource, clientID, secret string) (Token, error) {
	ep, err := ta.a.TokenEndpoint()
	if err != nil {
		return nil, err
	}
	c := clientCredentials{
		ep:  ep,
		cid: clientID,
		sec: secret,
		res: resource,
	}
	return c.acquire(ctx, ta.p)
}

// NewAuthenticationContext creates a context using the specified authority and pipeline.
func NewAuthenticationContext(auth Authority, p pipeline.Pipeline) AuthenticationContext {
	if auth == nil {
		panic("auth can't be nil")
	}
	if p == nil {
		panic("p can't be nil")
	}
	return authCtx{a: auth, p: p}
}
