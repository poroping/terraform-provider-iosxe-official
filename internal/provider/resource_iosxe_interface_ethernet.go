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
	"context"
	"fmt"
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

func NewInterfaceEthernetResource() resource.Resource {
	return &InterfaceEthernetResource{}
}

type InterfaceEthernetResource struct {
	clients map[string]*restconf.Client
}

func (r *InterfaceEthernetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet"
}

func (r *InterfaceEthernetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Interface Ethernet configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface type").AddStringEnumDescription("GigabitEthernet", "TwoGigabitEthernet", "FiveGigabitEthernet", "TenGigabitEthernet", "TwentyFiveGigE", "FortyGigabitEthernet", "HundredGigE", "TwoHundredGigE", "FourHundredGigE").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("GigabitEthernet", "TwoGigabitEthernet", "FiveGigabitEthernet", "TenGigabitEthernet", "TwentyFiveGigE", "FortyGigabitEthernet", "HundredGigE", "TwoHundredGigE", "FourHundredGigE"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(0|[1-9][0-9]*)(/(0|[1-9][0-9]*))*(\.[0-9]*)?`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"media_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Media type").AddStringEnumDescription("auto-select", "rj45", "sfp").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("auto-select", "rj45", "sfp"),
				},
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 200000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 200000000),
				},
			},
			"switchport": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface specific description").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
					stringvalidator.RegexMatches(regexp.MustCompile(`.*`), ""),
				},
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shutdown the selected interface").String,
				Optional:            true,
			},
			"ip_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable proxy ARP").String,
				Optional:            true,
			},
			"ip_redirects": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable sending ICMP Redirect messages").String,
				Optional:            true,
			},
			"ip_unreachables": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable sending ICMP Unreachable messages").String,
				Optional:            true,
			},
			"vrf_forwarding": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure forwarding table").String,
				Optional:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
				},
			},
			"ipv4_address_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
				},
			},
			"unnumbered": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IP processing without an explicit address").String,
				Optional:            true,
			},
			"encapsulation_dot1q_vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"channel_group_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 512).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 512),
				},
			},
			"channel_group_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Etherchannel Mode of the interface").AddStringEnumDescription("active", "auto", "desirable", "on", "passive").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("active", "auto", "desirable", "on", "passive"),
				},
			},
			"ip_dhcp_relay_source_interface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set source interface for relayed messages").String,
				Optional:            true,
			},
			"ip_access_group_in": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"ip_access_group_in_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("inbound packets").String,
				Optional:            true,
			},
			"ip_access_group_out": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"ip_access_group_out_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("outbound packets").String,
				Optional:            true,
			},
			"spanning_tree_guard": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Change an interface's spanning tree guard mode").AddStringEnumDescription("loop", "none", "root").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("loop", "none", "root"),
				},
			},
			"auto_qos_classify": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure classification for untrusted devices").String,
				Optional:            true,
			},
			"auto_qos_classify_police": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure QoS policing for untrusted devices").String,
				Optional:            true,
			},
			"auto_qos_trust": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the DSCP/CoS marking").String,
				Optional:            true,
			},
			"auto_qos_trust_cos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the CoS marking").String,
				Optional:            true,
			},
			"auto_qos_trust_dscp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the DSCP marking").String,
				Optional:            true,
			},
			"auto_qos_video_cts": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the QoS marking of the Cisco Telepresence System").String,
				Optional:            true,
			},
			"auto_qos_video_ip_camera": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the QoS marking of the Ip Video Surveillance camera").String,
				Optional:            true,
			},
			"auto_qos_video_media_player": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the Qos marking of the Cisco Media Player").String,
				Optional:            true,
			},
			"auto_qos_voip": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure AutoQoS for VoIP").String,
				Optional:            true,
			},
			"auto_qos_voip_cisco_phone": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the QoS marking of Cisco IP Phone").String,
				Optional:            true,
			},
			"auto_qos_voip_cisco_softphone": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the QoS marking of Cisco IP SoftPhone").String,
				Optional:            true,
			},
			"auto_qos_voip_trust": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Trust the DSCP/CoS marking").String,
				Optional:            true,
			},
			"trust_device": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("trusted device class").AddStringEnumDescription("cisco-phone", "cts", "ip-camera", "media-player").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("cisco-phone", "cts", "ip-camera", "media-player"),
				},
			},
			"helper_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a destination address for UDP broadcasts").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
							},
						},
						"global": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Helper-address is global").String,
							Optional:            true,
						},
						"vrf": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("VRF name for helper-address (if different from interface VRF)").String,
							Optional:            true,
						},
					},
				},
			},
			"source_template": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"template_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Required:            true,
						},
						"merge": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("merge option of binding").String,
							Optional:            true,
						},
					},
				},
			},
			"bfd_template": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BFD template").String,
				Optional:            true,
			},
			"bfd_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable BFD under the interface").String,
				Optional:            true,
			},
			"bfd_local_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The Source IP address to be used for BFD sessions over this interface.").String,
				Optional:            true,
			},
			"bfd_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(50, 9999).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(50, 9999),
				},
			},
			"bfd_interval_min_rx": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum receive interval capability").AddIntegerRangeDescription(50, 9999).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(50, 9999),
				},
			},
			"bfd_interval_multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Multiplier value used to compute holddown").AddIntegerRangeDescription(3, 50).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 50),
				},
			},
			"bfd_echo": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use echo adjunct as bfd detection mechanism").String,
				Optional:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IPv6 on interface").String,
				Optional:            true,
			},
			"ipv6_mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set IPv6 Maximum Transmission Unit").AddIntegerRangeDescription(1280, 9976).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1280, 9976),
				},
			},
			"ipv6_nd_ra_suppress_all": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Suppress all IPv6 RA").String,
				Optional:            true,
			},
			"ipv6_address_autoconfig_default": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Insert default route").String,
				Optional:            true,
			},
			"ipv6_address_dhcp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Obtain IPv6 address from DHCP server").String,
				Optional:            true,
			},
			"ipv6_link_local_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?`), ""),
							},
						},
						"link_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use link-local address").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`), ""),
								stringvalidator.RegexMatches(regexp.MustCompile(`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(/.+)`), ""),
							},
						},
						"eui_64": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use eui-64 interface identifier").String,
							Optional:            true,
						},
					},
				},
			},
			"arp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set ARP cache timeout").AddIntegerRangeDescription(0, 2147483).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483),
				},
			},
			"spanning_tree_link_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a link type for spanning tree tree protocol use").AddStringEnumDescription("point-to-point", "shared").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("point-to-point", "shared"),
				},
			},
			"spanning_tree_portfast_trunk": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable portfast on the interface even in trunk mode").String,
				Optional:            true,
			},
			"ip_arp_inspection_trust": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Trust state").String,
				Optional:            true,
			},
			"ip_arp_inspection_limit_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rate Limit").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"ip_dhcp_snooping_trust": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DHCP Snooping trust config").String,
				Optional:            true,
			},
			"speed_100": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("100 Mbps operation").String,
				Optional:            true,
			},
			"speed_1000": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("1000 Mbps operation").String,
				Optional:            true,
			},
			"speed_2500": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("2500 Mbps operation").String,
				Optional:            true,
			},
			"speed_5000": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("5000 Mbps operation").String,
				Optional:            true,
			},
			"speed_10000": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("10000 Mbps operation").String,
				Optional:            true,
			},
			"speed_25000": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("25000 Mbps operation").String,
				Optional:            true,
			},
			"speed_40000": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("40000 Mbps operation").String,
				Optional:            true,
			},
			"speed_100000": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("100000 Mbps operation").String,
				Optional:            true,
			},
			"negotiation_auto": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable link autonegotiation").String,
				Optional:            true,
			},
			"speed_nonegotiate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"authentication_host_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the Host mode for authentication on this interface").AddStringEnumDescription("multi-auth", "multi-domain", "multi-host", "single-host").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("multi-auth", "multi-domain", "multi-host", "single-host"),
				},
			},
			"authentication_order_dot1x": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method dot1x allowed").String,
				Optional:            true,
			},
			"authentication_order_dot1x_mab": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method mab allowed").String,
				Optional:            true,
			},
			"authentication_order_dot1x_webauth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method webauth allowed").String,
				Optional:            true,
			},
			"authentication_order_mab": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method mab allowed").String,
				Optional:            true,
			},
			"authentication_order_mab_dot1x": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method dot1x allowed").String,
				Optional:            true,
			},
			"authentication_order_mab_webauth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method webauth allowed").String,
				Optional:            true,
			},
			"authentication_order_webauth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method webauth allowed").String,
				Optional:            true,
			},
			"authentication_priority_dot1x": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method dot1x allowed").String,
				Optional:            true,
			},
			"authentication_priority_dot1x_mab": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method mab allowed").String,
				Optional:            true,
			},
			"authentication_priority_dot1x_webauth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method webauth allowed").String,
				Optional:            true,
			},
			"authentication_priority_mab": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method mab allowed").String,
				Optional:            true,
			},
			"authentication_priority_mab_dot1x": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method dot1x allowed").String,
				Optional:            true,
			},
			"authentication_priority_mab_webauth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method webauth allowed").String,
				Optional:            true,
			},
			"authentication_priority_webauth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication method webauth allowed").String,
				Optional:            true,
			},
			"authentication_port_control": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("set the port-control value").AddStringEnumDescription("auto", "force-authorized", "force-unauthorized").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "force-authorized", "force-unauthorized"),
				},
			},
			"authentication_periodic": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or Disable Reauthentication for this port").String,
				Optional:            true,
			},
			"authentication_timer_reauthenticate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter a value between 1 and 1073741823").AddIntegerRangeDescription(1, 1073741823).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1073741823),
				},
			},
			"authentication_timer_reauthenticate_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Obtain re-authentication timeout value from the server").String,
				Optional:            true,
			},
			"mab": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MAC Authentication Bypass Interface Config Commands").String,
				Optional:            true,
			},
			"mab_eap": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use EAP authentication for MAC Auth Bypass").String,
				Optional:            true,
			},
			"dot1x_pae": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set 802.1x interface pae type").AddStringEnumDescription("authenticator", "both", "supplicant").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("authenticator", "both", "supplicant"),
				},
			},
			"dot1x_timeout_auth_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout for authenticator reply").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_timeout_held_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout for authentication retries").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_timeout_quiet_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("QuietPeriod in Seconds").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_timeout_ratelimit_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ratelimit Period in seconds").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_timeout_server_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout for Radius Retries").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_timeout_start_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout for EAPOL-start retries").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_timeout_supp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout for supplicant reply").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_timeout_tx_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout for supplicant retries").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"dot1x_max_req": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Max No. of Retries").AddIntegerRangeDescription(1, 10).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
			},
			"dot1x_max_reauth_req": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Max No. of Reauthentication Attempts").AddIntegerRangeDescription(1, 10).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
			},
			"service_policy_input": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign policy-map to the input of an interface").String,
				Optional:            true,
			},
			"service_policy_output": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign policy-map to the output of an interface").String,
				Optional:            true,
			},
			"ip_flow_monitors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply a Flow Monitor").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User defined").String,
							Required:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("input", "output").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("input", "output"),
							},
						},
					},
				},
			},
			"load_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify interval for load calculation for an interface").AddIntegerRangeDescription(30, 600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(30, 600),
				},
			},
			"snmp_trap_link_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow SNMP LINKUP and LINKDOWN traps").String,
				Optional:            true,
			},
			"logging_event_link_status_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UPDOWN and CHANGE messages").String,
				Optional:            true,
			},
		},
	}
}

func (r *InterfaceEthernetResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *InterfaceEthernetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernet

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[plan.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range emptyLeafsDelete {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		_, err := r.clients[plan.Device.ValueString()].YangPatchData("", "1", "", edits)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object, got error: %s", err))
			return
		}
	} else {
		res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
		if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
			_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
		}
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
		for _, i := range emptyLeafsDelete {
			res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InterfaceEthernetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernet

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[state.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.clients[state.Device.ValueString()].GetData(state.Id.ValueString())
	if res.StatusCode == 404 {
		state = InterfaceEthernet{Device: state.Device, Id: state.Id}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.updateFromBody(ctx, res.Res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *InterfaceEthernetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEthernet

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[plan.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx)

	deletedItems := plan.getDeletedItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedItems))

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range deletedItems {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		for _, i := range emptyLeafsDelete {
			edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
		}
		_, err := r.clients[plan.Device.ValueString()].YangPatchData("", "1", "", edits)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	} else {
		res, err := r.clients[plan.Device.ValueString()].PatchData(plan.getPathShort(), body)
		if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
			_, err = r.clients[plan.Device.ValueString()].PutData(plan.getPath(), body)
		}
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
		for _, i := range deletedItems {
			res, err := r.clients[state.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
		for _, i := range emptyLeafsDelete {
			res, err := r.clients[plan.Device.ValueString()].DeleteData(i)
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InterfaceEthernetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernet

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := r.clients[state.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	deleteMode := "attributes"

	if deleteMode == "all" {
		res, err := r.clients[state.Device.ValueString()].DeleteData(state.Id.ValueString())
		if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
			return
		}
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		if YangPatch {
			edits := []restconf.YangPatchEdit{}
			for _, i := range deletePaths {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			_, err := r.clients[state.Device.ValueString()].YangPatchData("", "1", "", edits)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		} else {
			for _, i := range deletePaths {
				res, err := r.clients[state.Device.ValueString()].DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *InterfaceEthernetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
