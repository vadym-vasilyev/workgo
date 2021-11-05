package ping_controller
func Ping(c *gin.Context){
    c.String(http.StatusOK, format:"pong")
}