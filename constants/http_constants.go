package constants

const HTTP_GET string = "GET"
const HTTP_POST string = "POST"
const HTTP_PUT string = "PUT"
const HTTP_DELETE string = "DELETE"
const MESSAGE_OK string = "OK"
const STATUS_200 uint = 200
const STATUS_300 uint = 300
const STATUS_500 uint = 500
const AUTHENTICATION_FAILED string = "AUTHENTICATION_FAILED"

const NGROK_DOMAIN_NAME string = "stork-true-hedgehog.ngrok-free.app"

// Login API
// 200->login successfully
// 401->.invalid credentials
// 500-> otherwise …not registered
var IMS_LOGIN_URL string = "https://oryx-modern-carefully.ngrok-free.app/LogIn"

// IsStudentRegister
// 200-> yes
// 401 -> not registered.
// https://oryx-modern-carefully.ngrok-free.app/IsStudentRegister?RollNo=cs22m013
var IMS_IS_REGISTERED_URL string = "https://oryx-modern-carefully.ngrok-free.app/IsStudentRegister"

// GetStudentDetails
// https://oryx-modern-carefully.ngrok-free.app/GetStudentDetails?RollNo=cs22m013
// 200-> yes
// 401-> student not registered.
// 500->internal server
var IMS_GET_DETAILS_URL string = "https://oryx-modern-carefully.ngrok-free.app/GetStudentDetails"

var DEFAULT_PROFILE_IMAGE string = "https://images.unsplash.com/photo-1633332755192-727a05c4013d?q=80&w=1780&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
var DEFAULT_MOBILE_NUMBER string = "0000000000"
