package controllers

import (
	"context"
	"net/http"
	"path"
	"text/template"
	// "github.com/astaxie/beego"
)

// type MainController struct {
// 	beego.Controller
// }

// func (this *MainController) Get() {
// 	this.Data["Username"] = "astaxie"
// 	this.Data["Email"] = "astaxie@gmail.com"
// 	this.TplNames = "index.tpl"
// }

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/de/13.3.html

type Controller struct {
	Ctx       *context.Context
	Tpl       *template.Template
	Data      map[interface{}]interface{}
	ChildName string
	TplNames  string
	Layout    []string
	TplExt    string
}

type ControllerInterface interface {
	Init(ct *context.Context, cn string) //Initialize the context and subclass name
	Prepare()                            //some processing before execution begins
	Get()                                //method = GET processing
	Post()                               //method = POST processing
	Delete()                             //method = DELETE processing
	Put()                                //method = PUT handling
	Head()                               //method = HEAD processing
	Patch()                              //method = PATCH treatment
	Options()                            //method = OPTIONS processing
	Finish()                             //executed after completion of treatment
	Render() error                       //method executed after the corresponding method to render the page
}

func (c *Controller) Init(ctx *context.Context, cn string) {
	c.Data = make(map[interface{}]interface{})
	c.Layout = make([]string, 0)
	c.TplNames = ""
	c.ChildName = cn
	c.Ctx = ctx
	c.TplExt = "tpl"
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}

func (c *Controller) Get() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Head() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Patch() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Options() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Render() error {
	if len(c.Layout) > 0 {
		var filenames []string
		for _, file := range c.Layout {
			filenames = append(filenames, path.Join(ViewsPath, file))
		}
		t, err := template.ParseFiles(filenames...)
		if err != nil {
			Trace("template ParseFiles err:", err)
		}
		err = t.ExecuteTemplate(c.Ct.ResponseWriter, c.TplNames, c.Data)
		if err != nil {
			Trace("template Execute err:", err)
		}
	} else {
		if c.TplNames == "" {
			c.TplNames = c.ChildName + "/" + c.Ct.Request.Method + "." + c.TplExt
		}
		t, err := template.ParseFiles(path.Join(ViewsPath, c.TplNames))
		if err != nil {
			Trace("template ParseFiles err:", err)
		}
		err = t.Execute(c.Ct.ResponseWriter, c.Data)
		if err != nil {
			Trace("template Execute err:", err)
		}
	}
	return nil
}

func (c *Controller) Redirect(url string, code int) {
	c.Ct.Redirect(code, url)
}
