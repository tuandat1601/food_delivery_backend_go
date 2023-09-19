package common
type DataResponse struct{
	Status int         `json:status`
	Data interface{ }  `json:data`

}
func NewDataResponse(status int, data interface{}) * DataResponse{
	return &DataResponse{ Status: status, Data: data}
}