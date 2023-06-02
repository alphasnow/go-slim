package schema

type AppClientRegisterRes struct {
	AuthLoginRes
	ClientRes ClientRes `json:"client"`
}
type AppClientRegisterReq struct {
	DeviceID   string `json:"device_id" binding:"max=64"`
	DeviceUUID string `json:"device_uuid" binding:"max=64"`
	AppVersion string `json:"app_version" binding:"max=64"`
}
type ClientRes struct {
	DeviceUUID string `json:"device_uuid"`
	DeviceID   string `json:"device_id"`
}
type UserRes struct {
	UID       uint  `json:"uid"`
	CreatedAt int64 `json:"created_at"`
}
