package mail

import (
	"fmt"
	"os"
	"testing"

	"github.com/jhillyerd/enmime"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDecodeMail(t *testing.T) {
	Convey("TestMessage_GetMessages", t, func() {
		Convey("gen eml success", func() {
			file, err := os.Open("./testdata/回复：风景如此美丽.eml")
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			// 解析EML文件
			msg, err := enmime.ReadEnvelope(file)
			if err != nil {
				t.Fatal(err)
			}

			// 提取邮件信息
			fmt.Println("Subject:", msg.GetHeader("Subject"))
			fmt.Println("From:", msg.GetHeader("From"))
			fmt.Println("Text Body:", msg.Text)
			fmt.Println("HTML Body:", msg.HTML)
		})
	})
}
