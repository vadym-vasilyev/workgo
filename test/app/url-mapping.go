package app
func mapUrls(){
    router.GET("/ping", ping.Ping )

    router.GET("/users/:user_id", users.GetUser)
    router.POST("/users", users.CreateUser)
}