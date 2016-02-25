package securepassctl

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	testInstance SecurePass
	appLabel     string
)

func init() {
	testInstance = SecurePass{
		AppID:     "ce64dc90d88b11e5b001de2f4665c1f2@ci.secure-pass.net",
		AppSecret: "E2m6HawI743as61Kv0OhyPb6wAewXnwVkLLcF82rKOWe1SJ0Wd",
		Endpoint:  DefaultRemote,
	}
	appLabel = fmt.Sprintf("test_fixture_%d_%d", time.Now().Unix(), rand.Int())
}

func ExampleSecurePass() {
	fmt.Println(testInstance.AppID)
	fmt.Println(testInstance.AppSecret)
	fmt.Println(testInstance.Endpoint)
	// Output:
	// ce64dc90d88b11e5b001de2f4665c1f2@ci.secure-pass.net
	// E2m6HawI743as61Kv0OhyPb6wAewXnwVkLLcF82rKOWe1SJ0Wd
	// https://beta.secure-pass.net
}

func ExampleSecurePass_Ping() {
	resp, err := testInstance.Ping()
	fmt.Println(err)
	fmt.Println(resp.IPVersion)
	fmt.Println(resp.ErrorCode())
	fmt.Println(resp.ErrorMessage())
	// Output:
	// <nil>
	// 4
	// 0
	//
}

func ExampleSecurePass_AppAdd() {
	var (
		resp         APIResponse
		addResponse  *AppAddResponse
		infoResponse *AppInfoResponse
		fixtureAppID string
	)

	// Create a new app
	addResponse, _ = testInstance.AppAdd(&ApplicationDescriptor{
		Label: appLabel,
	})
	fixtureAppID = addResponse.AppID
	fmt.Println(addResponse.ErrorCode())
	fmt.Println(addResponse.ErrorMessage() == "")
	// Check for its existence
	resp, _ = testInstance.AppInfo(fixtureAppID)
	fmt.Println(resp.ErrorCode())
	// Modify it
	resp, _ = testInstance.AppMod(fixtureAppID, &ApplicationDescriptor{
		Write:   false,
		Label:   appLabel + "newLabel",
		Privacy: true,
	})
	fmt.Println(resp.ErrorCode())
	// Check whether the modifcations have been applied
	infoResponse, _ = testInstance.AppInfo(fixtureAppID)
	fmt.Println(infoResponse.Label == appLabel+"newLabel")
	// Remove it
	resp, _ = testInstance.AppDel(fixtureAppID)
	fmt.Println(resp.ErrorCode())
	// Check whether it does not longer exist
	resp, _ = testInstance.AppInfo(fixtureAppID)
	fmt.Println(resp.ErrorCode())
	// Output:
	// 0
	// true
	// 0
	// 0
	// true
	// 0
	// 10
}

func ExampleSecurePass_AppList() {
	var resp APIResponse
	resp, err := testInstance.AppList("")
	fmt.Println(resp.ErrorCode())
	fmt.Println(err)
	// Output:
	// 0
	// <nil>
}
