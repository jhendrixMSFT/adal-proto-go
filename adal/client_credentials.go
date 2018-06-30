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
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/jhendrixMSFT/adal-proto-go/internal/runtime"
)

type clientCredentials struct {
	ep  *url.URL
	cid string
	sec string
	res string
}

func (cc clientCredentials) acquire(ctx context.Context, p pipeline.Pipeline) (Token, error) {
	req, err := cc.acquirePreparer()
	if err != nil {
		return nil, err
	}
	resp, err := p.Do(ctx, runtime.NewResponderPolicyFactory(cc.acquireResponder), req)
	if err != nil {
		return nil, err
	}
	return resp.(Token), nil
}

func (cc clientCredentials) acquirePreparer() (pipeline.Request, error) {
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	v.Set("client_id", cc.cid)
	v.Set("resource", cc.res)
	v.Set("client_secret", cc.sec)
	s := v.Encode()
	req, err := pipeline.NewRequest(http.MethodPost, *cc.ep, strings.NewReader(s))
	if err != nil {
		return req, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func (cc clientCredentials) acquireResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := runtime.ValidateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	at := &accessToken{rawResp: resp.Response()}
	if err != nil {
		return at, err
	}
	return at, runtime.FromJSON(resp, at)
}
