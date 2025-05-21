package security

import (
	"encoding/base32"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotp(t *testing.T) {
	g := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: testSecret,
		ExpireSecond:                 30,
		Digits:                       6,
	}
	totp, err := g.Totp()

	assert.NoError(t, err, "g.Totp() should not return an error")
	if err == nil {
		assert.Len(t, totp, g.Digits, "TOTP code should have the length specified by g.Digits")
	}
	t.Log(totp)
}

func TestQr(t *testing.T) {
	g := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: testSecret,
		ExpireSecond:                 30,
		Digits:                       6,
	}
	accountName := "user@example.com"
	issuerName := "MyOrg"
	qrString := g.QrString(accountName, issuerName)
	t.Log(qrString)

	assert.Contains(t, qrString, "secret="+testSecret, "QR string should contain the secret")
	assert.Contains(t, qrString, "issuer="+issuerName, "QR string should contain the correct issuer") // issuerName "MyOrg" does not require QueryEscape
	assert.Contains(t, qrString, "digits="+strconv.Itoa(g.Digits), "QR string should contain correct digits")
	assert.Contains(t, qrString, "period="+strconv.FormatUint(g.ExpireSecond, 10), "QR string should contain correct period")
	// Expected label: url.PathEscape(issuerName + ":" + accountName) -> url.PathEscape("MyOrg:user@example.com") -> "MyOrg%3Auser%40example.com"
	assert.Contains(t, qrString, "otpauth://totp/MyOrg%3Auser%40example.com?", "QR string should contain the correct URL-escaped label")
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

	assert.NoError(t, err, "mfa.Totp() should not return an error")
	if err == nil {
		assert.Len(t, code, mfa.Digits, "TOTP code should have the length specified by mfa.Digits")
	}
	t.Log("使用这个code发送给用户邮箱手机,", "也可以使用这个code来,验证码收到的code == 服务器收到的code 是否一直", code)
}
