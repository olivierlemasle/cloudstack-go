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
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type CreateAutoScalePolicyParams struct {
	p map[string]interface{}
}

func (p *CreateAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["conditionids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("conditionids", vv)
	}
	if v, found := p.p["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := p.p["quiettime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quiettime", vv)
	}
	return u
}

func (p *CreateAutoScalePolicyParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *CreateAutoScalePolicyParams) SetConditionids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conditionids"] = v
}

func (p *CreateAutoScalePolicyParams) SetDuration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["duration"] = v
}

func (p *CreateAutoScalePolicyParams) SetQuiettime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiettime"] = v
}

// You should always use this function to get a new CreateAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScalePolicyParams(action string, conditionids []string, duration int) *CreateAutoScalePolicyParams {
	p := &CreateAutoScalePolicyParams{}
	p.p = make(map[string]interface{})
	p.p["action"] = action
	p.p["conditionids"] = conditionids
	p.p["duration"] = duration
	return p
}

// Creates an autoscale policy for a provision or deprovision action, the action is taken when the all the conditions evaluates to true for the specified duration. The policy is in effect once it is attached to a autscale vm group.
func (s *AutoScaleService) CreateAutoScalePolicy(p *CreateAutoScalePolicyParams) (*CreateAutoScalePolicyResponse, error) {
	resp, err := s.cs.newRequest("createAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScalePolicyResponse
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

type CreateAutoScalePolicyResponse struct {
	Account    string   `json:"account"`
	Action     string   `json:"action"`
	Conditions []string `json:"conditions"`
	Domain     string   `json:"domain"`
	Domainid   string   `json:"domainid"`
	Duration   int      `json:"duration"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Project    string   `json:"project"`
	Projectid  string   `json:"projectid"`
	Quiettime  int      `json:"quiettime"`
}

type CreateAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *CreateAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["maxmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmembers", vv)
	}
	if v, found := p.p["minmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmembers", vv)
	}
	if v, found := p.p["scaledownpolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaledownpolicyids", vv)
	}
	if v, found := p.p["scaleuppolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaleuppolicyids", vv)
	}
	if v, found := p.p["vmprofileid"]; found {
		u.Set("vmprofileid", v.(string))
	}
	return u
}

func (p *CreateAutoScaleVmGroupParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateAutoScaleVmGroupParams) SetInterval(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["interval"] = v
}

func (p *CreateAutoScaleVmGroupParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *CreateAutoScaleVmGroupParams) SetMaxmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxmembers"] = v
}

func (p *CreateAutoScaleVmGroupParams) SetMinmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minmembers"] = v
}

func (p *CreateAutoScaleVmGroupParams) SetScaledownpolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaledownpolicyids"] = v
}

func (p *CreateAutoScaleVmGroupParams) SetScaleuppolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaleuppolicyids"] = v
}

func (p *CreateAutoScaleVmGroupParams) SetVmprofileid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmprofileid"] = v
}

// You should always use this function to get a new CreateAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScaleVmGroupParams(lbruleid string, maxmembers int, minmembers int, scaledownpolicyids []string, scaleuppolicyids []string, vmprofileid string) *CreateAutoScaleVmGroupParams {
	p := &CreateAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	p.p["maxmembers"] = maxmembers
	p.p["minmembers"] = minmembers
	p.p["scaledownpolicyids"] = scaledownpolicyids
	p.p["scaleuppolicyids"] = scaleuppolicyids
	p.p["vmprofileid"] = vmprofileid
	return p
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *AutoScaleService) CreateAutoScaleVmGroup(p *CreateAutoScaleVmGroupParams) (*CreateAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("createAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScaleVmGroupResponse
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

type CreateAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type CreateAutoScaleVmProfileParams struct {
	p map[string]interface{}
}

func (p *CreateAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["autoscaleuserid"]; found {
		u.Set("autoscaleuserid", v.(string))
	}
	if v, found := p.p["counterparam"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("counterparam[%d].key", i), k)
			u.Set(fmt.Sprintf("counterparam[%d].value", i), m[k])
		}
	}
	if v, found := p.p["destroyvmgraceperiod"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("destroyvmgraceperiod", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["otherdeployparams"]; found {
		u.Set("otherdeployparams", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateAutoScaleVmProfileParams) SetAutoscaleuserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscaleuserid"] = v
}

func (p *CreateAutoScaleVmProfileParams) SetCounterparam(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterparam"] = v
}

func (p *CreateAutoScaleVmProfileParams) SetDestroyvmgraceperiod(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destroyvmgraceperiod"] = v
}

func (p *CreateAutoScaleVmProfileParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateAutoScaleVmProfileParams) SetOtherdeployparams(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["otherdeployparams"] = v
}

func (p *CreateAutoScaleVmProfileParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateAutoScaleVmProfileParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *CreateAutoScaleVmProfileParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new CreateAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScaleVmProfileParams(serviceofferingid string, templateid string, zoneid string) *CreateAutoScaleVmProfileParams {
	p := &CreateAutoScaleVmProfileParams{}
	p.p = make(map[string]interface{})
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a profile that contains information about the virtual machine which will be provisioned automatically by autoscale feature.
func (s *AutoScaleService) CreateAutoScaleVmProfile(p *CreateAutoScaleVmProfileParams) (*CreateAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newRequest("createAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScaleVmProfileResponse
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

type CreateAutoScaleVmProfileResponse struct {
	Account              string `json:"account"`
	Autoscaleuserid      string `json:"autoscaleuserid"`
	Destroyvmgraceperiod int    `json:"destroyvmgraceperiod"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Otherdeployparams    string `json:"otherdeployparams"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Serviceofferingid    string `json:"serviceofferingid"`
	Templateid           string `json:"templateid"`
	Zoneid               string `json:"zoneid"`
}

type CreateConditionParams struct {
	p map[string]interface{}
}

func (p *CreateConditionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["counterid"]; found {
		u.Set("counterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["relationaloperator"]; found {
		u.Set("relationaloperator", v.(string))
	}
	if v, found := p.p["threshold"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("threshold", vv)
	}
	return u
}

func (p *CreateConditionParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateConditionParams) SetCounterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterid"] = v
}

func (p *CreateConditionParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateConditionParams) SetRelationaloperator(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["relationaloperator"] = v
}

func (p *CreateConditionParams) SetThreshold(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["threshold"] = v
}

// You should always use this function to get a new CreateConditionParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateConditionParams(counterid string, relationaloperator string, threshold int64) *CreateConditionParams {
	p := &CreateConditionParams{}
	p.p = make(map[string]interface{})
	p.p["counterid"] = counterid
	p.p["relationaloperator"] = relationaloperator
	p.p["threshold"] = threshold
	return p
}

// Creates a condition
func (s *AutoScaleService) CreateCondition(p *CreateConditionParams) (*CreateConditionResponse, error) {
	resp, err := s.cs.newRequest("createCondition", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateConditionResponse
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

type CreateConditionResponse struct {
	Account            string   `json:"account"`
	Counter            []string `json:"counter"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Id                 string   `json:"id"`
	JobID              string   `json:"jobid"`
	Jobstatus          int      `json:"jobstatus"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Relationaloperator string   `json:"relationaloperator"`
	Threshold          int64    `json:"threshold"`
	Zoneid             string   `json:"zoneid"`
}

type CreateCounterParams struct {
	p map[string]interface{}
}

func (p *CreateCounterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["source"]; found {
		u.Set("source", v.(string))
	}
	if v, found := p.p["value"]; found {
		u.Set("value", v.(string))
	}
	return u
}

func (p *CreateCounterParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateCounterParams) SetSource(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["source"] = v
}

func (p *CreateCounterParams) SetValue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["value"] = v
}

// You should always use this function to get a new CreateCounterParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateCounterParams(name string, source string, value string) *CreateCounterParams {
	p := &CreateCounterParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["source"] = source
	p.p["value"] = value
	return p
}

// Adds metric counter
func (s *AutoScaleService) CreateCounter(p *CreateCounterParams) (*CreateCounterResponse, error) {
	resp, err := s.cs.newRequest("createCounter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateCounterResponse
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

type CreateCounterResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Source    string `json:"source"`
	Value     string `json:"value"`
	Zoneid    string `json:"zoneid"`
}

type DeleteAutoScalePolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAutoScalePolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScalePolicyParams(id string) *DeleteAutoScalePolicyParams {
	p := &DeleteAutoScalePolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a autoscale policy.
func (s *AutoScaleService) DeleteAutoScalePolicy(p *DeleteAutoScalePolicyParams) (*DeleteAutoScalePolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScalePolicyResponse
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

type DeleteAutoScalePolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScaleVmGroupParams(id string) *DeleteAutoScaleVmGroupParams {
	p := &DeleteAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a autoscale vm group.
func (s *AutoScaleService) DeleteAutoScaleVmGroup(p *DeleteAutoScaleVmGroupParams) (*DeleteAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScaleVmGroupResponse
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

type DeleteAutoScaleVmGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteAutoScaleVmProfileParams struct {
	p map[string]interface{}
}

func (p *DeleteAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAutoScaleVmProfileParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScaleVmProfileParams(id string) *DeleteAutoScaleVmProfileParams {
	p := &DeleteAutoScaleVmProfileParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a autoscale vm profile.
func (s *AutoScaleService) DeleteAutoScaleVmProfile(p *DeleteAutoScaleVmProfileParams) (*DeleteAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newRequest("deleteAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScaleVmProfileResponse
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

type DeleteAutoScaleVmProfileResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteConditionParams struct {
	p map[string]interface{}
}

func (p *DeleteConditionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteConditionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteConditionParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteConditionParams(id string) *DeleteConditionParams {
	p := &DeleteConditionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes a condition
func (s *AutoScaleService) DeleteCondition(p *DeleteConditionParams) (*DeleteConditionResponse, error) {
	resp, err := s.cs.newRequest("deleteCondition", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteConditionResponse
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

type DeleteConditionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteCounterParams struct {
	p map[string]interface{}
}

func (p *DeleteCounterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteCounterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteCounterParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteCounterParams(id string) *DeleteCounterParams {
	p := &DeleteCounterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a counter
func (s *AutoScaleService) DeleteCounter(p *DeleteCounterParams) (*DeleteCounterResponse, error) {
	resp, err := s.cs.newRequest("deleteCounter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteCounterResponse
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

type DeleteCounterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *DisableAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DisableAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DisableAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDisableAutoScaleVmGroupParams(id string) *DisableAutoScaleVmGroupParams {
	p := &DisableAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Disables an AutoScale Vm Group
func (s *AutoScaleService) DisableAutoScaleVmGroup(p *DisableAutoScaleVmGroupParams) (*DisableAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("disableAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableAutoScaleVmGroupResponse
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

type DisableAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type EnableAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *EnableAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new EnableAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewEnableAutoScaleVmGroupParams(id string) *EnableAutoScaleVmGroupParams {
	p := &EnableAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Enables an AutoScale Vm Group
func (s *AutoScaleService) EnableAutoScaleVmGroup(p *EnableAutoScaleVmGroupParams) (*EnableAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("enableAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableAutoScaleVmGroupResponse
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

type EnableAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type ListAutoScalePoliciesParams struct {
	p map[string]interface{}
}

func (p *ListAutoScalePoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["conditionid"]; found {
		u.Set("conditionid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["vmgroupid"]; found {
		u.Set("vmgroupid", v.(string))
	}
	return u
}

func (p *ListAutoScalePoliciesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListAutoScalePoliciesParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *ListAutoScalePoliciesParams) SetConditionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conditionid"] = v
}

func (p *ListAutoScalePoliciesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAutoScalePoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAutoScalePoliciesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAutoScalePoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAutoScalePoliciesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAutoScalePoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAutoScalePoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAutoScalePoliciesParams) SetVmgroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmgroupid"] = v
}

// You should always use this function to get a new ListAutoScalePoliciesParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScalePoliciesParams() *ListAutoScalePoliciesParams {
	p := &ListAutoScalePoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScalePolicyByID(id string, opts ...OptionFunc) (*AutoScalePolicy, int, error) {
	p := &ListAutoScalePoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScalePolicies(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.AutoScalePolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScalePolicy UUID: %s!", id)
}

// Lists autoscale policies.
func (s *AutoScaleService) ListAutoScalePolicies(p *ListAutoScalePoliciesParams) (*ListAutoScalePoliciesResponse, error) {
	resp, err := s.cs.newRequest("listAutoScalePolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScalePoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScalePoliciesResponse struct {
	Count             int                `json:"count"`
	AutoScalePolicies []*AutoScalePolicy `json:"autoscalepolicy"`
}

type AutoScalePolicy struct {
	Account    string   `json:"account"`
	Action     string   `json:"action"`
	Conditions []string `json:"conditions"`
	Domain     string   `json:"domain"`
	Domainid   string   `json:"domainid"`
	Duration   int      `json:"duration"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Project    string   `json:"project"`
	Projectid  string   `json:"projectid"`
	Quiettime  int      `json:"quiettime"`
}

type ListAutoScaleVmGroupsParams struct {
	p map[string]interface{}
}

func (p *ListAutoScaleVmGroupsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["vmprofileid"]; found {
		u.Set("vmprofileid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListAutoScaleVmGroupsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetPolicyid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyid"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetVmprofileid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmprofileid"] = v
}

func (p *ListAutoScaleVmGroupsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListAutoScaleVmGroupsParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScaleVmGroupsParams() *ListAutoScaleVmGroupsParams {
	p := &ListAutoScaleVmGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmGroupByID(id string, opts ...OptionFunc) (*AutoScaleVmGroup, int, error) {
	p := &ListAutoScaleVmGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScaleVmGroups(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.AutoScaleVmGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScaleVmGroup UUID: %s!", id)
}

// Lists autoscale vm groups.
func (s *AutoScaleService) ListAutoScaleVmGroups(p *ListAutoScaleVmGroupsParams) (*ListAutoScaleVmGroupsResponse, error) {
	resp, err := s.cs.newRequest("listAutoScaleVmGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScaleVmGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScaleVmGroupsResponse struct {
	Count             int                 `json:"count"`
	AutoScaleVmGroups []*AutoScaleVmGroup `json:"autoscalevmgroup"`
}

type AutoScaleVmGroup struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type ListAutoScaleVmProfilesParams struct {
	p map[string]interface{}
}

func (p *ListAutoScaleVmProfilesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["otherdeployparams"]; found {
		u.Set("otherdeployparams", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListAutoScaleVmProfilesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetOtherdeployparams(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["otherdeployparams"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *ListAutoScaleVmProfilesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListAutoScaleVmProfilesParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScaleVmProfilesParams() *ListAutoScaleVmProfilesParams {
	p := &ListAutoScaleVmProfilesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmProfileByID(id string, opts ...OptionFunc) (*AutoScaleVmProfile, int, error) {
	p := &ListAutoScaleVmProfilesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScaleVmProfiles(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.AutoScaleVmProfiles[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScaleVmProfile UUID: %s!", id)
}

// Lists autoscale vm profiles.
func (s *AutoScaleService) ListAutoScaleVmProfiles(p *ListAutoScaleVmProfilesParams) (*ListAutoScaleVmProfilesResponse, error) {
	resp, err := s.cs.newRequest("listAutoScaleVmProfiles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScaleVmProfilesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScaleVmProfilesResponse struct {
	Count               int                   `json:"count"`
	AutoScaleVmProfiles []*AutoScaleVmProfile `json:"autoscalevmprofile"`
}

type AutoScaleVmProfile struct {
	Account              string `json:"account"`
	Autoscaleuserid      string `json:"autoscaleuserid"`
	Destroyvmgraceperiod int    `json:"destroyvmgraceperiod"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Otherdeployparams    string `json:"otherdeployparams"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Serviceofferingid    string `json:"serviceofferingid"`
	Templateid           string `json:"templateid"`
	Zoneid               string `json:"zoneid"`
}

type ListConditionsParams struct {
	p map[string]interface{}
}

func (p *ListConditionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["counterid"]; found {
		u.Set("counterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	return u
}

func (p *ListConditionsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListConditionsParams) SetCounterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterid"] = v
}

func (p *ListConditionsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListConditionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListConditionsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListConditionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListConditionsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListConditionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListConditionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListConditionsParams) SetPolicyid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyid"] = v
}

// You should always use this function to get a new ListConditionsParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListConditionsParams() *ListConditionsParams {
	p := &ListConditionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetConditionByID(id string, opts ...OptionFunc) (*Condition, int, error) {
	p := &ListConditionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListConditions(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Conditions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Condition UUID: %s!", id)
}

// List Conditions for the specific user
func (s *AutoScaleService) ListConditions(p *ListConditionsParams) (*ListConditionsResponse, error) {
	resp, err := s.cs.newRequest("listConditions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListConditionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListConditionsResponse struct {
	Count      int          `json:"count"`
	Conditions []*Condition `json:"condition"`
}

type Condition struct {
	Account            string   `json:"account"`
	Counter            []string `json:"counter"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Id                 string   `json:"id"`
	JobID              string   `json:"jobid"`
	Jobstatus          int      `json:"jobstatus"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Relationaloperator string   `json:"relationaloperator"`
	Threshold          int64    `json:"threshold"`
	Zoneid             string   `json:"zoneid"`
}

type ListCountersParams struct {
	p map[string]interface{}
}

func (p *ListCountersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["source"]; found {
		u.Set("source", v.(string))
	}
	return u
}

func (p *ListCountersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListCountersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListCountersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListCountersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListCountersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListCountersParams) SetSource(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["source"] = v
}

// You should always use this function to get a new ListCountersParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListCountersParams() *ListCountersParams {
	p := &ListCountersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListCountersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListCounters(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Counters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Counters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterByName(name string, opts ...OptionFunc) (*Counter, int, error) {
	id, count, err := s.GetCounterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetCounterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterByID(id string, opts ...OptionFunc) (*Counter, int, error) {
	p := &ListCountersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListCounters(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Counters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Counter UUID: %s!", id)
}

// List the counters
func (s *AutoScaleService) ListCounters(p *ListCountersParams) (*ListCountersResponse, error) {
	resp, err := s.cs.newRequest("listCounters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCountersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCountersResponse struct {
	Count    int        `json:"count"`
	Counters []*Counter `json:"counter"`
}

type Counter struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Source    string `json:"source"`
	Value     string `json:"value"`
	Zoneid    string `json:"zoneid"`
}

type UpdateAutoScalePolicyParams struct {
	p map[string]interface{}
}

func (p *UpdateAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["conditionids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("conditionids", vv)
	}
	if v, found := p.p["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["quiettime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quiettime", vv)
	}
	return u
}

func (p *UpdateAutoScalePolicyParams) SetConditionids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conditionids"] = v
}

func (p *UpdateAutoScalePolicyParams) SetDuration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["duration"] = v
}

func (p *UpdateAutoScalePolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAutoScalePolicyParams) SetQuiettime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiettime"] = v
}

// You should always use this function to get a new UpdateAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScalePolicyParams(id string) *UpdateAutoScalePolicyParams {
	p := &UpdateAutoScalePolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing autoscale policy.
func (s *AutoScaleService) UpdateAutoScalePolicy(p *UpdateAutoScalePolicyParams) (*UpdateAutoScalePolicyResponse, error) {
	resp, err := s.cs.newRequest("updateAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScalePolicyResponse
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

type UpdateAutoScalePolicyResponse struct {
	Account    string   `json:"account"`
	Action     string   `json:"action"`
	Conditions []string `json:"conditions"`
	Domain     string   `json:"domain"`
	Domainid   string   `json:"domainid"`
	Duration   int      `json:"duration"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Project    string   `json:"project"`
	Projectid  string   `json:"projectid"`
	Quiettime  int      `json:"quiettime"`
}

type UpdateAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *UpdateAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	if v, found := p.p["maxmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmembers", vv)
	}
	if v, found := p.p["minmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmembers", vv)
	}
	if v, found := p.p["scaledownpolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaledownpolicyids", vv)
	}
	if v, found := p.p["scaleuppolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaleuppolicyids", vv)
	}
	return u
}

func (p *UpdateAutoScaleVmGroupParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateAutoScaleVmGroupParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAutoScaleVmGroupParams) SetInterval(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["interval"] = v
}

func (p *UpdateAutoScaleVmGroupParams) SetMaxmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxmembers"] = v
}

func (p *UpdateAutoScaleVmGroupParams) SetMinmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minmembers"] = v
}

func (p *UpdateAutoScaleVmGroupParams) SetScaledownpolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaledownpolicyids"] = v
}

func (p *UpdateAutoScaleVmGroupParams) SetScaleuppolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaleuppolicyids"] = v
}

// You should always use this function to get a new UpdateAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScaleVmGroupParams(id string) *UpdateAutoScaleVmGroupParams {
	p := &UpdateAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing autoscale vm group.
func (s *AutoScaleService) UpdateAutoScaleVmGroup(p *UpdateAutoScaleVmGroupParams) (*UpdateAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("updateAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScaleVmGroupResponse
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

type UpdateAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type UpdateAutoScaleVmProfileParams struct {
	p map[string]interface{}
}

func (p *UpdateAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["autoscaleuserid"]; found {
		u.Set("autoscaleuserid", v.(string))
	}
	if v, found := p.p["counterparam"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("counterparam[%d].key", i), k)
			u.Set(fmt.Sprintf("counterparam[%d].value", i), m[k])
		}
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["destroyvmgraceperiod"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("destroyvmgraceperiod", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	return u
}

func (p *UpdateAutoScaleVmProfileParams) SetAutoscaleuserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscaleuserid"] = v
}

func (p *UpdateAutoScaleVmProfileParams) SetCounterparam(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterparam"] = v
}

func (p *UpdateAutoScaleVmProfileParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateAutoScaleVmProfileParams) SetDestroyvmgraceperiod(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destroyvmgraceperiod"] = v
}

func (p *UpdateAutoScaleVmProfileParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateAutoScaleVmProfileParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAutoScaleVmProfileParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

// You should always use this function to get a new UpdateAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScaleVmProfileParams(id string) *UpdateAutoScaleVmProfileParams {
	p := &UpdateAutoScaleVmProfileParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing autoscale vm profile.
func (s *AutoScaleService) UpdateAutoScaleVmProfile(p *UpdateAutoScaleVmProfileParams) (*UpdateAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newRequest("updateAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScaleVmProfileResponse
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

type UpdateAutoScaleVmProfileResponse struct {
	Account              string `json:"account"`
	Autoscaleuserid      string `json:"autoscaleuserid"`
	Destroyvmgraceperiod int    `json:"destroyvmgraceperiod"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Otherdeployparams    string `json:"otherdeployparams"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Serviceofferingid    string `json:"serviceofferingid"`
	Templateid           string `json:"templateid"`
	Zoneid               string `json:"zoneid"`
}
