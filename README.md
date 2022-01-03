# k8s-api go types

## Why?

Its nice to have it all in a single package.

```
.
|-- pkg
|   `-- v1_20
```

## Want to contritube a k8s version?

1. Get a schema from a specific k8s release branch 
```
version='1.20'
curl https://raw.githubusercontent.com/kubernetes/kubernetes/release-$version/api/openapi-spec/swagger.json > definitions.json
```
2. Use  go-swagger to generate the models.
```
swagger generate model --struct-tags='json,yaml' -f ./definitions.json
```
3. Update pkg name to k8s release version. Example: `package v1_20`.
4. Create a PR : )
