package sdk

import (
	"context"
	"github.com/iimeta/fastapi-sdk/aliyun"
	"github.com/iimeta/fastapi-sdk/baidu"
	"github.com/iimeta/fastapi-sdk/consts"
	"github.com/iimeta/fastapi-sdk/deepseek"
	"github.com/iimeta/fastapi-sdk/google"
	"github.com/iimeta/fastapi-sdk/logger"
	"github.com/iimeta/fastapi-sdk/model"
	"github.com/iimeta/fastapi-sdk/openai"
	"github.com/iimeta/fastapi-sdk/xfyun"
	"github.com/iimeta/fastapi-sdk/zhipuai"
)

type Chat interface {
	ChatCompletion(ctx context.Context, request model.ChatCompletionRequest) (res model.ChatCompletionResponse, err error)
	ChatCompletionStream(ctx context.Context, request model.ChatCompletionRequest) (responseChan chan *model.ChatCompletionResponse, err error)
	Image(ctx context.Context, request model.ImageRequest) (res model.ImageResponse, err error)
}

func NewClient(ctx context.Context, corp, model, key, baseURL, path string, proxyURL ...string) Chat {

	logger.Infof(ctx, "NewClient corp: %s, model: %s, key: %s", corp, model, key)

	switch corp {
	case consts.CORP_OPENAI:
		return openai.NewClient(ctx, model, key, baseURL, path, proxyURL...)
	case consts.CORP_BAIDU:
		return baidu.NewClient(ctx, model, key, baseURL, path, proxyURL...)
	case consts.CORP_XFYUN:
		return xfyun.NewClient(ctx, model, key, baseURL, path, proxyURL...)
	case consts.CORP_ALIYUN:
		return aliyun.NewClient(ctx, model, key, baseURL, path, proxyURL...)
	case consts.CORP_ZHIPUAI:
		return zhipuai.NewClient(ctx, model, key, baseURL, path, proxyURL...)
	case consts.CORP_GOOGLE:
		return google.NewClient(ctx, model, key, baseURL, path, proxyURL...)
	case consts.CORP_DEEPSEEK:
		return deepseek.NewClient(ctx, model, key, baseURL, path, proxyURL...)
	}

	return openai.NewClient(ctx, model, key, baseURL, path, proxyURL...)
}
