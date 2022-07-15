package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Helper() //帮助函数 放在run上面
	assert:=assert.New(t)
	assert.Equal(3,Add(1,2),"error")
	//子测试
	t.Run("pos", func(t *testing.T) {
		if Add(2, 3) != 6 {
			t.Fatal("fail")
		}
	})
	t.Run("naa", func(t *testing.T) {
		if Add(2, 3) != 7 {
			t.Fatal("fail")
		}
	})
	//t.Error/t.Errorf   遇错不停   t.Fatal/t.Fatalf  遇错即停。
}

func TestMul(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})

	}
}

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

//如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func TestConn(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	//创建一个 ResponseRecorder (which satisfies http.ResponseWriter)来记录响应
	w := httptest.NewRecorder()
	helloHandler(w,req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)
		assert.Equal(t,"hello world",string(bytes)) // 使用assert提供的断言函数
}
//基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }


//并行测试
//RunParallel 测试并发性能
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(3) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

