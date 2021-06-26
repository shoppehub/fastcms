package site

import (
	_ "embed"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/shoppehub/fastapi/base"
	"github.com/sirupsen/logrus"
)

var localMenus = make(map[string][]Menu)

type Menu struct {
	base.BaseId    `bson,inline`
	ApplicationKey string `bson:"applicationKey,omitempty" json:"applicationKey,omitempty"`
	Key            string `bson:"key,omitempty" json:"key,omitempty"`
	Title          string `bson:"title,omitempty" json:"title,omitempty"`
	Href           string `bson:"href,omitempty" json:"href,omitempty"`
	Icon           string `bson:"icon,omitempty" json:"icon,omitempty"`
	Children       []Menu `bson:"children,omitempty" json:"children,omitempty"`
}

func init() {

	root := "./web/menus"

	files, _ := ioutil.ReadDir(root)

	for _, v := range files {

		if strings.HasSuffix(v.Name(), ".json") {

			strbs, err := ioutil.ReadFile(strings.Join([]string{root, v.Name()}, "/"))
			if err != nil {
				logrus.Panicln(err)
				return
			}
			var ms []Menu
			err = json.Unmarshal(strbs, &ms)
			if err != nil {
				logrus.Panicln(err)
				return
			}
			localMenus[strings.TrimSuffix(v.Name(), ".json")] = ms
		}

	}

}

// 初始化模板
func InitTemplate(views *jet.Set) {
	views.AddGlobalFunc("getMenu", func(a jet.Arguments) reflect.Value {
		applicationKey := "system"
		if a.NumOfArguments() != 0 {
			applicationKey = a.Get(0).String()
		}

		if menus, ok := localMenus[applicationKey]; ok {
			return reflect.ValueOf(menus)
		}
		return reflect.ValueOf([]Menu{})
	})
}

//vars.Set("menus", menu.GetAppMenus(menu.SystemApplicationKey))
