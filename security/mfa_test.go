package security

import (
	"encoding/base32"
	"testing"
)

func TestTotp(t *testing.T) {
	g := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: testSecret,
		ExpireSecond:                 30,
		Digits:                       6,
	}
	totp, err := g.Totp()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(totp)
}

func TestQr(t *testing.T) {
	g := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: testSecret,
		ExpireSecond:                 30,
		Digits:                       6,
	}
	qrString := g.QrString("TechBlog:mojotv.cn", "Eric Zhou")
	t.Log(qrString)
}

func TestMakePhoneOrEmail2FACode(t *testing.T) {
	yourAppSecret := "server_config_secret"
	phoneOrEmailOrUserId := "13312345678" //neochau@gmail.com
	//base 32 no padding encode 是必须的不能省去
	googleAuthenticatorKey := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(yourAppSecret + phoneOrEmailOrUserId))
	mfa := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: googleAuthenticatorKey,
		ExpireSecond:                 300, //五分钟之后失效
		Digits:                       4,   //数字
	}
	code, err := mfa.Totp()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("使用这个code发送给用户邮箱手机,", "也可以使用这个code来,验证码收到的code == 服务器收到的code 是否一直", code)
}
