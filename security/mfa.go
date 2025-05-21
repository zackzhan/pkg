package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"time"
)

// GoogleAuthenticator2FaSha1 只实现google authenticator sha1
type GoogleAuthenticator2FaSha1 struct {
	Base32NoPaddingEncodedSecret string //The base32NoPaddingEncodedSecret parameter is an arbitrary key value encoded in Base32 according to RFC 3548. The padding specified in RFC 3548 section 2.2 is not required and should be omitted.
	ExpireSecond                 uint64 //更新周期单位秒
	Digits                       int    //数字数量
}

// otpauth://totp/ACME%20Co:john@example.com?secret=HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ&issuer=ACME%20Co&algorithm=SHA1&digits=6&period=30
const testSecret = "HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ" //base32-no-padding-encoded-string

// Totp 计算Time-based One-time Password 数字
func (m *GoogleAuthenticator2FaSha1) Totp() (code string, err error) {
	count := uint64(time.Now().Unix()) / m.ExpireSecond
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(m.Base32NoPaddingEncodedSecret)
	if err != nil {
		return "", fmt.Errorf("failed to decode base32 secret: %w", err)
	}
	codeInt := hotp(key, count, m.Digits)
	intFormat := fmt.Sprintf("%%0%dd", m.Digits) //数字长度补零
	return fmt.Sprintf(intFormat, codeInt), nil
}

// QrString google authenticator 扫描二维码的二维码字符串
func (m *GoogleAuthenticator2FaSha1) QrString(accountName, issuerName string) (qr string) {
	// Explicitly define the unescaped label string for the path
	unscapedLabelInPath := issuerName + ":" + accountName
	// Then escape it
	escapedLabelInPath := url.PathEscape(unscapedLabelInPath)

	// Explicitly define and escape the issuer for the query parameter
	escapedIssuerQueryParam := url.QueryEscape(issuerName)

	//文档 https://github.com/google/google-authenticator/wiki/Key-Uri-Format
	//otpauth://totp/ACME%20Co:john.doe@email.com?secret=HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ&issuer=ACME%20Co&algorithm=SHA1&digits=6&period=30
	return fmt.Sprintf(`otpauth://totp/%s?secret=%s&issuer=%s&algorithm=SHA1&digits=%d&period=%d`,
		escapedLabelInPath, // Use the new variable for the path
		m.Base32NoPaddingEncodedSecret,
		escapedIssuerQueryParam, // Use the new variable for the issuer query parameter
		m.Digits,
		m.ExpireSecond)
}

func hotp(key []byte, counter uint64, digits int) int {
	//RFC 6238
	//只支持sha1
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	//取sha1的最后4byte
	//0x7FFFFFFF 是long int的最大值
	//math.MaxUint32 == 2^32-1
	//& 0x7FFFFFFF == 2^31  Set the first bit of truncatedHash to zero  //remove the most significant bit
	// len(sum)-1]&0x0F 最后 像登陆 (bytes.len-4)
	//取sha1 bytes的最后4byte 转换成 uint32
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	//取十进制的余数
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}
