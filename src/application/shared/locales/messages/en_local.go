package messages

var EnMessages = map[string]string{
	"SOMETHING_WENT_WRONG":           "Oh sorry, something went wrong with current action!",
	"RESOURCE_NOT_FOUND":             "Resource not found!",
	"SOME_PARAMETERS_ARE_MISSING":    "Some parameters are missing: %s.",
	"UNKNOWN_RESPONSE_STATUS":        "Response status from server unknown.",
	"TOOL_HAS_NOT_BEEN_INITIALIZED":  "The %s tool has not been initialized.",
	"PROCESSING_DATA_CLIENT_ERROR":   "Error processing http client data.",
	"DEPENDENCY_NOT_FOUND":           "%s not found in dependencies container.",
	"AUTHORIZATION_REQUIRED":         "Authorization is required.",
	"INVALID_USER_OR_PASSWORD":       "Invalid user or password.",
	"ERROR_CREATING_USER":            "Create user error.",
	"USER_WAS_CREATED":               "User was created.",
	"USER_WITH_EMAIL_ALREADY_EXISTS": "A user has already registered with the email address: %s.",
	"INVALID_EMAIL":                  "Invalid email.",
	"INVALID_PASSWORD":               "Invalid password.",
	"INVALID_SESSION":                "Invalid session.",
}
