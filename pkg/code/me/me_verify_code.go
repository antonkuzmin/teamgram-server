// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package me

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/pkg/code/conf"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	_smsURL = "http://127.0.0.1:8181/code?phone=%s&code=%s"
)

func New(c *conf.SmsVerifyCodeConfig) *meVerifyCode {
	return &meVerifyCode{
		code: c,
	}
}

type meVerifyCode struct {
	code *conf.SmsVerifyCodeConfig
}

func (m *meVerifyCode) SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string) (string, error) {
	baseUrl := m.code.SendCodeUrl
	if baseUrl == "" {
		baseUrl = "http://127.0.0.1:8181/code"
	}
	urlV := baseUrl + fmt.Sprintf("?phone=%s&code=%s", phoneNumber, code)
	logx.Infof("send me sms: %s", urlV)
	resp, err := http.Get(urlV)
	if err != nil {
		logx.Errorf("request verify code error: %v", err)
		return code, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("request verify code error: %v", err)
		return code, err
	}
	logx.Infof("result: %s", hack.String(body))
	return code, nil
}

func (m *meVerifyCode) VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error {
	if len(code) != 5 {
		return mtproto.ErrPhoneCodeInvalid
	}

	if code != extraData {
		return mtproto.ErrPhoneCodeInvalid
	}

	return nil
}
