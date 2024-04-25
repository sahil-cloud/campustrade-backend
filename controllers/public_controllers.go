package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/sahil-cloud/backend/blockchain"
	"github.com/sahil-cloud/backend/constants"
	"github.com/sahil-cloud/backend/helper"
)

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobileNumber"`
	ImgUrl       string `json:"img"`
	CryptoKey    string `json:"cryptoKey"`
	JWTToken     string `json:"jwtToken"`
	Email        string `json:"email"`
}

type IMSResp struct {
	RollNo       string `json:"rollno"`
	Name         string `json:"name"`
	Department   string `json:"department"`
	Email        string `json:"email"`
	CryptoKey    string `json:"cryptokey"`
	MobileNumber string `json:"mobilenumber"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type tokenStruct struct {
	JWTToken string `json:"jwttoken"`
}

// func getInfoFromIdentity(cred Credentials) (User, error) {
// 	var user User
// 	if cred.Username == "cs22m014" && cred.Password == "12345678" {
// 		// buyer
// 		user.ID = "cs22m014"
// 		user.Name = "Ankit Gupta"
// 		user.ImgUrl = "https://images.unsplash.com/photo-1599566150163-29194dcaad36?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
// 		user.MobileNumber = "9090909090"
// 		user.CryptoKey = "diehurwndiweuhirhqiwruoiyriuwfnfheuihiuwfbwoqertqwyr"
// 		user.Email = "cs22m014@gmail.com"
// 		user.Status = "true"
// 		return user, nil
// 	}
// 	if cred.Username == "cs22m075" && cred.Password == "12345678" {
// 		// seller
// 		user.ID = "cs22m075"
// 		user.Name = "Sahil Jsuja"
// 		user.ImgUrl = "https://images.unsplash.com/photo-1557862921-37829c790f19?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTR8fHBlcnNvbnxlbnwwfHwwfHx8MA%3D%3D"
// 		user.MobileNumber = "9191919191"
// 		user.CryptoKey = "tyuiwieyqwriyiuoeuoqueoqueoiqonckdhiuhi"
// 		user.Email = "cs22m075@gmail.com"
// 		user.Status = "true"
// 		return user, nil
// 	}
// 	if cred.Username == "cs22m002" && cred.Password == "12345678" {
// 		// seller
// 		user.ID = "cs22m002"
// 		user.Name = "AB"
// 		user.ImgUrl = "https://images.unsplash.com/photo-1557862921-37829c790f19?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTR8fHBlcnNvbnxlbnwwfHwwfHx8MA%3D%3D"
// 		user.MobileNumber = "9191919191"
// 		user.CryptoKey = "tyuiwjuuieyqwijbhbhbheuoqoquijijhheoiqonckdhiuhi"
// 		user.Email = "cs22m002@gmail.com"
// 		user.Status = "true"
// 		return user, nil
// 	}
// 	return user, fmt.Errorf("signup failed")
// }

// func SignUp() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// send the username and password to identity management and get cryptographic identity for that user from them
// 		// if failed to get token then signup failed

// 		var credentials Credentials
// 		resp := make(map[string]string)
// 		if err := ctx.BindJSON(&credentials); err != nil {
// 			resp["message"] = "Username or password not found"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}

// 		infoFromIdentity, err := getInfoFromIdentity(credentials)

// 		if err != nil {
// 			resp["message"] = "Username or password is incorrect"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}

// 		var user User
// 		user.ID = infoFromIdentity.ID

// 		user.Name = infoFromIdentity.Name
// 		user.MobileNumber = infoFromIdentity.MobileNumber
// 		user.CryptoKey = infoFromIdentity.CryptoKey
// 		user.ImgUrl = infoFromIdentity.ImgUrl
// 		user.Email = infoFromIdentity.Email
// 		user.Status = infoFromIdentity.Status

// 		// check crypto key exists in blockchain if exist means user already signed up
// 		evaluatedRespone, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_IS_USER_KEY_EXISTS, user.CryptoKey)
// 		if err != nil {
// 			resp["message"] = "Blockchain Transaction evaluation failed "
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		isKeyExists, err := strconv.ParseBool(string(evaluatedRespone))
// 		if err != nil {
// 			resp["message"] = "Internal Error boolean parsing failed"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		if isKeyExists {
// 			resp["message"] = "User already exist"
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}

// 		// key not found, set crypto key as a userkey in our blockchain db
// 		txnProposal, err := blockchain.Contract.NewProposal(constants.CONTRACT_CREATE_USER_KEY, client.WithArguments(user.CryptoKey, user.MobileNumber))
// 		if err != nil {
// 			resp["message"] = "Blockain Transaction proposal failed"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		txnEndorsed, err := txnProposal.Endorse()
// 		if err != nil {
// 			resp["message"] = "Blockain Transaction endorsement failed"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		txnCommitted, err := txnEndorsed.Submit()
// 		if err != nil {
// 			resp["message"] = "Blockain Transaction commit failed"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}

// 		// log here
// 		fmt.Printf("Signup Transaction ID : %s Response: %s\n", txnCommitted.TransactionID(), txnEndorsed.Result())

// 		// then generate a jwt token and send the details to UI
// 		token, err := helper.GenerateToken(user.Name, user.ID, user.CryptoKey)
// 		if err != nil {
// 			resp["message"] = "SignUp Success. Failed to generate token. Please login"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		user.JWTToken = token
// 		ctx.IndentedJSON(http.StatusOK, user)
// 	}
// }

// func Login() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// send the username and password to identity management and get cryptographic identity for that user from them
// 		// if failed to get token then login failed
// 		var credentials Credentials
// 		resp := make(map[string]string)
// 		if err := ctx.BindJSON(&credentials); err != nil {
// 			resp["message"] = "Username or password not found"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}

// 		infoFromIdentity, err := getInfoFromIdentity(credentials)

// 		if err != nil {
// 			resp["message"] = "Username or password is incorrect"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}

// 		var user User
// 		user.ID = infoFromIdentity.ID
// 		user.Name = infoFromIdentity.Name
// 		user.MobileNumber = infoFromIdentity.MobileNumber
// 		user.CryptoKey = infoFromIdentity.CryptoKey
// 		user.ImgUrl = infoFromIdentity.ImgUrl
// 		user.Email = infoFromIdentity.Email
// 		user.Status = infoFromIdentity.Status

// 		// check if user crypt key exist in our blockchain db if yes then continue to generate token
// 		// if not then login failed
// 		evaluatedRespone, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_IS_USER_KEY_EXISTS, user.CryptoKey)
// 		if err != nil {
// 			resp["message"] = "Blockchain Transaction evaluation failed "
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		isKeyExists, err := strconv.ParseBool(string(evaluatedRespone))
// 		if err != nil {
// 			resp["message"] = "Internal Error boolean parsing failed"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		if !isKeyExists {
// 			resp["message"] = "User not found. Please signup"
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}

// 		// then generate a jwt token and send the details to UI
// 		token, err := helper.GenerateToken(user.Name, user.ID, user.CryptoKey)
// 		if err != nil {
// 			resp["message"] = "Internal error"
// 			resp["error"] = fmt.Sprintf("%s", err)
// 			ctx.IndentedJSON(http.StatusInternalServerError, resp)
// 			return
// 		}
// 		user.JWTToken = token
// 		ctx.IndentedJSON(http.StatusOK, user)
// 	}
// }

func getInfoFromIMS(cred Credentials) (User, error) {
	var user User

	resp, err := http.PostForm(constants.IMS_LOGIN_URL, url.Values{
		"RollNo":   {cred.Username},
		"Password": {cred.Password},
	})

	if err != nil {
		return user, fmt.Errorf("login failed")
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		var postData tokenStruct

		if err := json.NewDecoder(resp.Body).Decode(&postData); err != nil {
			log.Printf("decoder error 1 %v", err)
			return user, fmt.Errorf("login failed")
		}

		var imsAuthToken string = postData.JWTToken

		client := http.Client{}
		var url string = constants.IMS_GET_DETAILS_URL + "?RollNo=" + cred.Username
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return user, fmt.Errorf("login failed")
		}

		req.Header = http.Header{
			"jwttoken": {imsAuthToken},
		}

		resp1, err := client.Do(req)
		if err != nil {
			return user, fmt.Errorf("login failed")
		}

		defer resp1.Body.Close()

		// resp1, err := http.Get(constants.IMS_GET_DETAILS_URL + "?RollNo=" + cred.Username)

		// if err != nil {
		// 	return user, fmt.Errorf("login failed")
		// }

		if resp1.StatusCode == 200 {
			var imsResp IMSResp

			if err := json.NewDecoder(resp1.Body).Decode(&imsResp); err != nil {
				log.Printf("decoder error %v", err)
				return user, fmt.Errorf("login failed")
			}

			user.ID = imsResp.RollNo
			user.Name = imsResp.Name
			user.ImgUrl = constants.DEFAULT_PROFILE_IMAGE
			user.MobileNumber = imsResp.MobileNumber
			user.CryptoKey = imsResp.CryptoKey
			user.Email = imsResp.Email

			return user, nil
		}

		return user, fmt.Errorf("login failed")

	}

	return user, fmt.Errorf("login failed")
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// send the username and password to identity management and get cryptographic identity for that user from them
		// if failed to get token then login failed
		var credentials Credentials
		resp := make(map[string]string)
		if err := ctx.BindJSON(&credentials); err != nil {
			resp["message"] = "Username or password not found"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		infoFromIdentity, err := getInfoFromIMS(credentials)

		if err != nil {
			resp["message"] = "Either Username / password is incorrect or You are not registered with IMS"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}

		var user User
		user.ID = infoFromIdentity.ID
		user.Name = infoFromIdentity.Name
		user.MobileNumber = infoFromIdentity.MobileNumber
		user.CryptoKey = infoFromIdentity.CryptoKey
		user.ImgUrl = infoFromIdentity.ImgUrl
		user.Email = infoFromIdentity.Email

		// check if user crypt key exist in our blockchain db if yes then continue to generate token
		// if not then login failed
		evaluatedRespone, err := blockchain.Contract.EvaluateTransaction(constants.CONTRACT_IS_USER_KEY_EXISTS, user.CryptoKey)
		if err != nil {
			resp["message"] = "Blockchain Transaction evaluation failed "
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		isKeyExists, err := strconv.ParseBool(string(evaluatedRespone))
		if err != nil {
			resp["message"] = "Internal Error boolean parsing failed"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		if !isKeyExists {
			// resp["message"] = "User not found. Please signup"
			// ctx.IndentedJSON(http.StatusInternalServerError, resp)
			// return

			// key not found, set crypto key as a userkey in our blockchain db
			txnProposal, err := blockchain.Contract.NewProposal(constants.CONTRACT_CREATE_USER_KEY, client.WithArguments(user.CryptoKey, user.MobileNumber))
			if err != nil {
				resp["message"] = "Blockain Transaction proposal failed"
				resp["error"] = fmt.Sprintf("%s", err)
				ctx.IndentedJSON(http.StatusInternalServerError, resp)
				return
			}
			txnEndorsed, err := txnProposal.Endorse()
			if err != nil {
				resp["message"] = "Blockain Transaction endorsement failed"
				resp["error"] = fmt.Sprintf("%s", err)
				ctx.IndentedJSON(http.StatusInternalServerError, resp)
				return
			}
			txnCommitted, err := txnEndorsed.Submit()
			if err != nil {
				resp["message"] = "Blockain Transaction commit failed"
				resp["error"] = fmt.Sprintf("%s", err)
				ctx.IndentedJSON(http.StatusInternalServerError, resp)
				return
			}
			// log here
			fmt.Printf("Signup Transaction ID : %s Response: %s\n", txnCommitted.TransactionID(), txnEndorsed.Result())
		}

		// then generate a jwt token and send the details to UI
		token, err := helper.GenerateToken(user.Name, user.ID, user.CryptoKey)
		if err != nil {
			resp["message"] = "Internal error"
			resp["error"] = fmt.Sprintf("%s", err)
			ctx.IndentedJSON(http.StatusInternalServerError, resp)
			return
		}
		user.JWTToken = token
		ctx.IndentedJSON(http.StatusOK, user)
	}
}
