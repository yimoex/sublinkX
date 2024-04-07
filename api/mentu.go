package api

import (
	"github.com/gin-gonic/gin"
)

type Meta struct {
	Title     string   `json:"title"`
	Icon      string   `json:"icon"`
	Hidden    bool     `json:"hidden"`
	Roles     []string `json:"roles"`
	KeepAlive bool     `json:"keepAlive,omitempty"`
}

type Child struct {
	Path      string `json:"path"`
	Component string `json:"component"`
	Name      string `json:"name"`
	Meta      Meta   `json:"meta"`
}

type Menu struct {
	Path      string  `json:"path"`
	Component string  `json:"component"`
	Redirect  string  `json:"redirect"`
	Name      string  `json:"name"`
	Meta      Meta    `json:"meta"`
	Children  []Child `json:"children"`
}

func GetMenus(c *gin.Context) {
	menus := []Menu{
		{
			Path:      "/system",
			Component: "Layout",
			Redirect:  "/system/user",
			Name:      "system",
			Meta: Meta{
				Title:  "system",
				Icon:   "system",
				Hidden: true,
				Roles:  []string{"ADMIN"},
			},
			Children: []Child{
				{
					Path:      "user/set",
					Component: "system/user/set",
					Name:      "Userset",
					Meta: Meta{
						Title:     "userset",
						Icon:      "role",
						Hidden:    true,
						Roles:     []string{"ADMIN"},
						KeepAlive: true,
					},
				},
			},
		},
		// 订阅管理
		{
			Path:      "/subcription",
			Component: "Layout",
			Redirect:  "/system/user",
			Name:      "subcription",
			Meta: Meta{
				Title:  "subcription",
				Icon:   "system",
				Hidden: false,
				Roles:  []string{"ADMIN"},
			},
			Children: []Child{
				{
					Path:      "subs",
					Component: "subcription/subs",
					Name:      "Subs",
					Meta: Meta{
						Title:     "subs",
						Icon:      "role",
						Hidden:    false,
						Roles:     []string{"ADMIN"},
						KeepAlive: true,
					},
				},
				{
					Path:      "nodes",
					Component: "subcription/nodes",
					Name:      "Nodes",
					Meta: Meta{
						Title:     "Nodes",
						Icon:      "role",
						Hidden:    false,
						Roles:     []string{"ADMIN"},
						KeepAlive: true,
					},
				},
				{
					Path:      "iplogs",
					Component: "subcription/iplogs",
					Name:      "Iplogs",
					Meta: Meta{
						Title:     "Iplogs",
						Icon:      "role",
						Hidden:    false,
						Roles:     []string{"ADMIN"},
						KeepAlive: true,
					},
				},
			},
		},
	}
	c.JSON(200, gin.H{
		"code": "00000",
		"data": menus,
		"msg":  "获取成功",
	})
}
