package wgin

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getDefaultInt64(value string, defaultValue ...int64) int64 {
	d := int64(0)
	if len(defaultValue) > 0 {
		d = defaultValue[0]
	}
	if value == "" {
		return d
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return d
	}
	return v
}
func getDefaultUint64(value string, defaultValue ...uint64) uint64 {
	d := uint64(0)
	if len(defaultValue) > 0 {
		d = defaultValue[0]
	}
	if value == "" {
		return d
	}
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return d
	}
	return v
}

func getDefaultInt(value string, defaultValue ...int) int {
	d := 0
	if len(defaultValue) > 0 {
		d = defaultValue[0]
	}
	if value == "" {
		return d
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return d
	}
	return v
}

type CommonContext struct {
	ctx *gin.Context
}

func (c *CommonContext) JSONResponse(data interface{}) {
	c.ctx.JSON(200, ReturnCode{
		Code: SuccessCode,
		Data: data,
	})
}

func (c *CommonContext) JSON(data interface{}, statusCode ...int) {
	code := 200
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.ctx.JSON(code, data)
}
func (c *CommonContext) Ok(msg ...string) {
	rc := ReturnCode{Code: SuccessCode, Msg: "ok"}
	if len(msg) > 0 {
		rc.Msg = msg[0]
	}
	c.ctx.JSON(200, rc)
}
func (c *CommonContext) AbortWithJSON(data interface{}, statusCode ...int) {
	code := 200
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.ctx.AbortWithStatusJSON(code, data)
}

func BuildContext(ctx *gin.Context) *CommonContext {
	return &CommonContext{
		ctx: ctx,
	}
}

func (c *CommonContext) Param(key string) string {
	return c.ctx.Param(key)
}

func (c *CommonContext) ParamInt(key string, defaultValue ...int) int {
	return getDefaultInt(c.ctx.Param(key), defaultValue...)
}

// if value is not int64 return 0 or default value
func (c *CommonContext) ParamInt64(key string, defaultValue ...int64) int64 {
	return getDefaultInt64(c.ctx.Param(key), defaultValue...)
}

// if value is not uint64 return 0 or default value
func (c *CommonContext) ParamUint64(key string, defaultValue ...uint64) uint64 {
	return getDefaultUint64(c.ctx.Param(key), defaultValue...)
}

// if value is not int64 return 0 or default value
func (c *CommonContext) URLParamInt64(key string, defaultValue ...int64) int64 {
	return getDefaultInt64(c.ctx.Query(key), defaultValue...)
}

func (c *CommonContext) URLParamUint64(key string, defaultValue ...uint64) uint64 {
	return getDefaultUint64(c.ctx.Query(key), defaultValue...)
}

func (c *CommonContext) URLParam(key string) string {
	return c.ctx.Query(key)
}

func (c *CommonContext) URLParamExists(key string) bool {
	_, ok := c.ctx.GetQuery(key)
	return ok
}

func (c *CommonContext) GetValue(key string) interface{} {
	any, ok := c.ctx.Get(key)
	if !ok {
		return nil
	}
	return any
}

func (c *CommonContext) SetValue(key string, value interface{}) {
	c.ctx.Set(key, value)
}

func (c *CommonContext) Next() {
	c.ctx.Next()
}
