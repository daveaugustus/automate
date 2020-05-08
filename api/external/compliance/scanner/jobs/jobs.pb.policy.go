// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/compliance/scanner/jobs/jobs.proto

package jobs

import policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Create", "compliance:scanner:jobs", "compliance:scannerJobs:create", "POST", "/api/v0/compliance/scanner/jobs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Job); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "type":
					return m.Type
				case "status":
					return m.Status
				case "recurrence":
					return m.Recurrence
				case "parent_id":
					return m.ParentId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Read", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:get", "GET", "/api/v0/compliance/scanner/jobs/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Update", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:update", "PUT", "/api/v0/compliance/scanner/jobs/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Job); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "type":
					return m.Type
				case "status":
					return m.Status
				case "recurrence":
					return m.Recurrence
				case "parent_id":
					return m.ParentId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Delete", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:delete", "DELETE", "/api/v0/compliance/scanner/jobs/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/List", "compliance:scanner:jobs", "compliance:scannerJobs:list", "POST", "/api/v0/compliance/scanner/jobs/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.scanner.jobs.v1.JobsService/Rerun", "compliance:scanner:jobs:{id}", "compliance:scannerJobs:rerun", "GET", "/api/v0/compliance/scanner/jobs/rerun/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
}
