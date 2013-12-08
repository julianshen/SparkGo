package SparkGo

import (
    "net/http"
    "net/url"
    "encoding/json"
    "log"
    "strconv"
)

type SparkCore struct {
    Access_token string
    Device_id string
}

type TinkerResult struct {
    Id string `json:id`
    Name string `json:name`
    Connected bool `json:connected`
    ReturnVal int `json:return_value`
}

const (
    END_POINT_URL="https://api.spark.io/v1/devices/"
)

func (core *SparkCore)post(command string, params string) (*TinkerResult, error) {
    api_url := END_POINT_URL + core.Device_id + "/" + command

    values := make(url.Values)
    values.Set("access_token", core.Access_token)
    values.Set("params", params)

    r, err := http.PostForm(api_url, values);

    if err != nil {
         log.Printf("error: %s", err)
         return nil, err
    }

    var result TinkerResult
    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(&result)

    if err != nil {
         return nil,err
    }

    return &result, nil
}

func (core *SparkCore)DigitalWrite(pin string, val bool) (int, error) {
    params := pin + ","

    if val {
        params+="HIGH"
    } else {
        params+="LOW"
    }

    r,err := core.post("digitalwrite", params)

    if err!= nil {
        return -1, err
    }

    return r.ReturnVal, nil
}

func (core *SparkCore)DigitalRead(pin string) (int, error) {
    params := pin

    r,err := core.post("digitalread", params)

    if err!= nil {
        return -1, err
    }

    return r.ReturnVal, nil
}

func (core *SparkCore)AnalogWrite(pin string, val int) (int, error) {
    params := pin + "," + strconv.Itoa(val)

    r,err := core.post("analogwrite", params)

    if err!= nil {
        return -1, err
    }

    return r.ReturnVal, nil
}

func (core *SparkCore)AnalogRead(pin string) (int, error) {
    params := pin

    r,err := core.post("analogread", params)

    if err!= nil {
        return -1, err
    }

    return r.ReturnVal, nil
}


