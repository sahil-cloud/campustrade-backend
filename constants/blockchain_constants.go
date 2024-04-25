package constants

const CRYPTO_PATH string = "/home/ccd075/fabric-samples/campustrade-backend/crypto-config/peerOrganizations/org1.example.com"
const CERT_PATH string = CRYPTO_PATH + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem"
const KEY_PATH string = CRYPTO_PATH + "/users/User1@org1.example.com/msp/keystore/"
const TLS_CERT_PATH string = CRYPTO_PATH + "/peers/peer0.org1.example.com/tls/ca.crt"

// name for the chaincode during channel run (EfE)
const CHAINCODE_NAME string = "CampusTrade"
const CHANNEL_NAME string = "mychannel"

// smart contract function names

// my contracts //
// get all products
// CreateUserKey (cryptokey, mobileNumber)
const CONTRACT_CREATE_USER_KEY string = "CreateUserKey"

// IsUserKeyExists (cryptokey)
const CONTRACT_IS_USER_KEY_EXISTS string = "IsUserKeyExists"
const CONTRACT_GET_ALL_PRODUCTS = "GetAllProducts"

const CONTRACT_GET_ALL_SOLD_PRODUCTS = "GetAllSoldProducts"

const CONTRACT_GET_ALL_UNSOLD_PRODUCTS = "GetAllUnsoldProducts"

const CONTRACT_ADD_PRODUCT = "AddProduct"

const CONTRACT_ADD_RATING = "AddRating"

const CONTRACT_GET_ALL_RATINGS = "GetAllRatings"

const CONTRACT_ADD_REVIEW = "AddReview"

const CONTRACT_GET_ALL_REVIEWS = "GetAllReviews"

const CONTRACT_GET_BUYER_ORDERED_PRODUCTS = "GetBuyerOrderedProducts"

const CONTRACT_GET_BUYER_DELIVERED_PRODUCTS = "GetTransactionByBuyerId"

const CONTRACT_GET_SELLER_SOLD_PRODUCTS = "GetTransactionBySellerId"

const CONTRACT_GET_SELLER_ORDERED_PRODUCTS = "GetSellerOrderedProducts"

const CONTRACT_ADD_TRANSACTION = "AddTransaction"

const CONTRACT_VERIFY_TRANSACTION = "VerifyTransaction"

const CONTRACT_GET_ALL_TRANSACTIONS = "GetAllTransactions"
