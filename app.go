package main

import (
	"fmt"

	"github.com/apache/plc4x/plc4go/pkg/api"
	"github.com/apache/plc4x/plc4go/pkg/api/drivers"
	"github.com/apache/plc4x/plc4go/pkg/api/transports"
)

func main() {

	driverManager := plc4go.NewPlcDriverManager()
	// Register the Transports
	transports.RegisterTcpTransport(driverManager)
	transports.RegisterUdpTransport(driverManager)

	drivers.RegisterOpcuaDriver(driverManager)

	connectionRequestChanel := driverManager.GetConnection("opcua:tcp://127.0.0.1:62541?discovery=true")
	connectionResult := <-connectionRequestChanel

	if connectionResult.GetErr() != nil {
		fmt.Printf("Error connecting to PLC: %s", connectionResult.GetErr().Error())
		return
	}

	connection := connectionResult.GetConnection()

	defer connection.Close()
}
