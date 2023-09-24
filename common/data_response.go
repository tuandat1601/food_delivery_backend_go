package common
type DataResponse struct{
	Status int         `json:status`
	Data interface{}  `json:data`

}
type DataResponseLogin struct{
	Status int         `json:status`
	AccessToken string  `json:access_token`

}
func NewDataResponse(status int, data interface{}) * DataResponse{
	return &DataResponse{ Status: status, Data: data}
}
func NewDataResponseLogin(status int, access_token string) * DataResponseLogin{
	return &DataResponseLogin{ Status: status, AccessToken: access_token}
}