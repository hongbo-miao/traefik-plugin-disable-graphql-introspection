package traefik_plugin_disable_graphql_introspection

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Config struct {
	GraphQLPath string
}

func CreateConfig() *Config {
	return &Config{
		GraphQLPath: "/graphql",
	}
}

type DisableGraphQLIntrospection struct {
	next        http.Handler
	name        string
	graphQLPath string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &DisableGraphQLIntrospection{
		next:        next,
		name:        name,
		graphQLPath: config.GraphQLPath,
	}, nil
}

func checkIfRequestIsIntrospection(body string) bool {
	if strings.Contains(body, "__schema") {
		return true
	}

	match, _ := regexp.MatchString("__type[^name]", body)

	if match {
		return true
	}

	return false
}

func (d *DisableGraphQLIntrospection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(`{
			"error": {
				"code": 400,
				"message": "Failed to read request body."
			}
		}`))
		return
	}
	if r.Method == "POST" && r.URL.Path == d.graphQLPath {
		if checkIfRequestIsIntrospection(string(body)) {
			rw.Header().Set("Content-Type", "application/json")
			rw.Write([]byte(`{
				"errors": [
					{
						"message": "GraphQL introspection is not allowed."
					}
				]
			}`))
			return
		}
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	d.next.ServeHTTP(rw, r)
}
