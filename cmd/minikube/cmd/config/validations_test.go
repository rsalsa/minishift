/*
Copyright (C) 2016 Red Hat, Inc.

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

import "testing"

type validationTest struct {
	value     string
	shouldErr bool
}

func runValidations(t *testing.T, tests []validationTest, name string, f func(string, string) error) {
	for _, tt := range tests {
		err := f(name, tt.value)
		if err != nil && !tt.shouldErr {
			t.Errorf("%s: %v", tt.value, err)
		}
		if err == nil && tt.shouldErr {
			t.Errorf("%s: %v", tt.value, err)
		}
	}
}

func TestDriver(t *testing.T) {

	var tests = []validationTest{
		{
			value:     "vkasdhfasjdf",
			shouldErr: true,
		},
		{
			value:     "",
			shouldErr: true,
		},
	}

	runValidations(t, tests, "vm-driver", IsValidDriver)

}

func TestValidCIDR(t *testing.T) {
	var tests = []validationTest{
		{
			value:     "0.0.0.0/0",
			shouldErr: false,
		},
		{
			value:     "1.1.1.1/32",
			shouldErr: false,
		},
		{
			value:     "192.168.0.0/16",
			shouldErr: false,
		},
		{
			value:     "255.255.255.255/1",
			shouldErr: false,
		},
		{
			value:     "8.8.8.8/33",
			shouldErr: true,
		},
		{
			value:     "12.1",
			shouldErr: true,
		},
		{
			value:     "1",
			shouldErr: true,
		},
		{
			value:     "a string!",
			shouldErr: true,
		},
		{
			value:     "192.168.1.1/8/",
			shouldErr: true,
		},
	}

	runValidations(t, tests, "cidr", IsValidCIDR)
}

func TestValidURL(t *testing.T) {

	var tests = []validationTest{
		{
			value:	   "",
			shouldErr: true,
		},
		{
			value:     "http/foo.com/minishift.tar.gz",
			shouldErr: true,
		},
		{
			value:      "foo/download/minishift.tar.gz",
			shouldErr:  true,
		},
		{
			value:      "/absolute/path/no/protocol/minishift.tar.gz",
			shouldErr:  false,
		},
		{
			value:	   "http://foo.com/minishift.tar.gz",
			shouldErr: false,
		},
		{
			value:     "file:///foo/download/minishift.tar.gz",
			shouldErr: false,
		},

	}
	runValidations(t, tests, "iso-url", IsValidUrl)
}
