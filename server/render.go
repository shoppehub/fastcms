package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
	"github.com/shoppehub/conf"
	"github.com/shoppehub/fastapi/crud"
	"github.com/shoppehub/fastapi/engine/template"
	"github.com/shoppehub/fastcms/server/list"
	"github.com/shoppehub/fastcms/server/menu"
	"github.com/sirupsen/logrus"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./web"),
	jet.InDevelopmentMode(), // remove in production
)

var resource *crud.Resource

func initResource() {
	if resource != nil {
		return
	}
	var err error
	resource, err = crud.NewDB(conf.GetString("mongodb.url"), conf.GetString("mongodb.dbname"))
	if err != nil {
		logrus.Error(err)
	}
}

func init() {

	menu.InitTemplate(views)

}

// 渲染模板
func RenderTemplate(c *gin.Context) {
	initResource()
	module := c.Params.ByName("module")
	if module == "" {
		module = "index"
	}
	page := c.Params.ByName("page")
	if page == "index" {
		page = ""
	}

	templ := c.Params.ByName("templ")
	if templ == "" {
		templ = "index.jet"
	} else {
		templ += ".jet"
	}

	root := strings.Join([]string{"pages", module, page}, "/")
	templatePath := strings.Join([]string{root, templ}, "/")
	view, err := views.GetTemplate(templatePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	vars := template.NewVars(resource, nil)

	list.InitTemplate(vars, root)

	// vars.Set("menus", menu.GetAppMenus(menu.SystemApplicationKey))
	vars.Set("namespace", strings.Join([]string{module, page}, "_"))

	vars.Set("path", c.Request.URL.Path)
	var curPage int64
	if c.Query("curPage") == "" {
		curPage = 1
	} else {
		curPage, _ = strconv.ParseInt(c.Query("curPage"), 10, 64)
	}

	vars.Set("curPage", curPage)

	vars.SetFunc("numArray", func(a jet.Arguments) reflect.Value {

		var total int
		k := a.Get(0).Kind()
		switch k {
		case reflect.Float64:
			total = int(a.Get(0).Float())
		case reflect.Int64:
			total = int(a.Get(0).Int())
		}

		nums := make([]int64, total)
		for i := 0; i < total; i++ {
			nums[i] = int64(i + 1)
		}
		return reflect.ValueOf(nums)
	})

	vars.SetFunc("getUrlPath", func(a jet.Arguments) reflect.Value {

		if !a.Get(0).IsValid() {
			return reflect.ValueOf("")
		}

		u, _ := url.Parse(a.Get(0).Interface().(string))
		return reflect.ValueOf(u.Path)
	})

	// vars.Set("showingAllDone", true)
	rerr := view.Execute(c.Writer, *vars, nil)
	if rerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": rerr.Error(),
		})
		return
	}
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
