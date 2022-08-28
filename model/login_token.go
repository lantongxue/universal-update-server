package model

type LoginToken struct {
	Expire int    // expired time, seconds
	Token  string // token
	Uid    int    // User ID
}

func NewLoginToken(uid, expire int) LoginToken {
	return LoginToken{
		Uid:    uid,
		Expire: expire,
		Token:  makeToken(),
	}
}

func makeToken() string {
	// token 算法
	// token = hmac_md5( uid + time + rand_string, uid )
	return ""
}

func (t *LoginToken) IsValidate() bool {
	return true
}
