package structuralpatterns

import (
	"fmt"
	"testing"
)

func BenchmarkGetInstance(t *testing.B) {
	ans := GetInstance()
	fmt.Println(ans)
}
