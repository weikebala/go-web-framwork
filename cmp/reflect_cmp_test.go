package cmp

import "testing"

//func TestReflect_cmp(t *testing.T) {
//	type args struct {
//		a interface{}
//		b interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		{
//			name: "比较通过",
//			args: args{
//				a:1,
//				b:1,
//			},
//			want: true,
//		},
//		{
//			name: "比较不通过通过",
//			args: args{
//				a:map[string]interface{}{
//					"a":1,
//					"b":[]int{1,2,3},
//				},
//				b:map[string]interface{}{
//					"a":1,
//					"b":[]int{1,2,3},
//				},
//			},
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Reflect_cmp(tt.args.a, tt.args.b); got != tt.want {
//				t.Errorf("Reflect_cmp() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func BenchmarkReflect_cmp(b *testing.B) {
	x := map[string]interface{}{
		"a":1,
		"b":[]int{1,2,3},
	}
	y:= map[string]interface{}{
		"a":1,
		"b":[]int{1,2,3},
		"c":1.25,
	}
	for i := 0; i < b.N; i++ {
		Reflect_cmp(x, y)
	}
}


func TestReflect_cmp(t *testing.T) {
	table := []struct{
		arg1 interface{}
		arg2 interface{}
		except bool
	}{
		{
			arg1: 1,
			arg2: 2,
			except: false,
		},
		{
			arg1: 1,
			arg2: 2,
			except: true,
		},
	}
	for _,item :=  range table {
		if result := Reflect_cmp(item.arg1,item.arg2); result != item.except {
			t.Errorf("arg1:%v arg2:%v except:%v get:%v", item.arg1, item.arg2, item.except, result)
		} else {
			t.Log(`通过`)
		}
	}
}