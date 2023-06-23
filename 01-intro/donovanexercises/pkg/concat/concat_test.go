package concat

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	tests := []struct {
		name    string
		ss      []string
		s       string
		want    string
		wantErr bool
	}{
		{
			name: "positive",
			ss:   []string{"1", "2", "3"},
			want: "1 2 3",
		},
		{
			name:    "empty slice",
			ss:      []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := basic(tt.ss, &tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("basic() error = %v", err)
			}
			if tt.s != tt.want {
				t.Errorf("got: %s want: %s", tt.s, tt.want)
			}
		})
	}
}

func TestCustom(t *testing.T) {
	tests := []struct {
		name string
		ss   []string
		sep  string
		want string
	}{
		{
			name: "positive",
			ss:   []string{"1", "2", "3"},
			sep:  " ",
			want: "1 2 3",
		},
		{
			name: "empty sep",
			ss:   []string{"1", "2", "3"},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := custom(tt.ss, tt.sep)
			if got != tt.want {
				t.Errorf("got: %s want: %s", got, tt.want)
			}
		})
	}
}

// BenchmarkBasic/basic_concat_0_elems-16 || 5117298 ||  223.4 ns/op
// BenchmarkBasic/basic_concat_1000_elems-16 || 2086 || 604986 ns/op
// BenchmarkBasic/basic_concat_1000000_elems-16 || 1 || 142943001100 ns/op
// Warning, explosion may take place!
func BenchmarkBasic(b *testing.B) {
	tests := []struct {
		sliceLen int
	}{
		{
			sliceLen: 0,
		},
		{
			sliceLen: 1000, // 10^3
		},
		{
			sliceLen: 1000000, // 10^6
		},
	}

	for _, tt := range tests {
		b.Run(fmt.Sprintf("basic concat %d elems", tt.sliceLen), func(b *testing.B) {
			var ss []string
			for i := 0; i < tt.sliceLen; i++ {
				ss = append(ss, "1")
			}

			var s string
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				// Надеюсь в бэнчмарке не обязательно хэндлить каждую error :)
				basic(ss, &s)
			}
		})
	}
}

// BenchmarkJoin/join_concat_0_elems-16 || 225141092 || 5.234 ns/op
// BenchmarkJoin/join_concat_1000_elems-16 || 75950 || 15912 ns/op
// BenchmarkJoin/join_concat_1000000_elems-16 || 70 || 16871424 ns/op
// BenchmarkJoin/join_concat_10000000_elems-16 || 6 || 174500350 ns/op
func BenchmarkJoin(b *testing.B) {
	tests := []struct {
		sliceLen int
	}{
		{
			sliceLen: 0,
		},
		{
			sliceLen: 1000, // 10^3
		},
		{
			sliceLen: 1000000, // 10^6
		},
		{
			sliceLen: 10000000, // 10^7
		},
	}

	for _, tt := range tests {
		b.Run(fmt.Sprintf("join concat %d elems", tt.sliceLen), func(b *testing.B) {
			var ss []string
			for i := 0; i < tt.sliceLen; i++ {
				ss = append(ss, "1")
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				strings.Join(ss, " ")
			}
		})
	}
}

// BenchmarkCustom/custom_concat_0_elems-16 || 766762873 || 1.407 ns/op
// BenchmarkCustom/custom_concat_1000_elems-16 || 96386 || 11381 ns/op
// BenchmarkCustom/custom_concat_1000000_elems-16 || 94 || 11436200 ns/op
// BenchmarkCustom/custom_concat_10000000_elems-16 || 9 || 124055578 ns/op
//
// Оказывается, если не передавать sep в функцию custom, а использовать какой-нибудь статичный, то
// производительность вырастет до следующих показателей для 10^7 элементов
// BenchmarkCustom/custom_concat_10000000_elems-16 || 12 || 94625000 ns/op
func BenchmarkCustom(b *testing.B) {
	tests := []struct {
		sliceLen int
	}{
		{
			sliceLen: 0,
		},
		{
			sliceLen: 1000, // 10^3
		},
		{
			sliceLen: 1000000, // 10^6
		},
		{
			sliceLen: 10000000, // 10^7
		},
	}

	for _, tt := range tests {
		b.Run(fmt.Sprintf("custom concat %d elems", tt.sliceLen), func(b *testing.B) {
			var ss []string
			for i := 0; i < tt.sliceLen; i++ {
				ss = append(ss, "1")
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				custom(ss, " ")
			}
		})
	}
}
