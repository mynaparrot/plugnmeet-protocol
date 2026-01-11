package nats

import (
	"fmt"
	"strings"

	"github.com/mynaparrot/plugnmeet-protocol/utils"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NatsInfo struct {
	NatsUrls                 []string         `yaml:"nats_urls"`
	NatsWSUrls               []string         `yaml:"nats_ws_urls"`
	Account                  string           `yaml:"account"`
	User                     string           `yaml:"user"`
	Password                 string           `yaml:"password"`
	Nkey                     *string          `yaml:"nkey"`
	AuthCalloutEnabled       *bool            `yaml:"auth_callout_enabled"`
	AuthCalloutIssuerPrivate string           `yaml:"auth_callout_issuer_private"`
	AuthCalloutXkeyPrivate   *string          `yaml:"auth_callout_xkey_private"`
	NumReplicas              int              `yaml:"num_replicas"`
	Subjects                 NatsSubjects     `yaml:"subjects"`
	Recorder                 NatsInfoRecorder `yaml:"recorder"`
}

type NatsSubjects struct {
	SystemApiWorker string `yaml:"system_api_worker"`
	SystemJsWorker  string `yaml:"system_js_worker"`
	SystemPublic    string `yaml:"system_public"`
	SystemPrivate   string `yaml:"system_private"`
	Chat            string `yaml:"chat"`
	Whiteboard      string `yaml:"whiteboard"`
	DataChannel     string `yaml:"data_channel"`
}

type NatsInfoRecorder struct {
	RecorderChannel string `yaml:"recorder_channel"`
	RecorderInfoKv  string `yaml:"recorder_info_kv"`
	TranscodingJobs string `yaml:"transcoding_jobs_subject"`
}

func NewNatsConnection(info *NatsInfo) (*nats.Conn, jetstream.JetStream, error) {
	var opt nats.Option
	var err error

	if info.Nkey != nil {
		opt, err = utils.NkeyOptionFromSeedText(*info.Nkey)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create nkey option: %w", err)
		}
	} else {
		opt = nats.UserInfo(info.User, info.Password)
	}

	nc, err := nats.Connect(strings.Join(info.NatsUrls, ","), opt)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to nats: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create jetstream: %w", err)
	}

	return nc, js, nil
}
