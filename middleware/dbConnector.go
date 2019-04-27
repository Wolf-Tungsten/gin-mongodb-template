package middleware

import(
	"github.com/gin-gonic/gin"
	"seu-job-mina-service/db"
)

func MongoConnect(context *gin.Context){
	context.Set("db", db.Client)
	context.Next()
}

