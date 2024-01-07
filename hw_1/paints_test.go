package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestPaint3(t *testing.T) {
	expected := func() string {
		f, _ := os.Open("./test/paint_3.txt")
		defer f.Close()
		r, _ := io.ReadAll(f)
		return string(r)
	}()
	actual := func() string {
		b := bytes.NewBuffer([]byte{})
		paint3(25, 25, b)
		return b.String()
	}()
	require.Equal(t, expected, actual)
}

func TestPaint6(t *testing.T) {
	expected := func() string {
		f, _ := os.Open("./test/paint_6.txt")
		defer f.Close()
		r, _ := io.ReadAll(f)
		return string(r)
	}()
	actual := func() string {
		b := bytes.NewBuffer([]byte{})
		paint6(25, 25, b)
		return b.String()
	}()
	require.Equal(t, expected, actual)
}

func TestPaint8(t *testing.T) {
	expected := func() string {
		f, _ := os.Open("./test/paint_8.txt")
		defer f.Close()
		r, _ := io.ReadAll(f)
		return string(r)
	}()
	actual := func() string {
		b := bytes.NewBuffer([]byte{})
		paint8(25, 25, b)
		return b.String()
	}()
	require.Equal(t, expected, actual)
}

func TestPaint11(t *testing.T) {
	expected := func() string {
		f, _ := os.Open("./test/paint_11.txt")
		defer f.Close()
		r, _ := io.ReadAll(f)
		return string(r)
	}()
	actual := func() string {
		b := bytes.NewBuffer([]byte{})
		paint11(25, 25, b)
		return b.String()
	}()
	require.Equal(t, expected, actual)
}

func TestPaint13(t *testing.T) {
	expected := func() string {
		f, _ := os.Open("./test/paint_13.txt")
		defer f.Close()
		r, _ := io.ReadAll(f)
		return string(r)
	}()
	actual := func() string {
		b := bytes.NewBuffer([]byte{})
		paint13(25, 25, b)
		return b.String()
	}()
	require.Equal(t, expected, actual)
}
