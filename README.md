# k8s-api go types

## Why?

Its nice to have it all in a single package.

```
.
|-- pkg
|   `-- v1_20
```

## Want to contritube a k8s version?

1. Get a schema from [here](https://github.com/cdk8s-team/cdk8s/tree/master/kubernetes-schemas).
2. Use  go-swagger to generate the models.
```
swagger generate model -f definition.json
```
3. Update pkg name to version. Example: `package v1_20`.
4. Create a PR : )
