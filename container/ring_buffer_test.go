package container

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRingBuffer(t *testing.T) {
	rb, _ := NewRingBuffer[int](5)
	if rb.Len() != 0 {
		t.Errorf("Expected length 0, got %d", rb.Len())
	}
	if rb.IsFull() {
		t.Error("New buffer should not be full")
	}
}

func TestReadWrite(t *testing.T) {
	rb, _ := NewRingBuffer[int](3)

	// Test basic write and read
	if !rb.Write(1) {
		t.Error("Write should succeed")
	}
	if rb.Len() != 1 {
		t.Errorf("Expected length 1, got %d", rb.Len())
	}

	val, ok := rb.Read()
	if !ok || val != 1 {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}
	if rb.Len() != 0 {
		t.Errorf("Expected length 0, got %d", rb.Len())
	}
}

func TestOverflow(t *testing.T) {
	rb, _ := NewRingBuffer[int](2)

	// Fill buffer
	if !rb.Write(1) {
		t.Error("Write should succeed")
	}
	if !rb.Write(2) {
		t.Error("Write should succeed")
	}
	if rb.Len() != 2 {
		t.Errorf("Expected length 2, got %d", rb.Len())
	}

	// Test overflow
	if rb.Write(3) {
		t.Error("Write should fail on full buffer")
	}

	// Read one and write again
	val, ok := rb.Read()
	if !ok || val != 1 {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}
	if !rb.Write(3) {
		t.Error("Write should succeed after read")
	}
}

func TestMustWrite(t *testing.T) {
	rb, _ := NewRingBuffer[int](2)

	// Fill buffer
	rb.MustWrite(1)
	rb.MustWrite(2)

	// Test overwrite
	rb.MustWrite(3) // Should overwrite the first element

	val, ok := rb.Read()
	if !ok || val != 3 {
		t.Errorf("Expected (3, true), got (%d, %v)", val, ok)
	}
	val, ok = rb.Read()
	if !ok || val != 2 {
		t.Errorf("Expected (2, true), got (%d, %v)", val, ok)
	}
}

func TestEmptyRead(t *testing.T) {
	rb, _ := NewRingBuffer[int](2)

	val, ok := rb.Read()
	if ok || val != 0 {
		t.Errorf("Expected (0, false), got (%d, %v)", val, ok)
	}
}

func TestConcurrentAccess(t *testing.T) {
	rb, _ := NewRingBuffer[int](30)
	flush := make(chan bool)
	done := make(chan bool)
	finished := make(chan bool)
	// Writer goroutine
	go func() {
		var i int
		for i < 1000 {
			if !rb.Write(i) {
				flush <- true
				continue
			}
			i++
		}
		flush <- true
		done <- true
	}()

	// Reader goroutine
	go func() {
		for {
			select {
			case <-done:
				finished <- true
				return
			case <-flush:
				for {
					data, ok := rb.Read()
					if !ok {
						break
					}
					fmt.Printf("goroutine1: read data %d\n", data)
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-done:
				finished <- true
				return
			case <-flush:
				for {
					data, ok := rb.Read()
					if !ok {
						break
					}
					fmt.Printf("goroutine2: read data %d\n", data)
				}
			}
		}
	}()

	// Wait for both goroutines to finish
	<-finished
	fmt.Println("exit")

	if rb.Len() != 0 {
		t.Errorf("Expected empty buffer, got length %d", rb.Len())
	}
}

func TestRingBuffer_ZeroSize(t *testing.T) {
	rb, err := NewRingBuffer[int](0)
	assert.Nil(t, rb)
	assert.NotNil(t, err)
}

func TestRingBuffer_ReadBatch(t *testing.T) {
	rb, _ := NewRingBuffer[int](5)
	for i := 0; i < 5; i++ {
		rb.Write(i)
	}

	l := rb.ReadBatch(3)
	assert.Equal(t, l, []int{0, 1, 2})

	l = rb.ReadBatch(2)
	assert.Equal(t, l, []int{3, 4})

	l = rb.ReadBatch(1)
	assert.True(t, len(l) == 0)
}
