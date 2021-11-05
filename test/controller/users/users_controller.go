package users
import (
"net/http"
)
var(
counter int
)
func CreateUser(c *Context){
    c.String(http.StatusNotImplemented, "implement me!")
}
func GetUser(c *Context){
    c.String(http.StatusNotImplemented, "implement me!")
}
