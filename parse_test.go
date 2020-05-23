package address_parse

import (
	"fmt"
	"testing"
)

func Test_loadData(t *testing.T) {
	p := Parse("北京市     朝阳区富康路姚家园3楼150-0000-0000150-0000-00001马云")
	fmt.Println(fmt.Sprintf("%#v",p))
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("北京市     朝阳区富康路姚家园3楼150-0000-0000150-0000-00001马云")
	}
}
