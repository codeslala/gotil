package authorize

import (
	"github.com/codeslala/gotil/api"
	"github.com/codeslala/gotil/errs"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthCheckMW(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	err := authorize(authorization)
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			api.Result400(c, errs.InvalidArgument.Wrap(err))
			return
		}
		if status.Code(err) == codes.Unauthenticated {
			api.Result400(c, errs.Unauthenticated.Wrap(err))
			return
		}
		api.Result400(c, errs.Internal.Wrap(err))
		return
	}
	c.Next()
}
