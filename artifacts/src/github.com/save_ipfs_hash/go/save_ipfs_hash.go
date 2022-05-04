package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"

)

type SmartContract struct {
}

type Item struct{
	Nft_id string `json:"nft_id"`
	Balance string `json:"balance"`
	Timestamp string `json:"timestamp"`
	TokenURI string `json:"tokenuri"`
	Token_id string `json:"token_id"`
	Minter string `json:"minter"`
	Fcn string `json:"fcn"`
	Title string `json:"title"`
	Category string `json:"category"`
	Nft_num string `json:"nft_num"`
	Nft_status string `json:"nft_status"`
}

type Item2 struct{
	Nft_read_id string `json:"nft_read_id"`
	Timestamp string `json:"timestamp"`
	Fcn string `json:"fcn"`
	Nft_id string `json:"nft_id"`
}

type Item3 struct{
	Nft_id string `json:"nft_id"`
	Timestamp string `json:"timestamp"`
	Fcn string `json:"fcn"`
	Nft_status string `json:"nft_status"`
}



func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("save_ipfs_hash_cc")

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %d", function)
	logger.Infof("Args length is : %d", len(args))

	switch function {
	case "query_item":
		return s.query_item(APIstub, args)
	case "initLedger":
		return s.initLedger(APIstub)
	case "Register_NFT":
		return s.Register_NFT(APIstub, args)
	case "Read_NFT":
		return s.Read_NFT(APIstub, args)
	case "BBBurn_NFT":
		return s.BBBurn_NFT(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}
}


func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response{

	items := []Item{
		Item{ Nft_id:"11", Balance: "22",Timestamp:"33",TokenURI:"44",Token_id:"55",Minter:"66",Fcn:"77",Title:"88",Category:"99",Nft_num:"10",Nft_status:"11" },
		Item{Nft_id:"11", Balance: "22",Timestamp:"33",TokenURI:"44",Token_id:"55",Minter:"66",Fcn:"77",Title:"88",Category:"99",Nft_num:"10",Nft_status:"11" },
	}

	items2 := []Item2{
		Item2{Nft_read_id:"11", Timestamp:"22", Fcn:"33",Nft_id:"44" },
		Item2{Nft_read_id:"11", Timestamp:"22", Fcn:"33",Nft_id:"44" },
	}

	items3 := []Item3{
		Item3{Nft_id:"11", Timestamp:"22", Fcn:"33",Nft_status:"44" },
		Item3{Nft_id:"11", Timestamp:"22", Fcn:"33",Nft_status:"44" },
	}

/*	create_certs := []Create_cert{
		Create_cert{Cid: "1111", Org_id:"2222",Create_timestamp:"3333",Expire_timestamp:"44444",Cert_name:"555",Status: struct {
			To string `json:"to"`
			From string `json:"from"`
			Requester string `json:"requester"`
			Status string `json:"status"`
			Timestamp string `json:"timestamp"`
		}{To:"111",
			From:"222",
			Requester:"333",
			Status:"223",
			Timestamp:"555",

			}, },
	}*/



	i := 0
	for i < len(items){
		itemAsBytes, _ :=json.Marshal(items[i])
		APIstub.PutState("ITEM"+strconv.Itoa(i),itemAsBytes)
		i = i + 1
	}

	j := 0
	for j < len(items2){
                itemAsBytes2, _ :=json.Marshal(items2[j])
                APIstub.PutState("ITEM"+strconv.Itoa(j),itemAsBytes2)
                j = j + 1
        }

	k := 0
	for k < len(items3){
		itemAsBytes3, _ :=json.Marshal(items3[k])
		APIstub.PutState("ITEM"+strconv.Itoa(k),itemAsBytes3)
		k = k + 1
	}


	return shim.Success(nil)
}


func (s *SmartContract) Register_NFT(APIstub shim.ChaincodeStubInterface, args[]string) sc.Response{

        var item = Item{Nft_id:args[6], Balance: args[1], Timestamp:args[2], TokenURI:args[3],Token_id:args[4],Minter:args[5],Fcn:args[0],Title:args[7],Category:args[8],Nft_num:args[9],Nft_status:args[10]}
        itemAsBytes, _ := json.Marshal(item)
        APIstub.PutState(args[0], itemAsBytes)

        indexName := "owner~key"
        colorNameIndexKey, err :=APIstub.CreateCompositeKey(indexName, []string{item.Nft_id, args[0]})
        if err != nil {
                return shim.Error(err.Error())
        }
        value := []byte{0x00}
        APIstub.PutState(colorNameIndexKey, value)
        return shim.Success(itemAsBytes)
}

func (s *SmartContract) Read_NFT(APIstub shim.ChaincodeStubInterface, args[]string) sc.Response{

        var item2 = Item2{Nft_read_id:args[0], Timestamp:args[1], Fcn:args[2],Nft_id:args[3]}
        itemAsBytes2, _ := json.Marshal(item2)
        APIstub.PutState(args[0], itemAsBytes2)

        indexName := "owner~key"
        colorNameIndexKey, err :=APIstub.CreateCompositeKey(indexName, []string{item2.Nft_read_id, args[0]})
        if err != nil {
                return shim.Error(err.Error())
        }
        value := []byte{0x00}
        APIstub.PutState(colorNameIndexKey, value)
        return shim.Success(itemAsBytes2)
}

func (s *SmartContract) BBBurn_NFT(APIstub shim.ChaincodeStubInterface, args[]string) sc.Response{

        var item3 = Item3{Nft_id:args[0], Timestamp:args[1], Fcn:args[2],Nft_status:args[3]}
        itemAsBytes3, _ := json.Marshal(item3)
        APIstub.PutState(args[0], itemAsBytes3)

        indexName := "owner~key"
        colorNameIndexKey, err :=APIstub.CreateCompositeKey(indexName, []string{item3.Nft_id, args[0]})
        if err != nil {
                return shim.Error(err.Error())
        }
        value := []byte{0x00}
        APIstub.PutState(colorNameIndexKey, value)
        return shim.Success(itemAsBytes3)
}

func (s *SmartContract) query_item(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	itemAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(itemAsBytes)
}

/*func (s *SmartContract) createcert(APIstub shim.ChaincodeStubInterface, args[]string) sc.Response{
	if len(args) !=11 {
		return shim.Error("Incorrect number of arguments. Expecting 5?? maybe 4")
	}

	var create_cert = Create_cert{Cid: args[1], Org_id:args[2], Create_timestamp:args[3],Expire_timestamp:args[4],Cert_name:args[5],Status:struct{
		To string `json:"to"`
		From string `json:"from"`
		Requester string `json:"requester"`
		Status string `json:"status"`
		Timestamp string `json:"timestamp"`
		}{
			To:args[6],
			From:args[7],
			Requester:args[8],
			Status:args[9],
			Timestamp:args[10],
			},}
	create_certAsBytes, _ := json.Marshal(create_cert)
	APIstub.PutState(args[0], create_certAsBytes)

	indexName := "owner~key"
	colorNameIndexKey, err :=APIstub.CreateCompositeKey(indexName, []string{create_cert.Cid, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	APIstub.PutState(colorNameIndexKey, value)
	return shim.Success(create_certAsBytes)
}*/

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
