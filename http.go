package sdk_sdk

import (
	"time"

	"github.com/gotomicro/ego/client/ehttp"
)

func (sm *SdkManager) initHttpClient(options ClientOptions) {
	if &options == nil {
		options = ClientOptions{}
	}
	sm.httpClient = NewRestyClientWithOptions(options)
}

type ClientOptions struct {
	Timeout          time.Duration
	RetryAttempts    int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
}

func NewRestyClientWithOptions(options ClientOptions) *ehttp.Component {
	client := ehttp.DefaultContainer().Build()
	if options.RetryAttempts > 0 {
		client.SetRetryCount(options.RetryAttempts)

		if options.RetryWaitTime != 0 {
			client.SetRetryWaitTime(options.RetryWaitTime)
		}

		if options.RetryMaxWaitTime != 0 {
			client.SetRetryMaxWaitTime(options.RetryMaxWaitTime)
		}
	}
	if options.Timeout != 0 {
		client.SetTimeout(options.Timeout)
	}

	return client
}
