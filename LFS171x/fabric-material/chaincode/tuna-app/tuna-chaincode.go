// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
*/

package main

/* Imports
* 4 utility libraries for handling bytes, reading and writing JSON,
formatting, and string manipulation
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Tuna structure, with 4 properties.
Structure tags are used by encoding/json library
*/
// type Tuna struct {
// 	Vessel string `json:"vessel"`
// 	Timestamp string `json:"timestamp"`
// 	Location  string `json:"location"`
// 	Holder  string `json:"holder"`
// }

type Product struct {
	Producing_area string  `json:"producing_area"`
	Product_type   string  `json:"product_type"`
	Product_unit   string  `json:"product_unit"`
	Unit_Price     float32 `json:"unit_price"`
	Quantity       float32 `json:"quantity"`
	Timestamp      string  `json:"timestamp"`
	Holder         string  `json:"holder"`
}

/*
  * The Init method *
  called when the Smart Contract "tuna-chaincode" is instantiated by the network
  * Best practice is to have any Ledger initialization in separate function
  -- see initLedger()
*/
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
  * The Invoke method *
  called when an application requests to run the Smart Contract "tuna-chaincode"
  The app also specifies the specific smart contract function to call with args
*/
// func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

// 	// Retrieve the requested Smart Contract function and arguments
// 	function, args := APIstub.GetFunctionAndParameters()
// 	// Route to the appropriate handler function to interact with the ledger
// 	if function == "queryTuna" {
// 		return s.queryTuna(APIstub, args)
// 	} else if function == "initLedger" {
// 		return s.initLedger(APIstub)
// 	} else if function == "recordTuna" {
// 		return s.recordTuna(APIstub, args)
// 	} else if function == "queryAllTuna" {
// 		return s.queryAllTuna(APIstub)
// 	} else if function == "changeTunaHolder" {
// 		return s.changeTunaHolder(APIstub, args)
// 	}

// 	return shim.Error("Invalid Smart Contract function name.")
// }

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryProduct" {
		return s.queryProduct(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordProduct" {
		return s.recordProduct(APIstub, args)
	} else if function == "queryAllProduct" {
		return s.queryAllProduct(APIstub)
	} else if function == "changeProductHolder" {
		return s.changeProductHolder(APIstub, args)
	} else if function == "getQueryResultForQueryString" {
		query := "{\"selector\":{\"product_type\":\"" + args[0] + "\"}}"
		return s.getQueryResultForQueryString(APIstub, query)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
  * The queryTuna method *
 Used to view the records of one particular tuna
 It takes one argument -- the key for the tuna in question
*/
// func (s *SmartContract) queryTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}

// 	tunaAsBytes, _ := APIstub.GetState(args[0])
// 	if tunaAsBytes == nil {
// 		return shim.Error("Could not locate tuna")
// 	}
// 	return shim.Success(tunaAsBytes)
// }

func (s *SmartContract) queryProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	if productAsBytes == nil {
		return shim.Error("Could not locate product")
	}
	return shim.Success(productAsBytes)
}

/*
  * The initLedger method *
 Will add test data (10 tuna catches)to our network
*/
// func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
// 	tuna := []Tuna{
// 		Tuna{Vessel: "923F", Location: "67.0006, -70.5476", Timestamp: "1504054225", Holder: "Miriam"},
// 		Tuna{Vessel: "M83T", Location: "91.2395, -49.4594", Timestamp: "1504057825", Holder: "Dave"},
// 		Tuna{Vessel: "T012", Location: "58.0148, 59.01391", Timestamp: "1493517025", Holder: "Igor"},
// 		Tuna{Vessel: "P490", Location: "-45.0945, 0.7949", Timestamp: "1496105425", Holder: "Amalea"},
// 		Tuna{Vessel: "S439", Location: "-107.6043, 19.5003", Timestamp: "1493512301", Holder: "Rafa"},
// 		Tuna{Vessel: "J205", Location: "-155.2304, -15.8723", Timestamp: "1494117101", Holder: "Shen"},
// 		Tuna{Vessel: "S22L", Location: "103.8842, 22.1277", Timestamp: "1496104301", Holder: "Leila"},
// 		Tuna{Vessel: "EI89", Location: "-132.3207, -34.0983", Timestamp: "1485066691", Holder: "Yuan"},
// 		Tuna{Vessel: "129R", Location: "153.0054, 12.6429", Timestamp: "1485153091", Holder: "Carlo"},
// 		Tuna{Vessel: "49W4", Location: "51.9435, 8.2735", Timestamp: "1487745091", Holder: "Fatima"},
// 	}

// 	i := 0
// 	for i < len(tuna) {
// 		fmt.Println("i is ", i)
// 		tunaAsBytes, _ := json.Marshal(tuna[i])
// 		APIstub.PutState(strconv.Itoa(i+1), tunaAsBytes)
// 		fmt.Println("Added", tuna[i])
// 		i = i + 1
// 	}

// 	return shim.Success(nil)
// }

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	product := []Product{
		Product{Producing_area: "남이면", Product_type: "옥수수", Product_unit: "kg", Unit_Price: 10000, Quantity: 5, Timestamp: "1504054225", Holder: "홍길동"},
		Product{Producing_area: "문의면", Product_type: "옥수수", Product_unit: "kg", Unit_Price: 11000, Quantity: 5, Timestamp: "1493517025", Holder: "김문의"},
		Product{Producing_area: "가덕면", Product_type: "감자", Product_unit: "kg", Unit_Price: 3000, Quantity: 10, Timestamp: "1496104301", Holder: "이가덕"},
		Product{Producing_area: "미원면", Product_type: "옥수수", Product_unit: "ea", Unit_Price: 3960, Quantity: 10, Timestamp: "1493512301", Holder: "박미원"},
		Product{Producing_area: "남이면", Product_type: "감자", Product_unit: "g", Unit_Price: 45, Quantity: 500, Timestamp: "1504057825", Holder: "김철수"},
		Product{Producing_area: "북이면", Product_type: "감자", Product_unit: "ea", Unit_Price: 400, Quantity: 10, Timestamp: "1487745091", Holder: "김북이"},
		Product{Producing_area: "문의면", Product_type: "옥수수", Product_unit: "kg", Unit_Price: 6500, Quantity: 3, Timestamp: "1494117101", Holder: "이청주"},
		Product{Producing_area: "강내면", Product_type: "양파", Product_unit: "kg", Unit_Price: 1800, Quantity: 1, Timestamp: "1485153091", Holder: "강길동"},
		Product{Producing_area: "강내면", Product_type: "양파", Product_unit: "g", Unit_Price: 20, Quantity: 1500, Timestamp: "1496105425", Holder: "강흥덕"},
		Product{Producing_area: "남이면", Product_type: "옥수수", Product_unit: "kg", Unit_Price: 9500, Quantity: 4, Timestamp: "1485066691", Holder: "홍길동"},
	}

	i := 0
	for i < len(product) {
		fmt.Println("i is ", i)
		productAsBytes, _ := json.Marshal(product[i])
		APIstub.PutState(strconv.Itoa(i+1), productAsBytes)
		fmt.Println("Added", product[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
  * The recordTuna method *
 Fisherman like Sarah would use to record each of her tuna catches.
 This method takes in five arguments (attributes to be saved in the ledger).
*/
// func (s *SmartContract) recordTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 5 {
// 		return shim.Error("Incorrect number of arguments. Expecting 5")
// 	}

// 	var tuna = Tuna{ Vessel: args[1], Location: args[2], Timestamp: args[3], Holder: args[4] }

// 	tunaAsBytes, _ := json.Marshal(tuna)
// 	err := APIstub.PutState(args[0], tunaAsBytes)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("Failed to record tuna catch: %s", args[0]))
// 	}

// 	return shim.Success(nil)
// }

func (s *SmartContract) recordProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	unitPrice, _ := strconv.ParseFloat(args[4], 32)
	quantity, _ := strconv.ParseFloat(args[5], 32)

	var product = Product{
		Producing_area: args[1],
		Product_type:   args[2],
		Product_unit:   args[3],
		Unit_Price:     float32(unitPrice),
		Quantity:       float32(quantity),
		Timestamp:      args[6],
		Holder:         args[7]}

	productAsBytes, _ := json.Marshal(product)
	err := APIstub.PutState(args[0], productAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record product harvest: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
  * The queryAllTuna method *
 allows for assessing all the records added to the ledger(all tuna catches)
 This method does not take any arguments. Returns JSON string containing results.
*/
// func (s *SmartContract) queryAllTuna(APIstub shim.ChaincodeStubInterface) sc.Response {

// 	startKey := "0"
// 	endKey := "999"

// 	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	defer resultsIterator.Close()

// 	// buffer is a JSON array containing QueryResults
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	bArrayMemberAlreadyWritten := false
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		// Add comma before array members,suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"Key\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(queryResponse.Key)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Record\":")
// 		// Record is a JSON object, so we write as-is
// 		buffer.WriteString(string(queryResponse.Value))
// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 	}
// 	buffer.WriteString("]")

// 	fmt.Printf("- queryAllTuna:\n%s\n", buffer.String())

// 	return shim.Success(buffer.Bytes())
// }

func (s *SmartContract) getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) sc.Response {
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
	resultsIterator, err := APIstub.GetQueryResult(queryString)
	defer resultsIterator.Close()
	if err != nil {
		return shim.Error(err.Error())
	}
	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse,
			err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllProduct(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllProduct:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
  * The changeTunaHolder method *
 The data in the world state can be updated with who has possession.
 This function takes in 2 arguments, tuna id and new holder name.
*/
// func (s *SmartContract) changeTunaHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}

// 	tunaAsBytes, _ := APIstub.GetState(args[0])
// 	if tunaAsBytes == nil {
// 		return shim.Error("Could not locate tuna")
// 	}
// 	tuna := Tuna{}

// 	json.Unmarshal(tunaAsBytes, &tuna)
// 	// Normally check that the specified argument is a valid holder of tuna
// 	// we are skipping this check for this example
// 	tuna.Holder = args[1]

// 	tunaAsBytes, _ = json.Marshal(tuna)
// 	err := APIstub.PutState(args[0], tunaAsBytes)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("Failed to change tuna holder: %s", args[0]))
// 	}

// 	return shim.Success(nil)
// }

func (s *SmartContract) changeProductHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	if productAsBytes == nil {
		return shim.Error("Could not locate product")
	}
	product := Product{}

	json.Unmarshal(productAsBytes, &product)
	// Normally check that the specified argument is a valid holder of product
	// we are skipping this check for this example
	product.Holder = args[1]

	productAsBytes, _ = json.Marshal(product)
	err := APIstub.PutState(args[0], productAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change product holder: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
  * main function *
 calls the Start function
 The main function starts the chaincode in the container during instantiation.
*/
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
