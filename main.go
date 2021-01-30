package main

import (
	"homework/internal/DbHelper"
	. "homework/internal/DbHelper"
	"net/http"
	"strconv" // Generated by https://quicktype.io

	"github.com/gin-gonic/gin"
)

var roles = []Role{{
	ID:      1,
	Name:    "阿修羅",
	Summary: "死國魖族最強者，在歷經多場戰役統一死國，直接挑戰天者權威，不服其領導，要求讓三族和平共處，因此被天者以增加資源為名，前往打通連接火宅佛獄、死國及苦境的莫汗走廊，未料在工程半途被楓岫主人所乘駕的隕石撞毀，阿修羅在天者刻意算計下意外身亡，身葬苦境與死國間的異次元。",
	Skills: []Skill{{
		ID:   1,
		Type: "武學",
		Name: "天之爆",
	},
		{
			ID:   2,
			Type: "武學",
			Name: "魔之狂",
		},
		{
			ID:   3,
			Type: "武學",
			Name: "天之渦",
		},
		{
			ID:   4,
			Type: "武學",
			Name: "闇之爆",
		},
		{
			ID:   5,
			Type: "法術",
			Name: "山河凝元·天地共引",
		},
		{
			ID:   6,
			Type: "法術",
			Name: "地之火·九天滅絕",
		}},
}, {
	ID:      2,
	Name:    "白塵子",
	Summary: "火宅佛獄凱旋侯的副體之一，本名黑枒君，臥底中原武林，向佛獄通風報信，最後被素還真所殺並冒充其身份一探佛獄之秘。",
	Skills: []Skill{{
		ID:   7,
		Type: "武學",
		Name: "凝宇化空",
	},
		{
			ID:   8,
			Type: "武學",
			Name: "反神源",
		},
	},
}}

func main() {

	router := gin.Default()

	router.GET("/role", func(c *gin.Context) {
		c.JSON(http.StatusOK, roles)
	})

	router.GET("/role/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		for i := 0; i < len(roles); i++ {
			if roles[i].ID == id {
				c.JSON(http.StatusOK, roles[i])
				break
			}
		}

		c.JSON(http.StatusNotFound, nil)
	})

	router.POST("/role", func(c *gin.Context) {
		var u Role
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		DbHelper.SetRoleSkillID(&roles, &u)
		roles = append(roles, u)

		c.JSON(http.StatusOK, u)
	})

	router.PUT("/role/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		var u Role
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		for i := 0; i < len(roles); i++ {
			if roles[i].ID == id {
				roles[i].Name = u.Name
				roles[i].Summary = u.Summary
				c.JSON(http.StatusOK, roles[i])
				break
			}
		}
	})

	router.DELETE("/role/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		for i, role := range roles {
			if role.ID == id {
				roles = append(roles[0:i], roles[i+1:]...)
				break
			}
		}

		c.Status(http.StatusNoContent)
	})

	router.Run(":8080")
}