package web

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

type Base struct {
	Status            string      `json:"status"`
	Config            interface{} `json:"config"`
	ServerProcessTime string      `json:"server_process_time"`
	ErrorMessage      []string    `json:"message_error,omitempty"`
	StatusMessage     []string    `json:"message_status,omitempty"`
}

type Response struct {
	Base
	Data interface{} `json:"data"`
}

func (mj *Base) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Base) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "status":`)
	fflib.WriteJsonString(buf, string(mj.Status))
	buf.WriteString(`,"config":`)
	/* Interface types must use runtime reflection. type=interface {} kind=interface */
	err = buf.Encode(mj.Config)
	if err != nil {
		return err
	}
	buf.WriteString(`,"server_process_time":`)
	fflib.WriteJsonString(buf, string(mj.ServerProcessTime))
	buf.WriteByte(',')
	if len(mj.ErrorMessage) != 0 {
		buf.WriteString(`"message_error":`)
		if mj.ErrorMessage != nil {
			buf.WriteString(`[`)
			for i, v := range mj.ErrorMessage {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.StatusMessage) != 0 {
		buf.WriteString(`"message_status":`)
		if mj.StatusMessage != nil {
			buf.WriteString(`[`)
			for i, v := range mj.StatusMessage {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

func (mj *Response) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (mj *Response) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "data":`)
	/* Interface types must use runtime reflection. type=interface {} kind=interface */
	err = buf.Encode(mj.Data)
	if err != nil {
		return err
	}
	buf.WriteString(`,"status":`)
	fflib.WriteJsonString(buf, string(mj.Status))
	buf.WriteString(`,"config":`)
	/* Interface types must use runtime reflection. type=interface {} kind=interface */
	err = buf.Encode(mj.Config)
	if err != nil {
		return err
	}
	buf.WriteString(`,"server_process_time":`)
	fflib.WriteJsonString(buf, string(mj.ServerProcessTime))
	buf.WriteByte(',')
	if len(mj.ErrorMessage) != 0 {
		buf.WriteString(`"message_error":`)
		if mj.ErrorMessage != nil {
			buf.WriteString(`[`)
			for i, v := range mj.ErrorMessage {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.StatusMessage) != 0 {
		buf.WriteString(`"message_status":`)
		if mj.StatusMessage != nil {
			buf.WriteString(`[`)
			for i, v := range mj.StatusMessage {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}