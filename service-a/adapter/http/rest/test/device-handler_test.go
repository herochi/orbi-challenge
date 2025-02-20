package rest

/*import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"finantia/application/device/ports/dto"
	"finantia/application/device/ports/viewmodel"
	"finantia/config"
	ds "finantia/infrastructure/datastore"
	"finantia/infrastructure/http/server"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)*/

/*var (
	deviceServiceMock = new(mock.DeviceServiceMock)
	deviceHandler     = rest.NewDeviceHandler(deviceServiceMock)
)*/

/*func TestGet(t *testing.T) {
	config.Init()

	db, err := ds.NewMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	g := gin.Default()
	testServer := server.NewServer(g, db)
	testServer.MapRoutes()
	testRouter := testServer.Router

	req, err := http.NewRequest("GET", "http://localhost/api/v1/devices?countryCode=MX", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestGetByClientId(t *testing.T) {
	config.Init()

	db, err := ds.NewMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	g := gin.Default()
	testServer := server.NewServer(g, db)
	testServer.MapRoutes()
	testRouter := testServer.Router

	req, err := http.NewRequest("GET", "http://localhost/api/v1/clients/5ee018f612488c1c37613951/devices", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestPatch(t *testing.T) {
	config.Init()

	db, err := ds.NewMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	devicePatch := dto.DevicePatch{}
	devicePatch.Imei = "520002735123"
	isActive := false
	devicePatch.IsActive = &isActive

	requestByte, _ := json.Marshal(devicePatch)
	requestReader := bytes.NewReader(requestByte)

	g := gin.Default()
	testServer := server.NewServer(g, db)
	testServer.MapRoutes()
	testRouter := testServer.Router

	req, err := http.NewRequest("PATCH", "http://localhost/api/v1/devices/5fc95548a786b86e464e2496", requestReader)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestCreate(t *testing.T) {
	config.Init()

	db, err := ds.NewMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	device := viewmodel.DeviceVM{}
	device.Imei = "52009991"
	device.Phone = "9991111"
	device.Iccid = "89011703278460072112"
	device.Brand = "ST4340"
	clientOId, err := primitive.ObjectIDFromHex("5f1d06d1420d9d46f0dee811")
	device.ClientID = clientOId
	device.SiafiRef = "5f1d06d1420d9d46f0dee811"
	device.CountryCode = "US"
	device.RentPrice = 25
	device.Settings = viewmodel.DeviceSettings{}
	device.Settings.Name = "ST4340-520026890"
	device.InstalledAt = time.Now()

	requestByte, _ := json.Marshal(device)
	requestReader := bytes.NewReader(requestByte)

	g := gin.Default()
	testServer := server.NewServer(g, db)
	testServer.MapRoutes()
	testRouter := testServer.Router

	req, err := http.NewRequest("POST", "http://localhost/api/v1/devices", requestReader)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 201, resp.Code)
}*/
