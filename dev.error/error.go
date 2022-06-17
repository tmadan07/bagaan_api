package aerrors

type Error struct {
	Code        int
	Message     string
	ActualError string
}

// to set actual error
func SetActualError(err error, e *Error) {
	e.ActualError = err.Error()
}
func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Is(tgt error) bool {
	target, ok := tgt.(*Error)
	if !ok {
		return false
	}
	return e.Code == target.Code
}

func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

//1-9 == authorization/token related
//20-29     merchantAccount
//30-39     BusinessCategories

// Different types of error returned by the VerifyToken function
var (

	//token
	ErrTokenCreate        = New(6, "Token could not be creatded")
	ErrInvalidToken       = New(2, "token is invalid")
	ErrExpiredToken       = New(3, "token has expired")
	ErrTokenRefreshFailed = New(11, "Token Refresh failed")

	//authorizaion Header
	ErrNoAuthorizationHeader        = New(4, "authorization header is not provided")
	ErrNotAuthorized                = New(7, "You donot have access rights")
	ErrInvalidAuthorizationFromat   = New(8, "invalid authorization header format")
	ErrUnsupportedAuthorizationType = New(9, "unsupported authorization type")

	//password
	ErrHashingFailed       = New(10, "Password hasing failed")
	ErrWrongPassword       = New(5, "Password is Incorrect")
	ErrIncorrectCredential = New(1, "Incorrect username or password")

	//permissions
	ErrPermissionNotFound   = New(100, "permission data not found ")
	ErrPermissionBinding    = New(101, "Invalid permission Json recieved")
	ErrPermissionInsertion  = New(102, "permission Not able to insert")
	ErrPermissionID         = New(103, "Invalid permission ID")
	ErrPermissionAssignFail = New(104, "All permission assign failed")

	//role
	ErrRoleNotFound  = New(150, "role data not found ")
	ErrRoleBinding   = New(151, "Invalid role Json recieved")
	ErrRoleInsertion = New(152, "role Not able to insert")
	ErrRoleID        = New(153, "Invalid role ID")

	//rule
	ErrRuleNotFound  = New(160, "rule data not found ")
	ErrRuleBinding   = New(161, "Invalid rule Json recieved")
	ErrRuleInsertion = New(162, "rule Not able to insert")
	ErrRuleID        = New(163, "Invalid rule ID")


	//User
	ErrUserNotFound  = New(200, "user data not found ")
	ErrUserBinding   = New(201, "Invalid user Json recieved")
	ErrUserInsertion = New(202, "user Not able to insert")
	ErrUserID        = New(203, "Invalid user ID")
	ErrAuthCodeNotInserted = New(204, "Authcode not inserted in database")
	ErrAuthCodeNotMatched  = New(205, "Authcode not matched with record")
	//Invalid Id  600
	ErrInvalidId = New(600, "provided id is invalid")

	//Options No content
	ErrNoContent = New(700, "No Content found")

	//Json Binding
	ErrJSONBinding = New(11, "Unable to Bind")

	//Merchant
	ErrMerchantNotCreated = New(20, "Unable To create")
	ErrMerchantNotFound   = New(21, "Unable To found merchant")
	ErrApproveMerchantFailed = New(22,"Unable to approve merchant")

	//terminal
	ErrTerminalNotFound  = New(190, "terminal data not found ")
	ErrTerminalBinding   = New(191, "Invalid terminal Json recieved")
	ErrTerminalInsertion = New(192, "terminal Not able to insert")
	ErrTerminalID        = New(193, "Invalid terminal ID")

	//Files class 210-219
	ErrNoFileRecieved      = New(210, "No File Recieved")
	ErrUploadFailed        = New(211, "Failed to upload file")
	ErrReadFailed          = New(211, "Failed to upload file")
	ErrFileDetailInsertion = New(212, "Failed to insert file details returned with status 0")

	//
	ErrMccNotFound = New(300, "MCC Not Found")

	//
	ErrDiscountCashBackCreationFailed = New(400, "failed to insert discountCashBack in table")
	ErrDiscountCashBackNotFound = New(401, "failed to find discountCashBack in table")

	//role_permissions
	ErrRolesPermissionInsertion = New(220, "Failed to create Role_permission")
	ErrRolesPermissionNotFound  = New(221, "Role_permission data not found")

	//permission_categories
	ErrPermissionCategoryTruncate  = New(230, "Failed to delete permission_category")
	ErrPermissionCategoryInsertion = New(231, "Failed to insert permission_category")

	//auditLog
	ErrAuditLogNotFound = New(260, "No Auditlog data found")
	ErrAuditLogCreation = New(261, "No Auditlog data found")

	//subscriptions
	ErrSubscriptionsNotFound  = New(180, "subsription data not found ")
	ErrSubscriptionsBinding   = New(181, "Invalid subscription Json recieved")
	ErrSubscriptionsInsertion = New(182, "subscription Not able to insert")
	ErrSubscriptionsID        = New(183, "Invalid subscription ID")


	//journals   300
	ErrJournalNotFound = New(300, "Joural data not found")

	//notification     400
	ErrNotificationNotFound = New(600, "Notification data not found")

	//payload 500
	ErrPayloadEmpty = New(500, "Payload is empty")






)
