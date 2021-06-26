package site

import "github.com/shoppehub/fastapi/base"

// 站点信息
type Site struct {
	base.BaseId    `bson,inline`
	ApplicationKey string `bson:"applicationKey,omitempty" json:"applicationKey,omitempty"`
	Key            string `bson:"key,omitempty" json:"key,omitempty"`
	Title          string `bson:"title,omitempty" json:"title,omitempty"`
	Desc           string `desc:"url,omitempty" json:"desc,omitempty"`
	Keyword        string `keyword:"url,omitempty" json:"keyword,omitempty"`
	Url            string `bson:"url,omitempty" json:"url,omitempty"`
	LogoUrl        string `bson:"logoUrl,omitempty" json:"logoUrl,omitempty"`
	Menus          []Menu `bson:"menus,omitempty" json:"menus,omitempty"`
}
