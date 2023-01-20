package traefik_plugin_disable_graphql_introspection

import (
	"testing"
)

func TestCheckIfRequestIsIntrospection(t *testing.T) {
	tests := []struct {
		desc   string
		result bool
		body   string
	}{
		{
			desc:   "user query without typename",
			result: false,
			body: `
				query NotIntrospectionQuery{
					users {
						id
						plans {
							id
							name
						}
					}
				}
      `,
		},
		{
			desc:   "user query with typename",
			result: false,
			body: `
				query NotIntrospectionQuery{
					users {
						id
						plans {
							id
							name
							__typename
						}
					}
				}
      `,
		},
		{
			desc:   "user query with type and schema fields",
			result: false,
			body: `
				query NotIntrospectionQuery{
					users {
						id
						plans {
							id
							name
							type
							schema
						}
					}
				}
      `,
		},
		{
			desc:   "schema introspection request",
			result: true,
			body: `
				__schema {
					types {
						name
					}
				}
      `,
		},
		{
			desc:   "type introspection request",
			result: true,
			body: `
				{
					__type(name: "CustomUserType") {
						name
						fields {
							name
						}
					}
				}
      `,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			result := checkIfRequestIsIntrospection(test.body)

			if result != test.result {
				t.Fatal("Check is failed")
			}
		})
	}
}
