# Disable GraphQL Introspection

A plugin for [Traefik](https://github.com/traefik/traefik) to disable GraphQL Introspection.

## Configuration


### Static

```yaml
pilot:
  token: xxx
experimental:
  plugins:
    traefik-plugin-disable-graphql-introspection:
      modulename: github.com/Hongbo-Miao/traefik-plugin-disable-graphql-introspection
      version: v0.1.0
```

### Dynamic

```yaml
http:
  routers:
    graphql-server-entrypoint:
      service: graphql-server-service
      entrypoints:
        - graphql-server-entrypoint
      rule: Host(`localhost`)
      middlewares:
        - my-traefik-plugin-disable-graphql-introspection
  services:
    graphql-server-service:
      loadBalancer:
        servers:
          - url: http://localhost:5000/
  middlewares:
    my-traefik-plugin-disable-graphql-introspection:
      plugin:
        traefik-plugin-disable-graphql-introspection:
          GraphQLPath: /graphql
```
