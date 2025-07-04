/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package config

import (
	"crypto/x509"
	"path/filepath"
	"strings"
)

func hostPaths(root, host string) (hosts []string) {
	ch := hostDirectory(host)
	if ch != host {
		hosts = append(hosts, filepath.Join(root, strings.ReplaceAll(ch, ":", "")))
	}

	hosts = append(hosts,
		filepath.Join(root, strings.ReplaceAll(host, ":", "")),
		filepath.Join(root, "_default"),
	)

	return
}

func rootSystemPool() (*x509.CertPool, error) {
	return x509.NewCertPool(), nil
}
