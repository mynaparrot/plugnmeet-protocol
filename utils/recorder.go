package utils

import (
	"fmt"
	"strings"

	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
)

const (
	TranscodingQueueStreamName = "pnm-RecorderTranscoderJobs"
	TranscoderConsumerDurable  = "transcoderWorker"
	// RecorderKeyPrefix format: recorder_<recorderId>-FIELD_<field>
	RecorderKeyPrefix = "recorder_"
	// RecorderKeyFieldPrefix is the separator between the recorderId and the field.
	RecorderKeyFieldPrefix = "-FIELD_"
)

type RecorderInfo struct {
	RecorderId      string
	MaxLimit        int64
	CurrentProgress int64
	LastPing        int64
	TotalCores      int64
	CpuScore        float64
	Variant         plugnmeet.CloudRecordingVariants
}

// GetTranscodingJobSubject returns the NATS subject for a recording variant's transcoding jobs.
func GetTranscodingJobSubject(streamName string, variant plugnmeet.CloudRecordingVariants) string {
	return fmt.Sprintf("%s.%d", streamName, int32(variant))
}

// GetTranscoderConsumerDurable returns the durable consumer name for a recording variant.
func GetTranscoderConsumerDurable(variant plugnmeet.CloudRecordingVariants) string {
	return fmt.Sprintf("%s_%d", TranscoderConsumerDurable, int32(variant))
}

// FormatRecorderKey generates a key for a specific recorder's field within the consolidated bucket.
// The format will be `recorder_<recorderId>-FIELD_<field>`.
func FormatRecorderKey(recorderId, field string) string {
	return RecorderKeyPrefix + recorderId + RecorderKeyFieldPrefix + field
}

// ParseRecorderKey extracts the recorderId and field from a consolidated key.
func ParseRecorderKey(key string) (recorderId, field string, ok bool) {
	if !strings.HasPrefix(key, RecorderKeyPrefix) {
		return "", "", false
	}
	parts := strings.SplitN(strings.TrimPrefix(key, RecorderKeyPrefix), RecorderKeyFieldPrefix, 2)
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}
