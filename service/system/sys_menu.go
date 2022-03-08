package system

import (
	"go-admin/global"
	"go-admin/model/system"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) GetMenuTree(authorityId string) (menus []system.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func (menuService *MenuService) getMenuTreeMap(authorityId string) (treeMap map[string][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	treeMap = make(map[string][]system.SysMenu)
	err = global.GA_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
