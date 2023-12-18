// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeAAAAuthorization(t *testing.T) {
	if os.Getenv("AAA") == "" {
		t.Skip("skipping test, set environment variable AAA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "execs.0.name", "EXEC1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "execs.0.a1_group", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "execs.0.a2_group", "GROUP2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "execs.0.a3_group", "GROUP3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "execs.0.a4_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "networks.0.id", "NET1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "networks.0.a1_group", "RGROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "networks.0.a2_group", "RGROUP2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "networks.0.a3_group", "RGROUP3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_aaa_authorization.test", "networks.0.a4_local", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeAAAAuthorizationConfig_minimum(),
			},
			{
				Config: testAccIosxeAAAAuthorizationConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_aaa_authorization.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:authorization",
			},
		},
	})
}

func testAccIosxeAAAAuthorizationConfig_minimum() string {
	config := `resource "iosxe_aaa_authorization" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeAAAAuthorizationConfig_all() string {
	config := `resource "iosxe_aaa_authorization" "test" {` + "\n"
	config += `	execs = [{` + "\n"
	config += `		name = "EXEC1"` + "\n"
	config += `		a1_group = "GROUP1"` + "\n"
	config += `		a2_group = "GROUP2"` + "\n"
	config += `		a3_group = "GROUP3"` + "\n"
	config += `		a4_local = true` + "\n"
	config += `	}]` + "\n"
	config += `	networks = [{` + "\n"
	config += `		id = "NET1"` + "\n"
	config += `		a1_group = "RGROUP1"` + "\n"
	config += `		a2_group = "RGROUP2"` + "\n"
	config += `		a3_group = "RGROUP3"` + "\n"
	config += `		a4_local = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
