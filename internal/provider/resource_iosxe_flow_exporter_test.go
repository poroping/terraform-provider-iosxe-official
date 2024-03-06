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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeFlowExporter(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_flow_exporter.test", "name", "EXPORTER1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_flow_exporter.test", "description", "My exporter"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_flow_exporter.test", "destination_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_flow_exporter.test", "source_loopback", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_flow_exporter.test", "transport_udp", "655"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_flow_exporter.test", "template_data_timeout", "60"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeFlowExporterConfig_minimum(),
			},
			{
				Config: testAccIosxeFlowExporterConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_flow_exporter.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/Cisco-IOS-XE-flow:flow/exporter=EXPORTER1",
			},
		},
	})
}

func testAccIosxeFlowExporterConfig_minimum() string {
	config := `resource "iosxe_flow_exporter" "test" {` + "\n"
	config += `	name = "EXPORTER1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeFlowExporterConfig_all() string {
	config := `resource "iosxe_flow_exporter" "test" {` + "\n"
	config += `	name = "EXPORTER1"` + "\n"
	config += `	description = "My exporter"` + "\n"
	config += `	destination_ip = "1.1.1.1"` + "\n"
	config += `	source_loopback = 123` + "\n"
	config += `	transport_udp = 655` + "\n"
	config += `	template_data_timeout = 60` + "\n"
	config += `}` + "\n"
	return config
}
