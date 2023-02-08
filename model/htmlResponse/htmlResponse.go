package htmlResponseType

import (
	userinfo "server/pack/model/userInfoModel"
)
type GetAllUser struct{
	Data []userinfo.UserInfo
	SingleData []userinfo.UserInfo
	ErrorMessage string
	ErrorStatus bool
}