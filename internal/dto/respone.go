package dto

const (
	// Success messages
	SuccessMessageSaveEN         = "data saved successfully"
	SuccessMessageGetEN          = "data retrieved successfully"
	SuccessMessageGetAllEN       = "all data retrieved successfully"
	SuccessMessageUpdateEN       = "data updated successfully"
	SuccessMessageDeleteEN       = "data deleted successfully"
	SuccessMessageUploadEN       = "data uploaded successfully"
	SuccessMessageDataNotFoundEN = "data not found"
	// Failed messages
	FailedMessageSaveEN   = "failed to save data"
	FailedMessageGetEN    = "failed to retrieve data"
	FailedMessageGetAllEN = "failed to retrieve all data"
	FailedMessageUpdateEN = "failed to update data"
	FailedMessageDeleteEN = "failed to delete data"
	FailedMessageUploadEN = "failed to upload data"
	// Success messages in Thai
	SuccessMessageSaveTH         = "บันทึกข้อมูลสำเร็จ"
	SuccessMessageGetTH          = "ดึงข้อมูลสำเร็จ"
	SuccessMessageGetAllTH       = "ดึงข้อมูลทั้งหมดสำเร็จ"
	SuccessMessageUpdateTH       = "อัปเดตข้อมูลสำเร็จ"
	SuccessMessageDeleteTH       = "ลบข้อมูลสำเร็จ"
	SuccessMessageUploadTH       = "อัปโหลดข้อมูลสำเร็จ"
	SuccessMessageDataNotFoundTH = "ไม่พบข้อมูล"
	// Failed messages in Thai
	FailedMessageSaveTH   = "บันทึกข้อมูลล้มเหลว"
	FailedMessageGetTH    = "ดึงข้อมูลล้มเหลว"
	FailedMessageGetAllTH = "ดึงข้อมูลทั้งหมดล้มเหลว"
	FailedMessageUpdateTH = "อัปเดตข้อมูลล้มเหลว"
	FailedMessageDeleteTH = "ลบข้อมูลล้มเหลว"
	FailedMessageUploadTH = "อัปโหลดข้อมูลล้มเหลว"
)

/*
e.Use(middleware.RequestID())

requestID := c.Response().Header().Get(echo.HeaderXRequestID)

	response := Response{
		Message:    "Success",
		StatusHTTP: http.StatusOK,
		StatusBool: true,
		Timestamp:  time.Now().Format(time.RFC3339),
		Path:       c.Path(),
		RequestID:  requestID,
	}
*/
type Response struct {
	Message    string      `json:"message,omitempty"`
	StatusHTTP int         `json:"status_http,omitempty"`
	StatusBool bool        `json:"status_bool,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	ErrorCode  string      `json:"error_code,omitempty"`
	Timestamp  string      `json:"timestamp,omitempty"`
	Path       string      `json:"path,omitempty"`
	RequestID  string      `json:"request_id,omitempty"`
}
