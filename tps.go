package main

import (
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"fmt"
	"time"
)

type RstBlockNum struct {
	Jsonrpc string `json:"jsonrpc"`
	Id uint8	`json:"id"`
	Result string	`json:"result"`
}


type Header struct {
	Difficulty  string       `json:"difficulty"`
	Extra       string         `json:"extraData"`
	GasLimit    string       `json:"gasLimit"`
	GasUsed     string       `json:"gasUsed"`
	Hash		string
	Bloom       string          `json:"logsBloom"`
	Coinbase    string `json:"miner"`
	MixDigest   string    `json:"mixHash"          gencodec:"required"`
	Nonce       string     `json:"nonce"            gencodec:"required"`
	Number      string       `json:"number"           gencodec:"required"`
	ParentHash  string    `json:"parentHash"       gencodec:"required"`
	ReceiptHash string    `json:"receiptsRoot"     gencodec:"required"`
	UncleHash   string    `json:"sha3Uncles"       gencodec:"required"`
	Size 		string	   `json:"size"`
	Root        string    `json:"stateRoot"        gencodec:"required"`
	Time        string       `json:"timestamp"        gencodec:"required"`
	Td 			string     `json:"td"`
	Transactions []string  `transactions`
	TxHash      string    `json:"transactionsRoot" gencodec:"required"`
	uncles       []string `uncle hash`
}

type RstBlock struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      uint8  `json:"id"`
	Result  Header `json:"result"`
}

func GetInfo(url string,startBlock,interval,step int64) (number,txnum,timestamp int64){
	sBlockNum := strconv.FormatInt(startBlock,16)
	//fmt.Printf("第%d块 ",startBlock)
	str3:= "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"0x"
	str4 := "\",false],\"id\":1}"
	payload := strings.NewReader(str3+sBlockNum+str4)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "8768224d-db4a-4fff-b48a-30f79e083fe2")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	var rstdata = new(RstBlock)
	if err := json.Unmarshal(body,&rstdata);err != nil{
		panic(err)
	}
	txnum = int64(len(rstdata.Result.Transactions))
	//fmt.Printf("交易个数%d ", txnum)
	timestamp,_ = strconv.ParseInt(string(rstdata.Result.Time), 0, 64)
	//fmt.Printf("出块时间 %d\n",timestamp)
	//fmt.Println(rstdata.Result.Transactions)
	return startBlock,txnum,timestamp
}

func main(){
	startBlock := int64(7000)
	txCount := int64(0)
	//for {
	//	time.Sleep(time.Second)
	for i:=int64(0);i<100;i++{
		_,txnum,_ := GetInfo("http://192.168.3.32:22000",startBlock+i,10,10)
		txCount = txCount+txnum
	}
	//}
	fmt.Println("交易总数",txCount)
	//fmt.Println("块号",num,"交易个数",txnum,"区块时间",time)
}

func main1(){
	url := "http://192.168.3.32:22000"

	payload := strings.NewReader("{\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\",\"params\":[],\"id\":1}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "91cd3040-0f68-418d-9577-8bf373a19004")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))

	var mdata = new(RstBlockNum)
	err := json.Unmarshal(body,&mdata)
	if err != nil{
		panic(err)
	}

	blocknum,_:= strconv.ParseInt(mdata.Result, 0, 64)
	fmt.Printf("第 %d 块 ",blocknum)

	sBlockNum := strconv.FormatInt(blocknum,16)
	str1 := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockTransactionCountByNumber\",\"params\":[\"0x"
	str2 := "\"],\"id\":1}"

	payload = strings.NewReader(str1+sBlockNum+str2)
	req, _ = http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "366f478a-dde8-425b-8cef-bb8588929752")
	res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body,&mdata)
	if err != nil{
		panic(err)
	}

	txnum0,_:= strconv.ParseInt(mdata.Result, 0, 64)
	fmt.Printf("共 %d 交易 ",txnum0)

	str3:= "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"0x"
	str4 := "\",false],\"id\":1}"
	payload = strings.NewReader(str3+sBlockNum+str4)
	req, _ = http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "8768224d-db4a-4fff-b48a-30f79e083fe2")
	res, _ = http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	var rstdata = new(RstBlock)
	err = json.Unmarshal(body,&rstdata)
	if err != nil{
		panic(err)
	}
	mtime0,_ := strconv.ParseInt(string(rstdata.Result.Time), 0, 64)
	fmt.Printf("出块时间 %d\n",mtime0)
	//////////////////////////////////////////////
	for {
		time.Sleep(time.Second * 500)
		blocknum = blocknum + 1
		fmt.Printf("第 %d 块 ",blocknum)

		sBlockNum = strconv.FormatInt(blocknum,16)
		payload = strings.NewReader(str1+sBlockNum+str2)
		req, _ = http.NewRequest("POST", url, payload)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Cache-Control", "no-cache")
		req.Header.Add("Postman-Token", "366f478a-dde8-425b-8cef-bb8588929752")
		res, _ = http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body,&mdata)
		if err != nil{
			panic(err)
		}

		txnum1,_:= strconv.ParseInt(mdata.Result, 0, 64)
		fmt.Printf("共 %d 交易 ",txnum1)

		payload = strings.NewReader(str3+sBlockNum+str4)
		req, _ = http.NewRequest("POST", url, payload)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Cache-Control", "no-cache")
		req.Header.Add("Postman-Token", "8768224d-db4a-4fff-b48a-30f79e083fe2")
		res, _ = http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)
		rstdata = new(RstBlock)
		err = json.Unmarshal(body,&rstdata)
		if err != nil{
			panic(err)
		}
		mtime1,_ := strconv.ParseInt(string(rstdata.Result.Time), 0, 64)
		fmt.Printf("出块时间 %d",mtime1)
		fmt.Printf("\n 交易速度 %d TPS\n\n", txnum0*1000000000/(mtime1-mtime0) )
		mtime0 = mtime1
	}
}
