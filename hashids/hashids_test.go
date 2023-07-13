package hashids

import (
	"fmt"
	"testing"

	gohashids "github.com/speps/go-hashids/v2"
)

// go-hashids 是一个用于生成短、易读、唯一 ID 的 Go 语言库，它基于 Hashids 算法。Hashids 算法是一种非加密算法，将数字 ID 转换为短字符串，该字符串包含自定义字符集中的字符。
// 根据 Hashids 算法的特性，如果输入的数字 ID 有限制，那么生成的字符串 ID 也会有限制。如果 ID 数字空间不够大，那么会存在重复的情况。
// 例如，如果您使用 go-hashids 生成 10 位字符串 ID，字符集包含大小写字母和数字，那么最多可以生成 62 的 10 次方个不同的 ID，即 839,299,365,868,340,224 个。如果您生成的 ID 数量超过了这个范围，那么就会出现重复的情况。
// 此外，在使用 go-hashids 时，还需要注意生成 ID 的随机性。如果使用相同的 salt 和输入值，每次生成的 ID 都会相同，因此在生成 ID 时要确保输入值的随机性和唯一性，以避免出现重复的情况。
func TestCase(t *testing.T) {
	hd := gohashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 10

	h, _ := gohashids.NewWithData(hd)

	e, _ := h.Encode([]int{456, 1, 34534, 534534, 5346346, 65645, 6546, 6456, 56456})
	fmt.Println(e)

	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}
