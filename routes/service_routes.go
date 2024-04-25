package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahil-cloud/backend/controllers"
)

func ServiceRoutes(apiRoutes *gin.RouterGroup) {

	//myt routes//

	apiRoutes.GET("/allproducts", controllers.GetAllProducts())
	apiRoutes.GET("/allsoldproducts", controllers.GetAllSoldProducts())
	apiRoutes.GET("/allunsoldproducts", controllers.GetAllUnSoldProducts())
	apiRoutes.POST("/addproduct", controllers.AddProduct())

	// apiRoutes.GET("/allratings", controllers.GetAllRatings())
	// apiRoutes.GET("/addrating", controllers.AddRating())
	apiRoutes.POST("/addreview", controllers.AddReview())
	apiRoutes.GET("/allreviews", controllers.GetAllReviews())

	apiRoutes.POST("/addtransaction", controllers.AddTransaction())
	apiRoutes.GET("/buyerdelivered", controllers.GetTransactionByBuyerId())
	apiRoutes.GET("/buyernotdelivered", controllers.GetBuyerOrderedProducts())
	apiRoutes.GET("/sellerdelivered", controllers.GetTransactionBySellerId())
	apiRoutes.GET("/sellernotdelivered", controllers.GetSellerOrderedProducts())
	apiRoutes.POST("/verifytransaction", controllers.VerifyTransaction())
	apiRoutes.GET("/alltransactions", controllers.VerifyTransaction())

	apiRoutes.POST("/deleteproduct", controllers.DeleteProduct())
}
