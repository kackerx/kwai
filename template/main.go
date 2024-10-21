package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

var tpl = `
type {{ .Name }}HttpServer struct {
	server {{.Name}}Server
	router gin.IRouter
}

func Register{{ .Name }}HttpServer(server {{ .Name }}Server, router gin.IRouter) {
	u := &{{ .Name }}Handler{server: server, router: router}
	u.RegisterService()
}

{{ range .Methods }}
func (u *{{ $.Name }}HttpServer) {{ .Name }}(c *gin.Context) {
	var in {{ .Request }}
	if err := c.BindJSON(&in); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	res, err := u.server.{{ .Name }}(c, &in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
{{ end }}

func (u *{{ .Name }}HttpServer) RegisterService() {
{{ range .Methods }}
 	u.router.Handle("{{ .Method }}", "{{ .Path }}", u.{{ .Name }})
{{ end }}
}
`

type serviceDesc struct {
	Name    string
	Methods []method
}

type method struct {
	Name    string
	Request string
	Reply   string

	// http
	Path   string
	Method string
	Body   string
}

func main() {
	buf := new(bytes.Buffer)

	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}

	s := serviceDesc{
		Name: "User",
		Methods: []method{
			{
				Name:    "Hello",
				Request: "HelloReq",
				Reply:   "HelloResp",
				Path:    "/hello",
				Method:  "post",
				Body:    "*",
			},
			{
				Name:    "Ping",
				Request: "PingReq",
				Reply:   "PingResp",
				Path:    "/ping",
				Method:  "get",
				Body:    "*",
			},
		},
	}

	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
