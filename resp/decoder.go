package resp

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	// "bytes"
	// "github.com/juju/errors"
	// "github.com/ngaut/log"
)

type decoder struct {
	r *bufio.Reader
}

func Decode(r *bufio.Reader) (Resp, error) {
	return decode(r)
}
func DecodeString(s string) (Resp, error) {
	reader := bufio.NewReader(strings.NewReader(s))
	return decode(reader)
}
func DecodeBytes(bs []byte) (Resp, error) {
	return DecodeString(string(bs))
}

func decode(r *bufio.Reader) (Resp, error) {
	d := &decoder{r}
	return d.decodeNext()
}

func (d *decoder) decodeNext() (Resp, error) {
	// s, err := d.r.ReadString('\n')
	s, _, err := d.r.ReadLine()

	if err != nil {
		return nil, fmt.Errorf("empty message")
	}
	token := string(s)
	specifier := token[0]
	switch string(specifier) {
	default:
		if strings.HasPrefix(token, "PING") {
			return NewString("PONG"), nil
		}
		return nil, fmt.Errorf("unknown token: %s", token)
	case "+":
		return d.decodeString(token)
	case ":":
		return d.decodeInt(token)
	case "$":
		return d.decodeBulkBytes(token)
	case "*":
		return d.decodeArray(token)
	}

}

func (d *decoder) decodeString(token string) (Resp, error) {
	return NewString(strings.TrimPrefix(token, "+")), nil
}

func (d *decoder) decodeArray(token string) (Resp, error) {
	arraySize, _ := strconv.Atoi(strings.TrimPrefix(token, "*"))
	arr := NewArray()
	for i := 0; i < arraySize; i++ {
		elem, err := d.decodeNext()
		if err != nil {
			return nil, err
		}
		arr.Append(elem)
	}
	return arr, nil
}

func (d *decoder) decodeBulkBytes(token string) (Resp, error) {
	bulkStringSize, _ := strconv.Atoi(strings.TrimPrefix(token, "$"))
	if bulkStringSize == -1 {
		return nil, nil
	}
	s, _, err := d.r.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("invalid bulk string after %s", s)
	}
	return NewBulkBytes(s), nil
}

func (d *decoder) decodeInt(token string) (Resp, error) {
	num, _ := strconv.Atoi(strings.TrimPrefix(token, ":"))
	return NewInt(int64(num)), nil
}
