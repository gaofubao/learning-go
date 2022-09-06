package main

import "fmt"

func main() {
	msg := "[caller_once]"

	// 解码
	decodeOp := decode(msg)
	decodedMsg, err := decodeOp(msg)
	if err != nil {
		return
	}

	// 执行业务逻辑
	actionOp := opAction(decodedMsg)
	opRes, err := actionOp(decodedMsg)
	if err != nil {
		return
	}

	// 编码
	encodeOp := encode(opRes)
	_, err = encodeOp(opRes)
	if err != nil {
		return
	}
}

// Op 定义对msg处理的函数集合
type Op func(msg any) (any, error)

// 解码远程消息
func decode(msg any) Op {
	return func(any) (any, error) {
		fmt.Printf("decoding ... %v\n", msg)
		decodedRes := fmt.Sprintf("decoded_%v", msg)
		fmt.Printf("decoded to parameter -> %v\n", decodedRes)
		return decodedRes, nil
	}
}

// 模拟服务提供方的处理逻辑
func opAction(parameter any) Op {
	return func(any) (any, error) {
		fmt.Printf("do opAction ... %v\n", parameter)
		opRes := fmt.Sprintf("opAction_%v", parameter)
		fmt.Printf("after opAction result -> %v\n", opRes)
		return opRes, nil
	}
}

// 编码处理结果
func encode(result any) Op {
	return func(any) (any, error) {
		fmt.Printf("encoding ... %v\n", result)
		encodedRes := fmt.Sprintf("encoded_%v", result)
		fmt.Printf("after encoded result -> %v\n", encodedRes)
		return encodedRes, nil
	}
}
