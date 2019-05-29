// Copyright 2018 The Terraformer Authors.
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

package datadog

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

const datadogProviderVersion = ">1.9.0"

// DatadogProvider ...
type DatadogProvider struct {
	terraform_utils.Provider
}

// Init check env params
func (p *DatadogProvider) Init(args []string) error {
	return nil
}

// GetName return string of provider name for Datadog
func (p *DatadogProvider) GetName() string {
	return "datadog"
}

// InitService ...
func (p *DatadogProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("datadog: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]string{})
	return nil
}

// GetSupportedService return map of support service for Datadog
func (p *DatadogProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"monitor": &MonitorGenerator{},
	}
}

// GetResourceConnections return map of resource connections for Datadog
func (DatadogProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

// GetProviderData return map of provider data for Datadog
func (p DatadogProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				"version": datadogProviderVersion,
			},
		},
	}
}
