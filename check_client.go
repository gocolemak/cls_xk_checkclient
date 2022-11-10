package cls_xk_checkclient

// ReportContent 报告单, 只包含检查的结果和状态
type ReportContent struct {
	Status int `json:"status" bson:"status"` // 调用成功 1
}

type Check struct {
	Address string
}

func (c Check) CallCheckAsync(user string, check_type string) ReportContent {
	// 1. 参数校验\解析
	// 2. 服务发现 c.address
	// 3. 服务调用
	// 4. 返回结果

	// TODO 远程调用 check_work, 发送rpc or http
	/*
		UserRpc.GetUser(l.ctx, &user.IdReq{
			Id: userId,
		})
	*/
	//
	return ReportContent{
		Status: 0,
	}
}

func CallCheckSync(user string, check_type string, address string) ReportContent {
	// TODO 远程调用 check_work , 发送rpc or http
	return ReportContent{
		Status: 0,
	}
}
