package main

const (
	CommandConn   = iota + 0x01 // 0x01，连接请求包
	CommandSubmit               // 0x02，消息请求包
)

const (
	CommandConnAck   = iota + 0x81 // 0x81，连接请求的响应包
	CommandSubmitAck               // 0x82，消息请求的响应包
)

func main() {
	println(CommandConn)
	println(CommandSubmit)

	println(CommandConnAck)
	println(CommandSubmitAck)
}
