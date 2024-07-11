package ext

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"time"
)

func GetRes() string {
	//获取代理
	resp, err := HttpRequest.SetTimeout(1*time.Minute).SetHeaders(map[string]string{
		"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Connection": "keep-alive",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}).SetCookies(map[string]string{
		"Name":  "bid",
		"Value": "Sr-dWSYEhXU",
	}).JSON().Post("https://read.douban.com/j/search_v2", `{
		"page": 1,
		"q": "斗2：打遍天下无敌手",
		"query": "\n    query getFilterWorksList($works_ids: [ID!]) {\n      worksList(worksIds: $works_ids) {\n        \n    \n    title\n    cover(useSmall: false)\n    url\n    isBundle\n    coverLabel(preferVip: true)\n  \n    \n  url\n  title\n\n    \n  author {\n    name\n    url\n  }\n  origAuthor {\n    name\n    url\n  }\n  translator {\n    name\n    url\n  }\n\n    \n  abstract\n  authorHighlight\n  editorHighlight\n\n    \n    isOrigin\n    kinds {\n      \n    name @skip(if: true)\n    shortName @include(if: true)\n    id\n  \n    }\n    ... on WorksBase @include(if: true) {\n      wordCount\n      wordCountUnit\n    }\n    ... on WorksBase @include(if: false) {\n      inLibraryCount\n    }\n    ... on WorksBase @include(if: false) {\n      \n    isEssay\n    \n    ... on EssayWorks {\n      favorCount\n    }\n  \n    \n    \n    averageRating\n    ratingCount\n    url\n    isColumn\n    isFinished\n  \n  \n  \n    }\n    ... on EbookWorks @include(if: false) {\n      \n    ... on EbookWorks {\n      book {\n        url\n        averageRating\n        ratingCount\n      }\n    }\n  \n    }\n    ... on WorksBase @include(if: false) {\n      isColumn\n      isEssay\n      onSaleTime\n      ... on ColumnWorks {\n        updateTime\n      }\n    }\n    ... on WorksBase @include(if: true) {\n      isColumn\n      ... on ColumnWorks {\n        isFinished\n      }\n    }\n    ... on EssayWorks {\n      essayActivityData {\n        \n    title\n    uri\n    tag {\n      name\n      color\n      background\n      icon2x\n      icon3x\n      iconSize {\n        height\n      }\n      iconPosition {\n        x y\n      }\n    }\n  \n      }\n    }\n    highlightTags {\n      name\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        tags {\n          id\n          name\n          url\n        }\n      }\n    }\n  \n    \n  ... on WorksBase {\n    copyrightInfo {\n      newlyAdapted\n      newlyPublished\n      adaptedName\n      publishedName\n    }\n  }\n\n    isInLibrary\n    ... on WorksBase @include(if: false) {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on EbookWorks {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on WorksBase @include(if: true) {\n      ... on EbookWorks {\n        id\n        isPurchased\n        isInWishlist\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        fandoms {\n          title\n          url\n        }\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        kudoCount\n      }\n    }\n  \n        id\n        isOrigin\n        isEssay\n      }\n    }\n  ",
	
	}`) //设置消息头，cooki和发送post（json）请求
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Close()

	body, err := resp.Body()
	if err != nil {
		return "body获取错误"
	}
	return string(body)

}
