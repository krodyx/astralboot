// Serve web interface
//go:generate go-bindata -debug asset/...
package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"strconv"
	"strings"
)

var MenuOrder = []string{"machines", "containers", "configuration", "cluster"}

type pageData struct {
	Menu []string
}

func SetupPage() (p *pageData) {
	p = &pageData{
		Menu: MenuOrder,
	}
	return p
}

// short data interface for templating
type td map[string]interface{}

func (wh *WebHandler) current(section string) (m td) {
	m = td{
		"Page":    wh.page,
		"Section": section,
		"Leases":  wh.store.ListActive(),
	}
	// pre rendered content for snippet
	return m
}

func (wh *WebHandler) content(name string, data interface{}) (st template.HTML) {
	buffer := new(bytes.Buffer)
	err := wh.uiTemplates.ExecuteTemplate(buffer, name, data)
	if err != nil {
		logger.Critical("Content Fail", err)
	}
	return template.HTML(buffer.String())
}

// WebInterface : provides a web interface for astralboot functions and monitoring
func (wh *WebHandler) WebInterface() {
	wh.page = SetupPage()
	// Bind the Index
	wh.router.GET("/", wh.Index)
	wh.router.GET("/static/*path", wh.Static)
	// Confiugre the Subsections
	wh.router.GET("/machines", wh.machines)
	wh.router.GET("/configuration", wh.configuration)
	wh.router.GET("/containers", wh.containers)
	wh.router.GET("/cluster", wh.cluster)

	// Load the templates
	// get the asset dir
	pages, err := AssetDir("asset/pages")
	if err != nil {
		logger.Error("Loading pages %s", err)
		return
	}
	templ := template.New("")
	for _, j := range pages {
		logger.Critical("%s", j)
		data, _ := Asset("asset/pages/" + j)
		_, err = templ.New(j).Parse(string(data))
		if err != nil {
			logger.Error("Loading pages %s", err)
			return
		}
	}
	fmt.Println(tmpl)
	wh.uiTemplates = templ
}

func (wh *WebHandler) Index(c *gin.Context) {
	logger.Debug("Index HIT")
	data := wh.current("index")
	wh.uiTemplates.ExecuteTemplate(c.Writer, "index.html", data)
}

func (wh *WebHandler) machines(c *gin.Context) {
	data := wh.current("machines")
	data["Content"] = wh.content("machines.html", wh.store.ListActive())
	wh.uiTemplates.ExecuteTemplate(c.Writer, "index.html", data)
}

func (wh *WebHandler) configuration(c *gin.Context) {
	data := wh.current("configuration")
	data["Content"] = wh.content("configuration.html", wh.store.ListActive())
	wh.uiTemplates.ExecuteTemplate(c.Writer, "index.html", data)
}

func (wh *WebHandler) containers(c *gin.Context) {
	data := wh.current("containers")
	data["Content"] = wh.content("containers.html", wh.store.ListActive())
	wh.uiTemplates.ExecuteTemplate(c.Writer, "index.html", data)
}

func (wh *WebHandler) cluster(c *gin.Context) {
	data := wh.current("cluster")
	data["Content"] = wh.content("cluster.html", wh.store.ListActive())
	wh.uiTemplates.ExecuteTemplate(c.Writer, "index.html", data)
}

func (wh *WebHandler) Static(c *gin.Context) {
	path := c.Params.ByName("path")
	logger.Debug(path)
	data, err := Asset("asset" + path)
	if err != nil {
		logger.Error("Asset Error ", err)
		c.AbortWithStatus(404)
	}
	if strings.HasSuffix(path, ".css") {
		c.Writer.Header().Set("Content-Type", "text/css")
	}
	if strings.HasSuffix(path, ".js") {
		c.Writer.Header().Set("Content-Type", "text/javascript")
	}
	size := int64(len(data))
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	io.Copy(c.Writer, bytes.NewReader(data))
}
