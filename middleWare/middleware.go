package middleWare

import (
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/pkg/ecode"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"io/ioutil"
	a4 "kratosL/app04/api"
)

var (
	Forbiden = ecode.New(403)
)

var ErrorCode = map[int]string{
	403: "Forbiden",
}

func AuthMiddleWare(c *bm.Context){
	ecode.Register(ErrorCode)
	var req a4.LoginReq
	r, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(r, &req)
	fmt.Println(string(r))
	if req.Name == "aa" && req.Password == "111"{
		token := c.Request.Header.Get("access-token")
		if token == "111"{
			return
		}else {
			c.JSON("token error", Forbiden)
			c.Abort()
		}
		return
	}else {
		c.JSON("no login", Forbiden)
		c.Abort()
	}
}