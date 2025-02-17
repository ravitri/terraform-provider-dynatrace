/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package pvc_test

import (
	"testing"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api/builtin/anomalydetection/kubernetes/pvc"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/testing/api"
)

func TestK8sPVCAnomalies(t *testing.T) {
	api.TestService(t, pvc.Service)
}

func TestAccK8sPVCAnomalies(t *testing.T) {
	api.TestAcc(t)
}
