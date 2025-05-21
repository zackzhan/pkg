package security

import (
	"testing"

	dongle "github.com/dromara/dongle"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerateFromPassword(t *testing.T) {
	originalPassword := []byte("123456")
	hashedPassword, err := bcrypt.GenerateFromPassword(originalPassword, bcrypt.DefaultCost)

	assert.NoError(t, err, "bcrypt.GenerateFromPassword should not return an error")
	t.Log("Generated hash:", string(hashedPassword)) // Log the hash

	// Only compare if generation was successful
	if err == nil {
		err = bcrypt.CompareHashAndPassword(hashedPassword, originalPassword)
		assert.NoError(t, err, "bcrypt.CompareHashAndPassword should not return an error, indicating hash matches password")
	}
}

func TestPublicKey(t *testing.T) {
	privateKey := `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNR4mWU8Qyk4sk
EPjCh82Dwt370TBMfJLhj6eZUKGz+G850Yf4R53BXk/mySrIrXjna7+XEvR0cYr4
bz3pnWC+Velw0edF+vxTLiZG97Lz+v3hC7PYpTHV6AIr0Ukgn19Wf6cql6cH0Y8K
BHaSoIM7ylJCvHn3JHAR/Eq0/LFbx9MT/uCKJe2I73AenbnvJe4lIYAR64FJHyDU
riw7w0AX8F+qVmgQDEsvYHCGM5+oaLAkirqIxab3hCh+DWwIB+9LVlrHJLcOwxkL
U6AAbVl3EsBxHym68vikeRwQMdind8/qa4p1Rl2Gw7ZjBA3NBlkSHluuS0H7MHvb
8P8Paa/BAgMBAAECggEBAIwJlU677kgiZvU/rrPmOMj+ErlKx4wmH31IpqjsXKON
gzMKQZuaH4BeeluSOGGWFbipxJWnFKSrIIbCPJY2Cl1N9mZEi2UuBxwni202ZRoT
y0XO2e2dJBYso+6IYz6h+kPj6YIg2o125AGNceYtIdmT9/IW6Y0zilB6r+wvxj45
Jg6XI6JnNL3uCXPdGJqPlGpbBUFcw4N4wjB9I4iX0Jr1Pmdz2QqA7T50sXPMHpcs
BsmFppt9ESlrk9yp8vtx06OWiRUWG58VD0sBA36hCOmZElCN0gvGzMFMhvYG2lGu
q95rtqV0Ix9PQWFmz/RV8qtn//REFgUJw8rwopdmesECgYEA2ZDxble6beHHpZBt
iULUL5r767xM2Qo30RhlU3VXtCNkjGrnZwD41Tm0NQ8pzdSwtthtGymPD9Qh6clG
MuOtkVYZT+hxEURLHEJjl8WOoj7ooGS02SLSQX8LKmo6mM+6P1bjporNa4RHEF5a
ZxaxJvEANYI0XjWDH+w6mRmLG20CgYEA8YrylvMHHVGdhbHQA0ds5EBxws4J0iLY
KYwNXU/Huqn1qFgiYweNW7cKqq9T7IOzCceLN+JjKH8t8oaGOSJPVRvdhzgIXbo/
pwgDk/hsA6K0E9qqqSH11fOmvqwAi+eJ6C2zhELplbWD57n+dScIio39K7ao0nM/
DD+rJMCE/SUCgYAuh5EayChW17Ka9gh3EpPOmpbGE79bq5PEZDNHZhEbtXYLkdFr
dgnBkFW6A0QtgQ9KC39KD9lIyVV3alRZDdhQ/Njs/gkwaCqoIIqsKtQWXt3hb6g9
x/SvjTwWFiHyDHsIkMyfEZBdLi0EkylCYtgAPAqPZwlMHKWubhcaJxxZAQKBgBWQ
fSYC3oC8l56nCFYym0Mpib6FnJTZLYlQCqVpCQDeSaNU8wT6WKF2jwo8CvcTrvlR
illb6dRkvlpnrjYNnR2RFk8LLpwHk8U6zu73/9S3QvJ+1DuzX3pRfXKKcqYK7uH+
qodFXgtSoEsg7NnSuU19bbvoW9wmE9vpF0N5uKXNAoGBALmHTyVWcjQGf0yMjxi8
lk+q1A5MkOBdLrU3RO7iN2qHZHMkKmV1Vr/9HPsViqEUZ1S2zlfMcywRQ2QEu/Mm
qY/6HzkExGKtugIV/piq3wPeojoOsoeqUU35WnfH6hPlqtem4qWBTwzLy+AV9Tgq
e+XyctAwL/5c6G0A0bkOlX6c
-----END PRIVATE KEY-----`

	pwd := "123456"
	// debug
	publicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzUeJllPEMpOLJBD4wofN\ng8Ld+9EwTHyS4Y+nmVChs/hvOdGH+EedwV5P5skqyK1452u/lxL0dHGK+G896Z1g\nvlXpcNHnRfr8Uy4mRvey8/r94Quz2KUx1egCK9FJIJ9fVn+nKpenB9GPCgR2kqCD\nO8pSQrx59yRwEfxKtPyxW8fTE/7giiXtiO9wHp257yXuJSGAEeuBSR8g1K4sO8NA\nF/BfqlZoEAxLL2BwhjOfqGiwJIq6iMWm94Qofg1sCAfvS1ZaxyS3DsMZC1OgAG1Z\ndxLAcR8puvL4pHkcEDHYp3fP6muKdUZdhsO2YwQNzQZZEh5brktB+zB72/D/D2mv\nwQIDAQAB\n-----END PUBLIC KEY-----\n"
	cipherText := dongle.Encrypt.FromString(pwd).ByRsa([]byte(publicKey))
	t.Log(cipherText.ToBase64String())

	cipherText2 := "wAesppuohmg/YBS65DAWK0nOJ8fblR8n86qbo6JLMgWGwyD+8gyhcJRm+jBvwJKT0yAvhYb5hNP7GkF6OF9W8omR1kjpVEImM74Y1xpIZA3bGWDxzlSjLcOKc0iytcqd3z7NgS5XO5vk5i+RPBXODLayw2XxvT6t/wTIMhCIDtA5K/cAyUttfQOUqIoRPsdikxlcKXdgC2+YiiN0A9szLE5wZArznD0qa95asgXdh4UbhegrRrpwEAV/pvYZ7pQcMNmzpzvHsac99+OgQW73q25NVBCRdAbuudTuDg6DcU1uX8QVkTeorQZkO+p+nwawsLwdxKb5+9DrvoTTJvAW0A=="
	toString := dongle.Decrypt.FromBase64String(cipherText2).ByRsa([]byte(privateKey)).ToString()
	t.Log(toString)

	assert.Equal(t, pwd, toString)
}
