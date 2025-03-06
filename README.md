# lite_bangumi_api

包装了部分BangumiAPI到go里面。采用http通信形式获取数据。

目前仅支持条目搜索查看、角色搜索查看、人物搜索查看、章节搜索查看、用户搜索查看。

详细API见下：

## 全局变量设置

liteBangumiAPI有两个全局变量，必须在调用前设置。填写全局变量如下：

``` go
lite_bangumi_api.Token = "YOUR_ACCESS_TOKEN"
lite_bangumi_api.UserAgent = "YOUR User-Agent"
```

说明：

1.bangumi使用OAuth 2.0，token格式已经在内部写好了，这里token直接填写你从bangumi获取的token就行。**【不需要写入Bearer】**。

2.UserAgent的形式，请参考https://github.com/bangumi/api/blob/master/docs-raw/user%20agent.md

## 支持的API：

```
/calendar
/v0/search/subjects
/v0/subjects/{subject_id}

/v0/episodes
/v0/episodes/{episode_id}

/v0/search/characters
/v0/characters/{character_id}
/v0/characters/{character_id}/collect（POST）
/v0/characters/{character_id}/collect（DELETE）

/v0/search/persons
/v0/persons/{person_id}
/v0/persons/{person_id}/collect（POST）
/v0/persons/{person_id}/collect（DELETE）

/v0/users/{username}
/v0/me

/v0/users/{username}/collections
/v0/users/{username}/collections/{subject_id}
/v0/users/-/collections/{subject_id}（POST）
/v0/users/-/collections/{subject_id}（PATCH）
/v0/users/-/collections/{subject_id}/episodes（GET）
/v0/users/-/collections/{subject_id}/episodes（PATCH）
/v0/users/-/collections/-/episodes/{episode_id}（GET）
/v0/users/-/collections/-/episodes/{episode_id}（PUT）
/v0/users/{username}/collections/-/characters
/v0/users/{username}/collections/-/characters/{character_id}
/v0/users/{username}/collections/-/persons
/v0/users/{username}/collections/-/persons/{person_id}

/v0/revisions/persons
/v0/revisions/persons/{revision_id}
/v0/revisions/persons/{revision_id}
/v0/revisions/characters/{revision_id}
/v0/revisions/subjects
/v0/revisions/subjects/{revision_id}
/v0/revisions/episodes
/v0/revisions/episodes/{revision_id}

/v0/indices
/v0/indices/{index_id}（GET）
/v0/indices/{index_id}（PUT）
/v0/indices/{index_id}/subjects（GET）
/v0/indices/{index_id}/subjects（POST）
/v0/indices/{index_id}/subjects/{subject_id}（PUT）
/v0/indices/{index_id}/subjects/{subject_id}（DELETE）
/v0/indices/{index_id}/collect（POST）
/v0/indices/{index_id}/collect（DELETE）

/search/subject/{keywords}
```

