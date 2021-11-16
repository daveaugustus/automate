// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/infra_proxy/infra_proxy.proto

package infra_proxy

import (
	policy "github.com/chef/automate/api/external/iam/v2/policy"
	request "github.com/chef/automate/api/external/infra_proxy/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServers", "infra:infraServers", "infra:infraServers:list", "GET", "/api/v0/infra/servers", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServerStatus", "infra:infraServers", "infra:infraServers:get", "GET", "/api/v0/infra/servers/server_status", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServerStatus); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "fqdn":
					return m.Fqdn
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServer", "infra:infraServers:{id}", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateServer", "infra:infraServers", "infra:infraServers:create", "POST", "/api/v0/infra/servers", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateServer", "infra:infraServers:{id}", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteServer", "infra:infraServers:{id}", "infra:infraServers:delete", "DELETE", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgs", "infra:infraServers:{server_id}:orgs", "infra:infraServersOrgs:list", "GET", "/api/v0/infra/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgs); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServersOrgs:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateOrg", "infra:infraServers:{server_id}:orgs", "infra:infraServersOrgs:create", "POST", "/api/v0/infra/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "admin_key":
					return m.AdminKey
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServersOrgs:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServersOrgs:delete", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/ResetOrgAdminKey", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{id}/reset-key", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ResetOrgAdminKey); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				case "admin_key":
					return m.AdminKey
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbooks", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServersOrgsCookbooks:list", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Cookbooks); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbookVersions", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServersOrgsCookbooks:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CookbookVersions); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbook", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServersOrgsCookbooks:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Cookbook); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbookFileContent", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServersOrgsCookbooks:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}/file-content", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CookbookFileContent); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "version":
					return m.Version
				case "url":
					return m.Url
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRoles", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServersOrgsRoles:list", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Roles); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServersOrgsRoles:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Role); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRoleExpandedRunList", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServersOrgsRoles:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}/runlist/{environment}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ExpandedRunList); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "environment":
					return m.Environment
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRoleEnvironments", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServersOrgsRoles:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}/environments", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Role); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServersOrgsRoles:create", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateRole); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServersOrgsRoles:delete", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Role); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServersOrgsRoles:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateRole); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetClients", "infra:infraServers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Clients); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetClient", "infra:servers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Client); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateClient", "infra:servers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:update", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateClient); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteClient", "infra:servers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Client); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/ResetClientKey", "infra:servers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients/{name}/reset", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ClientKey); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "key":
					return m.Key
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetDataBags", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBags:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBags); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetDataBagItems", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBagsItem:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBagItems); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBagsItem:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBagItem); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "item":
					return m.Item
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateDataBag", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBags:create", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateDataBag); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBagsItem:create", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateDataBagItem); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteDataBag", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBags:delete", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBag); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBagsItem:delete", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBagItem); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "item":
					return m.Item
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServersOrgsDataBagsItem:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateDataBagItem); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "item_id":
					return m.ItemId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetEnvironments", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServersOrgsEnvironments:list", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environments); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServersOrgsEnvironments:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServersOrgsEnvironments:create", "POST", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateEnvironment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "json_class":
					return m.JsonClass
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServersOrgsEnvironments:delete", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServersOrgsEnvironments:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateEnvironment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "json_class":
					return m.JsonClass
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetEnvironmentRecipes", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServersOrgsEnvironments:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}/recipes", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetNodes", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Nodes); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteNode", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Node); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateNode", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeDetails); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "environment":
					return m.Environment
				case "policy_name":
					return m.PolicyName
				case "policy_group":
					return m.PolicyGroup
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateNodeAttributes", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}/attributes", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateNodeAttributes); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetNode", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Node); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetNodeExpandedRunList", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/node/{name}/runlist/{environment}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeExpandedRunList); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "environment":
					return m.Environment
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateNodeTags", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}/tags", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateNodeTags); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "action":
					return m.Action
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateNodeEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}/environment", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateNodeEnvironment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "environment":
					return m.Environment
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicyfiles", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServersOrgsPolicyFiles:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Policyfiles); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicyfile", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServersOrgsPolicyFiles:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Policyfile); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "revision_id":
					return m.RevisionId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeletePolicyfile", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServersOrgsPolicyFiles:delete", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeletePolicyfile); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicyfileRevisions", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServersOrgsPolicyFiles:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles/{name}/revisions", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.PolicyfileRevisions); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicygroup", "infra:infraServers:{server_id}:orgs:{org_id}:policygroups", "infra:infraServersOrgsPolicyGroups:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policygroups/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Policygroup); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgUsersList", "infra:infraServers:{server_id}:orgs:{org_id}:users", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.OrgUsers); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServerUsersList", "infra:infraServers:{server_id}:users", "infra:infraServers:get", "POST", "/api/v0/infra/servers/{server_id}/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ServerUsers); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "server_id":
					return m.ServerId
				case "admin_name":
					return m.AdminName
				case "admin_key":
					return m.AdminKey
				default:
					return ""
				}
			})
		}
		return ""
	})
}
