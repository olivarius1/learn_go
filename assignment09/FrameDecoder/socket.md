# 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

- FixedLengthFramDecoder 固定长度解码器
发送方和接收方约定固定长度的包,不足的用空字符补全

无论一次收到多少消息，它都会按照固定的长度进行解码，如果是半包消息，它会缓存半包消息并且等待下个包到达之后再进行拼包合并，直到读取到一个完整的消息包。

有内存浪费


- delimiter based
用分割符(定界符)判断
FTP协议的LineBasedFrameDecoder和DelimiterBasedFrameDecoder

需要额外查找分隔符

- LengthFieldBasedFrameDecoder
"包长协议"
将消息分为消息头和消息体，消息头中包含表示消息总长度（或者消息体长度）的字段，通常设计思路为消息头的第一个字段使用int32来表示消息的总长度.

