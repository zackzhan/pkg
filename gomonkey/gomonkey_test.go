package gomonkey

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
)

var num = 10 //全局变量

func TestApplyGlobalVar(t *testing.T) {
	convey.Convey("TestApplyGlobalVar", t, func() {
		convey.Convey("change", func() {
			patches := gomonkey.ApplyGlobalVar(&num, 150)
			defer patches.Reset()
			convey.So(num, convey.ShouldEqual, 150)
		})

		convey.Convey("recover", func() {
			convey.So(num, convey.ShouldEqual, 10)
		})
	})
}

func networkCompute(a, b int) (int, error) {
	// do something in remote computer
	c := a + b

	return c, nil
}

func Compute(a, b int) (int, error) {
	sum, err := networkCompute(a, b)
	return sum, err
}

func TestFunc(t *testing.T) {
	// mock 了 networkCompute()，返回了计算结果2
	patches := gomonkey.ApplyFunc(networkCompute, func(a, b int) (int, error) {
		return 2, nil
	})

	defer patches.Reset()

	sum, err := Compute(1, 2)
	println("expected %v, got %v", 2, sum)
	if sum != 2 || err != nil {
		t.Errorf("expected %v, got %v", 2, sum)
	}
}
