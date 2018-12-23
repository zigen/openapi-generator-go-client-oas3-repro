# openapi-generator go client codegen repro

at [9e1cce2e](https://github.com/OpenAPITools/openapi-generator/tree/9e1cce2e7b2fefa742c4992215c59998fb9fb311)
## directories
`go-sample-api`: generated code by [composed-oneof.yaml](https://github.com/OpenAPITools/openapi-generator/blob/master/modules/openapi-generator/src/test/resources/3_0/composed-oneof.yaml)
`go-oneof.sh`: shell script to generate `go-sample-api`. put this file into `openapi-generator/bin` and then `./bin/go-oneof.sh`
`go-client-sample`: client code which uses `go-sample-api`

## output

```
$ go run main.go

# github.com/zigen/go-sample-api
../go-sample-api/api_default.go:32:78: undefined: UNKNOWN_BASE_TYPE
../go-sample-api/api_default.go:99:60: undefined: OneOfObjA
../go-sample-api/api_default.go:106:24: undefined: OneOfObjA
../go-sample-api/api_default.go:155:10: undefined: OneOfObjA

```




