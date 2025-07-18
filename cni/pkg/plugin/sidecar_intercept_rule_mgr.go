// Copyright Istio Authors
//
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

package plugin

// InterceptRuleMgr configures networking tables (e.g. iptables or nftables) for
// redirecting traffic to an Istio proxy.
type InterceptRuleMgr interface {
	Program(podName, netns string, redirect *Redirect) error
}

// Constructor for iptables InterceptRuleMgr
func IptablesInterceptRuleMgr() InterceptRuleMgr {
	return newIPTables()
}

// Constructor for nftables InterceptRuleMgr
func NftablesInterceptRuleMgr() InterceptRuleMgr {
	return newNFTables()
}
