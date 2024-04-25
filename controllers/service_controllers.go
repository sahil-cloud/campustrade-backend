package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/sahil-cloud/backend/blockchain"
	"github.com/sahil-cloud/backend/constants"
	"github.com/sahil-cloud/backend/helper"
)

//my functions //
//get all products

func GetAllProducts() gin.HandlerFunc {
	// log.Printf("here")
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction("GetAllProductsUser", userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}
		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		ctx.IndentedJSON(http.StatusOK, result)
	}
}

// get all sole products for a particukar seller
func GetAllSoldProducts() gin.HandlerFunc {
	// log.Printf("here")
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_ALL_SOLD_PRODUCTS, userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}

		var result interface{}

		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

// get all unsole products for a particukar seller
func GetAllUnSoldProducts() gin.HandlerFunc {
	// log.Printf("here")
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_ALL_UNSOLD_PRODUCTS, userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}

		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

type Product struct {
	ID                string `json:"ID"`
	ProductName       string `json:"productName"`
	ProductPrice      string `json:"ProductPrice"`
	ProductSubName    string `json:"ProductSubName"`
	ProductDesc       string `json:"ProductDesc"`
	ProductPrimaryImg string `json:"ProductPrimaryImg"`
	ProductSecImg1    string `json:"ProductSecImg1"`
	ProductSecImg2    string `json:"ProductSecImg2"`
	SellerEmail       string `json:"sellerEmail"`
	SellerUPI         string `json:"sellerUPI"`
	Flag              string `json:"flag"`
}

// add product
func AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		var pro Product

		if err := ctx.BindJSON(&pro); err != nil {
			resp["message"] = "Invalid details"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		//generate service ID
		sericeId1, err := uuid.NewV7()
		if err != nil {
			resp["message"] = "Failed to generate service id 1"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		serviceId2, err := uuid.NewV7()
		if err != nil {
			resp["message"] = "Failed to generate service id 2"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		serviceId := "product_" + helper.TodayDateTime() + "_" + sericeId1.String() + serviceId2.String()
		pro.ID = serviceId
		pro.Flag = "false"

		txnProposal, err := blockchain.Contract.NewProposal(constants.CONTRACT_ADD_PRODUCT, client.WithArguments(helper.UnpackStruct(pro)...))
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Proposal Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnEndorsed, err := txnProposal.Endorse()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Endorsement Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnCommitted, err := txnEndorsed.Submit()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Commit Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		log.Printf("NewService Transaction Commit ID %s\n", txnCommitted.TransactionID())
		ctx.IndentedJSON(http.StatusOK, pro)
	}
}

func GetAllReviews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_ALL_REVIEWS, userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}

		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

type Reviews struct {
	ID          string `json:"ID"`
	SellerEmail string `json:"sellerEmail"`
	BuyerEmail  string `json:"buyerEmail"`
	Review      string `json:"Review"`
	Rating      string `json:"Rating"`
}

// add review
func AddReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		var pro Reviews

		if err := ctx.BindJSON(&pro); err != nil {
			resp["message"] = "Invalid details"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		//generate service ID
		sericeId1, err := uuid.NewV7()
		if err != nil {
			resp["message"] = "Failed to generate service id 1"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		serviceId2, err := uuid.NewV7()
		if err != nil {
			resp["message"] = "Failed to generate service id 2"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		serviceId := "review_" + helper.TodayDateTime() + "_" + sericeId1.String() + serviceId2.String()
		pro.ID = serviceId

		txnProposal, err := blockchain.Contract.NewProposal(constants.CONTRACT_ADD_REVIEW, client.WithArguments(helper.UnpackStruct(pro)...))
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Proposal Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnEndorsed, err := txnProposal.Endorse()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Endorsement Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnCommitted, err := txnEndorsed.Submit()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Commit Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		log.Printf("NewService Transaction Commit ID %s\n", txnCommitted.TransactionID())
		ctx.IndentedJSON(http.StatusOK, pro)
	}
}

// 1
func GetTransactionByBuyerId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_BUYER_DELIVERED_PRODUCTS, userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}
		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

// 2
func GetBuyerOrderedProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_BUYER_ORDERED_PRODUCTS, userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}

		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

// 3
func GetTransactionBySellerId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_SELLER_SOLD_PRODUCTS, userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}

		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

// 4
func GetSellerOrderedProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_SELLER_ORDERED_PRODUCTS, userid)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}

		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

type Transactions struct {
	ID            string `json:"ID"`
	Amount        string `json:"amount"`
	BuyerEmail    string `json:"buyerEmail"`
	SellerEmail   string `json:"sellerEmail"`
	ProductID     string `json:"productID"`
	PaymentMode   string `json:"paymentMode"`
	PhoneNumber   string `json:"phoneNumber"`
	ProductName   string `json:"productName"`
	ProductImg    string `json:"productImg"`
	Date          string `json:"date"`
	TransactionID string `json:"transactionID"`
	Delivered     string `json:"delivered"`
	Otp           string `json:"otp"`
}

// 5
// add review
func AddTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		var pro Transactions
		// var tp string

		if err := ctx.BindJSON(&pro); err != nil {
			resp["message"] = "Invalid details"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		//generate service ID
		sericeId1, err := uuid.NewV7()
		if err != nil {
			resp["message"] = "Failed to generate service id 1"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		serviceId2, err := uuid.NewV7()
		if err != nil {
			resp["message"] = "Failed to generate service id 2"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		serviceId := "transaction_" + helper.TodayDateTime() + "_" + sericeId1.String() + serviceId2.String()
		pro.ID = serviceId
		pro.Delivered = "false"
		pro.Date = helper.TodayDateTime()
		pro.Otp, err = helper.GenerateOTP()

		txnProposal, err := blockchain.Contract.NewProposal(constants.CONTRACT_ADD_TRANSACTION, client.WithArguments(helper.UnpackStruct(pro)...))
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Proposal Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnEndorsed, err := txnProposal.Endorse()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Endorsement Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnCommitted, err := txnEndorsed.Submit()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Commit Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		log.Printf("NewService Transaction Commit ID %s\n", txnCommitted.TransactionID())
		ctx.IndentedJSON(http.StatusOK, pro)
	}
}

// 4
func VerifyTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		otp := ctx.Query("otp")
		userid := ctx.Query("id")

		txnProposal, err := blockchain.Contract.NewProposal(constants.CONTRACT_VERIFY_TRANSACTION, client.WithArguments(userid, otp))
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Proposal Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnEndorsed, err := txnProposal.Endorse()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Endorsement Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnCommitted, err := txnEndorsed.Submit()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Commit Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		log.Printf("NewService Transaction Commit ID %s\n", txnCommitted.TransactionID())
		ctx.IndentedJSON(http.StatusOK, "true")
	}
}

func GetAllTransactions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		// userid := ctx.Query("email")

		evaluatedResult, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_GET_ALL_TRANSACTIONS)
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Evaluation Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		if evaluatedResult == nil {
			ctx.IndentedJSON(http.StatusOK, []string{})
			return
		}

		var result interface{}
		err = json.Unmarshal(evaluatedResult, &result)
		if err != nil {
			resp["message"] = "Internal Error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		ctx.IndentedJSON(http.StatusOK, result)
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := make(map[string]string)
		userid := ctx.Query("id")

		txnProposal, err := blockchain.Contract.NewProposal("DeleteProduct", client.WithArguments(userid))
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Proposal Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnEndorsed, err := txnProposal.Endorse()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Endorsement Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		txnCommitted, err := txnEndorsed.Submit()
		if err != nil {
			resp["message"] = "Blockchain Trnasaction Commit Failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		log.Printf("NewService Transaction Commit ID %s\n", txnCommitted.TransactionID())
		ctx.IndentedJSON(http.StatusOK, "true")
	}
}
