package nsq

import (
	"context"
	"encoding/json"
)

const (
	nsqTopicAppBuilder      = "dynomo_build_app"      // for build app queue
	nsqTopicKeystoreBuilder = "dynomo_build_keystore" // for build keystore queue
)

func (r *Repository) PublishBuildApp(ctx context.Context, param BuildAppParam) error {
	b, err := json.Marshal(param)
	if err != nil {
		return err
	}

	return r.nsq.Publish(nsqTopicAppBuilder, b)
}

func (r *Repository) PublishBuildKeystore(ctx context.Context, param BuildKeystoreParam) error {
	b, err := json.Marshal(param)
	if err != nil {
		return err
	}

	return r.nsq.Publish(nsqTopicKeystoreBuilder, b)
}
