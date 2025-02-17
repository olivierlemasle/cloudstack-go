//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloudstack

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type LoginParams struct {
	p map[string]interface{}
}

func (p *LoginParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["domainId"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("domainId", vv)
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *LoginParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *LoginParams) SetDomainId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainId"] = v
}

func (p *LoginParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *LoginParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

// You should always use this function to get a new LoginParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLoginParams(password string, username string) *LoginParams {
	p := &LoginParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Logs a user into the CloudStack. A successful login attempt will generate a JSESSIONID cookie value that can be passed in subsequent Query command calls until the "logout" command has been issued or the session has expired.
func (s *AuthenticationService) Login(p *LoginParams) (*LoginResponse, error) {
	resp, err := s.cs.newRequest("login", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LoginResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LoginResponse struct {
	Account        string `json:"account"`
	Domainid       string `json:"domainid"`
	Firstname      string `json:"firstname"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Lastname       string `json:"lastname"`
	Registered     string `json:"registered"`
	Sessionkey     string `json:"sessionkey"`
	Timeout        int    `json:"timeout"`
	Timezone       string `json:"timezone"`
	Timezoneoffset string `json:"timezoneoffset"`
	Type           string `json:"type"`
	Userid         string `json:"userid"`
	Username       string `json:"username"`
}

type LogoutParams struct {
	p map[string]interface{}
}

func (p *LogoutParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new LogoutParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLogoutParams() *LogoutParams {
	p := &LogoutParams{}
	p.p = make(map[string]interface{})
	return p
}

// Logs out the user
func (s *AuthenticationService) Logout(p *LogoutParams) (*LogoutResponse, error) {
	resp, err := s.cs.newRequest("logout", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LogoutResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LogoutResponse struct {
	Description string `json:"description"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
}
