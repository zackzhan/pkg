package ldap

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-ldap/ldap/v3"
)

func TestLdap(t *testing.T) {
	// 设置 LDAP 服务器地址和端口
	ldapServer := "xx"
	ldapPort := 389

	// 设置 LDAP 绑定的用户名和密码
	bindUsername := "CN=Administrator,CN=Users,DC=testad,DC=com"
	bindPassword := "xx"

	// 创建 LDAP 连接
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// 进行 LDAP 绑定
	err = l.Bind(bindUsername, bindPassword)
	if err != nil {
		log.Fatal(err)
	}

	res, err := l.WhoAmI(nil)
	if err != nil {
		log.Fatalf("Failed to call WhoAmI(): %s\n", err)
	}
	fmt.Printf("I am: %+v \n", res)

	attrs := []string{"dn", "cn", "telephoneNumber", "mail", "sAMAccountName", "name", "department"}
	// 设置 LDAP 搜索请求
	searchRequest := ldap.NewSearchRequest(
		"ou=部门1,dc=testad,dc=com", // 搜索的基本 DN
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		//"(objectClass=*)", // 搜索过滤器
		//"(|(objectClass=organizationalUnit)(&(objectCategory=person)(objectClass=user)))", // 搜索过滤器
		"((&(objectCategory=person)(objectClass=user)))", // 搜索过滤器
		attrs, // 返回的属性
		nil,
	)

	// 执行 LDAP 搜索
	searchResult, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 处理搜索结果
	for _, entry := range searchResult.Entries {
		for _, attr := range attrs {
			fmt.Printf("%s:%v ", attr, entry.GetAttributeValue(attr))
		}
		fmt.Printf("\n")
	}
}
