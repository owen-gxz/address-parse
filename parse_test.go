package address_parse

import (
	"fmt"
	"testing"
)

func Test_loadData(t *testing.T) {
	p := Parse("北京市             朝阳区富康路姚家园3楼，马云，150-0000-0000")
	p1 := Parse("马云,北京市           朝阳区富康路姚家园3楼，150-0000-0000")
	p2 := Parse("150-0000-0000,马云,北京市    朝阳区富康路姚家园3楼")
	fmt.Println(fmt.Sprintf("%#v",p))
	fmt.Println(p1,p2)
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("北京市     朝阳区富康路姚家园3楼150-0000-0000150-0000-00001马云")
	}
}
