package main

import (
	"encoding/json"
	"eventAnalysis/pkg/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ()

type CodeMessage struct {
	Code    int    `json:"code"`
	Version string `json:"version"`
}

func (c *CodeMessage) ToJsonString() string {
	b, _ := json.Marshal(c)
	return string(b)
}
func main() {
	router := gin.Default()
	route.PullimageTime(router)
	http.ListenAndServe(":9999", router)
}
