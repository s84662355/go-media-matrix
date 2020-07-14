package validation

import (
    "github.com/kataras/iris"
    "gopkg.in/go-playground/validator.v9"
    "strings"
)

var Validate = validator.New()

type paginationRequest struct {
    Page    int `form:"page" json:"page"`
    PerPage int `form:"per_page" json:"per_page"`
}

func getRequestData(c iris.Context, data interface{}) error {
    if c.Method() == "GET" {
        _ = c.ReadForm(data)
    } else {
        if strings.Contains(c.GetContentTypeRequested(), "application/json") {
            _ = c.ReadJSON(data)
        } else {
            _ = c.ReadForm(data)
        }
    }
    return Validate.Struct(data)
}
