package sendcloud_test

import (
	"encoding/json"
	"testing"

	"github.com/afosto/sendcloud-go"
	"github.com/stretchr/testify/assert"
)

func TestGetPayload(t *testing.T) {
	tests := []struct {
		Name   string
		Params sendcloud.ParcelParams
	}{
		{
			Name:   "Should ignore empty weight",
			Params: sendcloud.ParcelParams{},
		},
		{
			Name: "Should include weight in request",
			Params: sendcloud.ParcelParams{
				Weight: "0.040",
			},
		},
	}
	for _, test := range tests {
		payload := test.Params.GetPayload()
		b, _ := json.Marshal(payload)
		var obj sendcloud.ParcelRequestContainer
		json.Unmarshal(b, &obj)
		if test.Params.Weight != "" {
			assert.Equal(t, test.Params.Weight, *obj.Parcel.Weight, test.Name)
		} else {
			assert.Nil(t, obj.Parcel.Weight, test.Name)
		}
	}
}

func TestGetResponse(t *testing.T) {
	tests := []struct {
		Name     string
		Response sendcloud.ParcelResponseContainer
		Out      sendcloud.Parcel
	}{
		{
			Name: "Should ignore nil weight",
			Response: sendcloud.ParcelResponseContainer{
				Parcel: sendcloud.ParcelResponse{},
			},
			Out: sendcloud.Parcel{},
		},
		{
			Name: "Should include weight",
			Response: sendcloud.ParcelResponseContainer{
				Parcel: sendcloud.ParcelResponse{
					Weight: sendcloud.String("0.05"),
				},
			},
			Out: sendcloud.Parcel{
				Weight: "0.05",
			},
		},
	}
	for _, test := range tests {
		res := test.Response.GetResponse()
		if test.Response.Parcel.Weight != nil {
			assert.Equal(t, test.Out.Weight, res.(*sendcloud.Parcel).Weight, test.Name)
		} else {
			assert.Equal(t, res.(*sendcloud.Parcel).Weight, "", test.Name)
		}
	}
}
