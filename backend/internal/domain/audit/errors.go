package audit

import "errors"

var (
	ErrAuditWriteFailed = errors.New("failed to write audit log")
)
