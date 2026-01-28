package utils

import (
	"strings"
)

const (
	// RecorderKeyPrefix format: recorder_<recorderId>-FIELD_<field>
	RecorderKeyPrefix = "recorder_"
	// RecorderKeyFieldPrefix is the separator between the recorderId and the field.
	RecorderKeyFieldPrefix = "-FIELD_"
)

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
