package main

import (
	"fmt"
	"strconv"

	"github.com/MasterMeng/calc"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcs, args := stub.GetFunctionAndParameters()

	if funcs == "add" {
		return s.add(stub, args)
	}
	return shim.Error("Unknown function")
}

func (s *SmartContract) add(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])

	c := calc.Add(a, b)
	return shim.Success([]byte(strconv.Itoa(c)))
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}

	// a := calc.Add(1, 3)
	// fmt.Println(a)
}
