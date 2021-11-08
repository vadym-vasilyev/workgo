package app
func mapUrls(){
    router.GET("/ping", ping.Ping )

    router.GET("/Tweets/:Tweet_id", Tweets.GetAllTweets)
    router.POST("/Tweets", Tweets.CreateTweets)
}