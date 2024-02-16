package functions

import (
	"bitbucket.org/nulljet/go_ignt_func_deps"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Function interface {
	Handle(params go_ignt_func_deps.ParamsInterface, r *http.Request) go_ignt_func_deps.ResponseInterface
	GetParams() []string
	GetName() string
	GetDescription() string
	GetHelp(params []string) string
	IsWeb() bool
	GetUrl() string
}

type Naming struct {
	Name        string
	Description string
	Params      []string
	Orig        Function
}

func (n *Naming) GetName() string {
	return n.Name
}

func (n *Naming) GetDescription() string {
	return n.Description
}

func (n *Naming) GetParams() []string {
	return n.Orig.GetParams()
}

func (n *Naming) GetHelp(params []string) string {
	paramsStr := " "
	for i, param := range params {
		paramsStr += " " + param + "=<value" + strconv.Itoa(i) + ">"
	}
	return "Usage: ignt run " + n.Name + "" + paramsStr
}

type NotWeb struct {
}
type IsWebOnly struct {
}

type WebOnly interface {
	IsWebOnly() bool
}

func (n *NotWeb) IsWeb() bool {
	return false
}

func (o *IsWebOnly) IsWebOnly() bool {
	return true
}
func (n *NotWeb) GetUrl() string {
	return "false"
}

type Accessing struct {
	Public bool
}

func (n *Accessing) IsPublic() bool {
	return n.Public
}

type ConsoleResponse struct {
	Data string
}

func (j ConsoleResponse) RenderConsoleResponse() {
	fmt.Print(j.Data)
}

func (j ConsoleResponse) RenderWebResponse() ([]byte, int) {
	return []byte(j.Data), 200
}

type JsonResponse struct {
	Data   string
	Status int
}

func (j JsonResponse) RenderConsoleResponse() {
	fmt.Print(PrettyString(j.Data))
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func (j JsonResponse) RenderWebResponse() ([]byte, int) {
	return []byte(j.Data), j.Status
}
