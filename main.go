package main

import (
	"fmt"
	"os"
	muxgo "github.com/muxinc/mux-go"
)

func main(){
	// API Client
	client:=muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"),os.Getenv("MUX_TOKEN_SECRET")),
	))

	// Create the Asset
	asset, err:=client.AssetsApi.CreateAsset(muxgo.CreateAssetRequest{
		Input:[]muxgo.InputSettings{
			muxgo.InputSettings{
				Url:
