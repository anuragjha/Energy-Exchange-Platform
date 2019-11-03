package uri_router

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Start",
		"GET",
		"/start",
		Start,
	},
	Route{
		"Register",
		"POST",
		"/register",
		Register,
	},
	Route{
		"GetAllDevices",
		"GET",
		"/getalldevices",
		GetAllDevices,

	},
	Route{
		"DeleteDevice",
		"DELETE",
		"/device/{deviceId}",
		DeleteDevice,
		//	localhost:48081/api/v1/device/id/ce13abf3-fd29-453b-9707-df679cbb60a5
	},
	Route{
		"SwitchButton",
		"PUT",
		"/api/v1/device/{deviceId}/SwitchButton",
		SwitchButton,
		//	localhost:48081/api/v1/device/id/ce13abf3-fd29-453b-9707-df679cbb60a5
	},

	Route{
		"AllDeviceProfiles",
		"GET",
		"/alldeviceprofiles",
		GetAllDeviceProfiles,
		//	localhost:48081/api/v1/deviceprofile
	},
	Route{
		"DeleteDeviceProfile",
		"DELETE",
		"/deviceprofile/{deviceId}",
		DeleteDeviceProfile,
		//	localhost:48081/api/v1/deviceprofile/id/33793ade-9e34-4e24-a8a1-7936ba693f7a
	},
	Route{
		"ReadDeviceData",
		"GET",
		"/readDeviceData/{deviceName}/{noOfReadings}",
		ReadDeviceData,
		//	http://localhost:6686/readDeviceData/Supply-Device-01/10
		//  http://localhost:6686/readDeviceData/Consume-Device01/10
	},
	Route{
		"ShowLatestDeviceData",
		"GET",
		"/showLatestDeviceData/{deviceName}/{resourceName}",
		ShowLatestDeviceData,
		//	localhost:6686/showLatestDeviceData/Supply-Device-01/randomsuppliernumber
		//	localhost:6686/showLatestDeviceData/Consume-Device01/randomconsumenumber
	},
	Route{
		"ShowAllLatestDeviceData",
		"GET",
		"/showAllLatestDeviceData",
		ShowAllLatestDeviceData,
		//	localhost:6686/showAllLatestDeviceData
	},
	//Route{
	//	"MakeDecision",
	//	"GET",
	//	"/makeDecision",
	//	MakeDecision,
	//	//	localhost:6686/makeDecision
	//},

	Route{
		"GetTaskManager",
		"GET",
		"/gettaskmanager",
		TaskManagerFrontend,
		//	localhost:6686/gettaskmanager
	},
	Route{
		"GetDeviceFront",
		"GET",
		"/deviceFront",
		DeviceFront,
		//	localhost:6686/gettaskmanager
	},
	Route{
		"PostEvent",
		"POST",
		"/posttaskmanager",
		TaskManagerFrontend,
		//	localhost:6686/posttaskmanager
	},
	Route{
		"MakeDecision",
		"GET",
		"/makeDecision",
		MakeDecision,
		//	localhost:6686/makeDecision
	},
	///////////////////// !!!!!!!!! ///////////////////
	//device manager apis
	///////////////////// !!!!!!!!! ///////////////////
	Route{
		"On",
		"GET",
		"/on",
		On,
		//	localhost:9999/on
	},
	Route{
		"SendDeviceList",
		"GET",
		"/sendDeviceList",
		SendDeviceList,
		//	localhost:9999/sendDeviceList
	},
	Route{
		"SupplierTx",
		"POST",
		"/suppliertx",
		SupplierTx,
		//	localhost:9999/suppliertx
	},
	Route{
		"ConsumerTx",
		"POST",
		"/consumertx",
		ConsumerTx,
		//	localhost:9999/consumertx
	},
}
