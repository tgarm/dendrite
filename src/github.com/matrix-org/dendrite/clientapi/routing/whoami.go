// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package routing

import (
	"net/http"

	"github.com/matrix-org/dendrite/clientapi/auth/authtypes"
	"github.com/matrix-org/util"
)

// whoamiResponse represents an response for a `whoami` request
type whoamiResponse struct {
	UserID string `json:"user_id"`
}

// Whoami implements `/account/whoami` which enables client to query their account user id.
// https://matrix.org/docs/spec/client_server/r0.3.0.html#get-matrix-client-r0-account-whoami
func Whoami(req *http.Request, device *authtypes.Device) util.JSONResponse {
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: whoamiResponse{UserID: device.UserID},
	}
}