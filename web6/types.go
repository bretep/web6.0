package web6

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type UdnData struct {
	WebProtocolAction string
	Data              map[string]interface{}
	Temp              map[string]interface{}
	Output            map[string]interface{} // Staging output goes here, can share them with appending as well.
	Arg               map[string]interface{} // Every function call blows this away, and sets the args in it's data, so it's accessable
	FunctionArg       map[string]interface{} // Function arguments, from Stored UDN Function __function, sets the incoming function args
	Page              map[string]interface{} //TODO(g):NAMING: __widget is access here, and not from "widget", this can be changed, since thats what it is...
	SetApiInput       map[string]interface{}
	SetApiResult      map[string]interface{} // If this is an API call, set values in here, which will be encoded in JSON and sent back to the client on return
	SetCookies        map[string]interface{} // Set Cookies.  Any data set in here goes into a cookie.  Will use standard expiration and domain for now.
	SetHeaders        map[string]interface{} // Set HTTP Headers.
	SetHttpOptions    map[string]interface{} // Any other things we want to control from UDN, we put in here to be processed.  Can be anything, not based on a specific standard.
	Param             map[string]interface{}
	ApiInput          map[string]interface{}
	JsonBody          map[string]interface{}
	Cookies           map[string]interface{}
	Headers           map[string]interface{}
	Session           map[string]interface{}
	User              map[string]interface{}
	UserData          map[string]interface{}
	WebSite           map[string]interface{}
	BaseWidget        map[string]interface{}
	SchemaUDN         map[string]interface{}
	WebSitePage       WebSiteAPI
	UUID              string
	Uri               string

	HttpResponseCode int
}

func NewUdnData() UdnData {
	udn_data := UdnData{}
	udn_data.HttpResponseCode = http.StatusOK
	udn_data.WebSitePage.AllowLogging = true
	return udn_data
}

type WebSiteAPI struct {
	AllowLogging bool
}

type WebSitePage struct {
	Id int
}

func (s *WebSitePage) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
