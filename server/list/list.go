package list

import (
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/shoppehub/commons"
	"github.com/shoppehub/conf"
	"github.com/shoppehub/fastapi/base"
	"gopkg.in/yaml.v2"
)

// 表格模型
type List struct {
	base.BaseId `bson,inline`
	Key         string     `bson:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Title       string     `bson:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Items       []ListItem `bson:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	// 绑定主数据源的key
	DataSouceKey string           `bson:"dataSouceKey,omitempty" json:"dataSouceKey,omitempty" yaml:"dataSouceKey,omitempty" `
	ItemActions  []ListItemAction `bson:"itemActions,omitempty" json:"itemActions,omitempty" yaml:"itemActions,omitempty"`
	Values       *commons.PagingResponse
}

type ListFilter struct {
	// 列的key
	Key string `bson:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	// 是否隐藏该列，默认是false，隐藏列为了满足某种特殊显示需求
	Hidden bool `bson:"hidden,omitempty" json:"hidden,omitempty" yaml:"hidden,omitempty"`
}

//表格的每一列
type ListItem struct {
	// 列的key
	Key   string `bson:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Title string `bson:"title,omitempty"  json:"title,omitempty" yaml:"title,omitempty"`
	// 是否隐藏该列，默认是false，隐藏列为了满足某种特殊显示需求
	Hidden bool `bson:"hidden,omitempty" json:"hidden,omitempty" yaml:"hidden,omitempty"`
	// 是否参与排序
	Sort bool `bson:"sort,omitempty" json:"sort,omitempty" yaml:"sort,omitempty"`
	// 显示类型，默认是 text，按文本显示
	Type string `bson:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	// 显示的值，默认不用填写，需要填写的时候就是一个表达式，比如 {{.Items[0].Name}}
	Value    string `bson:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Template string `bson:"template,omitempty" json:"template,omitempty" yaml:"template,omitempty"`
}

type ListItemAction struct {
	Key   string `bson:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Title string `bson:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
}

// 初始化模板
func InitTemplate(vars *jet.VarMap, root string) {
	vars.SetFunc("getListConfig", func(a jet.Arguments) reflect.Value {
		key := "config"
		if a.NumOfArguments() != 0 {
			key = a.Get(0).String()
		}

		key += ".yaml"

		configPath := strings.Join([]string{"web", root, key}, "/")

		if conf.Exists(configPath) {
			strb, cerr := ioutil.ReadFile(configPath)
			if cerr != nil {
				return reflect.ValueOf(List{})
			}
			var result List
			yamlerr := yaml.Unmarshal(strb, &result)
			if yamlerr != nil {
				return reflect.ValueOf(List{})
			}
			return reflect.ValueOf(&result)
		} else {
			return reflect.ValueOf(List{})
		}

	})
}
