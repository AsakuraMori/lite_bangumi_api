/**
 * @file 	api_subjects.go
 * @brief 	关于subjects的API接口
 * @author 	AsakuraMori
 * @version 0.3.0
 * @date 	2025-03-25
 */

package lite_bangumi_api

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
SearchSubjectsByName

  - @brief 通过字符串搜索条目。

    API：/v0/search/subjects

  - @param

    【limit】：当前页最大数量

    【offset】：起始位置

    【requestBody】：请求体，结构如下：

    {
    "keyword": "string",
    "sort": "rank",
    "filter": {
    "type": [
    2
    ]
    }
    }

    【client】：http.Client对象

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchSubjectsByName(limit, offset string, requestBody string, client *http.Client) ([]byte, error) {

	baseURL := "https://api.bgm.tv/v0/search/subjects"

	params := url.Values{}
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	jsonData, err := getJsonDataFromURL("POST", apiURL, requestBody, client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchSubjectsById

  - @brief 通过搜索条目ID名获取ID信息。

    API：/v0/subjects/{subject_id}

  - @param

    【subID】：条目ID。

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchSubjectsById(subID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/subjects/"
	apiURL := fmt.Sprintf("%s%s", baseURL, subID)
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchAllSubjectsByName

  - @brief 通过关键字搜索所有信息。

    API：/search/subject/{keywords}

  - @param 【keyWord】：关键字。

    【typeName】：条目类型（只能是以下字符串：书籍、动漫、音乐、游戏、三次元。如果typeName不满足以上字符串，则将全局搜索）

    【responseGroup】：返回数据大小（只能指定为small、medium、large）

    【start】：开始的条目

    【nmaxResults】：每页最大数量

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchAllSubjectsByName(keyWord, typeName, responseGroup, start, nmaxResults string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/search/subject/"
	params := url.Values{}

	sType := 0
	switch typeName {
	case "书籍":
		sType = 1
	case "动漫":
		sType = 2
	case "音乐":
		sType = 3
	case "游戏":
		sType = 4
	case "三次元":
		sType = 6
	default:
		sType = 0
	}

	params.Add("type", fmt.Sprintf("%d", sType))
	params.Add("responseGroup", fmt.Sprintf("%s", responseGroup))
	params.Add("start", fmt.Sprintf("%s", start))
	params.Add("max_results", fmt.Sprintf("%s", nmaxResults))
	apiURL := fmt.Sprintf("%s%s?%s", baseURL, keyWord, params.Encode())

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
GetCalender

  - @brief 获取每日放送。

    API：/calendar

  - @param

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func GetCalender(client *http.Client) ([]byte, error) {
	apiURL := "https://api.bgm.tv/calendar"

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
