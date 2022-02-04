
# http-swagger (1.2.0)
with swag 1.7.8
----------
```shell
go install github.com/swaggo/swag/cmd/swag@1.7.8
swag init
2022/02/04 XX:XX:XX Generate swagger docs....
2022/02/04 XX:XX:XX Generate general API Info, search dir:./
2022/02/04 XX:XX:XX Generating api.Response
2022/02/04 XX:XX:XX create docs.go at  docs/docs.go
2022/02/04 XX:XX:XX create swagger.json at  docs/swagger.json
2022/02/04 XX:XX:XX create swagger.yaml at  docs/swagger.yaml

go build
```

with swag 1.7.9
----------
```shell
go install github.com/swaggo/swag/cmd/swag@1.7.9
swag init
2022/02/04 XX:XX:XX Generate swagger docs....
2022/02/04 XX:XX:XX Generate general API Info, search dir:./
2022/02/04 XX:XX:XX Generating api.Response
2022/02/04 XX:XX:XX create docs.go at  docs/docs.go
2022/02/04 XX:XX:XX create swagger.json at  docs/swagger.json
2022/02/04 XX:XX:XX create swagger.yaml at  docs/swagger.yaml
go build
# github.com/ubogdan/http-swagger-example/docs
docs/docs.go:61:28: undefined: "github.com/swaggo/swag".Spec
```
