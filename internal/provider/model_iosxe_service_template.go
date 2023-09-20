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
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ServiceTemplate struct {
	Device                  types.String                       `tfsdk:"device"`
	Id                      types.String                       `tfsdk:"id"`
	Word                    types.String                       `tfsdk:"word"`
	AccessGroup             []ServiceTemplateAccessGroup       `tfsdk:"access_group"`
	InactivityTimerValue    types.Int64                        `tfsdk:"inactivity_timer_value"`
	InactivityTimerProbe    types.Bool                         `tfsdk:"inactivity_timer_probe"`
	Vlan                    types.Int64                        `tfsdk:"vlan"`
	VoiceVlan               types.Bool                         `tfsdk:"voice_vlan"`
	LinksecPolicy           types.String                       `tfsdk:"linksec_policy"`
	Sgt                     types.Int64                        `tfsdk:"sgt"`
	AbsoluteTimer           types.Int64                        `tfsdk:"absolute_timer"`
	Description             types.String                       `tfsdk:"description"`
	InterfaceTemplate       []ServiceTemplateInterfaceTemplate `tfsdk:"interface_template"`
	TunnelCapwapName        types.String                       `tfsdk:"tunnel_capwap_name"`
	Vnid                    types.String                       `tfsdk:"vnid"`
	RedirectAppendClientMac types.String                       `tfsdk:"redirect_append_client_mac"`
	RedirectAppendSwitchMac types.String                       `tfsdk:"redirect_append_switch_mac"`
	RedirectUrlUrlName      types.String                       `tfsdk:"redirect_url_url_name"`
	RedirectUrlMatchAclName types.String                       `tfsdk:"redirect_url_match_acl_name"`
	RedirectUrlMatchAction  types.String                       `tfsdk:"redirect_url_match_action"`
	DnsAclPreauth           types.String                       `tfsdk:"dns_acl_preauth"`
	ServicePolicyQosInput   types.String                       `tfsdk:"service_policy_qos_input"`
	ServicePolicyQosOutput  types.String                       `tfsdk:"service_policy_qos_output"`
	TagConfig               []ServiceTemplateTagConfig         `tfsdk:"tag_config"`
	MdnsServicePolicy       types.String                       `tfsdk:"mdns_service_policy"`
}

type ServiceTemplateData struct {
	Device                  types.String                       `tfsdk:"device"`
	Id                      types.String                       `tfsdk:"id"`
	Word                    types.String                       `tfsdk:"word"`
	AccessGroup             []ServiceTemplateAccessGroup       `tfsdk:"access_group"`
	InactivityTimerValue    types.Int64                        `tfsdk:"inactivity_timer_value"`
	InactivityTimerProbe    types.Bool                         `tfsdk:"inactivity_timer_probe"`
	Vlan                    types.Int64                        `tfsdk:"vlan"`
	VoiceVlan               types.Bool                         `tfsdk:"voice_vlan"`
	LinksecPolicy           types.String                       `tfsdk:"linksec_policy"`
	Sgt                     types.Int64                        `tfsdk:"sgt"`
	AbsoluteTimer           types.Int64                        `tfsdk:"absolute_timer"`
	Description             types.String                       `tfsdk:"description"`
	InterfaceTemplate       []ServiceTemplateInterfaceTemplate `tfsdk:"interface_template"`
	TunnelCapwapName        types.String                       `tfsdk:"tunnel_capwap_name"`
	Vnid                    types.String                       `tfsdk:"vnid"`
	RedirectAppendClientMac types.String                       `tfsdk:"redirect_append_client_mac"`
	RedirectAppendSwitchMac types.String                       `tfsdk:"redirect_append_switch_mac"`
	RedirectUrlUrlName      types.String                       `tfsdk:"redirect_url_url_name"`
	RedirectUrlMatchAclName types.String                       `tfsdk:"redirect_url_match_acl_name"`
	RedirectUrlMatchAction  types.String                       `tfsdk:"redirect_url_match_action"`
	DnsAclPreauth           types.String                       `tfsdk:"dns_acl_preauth"`
	ServicePolicyQosInput   types.String                       `tfsdk:"service_policy_qos_input"`
	ServicePolicyQosOutput  types.String                       `tfsdk:"service_policy_qos_output"`
	TagConfig               []ServiceTemplateTagConfig         `tfsdk:"tag_config"`
	MdnsServicePolicy       types.String                       `tfsdk:"mdns_service_policy"`
}
type ServiceTemplateAccessGroup struct {
	Name types.String `tfsdk:"name"`
}
type ServiceTemplateInterfaceTemplate struct {
	Name types.String `tfsdk:"name"`
}
type ServiceTemplateTagConfig struct {
	Name types.String `tfsdk:"name"`
}

func (data ServiceTemplate) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/Cisco-IOS-XE-switch:service-template=%v", url.QueryEscape(fmt.Sprintf("%v", data.Word.ValueString())))
}

func (data ServiceTemplateData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/Cisco-IOS-XE-switch:service-template=%v", url.QueryEscape(fmt.Sprintf("%v", data.Word.ValueString())))
}

// if last path element has a key -> remove it
func (data ServiceTemplate) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data ServiceTemplate) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Word.IsNull() && !data.Word.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"word", data.Word.ValueString())
	}
	if !data.InactivityTimerValue.IsNull() && !data.InactivityTimerValue.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"inactivity-timer.value", strconv.FormatInt(data.InactivityTimerValue.ValueInt64(), 10))
	}
	if !data.InactivityTimerProbe.IsNull() && !data.InactivityTimerProbe.IsUnknown() {
		if data.InactivityTimerProbe.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"inactivity-timer.probe", map[string]string{})
		}
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan", strconv.FormatInt(data.Vlan.ValueInt64(), 10))
	}
	if !data.VoiceVlan.IsNull() && !data.VoiceVlan.IsUnknown() {
		if data.VoiceVlan.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"voice.vlan", map[string]string{})
		}
	}
	if !data.LinksecPolicy.IsNull() && !data.LinksecPolicy.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"linksec.policy", data.LinksecPolicy.ValueString())
	}
	if !data.Sgt.IsNull() && !data.Sgt.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"sgt", strconv.FormatInt(data.Sgt.ValueInt64(), 10))
	}
	if !data.AbsoluteTimer.IsNull() && !data.AbsoluteTimer.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"absolute-timer", strconv.FormatInt(data.AbsoluteTimer.ValueInt64(), 10))
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.TunnelCapwapName.IsNull() && !data.TunnelCapwapName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tunnel.type.capwap.name", data.TunnelCapwapName.ValueString())
	}
	if !data.Vnid.IsNull() && !data.Vnid.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vnid", data.Vnid.ValueString())
	}
	if !data.RedirectAppendClientMac.IsNull() && !data.RedirectAppendClientMac.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"redirect.append.client-mac", data.RedirectAppendClientMac.ValueString())
	}
	if !data.RedirectAppendSwitchMac.IsNull() && !data.RedirectAppendSwitchMac.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"redirect.append.switch-mac", data.RedirectAppendSwitchMac.ValueString())
	}
	if !data.RedirectUrlUrlName.IsNull() && !data.RedirectUrlUrlName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"redirect.url.url_name", data.RedirectUrlUrlName.ValueString())
	}
	if !data.RedirectUrlMatchAclName.IsNull() && !data.RedirectUrlMatchAclName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"redirect.url.match.acl_name", data.RedirectUrlMatchAclName.ValueString())
	}
	if !data.RedirectUrlMatchAction.IsNull() && !data.RedirectUrlMatchAction.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"redirect.url.match.action", data.RedirectUrlMatchAction.ValueString())
	}
	if !data.DnsAclPreauth.IsNull() && !data.DnsAclPreauth.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"dns-acl.preauth", data.DnsAclPreauth.ValueString())
	}
	if !data.ServicePolicyQosInput.IsNull() && !data.ServicePolicyQosInput.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"service-policy.qos.input", data.ServicePolicyQosInput.ValueString())
	}
	if !data.ServicePolicyQosOutput.IsNull() && !data.ServicePolicyQosOutput.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"service-policy.qos.output", data.ServicePolicyQosOutput.ValueString())
	}
	if !data.MdnsServicePolicy.IsNull() && !data.MdnsServicePolicy.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mdns-service-policy", data.MdnsServicePolicy.ValueString())
	}
	if len(data.AccessGroup) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-group-config", []interface{}{})
		for index, item := range data.AccessGroup {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"access-group-config"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
		}
	}
	if len(data.InterfaceTemplate) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interface-template", []interface{}{})
		for index, item := range data.InterfaceTemplate {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"interface-template"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
		}
	}
	if len(data.TagConfig) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tag-config", []interface{}{})
		for index, item := range data.TagConfig {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"tag-config"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
		}
	}
	return body
}

func (data *ServiceTemplate) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "word"); value.Exists() && !data.Word.IsNull() {
		data.Word = types.StringValue(value.String())
	} else {
		data.Word = types.StringNull()
	}
	for i := range data.AccessGroup {
		keys := [...]string{"name"}
		keyValues := [...]string{data.AccessGroup[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "access-group-config").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.AccessGroup[i].Name.IsNull() {
			data.AccessGroup[i].Name = types.StringValue(value.String())
		} else {
			data.AccessGroup[i].Name = types.StringNull()
		}
	}
	if value := res.Get(prefix + "inactivity-timer.value"); value.Exists() && !data.InactivityTimerValue.IsNull() {
		data.InactivityTimerValue = types.Int64Value(value.Int())
	} else {
		data.InactivityTimerValue = types.Int64Null()
	}
	if value := res.Get(prefix + "inactivity-timer.probe"); !data.InactivityTimerProbe.IsNull() {
		if value.Exists() {
			data.InactivityTimerProbe = types.BoolValue(true)
		} else {
			data.InactivityTimerProbe = types.BoolValue(false)
		}
	} else {
		data.InactivityTimerProbe = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan"); value.Exists() && !data.Vlan.IsNull() {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
	if value := res.Get(prefix + "voice.vlan"); !data.VoiceVlan.IsNull() {
		if value.Exists() {
			data.VoiceVlan = types.BoolValue(true)
		} else {
			data.VoiceVlan = types.BoolValue(false)
		}
	} else {
		data.VoiceVlan = types.BoolNull()
	}
	if value := res.Get(prefix + "linksec.policy"); value.Exists() && !data.LinksecPolicy.IsNull() {
		data.LinksecPolicy = types.StringValue(value.String())
	} else {
		data.LinksecPolicy = types.StringNull()
	}
	if value := res.Get(prefix + "sgt"); value.Exists() && !data.Sgt.IsNull() {
		data.Sgt = types.Int64Value(value.Int())
	} else {
		data.Sgt = types.Int64Null()
	}
	if value := res.Get(prefix + "absolute-timer"); value.Exists() && !data.AbsoluteTimer.IsNull() {
		data.AbsoluteTimer = types.Int64Value(value.Int())
	} else {
		data.AbsoluteTimer = types.Int64Null()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.InterfaceTemplate {
		keys := [...]string{"name"}
		keyValues := [...]string{data.InterfaceTemplate[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "interface-template").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.InterfaceTemplate[i].Name.IsNull() {
			data.InterfaceTemplate[i].Name = types.StringValue(value.String())
		} else {
			data.InterfaceTemplate[i].Name = types.StringNull()
		}
	}
	if value := res.Get(prefix + "tunnel.type.capwap.name"); value.Exists() && !data.TunnelCapwapName.IsNull() {
		data.TunnelCapwapName = types.StringValue(value.String())
	} else {
		data.TunnelCapwapName = types.StringNull()
	}
	if value := res.Get(prefix + "vnid"); value.Exists() && !data.Vnid.IsNull() {
		data.Vnid = types.StringValue(value.String())
	} else {
		data.Vnid = types.StringNull()
	}
	if value := res.Get(prefix + "redirect.append.client-mac"); value.Exists() && !data.RedirectAppendClientMac.IsNull() {
		data.RedirectAppendClientMac = types.StringValue(value.String())
	} else {
		data.RedirectAppendClientMac = types.StringNull()
	}
	if value := res.Get(prefix + "redirect.append.switch-mac"); value.Exists() && !data.RedirectAppendSwitchMac.IsNull() {
		data.RedirectAppendSwitchMac = types.StringValue(value.String())
	} else {
		data.RedirectAppendSwitchMac = types.StringNull()
	}
	if value := res.Get(prefix + "redirect.url.url_name"); value.Exists() && !data.RedirectUrlUrlName.IsNull() {
		data.RedirectUrlUrlName = types.StringValue(value.String())
	} else {
		data.RedirectUrlUrlName = types.StringNull()
	}
	if value := res.Get(prefix + "redirect.url.match.acl_name"); value.Exists() && !data.RedirectUrlMatchAclName.IsNull() {
		data.RedirectUrlMatchAclName = types.StringValue(value.String())
	} else {
		data.RedirectUrlMatchAclName = types.StringNull()
	}
	if value := res.Get(prefix + "redirect.url.match.action"); value.Exists() && !data.RedirectUrlMatchAction.IsNull() {
		data.RedirectUrlMatchAction = types.StringValue(value.String())
	} else {
		data.RedirectUrlMatchAction = types.StringNull()
	}
	if value := res.Get(prefix + "dns-acl.preauth"); value.Exists() && !data.DnsAclPreauth.IsNull() {
		data.DnsAclPreauth = types.StringValue(value.String())
	} else {
		data.DnsAclPreauth = types.StringNull()
	}
	if value := res.Get(prefix + "service-policy.qos.input"); value.Exists() && !data.ServicePolicyQosInput.IsNull() {
		data.ServicePolicyQosInput = types.StringValue(value.String())
	} else {
		data.ServicePolicyQosInput = types.StringNull()
	}
	if value := res.Get(prefix + "service-policy.qos.output"); value.Exists() && !data.ServicePolicyQosOutput.IsNull() {
		data.ServicePolicyQosOutput = types.StringValue(value.String())
	} else {
		data.ServicePolicyQosOutput = types.StringNull()
	}
	for i := range data.TagConfig {
		keys := [...]string{"name"}
		keyValues := [...]string{data.TagConfig[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "tag-config").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.TagConfig[i].Name.IsNull() {
			data.TagConfig[i].Name = types.StringValue(value.String())
		} else {
			data.TagConfig[i].Name = types.StringNull()
		}
	}
	if value := res.Get(prefix + "mdns-service-policy"); value.Exists() && !data.MdnsServicePolicy.IsNull() {
		data.MdnsServicePolicy = types.StringValue(value.String())
	} else {
		data.MdnsServicePolicy = types.StringNull()
	}
}

func (data *ServiceTemplateData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "access-group-config"); value.Exists() {
		data.AccessGroup = make([]ServiceTemplateAccessGroup, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceTemplateAccessGroup{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			data.AccessGroup = append(data.AccessGroup, item)
			return true
		})
	}
	if value := res.Get(prefix + "inactivity-timer.value"); value.Exists() {
		data.InactivityTimerValue = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "inactivity-timer.probe"); value.Exists() {
		data.InactivityTimerProbe = types.BoolValue(true)
	} else {
		data.InactivityTimerProbe = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan"); value.Exists() {
		data.Vlan = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "voice.vlan"); value.Exists() {
		data.VoiceVlan = types.BoolValue(true)
	} else {
		data.VoiceVlan = types.BoolValue(false)
	}
	if value := res.Get(prefix + "linksec.policy"); value.Exists() {
		data.LinksecPolicy = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "sgt"); value.Exists() {
		data.Sgt = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "absolute-timer"); value.Exists() {
		data.AbsoluteTimer = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "interface-template"); value.Exists() {
		data.InterfaceTemplate = make([]ServiceTemplateInterfaceTemplate, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceTemplateInterfaceTemplate{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			data.InterfaceTemplate = append(data.InterfaceTemplate, item)
			return true
		})
	}
	if value := res.Get(prefix + "tunnel.type.capwap.name"); value.Exists() {
		data.TunnelCapwapName = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vnid"); value.Exists() {
		data.Vnid = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "redirect.append.client-mac"); value.Exists() {
		data.RedirectAppendClientMac = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "redirect.append.switch-mac"); value.Exists() {
		data.RedirectAppendSwitchMac = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "redirect.url.url_name"); value.Exists() {
		data.RedirectUrlUrlName = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "redirect.url.match.acl_name"); value.Exists() {
		data.RedirectUrlMatchAclName = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "redirect.url.match.action"); value.Exists() {
		data.RedirectUrlMatchAction = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "dns-acl.preauth"); value.Exists() {
		data.DnsAclPreauth = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "service-policy.qos.input"); value.Exists() {
		data.ServicePolicyQosInput = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "service-policy.qos.output"); value.Exists() {
		data.ServicePolicyQosOutput = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "tag-config"); value.Exists() {
		data.TagConfig = make([]ServiceTemplateTagConfig, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceTemplateTagConfig{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			data.TagConfig = append(data.TagConfig, item)
			return true
		})
	}
	if value := res.Get(prefix + "mdns-service-policy"); value.Exists() {
		data.MdnsServicePolicy = types.StringValue(value.String())
	}
}

func (data *ServiceTemplate) getDeletedItems(ctx context.Context, state ServiceTemplate) []string {
	deletedItems := make([]string, 0)
	for i := range state.AccessGroup {
		stateKeyValues := [...]string{state.AccessGroup[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.AccessGroup[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.AccessGroup {
			found = true
			if state.AccessGroup[i].Name.ValueString() != data.AccessGroup[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/access-group-config=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	if !state.InactivityTimerValue.IsNull() && data.InactivityTimerValue.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/inactivity-timer/value", state.getPath()))
	}
	if !state.InactivityTimerProbe.IsNull() && data.InactivityTimerProbe.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/inactivity-timer/probe", state.getPath()))
	}
	if !state.Vlan.IsNull() && data.Vlan.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan", state.getPath()))
	}
	if !state.VoiceVlan.IsNull() && data.VoiceVlan.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/voice/vlan", state.getPath()))
	}
	if !state.LinksecPolicy.IsNull() && data.LinksecPolicy.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/linksec/policy", state.getPath()))
	}
	if !state.Sgt.IsNull() && data.Sgt.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/sgt", state.getPath()))
	}
	if !state.AbsoluteTimer.IsNull() && data.AbsoluteTimer.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/absolute-timer", state.getPath()))
	}
	if !state.Description.IsNull() && data.Description.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/description", state.getPath()))
	}
	for i := range state.InterfaceTemplate {
		stateKeyValues := [...]string{state.InterfaceTemplate[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.InterfaceTemplate[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.InterfaceTemplate {
			found = true
			if state.InterfaceTemplate[i].Name.ValueString() != data.InterfaceTemplate[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/interface-template=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	if !state.TunnelCapwapName.IsNull() && data.TunnelCapwapName.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/tunnel/type/capwap/name", state.getPath()))
	}
	if !state.Vnid.IsNull() && data.Vnid.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vnid", state.getPath()))
	}
	if !state.RedirectAppendClientMac.IsNull() && data.RedirectAppendClientMac.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/redirect/append/client-mac", state.getPath()))
	}
	if !state.RedirectAppendSwitchMac.IsNull() && data.RedirectAppendSwitchMac.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/redirect/append/switch-mac", state.getPath()))
	}
	if !state.RedirectUrlUrlName.IsNull() && data.RedirectUrlUrlName.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/redirect/url/url_name", state.getPath()))
	}
	if !state.RedirectUrlMatchAclName.IsNull() && data.RedirectUrlMatchAclName.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/redirect/url/match/acl_name", state.getPath()))
	}
	if !state.RedirectUrlMatchAction.IsNull() && data.RedirectUrlMatchAction.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/redirect/url/match/action", state.getPath()))
	}
	if !state.DnsAclPreauth.IsNull() && data.DnsAclPreauth.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/dns-acl/preauth", state.getPath()))
	}
	if !state.ServicePolicyQosInput.IsNull() && data.ServicePolicyQosInput.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/service-policy/qos/input", state.getPath()))
	}
	if !state.ServicePolicyQosOutput.IsNull() && data.ServicePolicyQosOutput.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/service-policy/qos/output", state.getPath()))
	}
	for i := range state.TagConfig {
		stateKeyValues := [...]string{state.TagConfig[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.TagConfig[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TagConfig {
			found = true
			if state.TagConfig[i].Name.ValueString() != data.TagConfig[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/tag-config=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	if !state.MdnsServicePolicy.IsNull() && data.MdnsServicePolicy.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/mdns-service-policy", state.getPath()))
	}
	return deletedItems
}

func (data *ServiceTemplate) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	if !data.InactivityTimerProbe.IsNull() && !data.InactivityTimerProbe.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/inactivity-timer/probe", data.getPath()))
	}
	if !data.VoiceVlan.IsNull() && !data.VoiceVlan.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/voice/vlan", data.getPath()))
	}

	return emptyLeafsDelete
}

func (data *ServiceTemplate) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.AccessGroup {
		keyValues := [...]string{data.AccessGroup[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/access-group-config=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.InactivityTimerValue.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/inactivity-timer/value", data.getPath()))
	}
	if !data.InactivityTimerProbe.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/inactivity-timer/probe", data.getPath()))
	}
	if !data.Vlan.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan", data.getPath()))
	}
	if !data.VoiceVlan.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/voice/vlan", data.getPath()))
	}
	if !data.LinksecPolicy.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/linksec/policy", data.getPath()))
	}
	if !data.Sgt.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/sgt", data.getPath()))
	}
	if !data.AbsoluteTimer.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/absolute-timer", data.getPath()))
	}
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	for i := range data.InterfaceTemplate {
		keyValues := [...]string{data.InterfaceTemplate[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/interface-template=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.TunnelCapwapName.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tunnel/type/capwap/name", data.getPath()))
	}
	if !data.Vnid.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vnid", data.getPath()))
	}
	if !data.RedirectAppendClientMac.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/redirect/append/client-mac", data.getPath()))
	}
	if !data.RedirectAppendSwitchMac.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/redirect/append/switch-mac", data.getPath()))
	}
	if !data.RedirectUrlUrlName.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/redirect/url/url_name", data.getPath()))
	}
	if !data.RedirectUrlMatchAclName.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/redirect/url/match/acl_name", data.getPath()))
	}
	if !data.RedirectUrlMatchAction.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/redirect/url/match/action", data.getPath()))
	}
	if !data.DnsAclPreauth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dns-acl/preauth", data.getPath()))
	}
	if !data.ServicePolicyQosInput.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/service-policy/qos/input", data.getPath()))
	}
	if !data.ServicePolicyQosOutput.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/service-policy/qos/output", data.getPath()))
	}
	for i := range data.TagConfig {
		keyValues := [...]string{data.TagConfig[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/tag-config=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.MdnsServicePolicy.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/mdns-service-policy", data.getPath()))
	}
	return deletePaths
}
