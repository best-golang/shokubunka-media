package infrastructure

import (
	"api/domain"
	"api/interfaces/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qor/admin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	mux := http.NewServeMux()

	sqlHandler := NewSqlHandler()
	Migrate(sqlHandler.Conn)

	// Admin config & routing
	Admin := admin.New(&admin.AdminConfig{
		DB:       sqlHandler.Conn,
		SiteName: "Hungry Researchers",
	})
	user := Admin.AddResource(&domain.User{})
	defineUserMetaInfo(user)
	Admin.MountTo("/admin", mux)
	router.Any("/admin/*resources", gin.WrapH(mux))

	// controller
	userController := controllers.NewUserController(sqlHandler)

	// Grouping route
	api := router.Group("/api")
	v1 := api.Group("/v1")

	// Define routes
	v1.POST("/users", func(c *gin.Context) {
		b := binding.Default(c.Request.Method, c.ContentType())
		userController.Create(c, b)
	})
	Router = router
}
