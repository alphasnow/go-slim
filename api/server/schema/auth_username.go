package schema

type AuthUsernameLoginReq struct {
	Username   string `json:"username" binding:"required,min=6,max=24"`
	Password   string `json:"password" binding:"required,min=6,max=24"`
	VerifyCode string `json:"verify_code" binding:"required,len=4"`
	DeviceUUID string `json:"device_uuid" binding:"required,len=4"`
	IsRemember int8   `json:"is_remember" binding:"required"`
}
type AuthUsernameRegisterReq struct {
	Username   string `json:"username" binding:"required,min=6,max=24"`
	Password   string `json:"password" binding:"required,min=6,max=24"`
	VerifyCode string `json:"verify_code" binding:"required,len=4"`
	DeviceUUID string `json:"device_uuid" binding:"required,len=4"`
	IsAgree    int8   `json:"is_agree" binding:"required"`
}
