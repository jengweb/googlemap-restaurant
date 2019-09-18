package routes

import (
	"googlemapsearch/app/controllers"

	"github.com/gin-gonic/gin"
)

var GoogleMapSearchController controllers.GoogleMapSearchInterface = &controllers.GoogleMapSearchController{}

func PostGoogleMapSearch(context *gin.Context) {
	GoogleMapSearchController.PostGoogleMapSearch(context, context.Request)
}
