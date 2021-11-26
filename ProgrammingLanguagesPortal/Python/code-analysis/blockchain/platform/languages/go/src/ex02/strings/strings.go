package main

import (
	"fmt"
	"bytes"
	"strings"
	"strconv"
	"unicode/utf8"
)

func main(){

	s := "hello, golang"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Println(s[0:5])
	fmt.Println(s[:5]) 
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	literals := "салем казақща! 您好，哈萨克语！Hello Kazakh Language!"
	fmt.Println(literals);

	fmt.Println("世界","\xe4\xb8\x96\xe7\x95\x8c","\u4e16\u754c","\U00004e16\U0000754c")

	s2 := "Hello, 首都图书馆"
	fmt.Println(len(s2)) 
	fmt.Println(utf8.RuneCountInString(s2))

	s = s2
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}

	for i,r := range s2 {
		fmt.Printf("%d\t%c\t%X\n", i, r, r)
	}

	s3 := "1913 жыл, 5 қаңтар"
	fmt.Println(s3)

	s4 := "Астанада 500 домбырашы бір уақытта күй тартты"
	fmt.Println(s4)

	s5 := ` تورابىمىزداعى مازمۇنداردىڭ مەنشىك ۇقىعى شينجياڭ ۇيعۇر اۆتونوميالى رايوندىق حالىق ۇكىمەتى اقپارات كەڭسەسىنە ءتان، جاريالانعان مازمۇنداردى سىلتەمەسىز كوشىرۋگە بولمايدى. كوشىرىپ پايدالانۋعا تۋرا كەلگەندە، «ءتاڭىرتاۋ تورابىنان» الىنعانى انىق ەسكەرتىلۋى ءتيىس، بولماسا زاڭدىق جاۋاپكەرلىگى قۋزاستىرىلادى.`
	fmt.Println(s5)

	s6 := "جىل باسىنان بەرى قازاقستاندا باسپانا قىمباتتادى"
	fmt.Println(s6)

	// "program" in Japanese katakana
	s7 := "このセクションは、ハッカー向けのマニュアルです。 自分の手を汚して何かをしたい人、PHP の内部構造を知ることで自分の理解を深めたい人、 イケてる拡張モジュールを自作したい人。そんな人たちのために用意しました。 このセクションでは、PHP の内部構造を深く掘り下げたり拡張モジュールの書きかたを説明したり、 複雑怪奇なマクロだらけのコードを読み解くヒントを提供したりします。 重要な内部機能はすべて取り上げますが、きちんと理解するには結局はソースを読むことです。"
	fmt.Printf("% c\n", s7) 
	fmt.Printf("% x\n", s7) 
	r2 := []rune(s7)
	fmt.Printf("%x\n", r2) 

	s8 := "巴特尔·马哈比尔"
	fmt.Printf("% c\n", s8) 
	fmt.Printf("% x\n", s8) 
	r3 := []rune(s8)
	fmt.Printf("%x\n", r3) 
	fmt.Println(string(r3))

	fmt.Println(string(1234567)) 

	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"

	fmt.Println(basename02("a/b/c.go")) // "c"
	fmt.Println(basename02("c.d.go"))   // "c.d"
	fmt.Println(basename02("abc"))      // "abc"

	s9 := "31415926"
	fmt.Println(s9,comma(s9))

	s10 := "abc"
	b2 := []byte(s10)
	s11 := string(b2)
	fmt.Println(s10,b2,s11)

	fmt.Println(intsToString([]int{1, 2, 3}))

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) 

	fmt.Println(strconv.FormatInt(int64(x), 2)) 
	fmt.Println(fmt.Sprintf("x=%b", x))
	
	x2, err1 := strconv.Atoi("123") 
	y2, err2 := strconv.ParseInt("123", 10, 64)
	fmt.Println(x2,y2,err1,err2)

	
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		} 
	}
	return false
}


     
// basename removes directory components and a .suffix.
     
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
     
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
    	if s[i] == '/' {
        	s = s[i+1:]
			break
		} 
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
    	if s[i] == '.' {
        	s = s[:i]
			break
		} 
	}
	return s
}

func basename02(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
         
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is like fmt.Sprintf(values) but adds commas.
     
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
    	if i > 0 {
        	buf.WriteString(", ")
    	}
    	fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
