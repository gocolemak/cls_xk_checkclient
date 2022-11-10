package cls_xk_checkclientv3

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReportContent 报告单, 只包含检查的结果和状态
type ReportContent struct {
	Status int `json:"status" bson:"status"` // 调用成功 1
}

type Check struct {
	Address   string
	user      string
	checkType string
	isSync    int
}

func (c Check) CallCheckAsync(user string, checkType string) (ReportContent, error) {
	c.user = user
	c.checkType = checkType
	c.isSync = 0
	return c.callCheckHandler()
}

func (c Check) CallCheckSync(user string, checkType string) (ReportContent, error) {
	// TODO 远程调用 check_work , 发送rpc or http
	c.user = user
	c.checkType = checkType
	c.isSync = 1
	return c.callCheckHandler()
}

func (c Check) callCheckHandler() (ReportContent, error) {
	// 1. 参数校验并解析
	// 1.1 判空,不支持为空, 为空则直接返回
	if c.user == "" {
		return ReportContent{Status: 0}, errors.New("用户id不能为空")
	}
	if c.checkType == "" {
		return ReportContent{Status: 0}, errors.New("检查项标识不能为空")
	}
	// 2. 服务发现 c.address
	url := c.Address + "/xk-api/checkwork/do?user=" + c.user + "&" + "check_type=" + c.checkType + "is_sync=" + fmt.Sprintf("%d", c.isSync)
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ReportContent{Status: 0}, errors.New("触发失败")
	}
	// 3. 服务调用
	response, _ := client.Do(reqest)
	bodyRes, err := ioutil.ReadAll(response.Body)

	var rc ReportContent
	json.Unmarshal(bodyRes, &rc)
	// 4. 返回结果
	return rc, nil
}
