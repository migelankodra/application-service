//
// Copyright (c) 2021 Intel Corporation
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
//

package transforms

import (
	"os"
	"testing"

	"github.com/migelankodra/application-service/v2/internal/appfunction"
	"github.com/migelankodra/application-service/v2/internal/bootstrap/container"
	"github.com/migelankodra/application-service/v2/internal/common"

	bootstrapContainer "github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
)

var lc logger.LoggingClient
var dic *di.Container
var ctx *appfunction.Context

func TestMain(m *testing.M) {
	lc = logger.NewMockClient()
	eventClient := http.NewEventClient("http://test")

	config := &common.ConfigurationStruct{}

	dic = di.NewContainer(di.ServiceConstructorMap{
		container.ConfigurationName: func(get di.Get) interface{} {
			return config
		},
		container.EventClientName: func(get di.Get) interface{} {
			return eventClient
		},
		bootstrapContainer.LoggingClientInterfaceName: func(get di.Get) interface{} {
			return lc
		},
	})

	ctx = appfunction.NewContext("123", dic, "")

	os.Exit(m.Run())
}
