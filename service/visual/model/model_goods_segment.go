package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type GoodsSegmentData struct {
	ImgUrl string `json:"img_url"`
}

type GoodsSegmentResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *GoodsSegmentData      `json:"data"`
}