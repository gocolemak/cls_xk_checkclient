package cls_xk_checkclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Reply 应答, 只包含发起检查的状态
type Reply struct {
	Status int `json:"status" bson:"status"` // 调用成功 1
}

// CheckClient 客户端类
type CheckClient struct {
	Address   string
	user      int
	checkType string
	isSync    int
}

func (c CheckClient) CallCheckAsync(user int, checkType string) (Reply, error) {
	c.user = user
	c.checkType = checkType
	c.isSync = 0
	return c.callCheckHandler()
}

func (c CheckClient) CallCheckSync(user int, checkType string) (Reply, error) {
	c.user = user
	c.checkType = checkType
	c.isSync = 1
	return c.callCheckHandler()
}

func (c CheckClient) callCheckHandler() (Reply, error) {
	// 1. 参数校验并解析
	// 1.1 判空,不支持为空, 为空则直接返回
	if c.checkType == "" {
		return Reply{Status: 0}, errors.New("检查项标识不能为空")
	}
	// 2. 服务发现 c.address
	// 远程调用 check_work , 发送rpc or http
	url := "http://" + c.Address + "/xk-api/checkwork/do?user=" + fmt.Sprintf("%d", c.user) + "&check_type=" + c.checkType + "&is_sync=" + fmt.Sprintf("%d", c.isSync)
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Reply{Status: 0}, errors.New("触发失败")
	}
	// 3. 服务调用
	response, _ := client.Do(reqest)
	bodyRes, err := ioutil.ReadAll(response.Body)

	var rc Reply
	json.Unmarshal(bodyRes, &rc)
	// 4. 返回结果
	return rc, nil
}
