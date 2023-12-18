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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Line struct {
	Device     types.String  `tfsdk:"device"`
	Id         types.String  `tfsdk:"id"`
	DeleteMode types.String  `tfsdk:"delete_mode"`
	Console    []LineConsole `tfsdk:"console"`
	Vty        []LineVty     `tfsdk:"vty"`
}

type LineData struct {
	Device  types.String  `tfsdk:"device"`
	Id      types.String  `tfsdk:"id"`
	Console []LineConsole `tfsdk:"console"`
	Vty     []LineVty     `tfsdk:"vty"`
}
type LineConsole struct {
	First               types.String `tfsdk:"first"`
	ExecTimeoutMinutes  types.Int64  `tfsdk:"exec_timeout_minutes"`
	ExecTimeoutSeconds  types.Int64  `tfsdk:"exec_timeout_seconds"`
	LoginLocal          types.Bool   `tfsdk:"login_local"`
	LoginAuthentication types.String `tfsdk:"login_authentication"`
	PrivilegeLevel      types.Int64  `tfsdk:"privilege_level"`
	Stopbits            types.String `tfsdk:"stopbits"`
	PasswordLevel       types.Int64  `tfsdk:"password_level"`
	PasswordType        types.String `tfsdk:"password_type"`
	Password            types.String `tfsdk:"password"`
}
type LineVty struct {
	First                      types.Int64            `tfsdk:"first"`
	Last                       types.Int64            `tfsdk:"last"`
	AccessClasses              []LineVtyAccessClasses `tfsdk:"access_classes"`
	ExecTimeoutMinutes         types.Int64            `tfsdk:"exec_timeout_minutes"`
	ExecTimeoutSeconds         types.Int64            `tfsdk:"exec_timeout_seconds"`
	PasswordLevel              types.Int64            `tfsdk:"password_level"`
	PasswordType               types.String           `tfsdk:"password_type"`
	Password                   types.String           `tfsdk:"password"`
	LoginAuthentication        types.String           `tfsdk:"login_authentication"`
	TransportPreferredProtocol types.String           `tfsdk:"transport_preferred_protocol"`
	EscapeCharacter            types.String           `tfsdk:"escape_character"`
	AuthorizationExec          types.String           `tfsdk:"authorization_exec"`
	AuthorizationExecDefault   types.Bool             `tfsdk:"authorization_exec_default"`
	TransportInputAll          types.Bool             `tfsdk:"transport_input_all"`
	TransportInputNone         types.Bool             `tfsdk:"transport_input_none"`
	TransportInput             types.String           `tfsdk:"transport_input"`
}
type LineVtyAccessClasses struct {
	Direction  types.String `tfsdk:"direction"`
	AccessList types.String `tfsdk:"access_list"`
	VrfAlso    types.Bool   `tfsdk:"vrf_also"`
}

func (data Line) getPath() string {
	return "Cisco-IOS-XE-native:native/line"
}

func (data LineData) getPath() string {
	return "Cisco-IOS-XE-native:native/line"
}

// if last path element has a key -> remove it
func (data Line) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data Line) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if len(data.Console) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console", []interface{}{})
		for index, item := range data.Console {
			if !item.First.IsNull() && !item.First.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"first", item.First.ValueString())
			}
			if !item.ExecTimeoutMinutes.IsNull() && !item.ExecTimeoutMinutes.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"exec-timeout.minutes", strconv.FormatInt(item.ExecTimeoutMinutes.ValueInt64(), 10))
			}
			if !item.ExecTimeoutSeconds.IsNull() && !item.ExecTimeoutSeconds.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"exec-timeout.seconds", strconv.FormatInt(item.ExecTimeoutSeconds.ValueInt64(), 10))
			}
			if !item.LoginLocal.IsNull() && !item.LoginLocal.IsUnknown() {
				if item.LoginLocal.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"login.local", map[string]string{})
				}
			}
			if !item.LoginAuthentication.IsNull() && !item.LoginAuthentication.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"login.authentication", item.LoginAuthentication.ValueString())
			}
			if !item.PrivilegeLevel.IsNull() && !item.PrivilegeLevel.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"privilege.level.number", strconv.FormatInt(item.PrivilegeLevel.ValueInt64(), 10))
			}
			if !item.Stopbits.IsNull() && !item.Stopbits.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"stopbits", item.Stopbits.ValueString())
			}
			if !item.PasswordLevel.IsNull() && !item.PasswordLevel.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"password.level", strconv.FormatInt(item.PasswordLevel.ValueInt64(), 10))
			}
			if !item.PasswordType.IsNull() && !item.PasswordType.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"password.type", item.PasswordType.ValueString())
			}
			if !item.Password.IsNull() && !item.Password.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"console"+"."+strconv.Itoa(index)+"."+"password.secret", item.Password.ValueString())
			}
		}
	}
	if len(data.Vty) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty", []interface{}{})
		for index, item := range data.Vty {
			if !item.First.IsNull() && !item.First.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"first", strconv.FormatInt(item.First.ValueInt64(), 10))
			}
			if !item.Last.IsNull() && !item.Last.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"last", strconv.FormatInt(item.Last.ValueInt64(), 10))
			}
			if !item.ExecTimeoutMinutes.IsNull() && !item.ExecTimeoutMinutes.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"exec-timeout.minutes", strconv.FormatInt(item.ExecTimeoutMinutes.ValueInt64(), 10))
			}
			if !item.ExecTimeoutSeconds.IsNull() && !item.ExecTimeoutSeconds.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"exec-timeout.seconds", strconv.FormatInt(item.ExecTimeoutSeconds.ValueInt64(), 10))
			}
			if !item.PasswordLevel.IsNull() && !item.PasswordLevel.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"password.level", strconv.FormatInt(item.PasswordLevel.ValueInt64(), 10))
			}
			if !item.PasswordType.IsNull() && !item.PasswordType.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"password.type", item.PasswordType.ValueString())
			}
			if !item.Password.IsNull() && !item.Password.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"password.secret", item.Password.ValueString())
			}
			if !item.LoginAuthentication.IsNull() && !item.LoginAuthentication.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"login.authentication", item.LoginAuthentication.ValueString())
			}
			if !item.TransportPreferredProtocol.IsNull() && !item.TransportPreferredProtocol.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"transport.preferred.protocol", item.TransportPreferredProtocol.ValueString())
			}
			if !item.EscapeCharacter.IsNull() && !item.EscapeCharacter.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"escape-character.char", item.EscapeCharacter.ValueString())
			}
			if !item.AuthorizationExec.IsNull() && !item.AuthorizationExec.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"authorization.exec.authorization-name", item.AuthorizationExec.ValueString())
			}
			if !item.AuthorizationExecDefault.IsNull() && !item.AuthorizationExecDefault.IsUnknown() {
				if item.AuthorizationExecDefault.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"authorization.exec.default", map[string]string{})
				}
			}
			if !item.TransportInputAll.IsNull() && !item.TransportInputAll.IsUnknown() {
				if item.TransportInputAll.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"transport.input.all", map[string]string{})
				}
			}
			if !item.TransportInputNone.IsNull() && !item.TransportInputNone.IsUnknown() {
				if item.TransportInputNone.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"transport.input.none", map[string]string{})
				}
			}
			if !item.TransportInput.IsNull() && !item.TransportInput.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"transport.input.input", item.TransportInput.ValueString())
			}
			if len(item.AccessClasses) > 0 {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"access-class.acccess-list", []interface{}{})
				for cindex, citem := range item.AccessClasses {
					if !citem.Direction.IsNull() && !citem.Direction.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"access-class.acccess-list"+"."+strconv.Itoa(cindex)+"."+"direction", citem.Direction.ValueString())
					}
					if !citem.AccessList.IsNull() && !citem.AccessList.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"access-class.acccess-list"+"."+strconv.Itoa(cindex)+"."+"access-list", citem.AccessList.ValueString())
					}
					if !citem.VrfAlso.IsNull() && !citem.VrfAlso.IsUnknown() {
						if citem.VrfAlso.ValueBool() {
							body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vty"+"."+strconv.Itoa(index)+"."+"access-class.acccess-list"+"."+strconv.Itoa(cindex)+"."+"vrf-also", map[string]string{})
						}
					}
				}
			}
		}
	}
	return body
}

func (data *Line) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	for i := range data.Console {
		keys := [...]string{"first"}
		keyValues := [...]string{data.Console[i].First.ValueString()}

		var r gjson.Result
		res.Get(prefix + "console").ForEach(
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
		if value := r.Get("first"); value.Exists() && !data.Console[i].First.IsNull() {
			data.Console[i].First = types.StringValue(value.String())
		} else {
			data.Console[i].First = types.StringNull()
		}
		if value := r.Get("exec-timeout.minutes"); value.Exists() && !data.Console[i].ExecTimeoutMinutes.IsNull() {
			data.Console[i].ExecTimeoutMinutes = types.Int64Value(value.Int())
		} else {
			data.Console[i].ExecTimeoutMinutes = types.Int64Null()
		}
		if value := r.Get("exec-timeout.seconds"); value.Exists() && !data.Console[i].ExecTimeoutSeconds.IsNull() {
			data.Console[i].ExecTimeoutSeconds = types.Int64Value(value.Int())
		} else {
			data.Console[i].ExecTimeoutSeconds = types.Int64Null()
		}
		if value := r.Get("login.local"); !data.Console[i].LoginLocal.IsNull() {
			if value.Exists() {
				data.Console[i].LoginLocal = types.BoolValue(true)
			} else {
				data.Console[i].LoginLocal = types.BoolValue(false)
			}
		} else {
			data.Console[i].LoginLocal = types.BoolNull()
		}
		if value := r.Get("login.authentication"); value.Exists() && !data.Console[i].LoginAuthentication.IsNull() {
			data.Console[i].LoginAuthentication = types.StringValue(value.String())
		} else {
			data.Console[i].LoginAuthentication = types.StringNull()
		}
		if value := r.Get("privilege.level.number"); value.Exists() && !data.Console[i].PrivilegeLevel.IsNull() {
			data.Console[i].PrivilegeLevel = types.Int64Value(value.Int())
		} else {
			data.Console[i].PrivilegeLevel = types.Int64Null()
		}
		if value := r.Get("stopbits"); value.Exists() && !data.Console[i].Stopbits.IsNull() {
			data.Console[i].Stopbits = types.StringValue(value.String())
		} else {
			data.Console[i].Stopbits = types.StringNull()
		}
		if value := r.Get("password.level"); value.Exists() && !data.Console[i].PasswordLevel.IsNull() {
			data.Console[i].PasswordLevel = types.Int64Value(value.Int())
		} else {
			data.Console[i].PasswordLevel = types.Int64Null()
		}
		if value := r.Get("password.type"); value.Exists() && !data.Console[i].PasswordType.IsNull() {
			data.Console[i].PasswordType = types.StringValue(value.String())
		} else {
			data.Console[i].PasswordType = types.StringNull()
		}
		if value := r.Get("password.secret"); value.Exists() && !data.Console[i].Password.IsNull() {
			data.Console[i].Password = types.StringValue(value.String())
		} else {
			data.Console[i].Password = types.StringNull()
		}
	}
	for i := range data.Vty {
		keys := [...]string{"first"}
		keyValues := [...]string{strconv.FormatInt(data.Vty[i].First.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "vty").ForEach(
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
		if value := r.Get("first"); value.Exists() && !data.Vty[i].First.IsNull() {
			data.Vty[i].First = types.Int64Value(value.Int())
		} else {
			data.Vty[i].First = types.Int64Null()
		}
		if value := r.Get("last"); value.Exists() && !data.Vty[i].Last.IsNull() {
			data.Vty[i].Last = types.Int64Value(value.Int())
		} else {
			data.Vty[i].Last = types.Int64Null()
		}
		for ci := range data.Vty[i].AccessClasses {
			keys := [...]string{"direction"}
			keyValues := [...]string{data.Vty[i].AccessClasses[ci].Direction.ValueString()}

			var cr gjson.Result
			r.Get("access-class.acccess-list").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("direction"); value.Exists() && !data.Vty[i].AccessClasses[ci].Direction.IsNull() {
				data.Vty[i].AccessClasses[ci].Direction = types.StringValue(value.String())
			} else {
				data.Vty[i].AccessClasses[ci].Direction = types.StringNull()
			}
			if value := cr.Get("access-list"); value.Exists() && !data.Vty[i].AccessClasses[ci].AccessList.IsNull() {
				data.Vty[i].AccessClasses[ci].AccessList = types.StringValue(value.String())
			} else {
				data.Vty[i].AccessClasses[ci].AccessList = types.StringNull()
			}
			if value := cr.Get("vrf-also"); !data.Vty[i].AccessClasses[ci].VrfAlso.IsNull() {
				if value.Exists() {
					data.Vty[i].AccessClasses[ci].VrfAlso = types.BoolValue(true)
				} else {
					data.Vty[i].AccessClasses[ci].VrfAlso = types.BoolValue(false)
				}
			} else {
				data.Vty[i].AccessClasses[ci].VrfAlso = types.BoolNull()
			}
		}
		if value := r.Get("exec-timeout.minutes"); value.Exists() && !data.Vty[i].ExecTimeoutMinutes.IsNull() {
			data.Vty[i].ExecTimeoutMinutes = types.Int64Value(value.Int())
		} else {
			data.Vty[i].ExecTimeoutMinutes = types.Int64Null()
		}
		if value := r.Get("exec-timeout.seconds"); value.Exists() && !data.Vty[i].ExecTimeoutSeconds.IsNull() {
			data.Vty[i].ExecTimeoutSeconds = types.Int64Value(value.Int())
		} else {
			data.Vty[i].ExecTimeoutSeconds = types.Int64Null()
		}
		if value := r.Get("password.level"); value.Exists() && !data.Vty[i].PasswordLevel.IsNull() {
			data.Vty[i].PasswordLevel = types.Int64Value(value.Int())
		} else {
			data.Vty[i].PasswordLevel = types.Int64Null()
		}
		if value := r.Get("password.type"); value.Exists() && !data.Vty[i].PasswordType.IsNull() {
			data.Vty[i].PasswordType = types.StringValue(value.String())
		} else {
			data.Vty[i].PasswordType = types.StringNull()
		}
		if value := r.Get("password.secret"); value.Exists() && !data.Vty[i].Password.IsNull() {
			data.Vty[i].Password = types.StringValue(value.String())
		} else {
			data.Vty[i].Password = types.StringNull()
		}
		if value := r.Get("login.authentication"); value.Exists() && !data.Vty[i].LoginAuthentication.IsNull() {
			data.Vty[i].LoginAuthentication = types.StringValue(value.String())
		} else {
			data.Vty[i].LoginAuthentication = types.StringNull()
		}
		if value := r.Get("transport.preferred.protocol"); value.Exists() && !data.Vty[i].TransportPreferredProtocol.IsNull() {
			data.Vty[i].TransportPreferredProtocol = types.StringValue(value.String())
		} else {
			data.Vty[i].TransportPreferredProtocol = types.StringNull()
		}
		if value := r.Get("escape-character.char"); value.Exists() && !data.Vty[i].EscapeCharacter.IsNull() {
			data.Vty[i].EscapeCharacter = types.StringValue(value.String())
		} else {
			data.Vty[i].EscapeCharacter = types.StringNull()
		}
		if value := r.Get("authorization.exec.authorization-name"); value.Exists() && !data.Vty[i].AuthorizationExec.IsNull() {
			data.Vty[i].AuthorizationExec = types.StringValue(value.String())
		} else {
			data.Vty[i].AuthorizationExec = types.StringNull()
		}
		if value := r.Get("authorization.exec.default"); !data.Vty[i].AuthorizationExecDefault.IsNull() {
			if value.Exists() {
				data.Vty[i].AuthorizationExecDefault = types.BoolValue(true)
			} else {
				data.Vty[i].AuthorizationExecDefault = types.BoolValue(false)
			}
		} else {
			data.Vty[i].AuthorizationExecDefault = types.BoolNull()
		}
		if value := r.Get("transport.input.all"); !data.Vty[i].TransportInputAll.IsNull() {
			if value.Exists() {
				data.Vty[i].TransportInputAll = types.BoolValue(true)
			} else {
				data.Vty[i].TransportInputAll = types.BoolValue(false)
			}
		} else {
			data.Vty[i].TransportInputAll = types.BoolNull()
		}
		if value := r.Get("transport.input.none"); !data.Vty[i].TransportInputNone.IsNull() {
			if value.Exists() {
				data.Vty[i].TransportInputNone = types.BoolValue(true)
			} else {
				data.Vty[i].TransportInputNone = types.BoolValue(false)
			}
		} else {
			data.Vty[i].TransportInputNone = types.BoolNull()
		}
		if value := r.Get("transport.input.input"); value.Exists() && !data.Vty[i].TransportInput.IsNull() {
			data.Vty[i].TransportInput = types.StringValue(value.String())
		} else {
			data.Vty[i].TransportInput = types.StringNull()
		}
	}
}

func (data *LineData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "console"); value.Exists() {
		data.Console = make([]LineConsole, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LineConsole{}
			if cValue := v.Get("first"); cValue.Exists() {
				item.First = types.StringValue(cValue.String())
			}
			if cValue := v.Get("exec-timeout.minutes"); cValue.Exists() {
				item.ExecTimeoutMinutes = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("exec-timeout.seconds"); cValue.Exists() {
				item.ExecTimeoutSeconds = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("login.local"); cValue.Exists() {
				item.LoginLocal = types.BoolValue(true)
			} else {
				item.LoginLocal = types.BoolValue(false)
			}
			if cValue := v.Get("login.authentication"); cValue.Exists() {
				item.LoginAuthentication = types.StringValue(cValue.String())
			}
			if cValue := v.Get("privilege.level.number"); cValue.Exists() {
				item.PrivilegeLevel = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("stopbits"); cValue.Exists() {
				item.Stopbits = types.StringValue(cValue.String())
			}
			if cValue := v.Get("password.level"); cValue.Exists() {
				item.PasswordLevel = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("password.type"); cValue.Exists() {
				item.PasswordType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("password.secret"); cValue.Exists() {
				item.Password = types.StringValue(cValue.String())
			}
			data.Console = append(data.Console, item)
			return true
		})
	}
	if value := res.Get(prefix + "vty"); value.Exists() {
		data.Vty = make([]LineVty, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LineVty{}
			if cValue := v.Get("first"); cValue.Exists() {
				item.First = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("last"); cValue.Exists() {
				item.Last = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("access-class.acccess-list"); cValue.Exists() {
				item.AccessClasses = make([]LineVtyAccessClasses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := LineVtyAccessClasses{}
					if ccValue := cv.Get("direction"); ccValue.Exists() {
						cItem.Direction = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("access-list"); ccValue.Exists() {
						cItem.AccessList = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("vrf-also"); ccValue.Exists() {
						cItem.VrfAlso = types.BoolValue(true)
					} else {
						cItem.VrfAlso = types.BoolValue(false)
					}
					item.AccessClasses = append(item.AccessClasses, cItem)
					return true
				})
			}
			if cValue := v.Get("exec-timeout.minutes"); cValue.Exists() {
				item.ExecTimeoutMinutes = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("exec-timeout.seconds"); cValue.Exists() {
				item.ExecTimeoutSeconds = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("password.level"); cValue.Exists() {
				item.PasswordLevel = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("password.type"); cValue.Exists() {
				item.PasswordType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("password.secret"); cValue.Exists() {
				item.Password = types.StringValue(cValue.String())
			}
			if cValue := v.Get("login.authentication"); cValue.Exists() {
				item.LoginAuthentication = types.StringValue(cValue.String())
			}
			if cValue := v.Get("transport.preferred.protocol"); cValue.Exists() {
				item.TransportPreferredProtocol = types.StringValue(cValue.String())
			}
			if cValue := v.Get("escape-character.char"); cValue.Exists() {
				item.EscapeCharacter = types.StringValue(cValue.String())
			}
			if cValue := v.Get("authorization.exec.authorization-name"); cValue.Exists() {
				item.AuthorizationExec = types.StringValue(cValue.String())
			}
			if cValue := v.Get("authorization.exec.default"); cValue.Exists() {
				item.AuthorizationExecDefault = types.BoolValue(true)
			} else {
				item.AuthorizationExecDefault = types.BoolValue(false)
			}
			if cValue := v.Get("transport.input.all"); cValue.Exists() {
				item.TransportInputAll = types.BoolValue(true)
			} else {
				item.TransportInputAll = types.BoolValue(false)
			}
			if cValue := v.Get("transport.input.none"); cValue.Exists() {
				item.TransportInputNone = types.BoolValue(true)
			} else {
				item.TransportInputNone = types.BoolValue(false)
			}
			if cValue := v.Get("transport.input.input"); cValue.Exists() {
				item.TransportInput = types.StringValue(cValue.String())
			}
			data.Vty = append(data.Vty, item)
			return true
		})
	}
}

func (data *Line) getDeletedItems(ctx context.Context, state Line) []string {
	deletedItems := make([]string, 0)
	for i := range state.Console {
		stateKeyValues := [...]string{state.Console[i].First.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Console[i].First.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Console {
			found = true
			if state.Console[i].First.ValueString() != data.Console[j].First.ValueString() {
				found = false
			}
			if found {
				if !state.Console[i].ExecTimeoutMinutes.IsNull() && data.Console[j].ExecTimeoutMinutes.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/exec-timeout/minutes", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].ExecTimeoutSeconds.IsNull() && data.Console[j].ExecTimeoutSeconds.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/exec-timeout/seconds", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].LoginLocal.IsNull() && data.Console[j].LoginLocal.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/login/local", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].LoginAuthentication.IsNull() && data.Console[j].LoginAuthentication.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/login/authentication", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].PrivilegeLevel.IsNull() && data.Console[j].PrivilegeLevel.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/privilege/level/number", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].Stopbits.IsNull() && data.Console[j].Stopbits.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/stopbits", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].PasswordLevel.IsNull() && data.Console[j].PasswordLevel.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/password/level", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].PasswordType.IsNull() && data.Console[j].PasswordType.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/password/type", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Console[i].Password.IsNull() && data.Console[j].Password.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v/password/secret", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/console=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Vty {
		stateKeyValues := [...]string{strconv.FormatInt(state.Vty[i].First.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Vty[i].First.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Vty {
			found = true
			if state.Vty[i].First.ValueInt64() != data.Vty[j].First.ValueInt64() {
				found = false
			}
			if found {
				if !state.Vty[i].Last.IsNull() && data.Vty[j].Last.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/last", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				for ci := range state.Vty[i].AccessClasses {
					cstateKeyValues := [...]string{state.Vty[i].AccessClasses[ci].Direction.ValueString()}

					cemptyKeys := true
					if !reflect.ValueOf(state.Vty[i].AccessClasses[ci].Direction.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.Vty[j].AccessClasses {
						found = true
						if state.Vty[i].AccessClasses[ci].Direction.ValueString() != data.Vty[j].AccessClasses[cj].Direction.ValueString() {
							found = false
						}
						if found {
							if !state.Vty[i].AccessClasses[ci].AccessList.IsNull() && data.Vty[j].AccessClasses[cj].AccessList.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/access-class/acccess-list=%v/access-list", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Vty[i].AccessClasses[ci].VrfAlso.IsNull() && data.Vty[j].AccessClasses[cj].VrfAlso.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/access-class/acccess-list=%v/vrf-also", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							break
						}
					}
					if !found {
						deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/access-class/acccess-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
					}
				}
				if !state.Vty[i].ExecTimeoutMinutes.IsNull() && data.Vty[j].ExecTimeoutMinutes.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/exec-timeout/minutes", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].ExecTimeoutSeconds.IsNull() && data.Vty[j].ExecTimeoutSeconds.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/exec-timeout/seconds", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].PasswordLevel.IsNull() && data.Vty[j].PasswordLevel.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/password/level", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].PasswordType.IsNull() && data.Vty[j].PasswordType.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/password/type", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].Password.IsNull() && data.Vty[j].Password.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/password/secret", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].LoginAuthentication.IsNull() && data.Vty[j].LoginAuthentication.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/login/authentication", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].TransportPreferredProtocol.IsNull() && data.Vty[j].TransportPreferredProtocol.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/transport/preferred/protocol", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].EscapeCharacter.IsNull() && data.Vty[j].EscapeCharacter.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/escape-character/char", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].AuthorizationExec.IsNull() && data.Vty[j].AuthorizationExec.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/authorization/exec/authorization-name", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].AuthorizationExecDefault.IsNull() && data.Vty[j].AuthorizationExecDefault.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/authorization/exec/default", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].TransportInputAll.IsNull() && data.Vty[j].TransportInputAll.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/transport/input/all", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].TransportInputNone.IsNull() && data.Vty[j].TransportInputNone.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/transport/input/none", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vty[i].TransportInput.IsNull() && data.Vty[j].TransportInput.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v/transport/input/input", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/vty=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *Line) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.Console {
		keyValues := [...]string{data.Console[i].First.ValueString()}
		if !data.Console[i].LoginLocal.IsNull() && !data.Console[i].LoginLocal.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/console=%v/login/local", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}

	for i := range data.Vty {
		keyValues := [...]string{strconv.FormatInt(data.Vty[i].First.ValueInt64(), 10)}

		for ci := range data.Vty[i].AccessClasses {
			ckeyValues := [...]string{data.Vty[i].AccessClasses[ci].Direction.ValueString()}
			if !data.Vty[i].AccessClasses[ci].VrfAlso.IsNull() && !data.Vty[i].AccessClasses[ci].VrfAlso.ValueBool() {
				emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vty=%v/access-class/acccess-list=%v/vrf-also", data.getPath(), strings.Join(keyValues[:], ","), strings.Join(ckeyValues[:], ",")))
			}
		}
		if !data.Vty[i].AuthorizationExecDefault.IsNull() && !data.Vty[i].AuthorizationExecDefault.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vty=%v/authorization/exec/default", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Vty[i].TransportInputAll.IsNull() && !data.Vty[i].TransportInputAll.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vty=%v/transport/input/all", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Vty[i].TransportInputNone.IsNull() && !data.Vty[i].TransportInputNone.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vty=%v/transport/input/none", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *Line) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	return deletePaths
}
