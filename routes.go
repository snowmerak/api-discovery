package api_discovery

import "encoding/json"

type ApiSpec struct {
	Route  string
	Method string
}

func NewApiSpec(route string, method string) *ApiSpec {
	return &ApiSpec{
		Route:  route,
		Method: method,
	}
}

type ApiMap struct {
	m map[string]*ApiSpec
}

func NewApiMap() *ApiMap {
	return &ApiMap{
		m: make(map[string]*ApiSpec),
	}
}

func (am *ApiMap) AddApiSpec(route string, method string) {
	apiSpec := NewApiSpec(route, method)
	am.m[route] = apiSpec
}

func (am *ApiMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(am.m)
}

func (am *ApiMap) UnmarshalJSON(bytes []byte) error {
	var apiSpecs map[string]*ApiSpec
	if err := json.Unmarshal(bytes, &apiSpecs); err != nil {
		return err
	}
	am.m = apiSpecs
	return nil
}
