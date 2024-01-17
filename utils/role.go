package utils

const (
	ADMIN = "ADMIN"
	USER  = "USER"
)

const (
	PENDING  = "PENDING"
	ACTIVE   = "ACTIVE"
	COMPLETE = "COMPLETE"
)

func IsAdmin(role string) bool {
	return role == ADMIN
}

func IsUser(role string) bool {
	return role == USER
}


func IsPending(status string) bool {
	return status == PENDING
}

func IsActive(status string) bool {
	return status == ACTIVE
}

func IsComplete(status string) bool {
	return status == COMPLETE
}

