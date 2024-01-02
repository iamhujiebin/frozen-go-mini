package mycontext

import (
	"context"
	"frozen-go-mini/common/mylogrus"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	TRACEID     = "traceId"
	USERID      = "userId"
	DEVICETYPE  = "deviceType"
	APP_VERSION = "appVersion"
	URL         = "url"
	METHOD      = "method"
)

/**
 * 主要是完成日志打印
 * @param
 * @return
 **/

type MyContext struct {
	context.Context
	Log *logrus.Entry
	Cxt map[string]interface{}
}

func CreateMyContext(ctx map[string]interface{}) *MyContext {
	var traceId string
	if traceIdTemp, ok := ctx[TRACEID]; ok {
		traceId, ok = traceIdTemp.(string)
	} else {
		traceId = strings.Replace(uuid.NewV4().String(), "-", "", -1)
	}

	var userId string
	if userIdTemp, ok := ctx[USERID]; ok {
		userId = strconv.FormatUint(userIdTemp.(uint64), 10)
	}

	var deviceTypeStr string
	if deviceTypeTemp, ok := ctx[DEVICETYPE]; ok {
		deviceTypeStr, ok = deviceTypeTemp.(string)
	}

	var sAppVersion string
	if appVersionTmp, ok := ctx[APP_VERSION]; ok {
		sAppVersion, ok = appVersionTmp.(string)
	}

	var url string
	if urlTmp, ok := ctx[URL]; ok {
		url, ok = urlTmp.(string)
	}

	var method string
	if methodTmp, ok := ctx[METHOD]; ok {
		method, ok = methodTmp.(string)
	}
	_ctx := context.WithValue(context.Background(), "traceId", traceId)
	_ctx = context.WithValue(_ctx, "userId", userId)
	return &MyContext{
		Context: _ctx,
		Log:     CreateContextLog(userId, traceId, deviceTypeStr, sAppVersion, url, method),
		Cxt:     ctx,
	}
}

/**
 * 创建上下文的日志
 **/
func CreateContextLog(userId string, traceId string, deviceType string, deviceVersion string, url string, method string) *logrus.Entry {
	return mylogrus.MyLog.WithFields(logrus.Fields{
		USERID:      userId,
		TRACEID:     traceId,
		DEVICETYPE:  deviceType,
		APP_VERSION: deviceVersion,
		URL:         url,
		METHOD:      method,
	})
}
