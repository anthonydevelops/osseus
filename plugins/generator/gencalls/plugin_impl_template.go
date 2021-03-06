// Copyright (c) 2019 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gencalls

// text/template for plugin_impl.go for custom plugin files
const pluginImplTemplate = `
/*** This is an autogenerated file; This file can be edited but changes should be saved***/

package {{.PackageName}}

import (
    "github.com/ligato/cn-infra/infra"
    "github.com/ligato/cn-infra/logging"
	// todo: add any necessary imports for your plugin
)

// RegisterFlags registers command line flags.
func RegisterFlags() {
    // todo: add command line flags here if needed
}

func init() {
    RegisterFlags()
}

// Plugin holds the internal data structures of the Rest Plugin
type Plugin struct {
    Deps
}

// Deps groups the dependencies of the Rest Plugin.
type Deps struct {
    infra.PluginDeps
	// todo: add any additional dependencies here
}

// Init initializes the Plugin
func (p *Plugin) Init() error {
    p.Log.SetLevel(logging.DebugLevel)
    return nil
}

// AfterInit can be used to run plugin functionality.
func (p *Plugin) AfterInit() (err error) {
    return nil
}

// Close is NOOP.
func (p *Plugin) Close() error {
    return nil
}`