/**
 * @file 	api_episodes.go
 * @brief 	关于episodes的API接口
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
SearchEpisodesByEpisodesName

  - @brief 通过搜索条目ID名获取章节信息。

    API：/v0/episodes

  - @param

    【sbjID】：条目ID。

    【typeName】：类型（只能是以下类型：本篇、特别篇、OP、ED、预告/宣传/广告、MAD、其他。如果不是以上的字符串，则全局搜索）

    【limit】：当前页面显示条目最大数量

    【offset】：开始的条目位置

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchEpisodesByEpisodesName(sbjID, typeName, limit, offset string, client *http.Client) ([]byte, error) {

	baseURL := "https://api.bgm.tv/v0/episodes"

	params := url.Values{}
	sType := 0
	switch typeName {
	case "本篇":
		sType = 0
	case "特别篇":
		sType = 1
	case "OP":
		sType = 2
	case "ED":
		sType = 3
	case "预告/宣传/广告":
		sType = 4
	case "MAD":
		sType = 5
	case "其他":
		sType = 6
	default:
		sType = 0
	}
	params.Add("subject_id", fmt.Sprintf("%s", sbjID))
	params.Add("type", fmt.Sprintf("%d", sType))
	params.Add("limit", fmt.Sprintf("%s", limit))
	params.Add("offset", fmt.Sprintf("%s", offset))

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
SearchEpisodesByEpisodesId

  - @brief 通过搜索章节ID名章节ID信息。

    API：/v0/episodes/{episode_id}

  - @param

    【epiID】：章节ID。

    【client】：http.Client对象。

  - @return 返回一个[]byte和一个err。

  - @retval []byte是返回体，err表示错误。如果err为nil，则没有错误。
*/
func SearchEpisodesByEpisodesId(epiID string, client *http.Client) ([]byte, error) {
	baseURL := "https://api.bgm.tv/v0/episodes/"
	apiURL := fmt.Sprintf("%s%s", baseURL, epiID)

	jsonData, err := getJsonDataFromURL("GET", apiURL, "", client)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
