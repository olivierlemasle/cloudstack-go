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

type AddNiciraNvpDeviceParams struct {
	p map[string]interface{}
}

func (p *AddNiciraNvpDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["l2gatewayserviceuuid"]; found {
		u.Set("l2gatewayserviceuuid", v.(string))
	}
	if v, found := p.p["l3gatewayserviceuuid"]; found {
		u.Set("l3gatewayserviceuuid", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["transportzoneuuid"]; found {
		u.Set("transportzoneuuid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddNiciraNvpDeviceParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *AddNiciraNvpDeviceParams) SetL2gatewayserviceuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["l2gatewayserviceuuid"] = v
}

func (p *AddNiciraNvpDeviceParams) SetL3gatewayserviceuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["l3gatewayserviceuuid"] = v
}

func (p *AddNiciraNvpDeviceParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddNiciraNvpDeviceParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddNiciraNvpDeviceParams) SetTransportzoneuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["transportzoneuuid"] = v
}

func (p *AddNiciraNvpDeviceParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

// You should always use this function to get a new AddNiciraNvpDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewAddNiciraNvpDeviceParams(hostname string, password string, physicalnetworkid string, transportzoneuuid string, username string) *AddNiciraNvpDeviceParams {
	p := &AddNiciraNvpDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["transportzoneuuid"] = transportzoneuuid
	p.p["username"] = username
	return p
}

// Adds a Nicira NVP device
func (s *NiciraNVPService) AddNiciraNvpDevice(p *AddNiciraNvpDeviceParams) (*AddNiciraNvpDeviceResponse, error) {
	resp, err := s.cs.newRequest("addNiciraNvpDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNiciraNvpDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AddNiciraNvpDeviceResponse struct {
	Hostname             string `json:"hostname"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	L2gatewayserviceuuid string `json:"l2gatewayserviceuuid"`
	L3gatewayserviceuuid string `json:"l3gatewayserviceuuid"`
	Niciradevicename     string `json:"niciradevicename"`
	Nvpdeviceid          string `json:"nvpdeviceid"`
	Physicalnetworkid    string `json:"physicalnetworkid"`
	Provider             string `json:"provider"`
	Transportzoneuuid    string `json:"transportzoneuuid"`
}

type DeleteNiciraNvpDeviceParams struct {
	p map[string]interface{}
}

func (p *DeleteNiciraNvpDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
	}
	return u
}

func (p *DeleteNiciraNvpDeviceParams) SetNvpdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nvpdeviceid"] = v
}

// You should always use this function to get a new DeleteNiciraNvpDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewDeleteNiciraNvpDeviceParams(nvpdeviceid string) *DeleteNiciraNvpDeviceParams {
	p := &DeleteNiciraNvpDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["nvpdeviceid"] = nvpdeviceid
	return p
}

//  delete a nicira nvp device
func (s *NiciraNVPService) DeleteNiciraNvpDevice(p *DeleteNiciraNvpDeviceParams) (*DeleteNiciraNvpDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteNiciraNvpDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNiciraNvpDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteNiciraNvpDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNiciraNvpDevicesParams struct {
	p map[string]interface{}
}

func (p *ListNiciraNvpDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListNiciraNvpDevicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNiciraNvpDevicesParams) SetNvpdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nvpdeviceid"] = v
}

func (p *ListNiciraNvpDevicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNiciraNvpDevicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNiciraNvpDevicesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

// You should always use this function to get a new ListNiciraNvpDevicesParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewListNiciraNvpDevicesParams() *ListNiciraNvpDevicesParams {
	p := &ListNiciraNvpDevicesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Nicira NVP devices
func (s *NiciraNVPService) ListNiciraNvpDevices(p *ListNiciraNvpDevicesParams) (*ListNiciraNvpDevicesResponse, error) {
	resp, err := s.cs.newRequest("listNiciraNvpDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNiciraNvpDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNiciraNvpDevicesResponse struct {
	Count            int                `json:"count"`
	NiciraNvpDevices []*NiciraNvpDevice `json:"niciranvpdevice"`
}

type NiciraNvpDevice struct {
	Hostname             string `json:"hostname"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	L2gatewayserviceuuid string `json:"l2gatewayserviceuuid"`
	L3gatewayserviceuuid string `json:"l3gatewayserviceuuid"`
	Niciradevicename     string `json:"niciradevicename"`
	Nvpdeviceid          string `json:"nvpdeviceid"`
	Physicalnetworkid    string `json:"physicalnetworkid"`
	Provider             string `json:"provider"`
	Transportzoneuuid    string `json:"transportzoneuuid"`
}
