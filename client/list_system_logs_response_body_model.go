// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSystemLogs(v []*ListSystemLogsResponseBodySystemLogs) *ListSystemLogsResponseBody
	GetSystemLogs() []*ListSystemLogsResponseBodySystemLogs
	SetTotalCount(v int64) *ListSystemLogsResponseBody
	GetTotalCount() *int64
}

type ListSystemLogsResponseBody struct {
	SystemLogs []*ListSystemLogsResponseBodySystemLogs `json:"SystemLogs,omitempty" xml:"SystemLogs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSystemLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemLogsResponseBody) GetSystemLogs() []*ListSystemLogsResponseBodySystemLogs {
	return s.SystemLogs
}

func (s *ListSystemLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSystemLogsResponseBody) SetSystemLogs(v []*ListSystemLogsResponseBodySystemLogs) *ListSystemLogsResponseBody {
	s.SystemLogs = v
	return s
}

func (s *ListSystemLogsResponseBody) SetTotalCount(v int64) *ListSystemLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSystemLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSystemLogsResponseBodySystemLogs struct {
	// example:
	//
	// You are using******
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2020-10-08T16:00:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// Error。
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s ListSystemLogsResponseBodySystemLogs) String() string {
	return dara.Prettify(s)
}

func (s ListSystemLogsResponseBodySystemLogs) GoString() string {
	return s.String()
}

func (s *ListSystemLogsResponseBodySystemLogs) GetContent() *string {
	return s.Content
}

func (s *ListSystemLogsResponseBodySystemLogs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListSystemLogsResponseBodySystemLogs) GetLevel() *string {
	return s.Level
}

func (s *ListSystemLogsResponseBodySystemLogs) SetContent(v string) *ListSystemLogsResponseBodySystemLogs {
	s.Content = &v
	return s
}

func (s *ListSystemLogsResponseBodySystemLogs) SetGmtCreateTime(v string) *ListSystemLogsResponseBodySystemLogs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSystemLogsResponseBodySystemLogs) SetLevel(v string) *ListSystemLogsResponseBodySystemLogs {
	s.Level = &v
	return s
}

func (s *ListSystemLogsResponseBodySystemLogs) Validate() error {
	return dara.Validate(s)
}
