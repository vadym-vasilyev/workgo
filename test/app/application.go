package app
var
( router = gin.Default()
)
func StartApplication(){
    mapUrls()
    router.Run(":8082")

}