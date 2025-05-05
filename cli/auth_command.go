// Copyright 2023-2024 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

func configureAuthCommand(app commandHost) {
	auth := app.Command("auth", "NATS Decentralized Authentication")
	addCheat("auth", auth)

	// todo:
	//  - Improve maintaining pub/sub permissions for a user, perhaps allow interactive edits of yaml?

	auth.HelpLong("WARNING: This is experimental and subject to change, do not use yet for production deployment. ")

	configureAuthOperatorCommand(auth)
	configureAuthAccountCommand(auth)
	configureAuthUserCommand(auth)
	configureAuthNkeyCommand(auth)
}

func init() {
	registerCommand("auth", 0, configureAuthCommand)
}
