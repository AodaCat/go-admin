package v1

import (
	"go-admin/api/v1/autocode"
	"go-admin/api/v1/example"
	"go-admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
