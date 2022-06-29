package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWord(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name : "amggit",
			request : "amggit",
			expected : "Hello amggit",
		},
		{
			name : "Eko",
			request : "Eko",
			expected : "Hello Eko",
		},
		{
			name : "Owen",
			request : "Owen",
			expected : "Hello Owen",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Anggit", func(t *testing.T) {
		result := HelloWorld("Anggit")
		require.Equal(t, "Hello Anggit", result, "Result must be 'Hello Anggit'") 
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'") 
	})
}

func TestMain(m *testing.M){
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit test")
}

func TestHelloWorldSkip(t *testing.T) {
	fmt.Println(runtime.GOOS) 

	if runtime.GOOS == "windows"{
		t.Skip("Unit test cant work in windows")
	}
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Anggit")
	require.Equal(t, "Hello Anggit", result, "Result must be 'Hello Anggit'") 
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Anggit")
	assert.Equal(t, "Hello Anggit", result, "Result must be 'Hello Anggit'") 
}

func TestHelloWorldAnggit(t *testing.T) {
	result := HelloWorld("Anggit")
	if result != "Hello Anggit" {
		//unit test failed
		t.Error("Result must be 'Hello Anggit'")
	}
}

func TestHelloWorldMare(t *testing.T) {
	result := HelloWorld("Mare")
	if result != "Hello Mare" {
		//unit test failed
		t.Fatal("Result must be 'Hello Mare'")
	}
}
