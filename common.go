/**
 * @file 	common.go
 * @brief 	API包装函数以及全局变量
 * @author 	AsakuraMori
 * @version 0.3.0
 * @date 	2025-03-25
 */

package lite_bangumi_api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
)

/**
 * @brief 定义Token和UserAgent
 */
var (
	Token     string
	UserAgent string
)

/*
*
  - @brief 从URL获取Json数据
  - @param 【method】：方法
    【url】：地址
    【requestBody】：请求体
    【client】：http.Client对象
  - @return 返回一个[]byte和一个err。
  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func getJsonDataFromURL(method, url, requestBody string, client *http.Client) ([]byte, error) {
	var req *http.Request
	var err error
	if len(requestBody) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(requestBody)))
	}
	if err != nil {
		errMsg := errors.New("getJsonDataFromURL：不正确的请求")
		return nil, errMsg
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		errMsg := errors.New("getJsonDataFromURL：连接失败或超时")
		return nil, errMsg
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := errors.New("getJsonDataFromURL：错误的返回码:" + strconv.Itoa(resp.StatusCode))
		return nil, errMsg
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := errors.New("getJsonDataFromURL：获取信息失败")
		return nil, errMsg
	}

	return body, nil
}

/*
*
  - @brief 从URL获取Bool数据
  - @param 【method】：方法
    【url】：地址
    【requestBody】：请求体
    【client】：http.Client对象
  - @return 返回一个bool和一个err。
  - @retval  如果bool为true，err为nil。如果bool为false，err表示错误信息
*/
func getBoolDataFromURL(method, url, requestBody string, client *http.Client) error {
	var req *http.Request
	var err error
	if len(requestBody) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(requestBody)))
	}
	if err != nil {
		errMsg := errors.New("getBoolDataFromURL：不正确的请求")
		return errMsg
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		errMsg := errors.New("getBoolDataFromURL：连接失败或超时")
		return errMsg
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		errMsg := errors.New("getBoolDataFromURL：错误的返回码:" + strconv.Itoa(resp.StatusCode))
		return errMsg
	}

	return nil
}
