package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
	"github.com/shoppehub/fastcms/server/menu"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./web"),
	jet.InDevelopmentMode(), // remove in production
)

// 渲染模板
func RenderTemplate(c *gin.Context) {
	module := c.Params.ByName("module")
	if module == "" {
		module = "index"
	}
	page := c.Params.ByName("page")
	if page == "" {
		page = "index"
	}

	path := strings.Join([]string{module, page}, "/")
	view, err := views.GetTemplate("pages/" + path + ".jet")
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	vars := make(jet.VarMap)
	vars.Set("menus", menu.GetAppMenus(menu.SystemApplicationKey))
	// vars.Set("showingAllDone", true)
	view.Execute(c.Writer, vars, nil)
	c.Status(200)
}

//将request转发给 http://127.0.0.1:4001
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	trueServer := "http://127.0.0.1:" + fmt.Sprint(Port+1)
	url, err := url.Parse(trueServer)
	if err != nil {
		log.Println(err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}
