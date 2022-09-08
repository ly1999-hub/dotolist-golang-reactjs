package response

const (
	CommonSuccess    = "common_success"
	CommonNotFound   = "common_not_found"
	CommonBadRequest = "common_bad_request"
)

var common = []Code{
	{
		Key:     CommonSuccess,
		Message: "thành công",
		Code:    1,
	},
	{
		Key:     CommonBadRequest,
		Message: "dữ liệu không hợp lệ",
		Code:    2,
	},
	{
		Key:     CommonNotFound,
		Message: "dữ liệu không tìm thấy",
		Code:    4,
	},
}
