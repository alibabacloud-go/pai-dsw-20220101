// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *GetInstanceRequest
	GetToken() *string
}

type GetInstanceRequest struct {
	// The sharing token information.
	//
	// example:
	//
	// WUzWCMr325LV0bH2JH4C4HoDaKIU6C4S
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) GetToken() *string {
	return s.Token
}

func (s *GetInstanceRequest) SetToken(v string) *GetInstanceRequest {
	s.Token = &v
	return s
}

func (s *GetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
