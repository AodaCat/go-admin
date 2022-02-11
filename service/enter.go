package service

import "go-admin/service/system"

type Group struct {
	SystemGroup system.Group
}

var GroupApp = new(Group)
