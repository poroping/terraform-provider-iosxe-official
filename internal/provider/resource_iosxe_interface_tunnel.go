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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

func NewInterfaceTunnelResource() resource.Resource {
	return &InterfaceTunnelResource{}
}

type InterfaceTunnelResource struct {
	clients map[string]*restconf.Client
}

func (r *InterfaceTunnelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_tunnel"
}

func (r *InterfaceTunnelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Interface Tunnel configuration.",

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
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"name": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(0, 4294967295).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
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
			"ipv6_address_prefix_lists": schema.ListNestedAttribute{
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
			"tunnel_source": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("source of tunnel packets").String,
				Optional:            true,
			},
			"tunnel_destination_ipv4": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ip address or host name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`), ""),
				},
			},
			"tunnel_protection_ipsec_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Determine the ipsec policy profile to use.").String,
				Optional:            true,
			},
			"crypto_ipsec_df_bit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Handling of encapsulated DF bit.").AddStringEnumDescription("clear", "copy", "set").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("clear", "copy", "set"),
				},
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
			"tunnel_mode_ipsec_ipv4": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("over IPv4").String,
				Optional:            true,
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
		},
	}
}

func (r *InterfaceTunnelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (r *InterfaceTunnelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceTunnel

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

func (r *InterfaceTunnelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceTunnel

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
		state = InterfaceTunnel{Device: state.Device, Id: state.Id}
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

func (r *InterfaceTunnelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceTunnel

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

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	if YangPatch {
		edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
		for _, i := range deletedListItems {
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
		for _, i := range deletedListItems {
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

func (r *InterfaceTunnelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceTunnel

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
	deleteMode := "all"
	if state.DeleteMode.ValueString() == "all" {
		deleteMode = "all"
	} else if state.DeleteMode.ValueString() == "attributes" {
		deleteMode = "attributes"
	}

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

func (r *InterfaceTunnelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
