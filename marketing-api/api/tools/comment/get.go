package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// Get 获取评论列表
func Get(clt *core.SDKClient, accessToken string, req *comment.GetRequest) (*comment.GetResponseData, error) {
	var resp comment.GetResponse
	if err := clt.Get("v3.0/tools/comment/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
