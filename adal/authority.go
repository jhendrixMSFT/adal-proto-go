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
	"fmt"
	"net/url"
)

const (
	MSOnlineAuthority = "https://login.microsoftonline.com/"
	endpointTemplate  = "%s/oauth2/%s?api-version=%s"
)

type Authority interface {
	AuthorizeEndpoint() (*url.URL, error)
	DeviceCodeEndpoint() (*url.URL, error)
	TokenEndpoint() (*url.URL, error)
}

// NewTenantAuthority returns an Authority based on the specified URL and tenant ID.
func NewTenantAuthority(authURL, tenantID string) Authority {
	return &tenantAuth{u: authURL, t: tenantID}
}

type tenantAuth struct {
	u string
	t string
}

func (t tenantAuth) AuthorizeEndpoint() (*url.URL, error) {
	return t.buildURL("authorize")
}

func (t tenantAuth) DeviceCodeEndpoint() (*url.URL, error) {
	return t.buildURL("devicecode")
}

func (t tenantAuth) TokenEndpoint() (*url.URL, error) {
	return t.buildURL("token")
}

func (t tenantAuth) buildURL(endpoint string) (*url.URL, error) {
	u, err := url.Parse(t.u)
	if err != nil {
		return nil, err
	}
	return u.Parse(fmt.Sprintf(endpointTemplate, t.t, endpoint, "1.0"))
}
