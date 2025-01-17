// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazahttp

import (
	"io"
	"net/http"
	"os"
	"time"
)

// 获取 http 文件保存至本地
func DownloadHttpFile(url string, saveTo string, timeoutMSec int) (int64, error) {
	var c http.Client
	if timeoutMSec > 0 {
		c.Timeout = time.Duration(timeoutMSec) * time.Millisecond
	}
	resp, err := c.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	fp, err := os.Create(saveTo)
	if err != nil {
		return -1, err
	}
	defer fp.Close()

	return io.Copy(fp, resp.Body)
}
