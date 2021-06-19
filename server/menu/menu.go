package menu

import (
	_ "embed"
	"encoding/json"

	"github.com/shoppehub/fastapi/base"
)

//go:embed menu.json
var defaultMenuStr []byte

var SystemApplicationKey = "system"
var systemMenus []Menu

func init() {
	json.Unmarshal(defaultMenuStr, &systemMenus)
}

type Menu struct {
	base.BaseId    `bson,inline`
	ApplicationKey string `bson:"applicationKey,omitempty" json:"applicationKey,omitempty"`
	Key            string `bson:"key,omitempty" json:"key,omitempty"`
	Title          string `bson:"title,omitempty" json:"title,omitempty"`
	Href           string `bson:"href,omitempty" json:"href,omitempty"`
	Icon           string `bson:"icon,omitempty" json:"icon,omitempty"`
	Children       []Menu `bson:"children,omitempty" json:"children,omitempty"`
}

// 获取menu
func GetAppMenus(applicationKey string) []Menu {

	if applicationKey == SystemApplicationKey {
		return systemMenus
	}

	return nil
}
