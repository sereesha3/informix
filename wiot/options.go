package wiot

import (
    "fmt"
    "strings"
)

func NewOptions(org, typ, id, tok string) *Options {
    if anyEmpty(org, typ, id, tok) {
        return nil
    }

    o := &Options{}

    o.Broker = fmt.Sprintf(
        "tcp://%s.messaging.internetofthings.ibmcloud.com:1883",
        org,
    )

    o.ClientId = fmt.Sprintf("g:%s:%s:%s", org, typ, id)
    o.Username = "use-token-auth"
    o.Password = tok

    return o
}

// anyEmpty returns true if any of the provided args have length zero.
func anyEmpty(args ...string) bool {
    for _, arg := range args {
        if len(arg) == 0 {
            return true
        }
    }

    return false
}

// Options for creating a new Watson IoT Platform MQTT client.
type Options struct {
    Broker      string
    ClientId    string
    Username    string
    Password    string
}

func (o *Options) DeviceType() string {
    return strings.Split(o.ClientId, ":")[2]
}

func (o *Options) DeviceId() string {
    return strings.Split(o.ClientId, ":")[3]
}
