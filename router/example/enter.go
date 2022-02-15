package example

import "github.com/gin-gonic/gin"

type Group struct{}

func (*Group) RegisterRouter(public *gin.RouterGroup, private *gin.RouterGroup) {

}
