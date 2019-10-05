package main

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestCallService_Get(t *testing.T) {
	type fields struct {
		Writer   io.Writer
		Response *http.Response
		err      error
	}
	type args struct {
		url   string
		model interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &CallService{
				Writer:   tt.fields.Writer,
				Response: tt.fields.Response,
				err:      tt.fields.err,
			}
			got, err := e.Get(tt.args.url, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
