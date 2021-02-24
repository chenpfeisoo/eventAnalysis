package route

import (
	"eventAnalysis/pkg/model"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
	"net/http"
)

func PullimageTime(r *gin.Engine) {
	r.POST("/show", func(c *gin.Context) {
		var pinfo *model.PullImageInfo
		if err := c.ShouldBindJSON(&pinfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if pinfo != nil {
			klog.Info(pinfo.ToJsonString())
			model.InsertEvent(pinfo)
		} else {
			klog.Info("has not data")
		}
	})
	r.LoadHTMLGlob("/Users/felixchen/go/src/eventAnalysis/templates/index.tmpl")
	r.GET("/show", func(c *gin.Context) {
		klog.Info(model.Show().ToJsonString())
		//c.String(200,model.Show())
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"events": model.Show()})
	})
}
