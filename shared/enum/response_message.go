package enum

// ResponseMessage contains common response messages used in the application
type ResponseMessage string

const (
	Success      ResponseMessage = "Success"
	Failure      ResponseMessage = "Fail"
	DataNotFound ResponseMessage = "Data not found"
)
