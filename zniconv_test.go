package zniconv

import (
	"bytes"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	want := "一二三四五六七八九十"
	for i, bufSize := range []int{3, 5, 7, 11, 13, 15} {
		in := bytes.NewBufferString(want)
		r, err := NewReader(Options{From: "utf-8", BufSize: bufSize}, in)
		if err != nil {
			t.Fatalf("newreader err=%v", err)
		}
		b := make([]byte, 10)
		n, err := r.Read(b)
		t.Logf("read n=%d err=%v", n, err)
		if err != nil {
			t.Fatalf("%d read got err=%v", i, err)
		}
		if string(b[:n]) != want[:10/3*3] {
			t.Fatalf("%d read want=%s got=%s", i, want[:10/3*3], b[:n])
		}
		if err = r.Close(); err != nil {
			t.Fatalf("%d close reader err=%v", i, err)
		}
	}

	r, _ := NewReader(Options{From: "utf-8"}, bytes.NewBufferString(want))
	out := new(bytes.Buffer)
	if _, err := io.Copy(out, r); err != nil {
		t.Fatalf("copy err=%v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("close err=%v", err)
	}
	if out.String() != want {
		t.Fatalf("copy want=%s got=%s", want, out.String())
	}
}

func TestReaderErr(t *testing.T) {
	want := []byte("一二三四五六七八九十")
	got := make([]byte, len(want))
	for _, i := range []int{0, 4, 6, 7, 11} {
		bad := append([]byte{}, want...)
		bad[i] = 0xff
		r, err := NewReader(Options{From: "utf-8"}, bytes.NewBuffer(bad))
		if err != nil {
			t.Fatalf("%d newreader err=%v", i, err)
		}
		n, err := r.Read(got)
		if err == nil {
			t.Fatalf("%d read expect err got=%d", i, n)
		}
		t.Logf("%d read err=%v", i, err)
		if !bytes.Equal(got[:n], want[:n]) {
			t.Fatalf("%d read want=%s got=%s", i, want[:n], got[:n])
		}

		r.Reset(bytes.NewBuffer(want))
		b, err := r.ReadAll()
		if err != nil {
			t.Fatalf("readall err=%v after reset", err)
		}
		if !bytes.Equal(b, want) {
			t.Fatalf("readall after reset want=%s got=%s", want, b)
		}

		if err = r.Close(); err != nil {
			t.Fatalf("%d close err=%v", i, err)
		}
	}

	bad := append([]byte{}, want...)
	bad = bad[:len(bad)-1]
	r, err := NewReader(Options{From: "utf-8"}, bytes.NewBuffer(bad))
	if err != nil {
		t.Fatalf("newreader err=%v", err)
	}
	n, err := r.Read(got)
	if err == nil {
		t.Fatalf("read expect err got=%d", n)
	}
	t.Logf("read err=%v", err)
	if !bytes.Equal(got[:n], want[:n]) {
		t.Fatalf("read want=%s got=%s", want[:n], got[:n])
	}
	if err = r.Close(); err != nil {
		t.Fatalf("close err=%v", err)
	}
}

func TestToAndFrom(t *testing.T) {
	want := []byte("1.一二三四五六七八九十")
	gbk, err := Convert("utf8", "gbk", want)
	if err != nil {
		t.Fatalf("to gbk err=%v", err)
	}
	u8, err := Convert("gbk", "utf8", gbk)
	if err != nil {
		t.Fatalf("from gbk err=%v", err)
	}
	if !bytes.Equal(u8, want) {
		t.Fatalf("want=%s got=%s", want, u8)
	}
}
