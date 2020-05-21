package address_parse

import "testing"

//func Test_loadData(t *testing.T) {
//	Parse("北京市     朝阳区富康路姚家园3楼150-0000-0000150-0000-00001马云")
//}

func BenchmarkParse(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Parse("北京市     朝阳区富康路姚家园3楼150-0000-0000150-0000-00001马云")
	}
}
