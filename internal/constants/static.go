package constants

import "os"

var IS_DEBUG = os.Getenv("DEBUG") == "1" || os.Getenv("DEBUG") == "true"

const JWT_ACCESS_TOKEN_VALIDITY_TIME = 5 * 60             // 5 minutes
const JWT_REFRESH_TOKEN_VALIDITY_TIME = 24 * 60 * 60 * 30 // 30 days
const JWT_ACCESS_TOKEN_COOKIE_NAME = "access_token"
const JWT_REFRESH_TOKEN_COOKIE_NAME = "refresh_token"

var JWT_SECRET = os.Getenv("JWT_SECRET")
