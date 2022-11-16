package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type expert struct {
	ENT_ID                     string `json "ENT_ID"`
	BLACK_NO                   string `json "BLACK_NO"`
	EXPERT_ID                  string `json "EXPERT_ID"`
	EXPERT_TYPE                string `json "EXPERT_TYPE"`
	EXPERT_FIRST_NAME          string `json "EXPERT_FIRST_NAME"`
	EXPERT_LAST_NAME           string `json "EXPERT_LAST_NAME"`
	EXPERT_LANG                string `json "EXPERT_LANG"`
	EXPERT_TRAVELING_ALLOWANCE string `json "EXPERT_TRAVELING_ALLOWANCE"`
	EXPERT_HOUSING_ALLOWANCE   string `json "EXPERT_HOUSING_ALLOWANCE"`
	EXPERT_OTH_FEE             string `json "EXPERT_OTH_FEE"`
	UPDATE_USER                string `json "UPDATE_USER"`
	UPDATE_TIME                string `json "UPDATE_TIME"`
}

var experts = []expert{
	{ENT_ID: "3020701", BLACK_NO: "ไม่รู้", EXPERT_ID: "ไม่รู้", EXPERT_TYPE: "ไม่รู้", EXPERT_FIRST_NAME: "ไม่รู้", EXPERT_LAST_NAME: "ไม่รู้",
		EXPERT_LANG: "ไม่รู้", EXPERT_TRAVELING_ALLOWANCE: "ไม่รู้", EXPERT_HOUSING_ALLOWANCE: "ไม่รู้", EXPERT_OTH_FEE: "ไม่รู้", UPDATE_USER: "ไม่รู้", UPDATE_TIME: "ไม่รู้"},
}

func main() {
	// router := gin.Default()
	// router.Run("localhost:9000")
	fmt.Println("Start API")
	// fmt.Printf("%v", details[0].mock_1)
	r := gin.New()
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, experts)
	})
	r.Run()

}
