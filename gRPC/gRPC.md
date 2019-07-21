### Теория

gRPC — это высокопроизводительный фреймворк разработанный компанией Google для вызов удаленных процедур (RPC), работает поверх **HTTP/2**. Для сериализации использует **Protocol Buffers**.

###### Protocol Buffers

В общем виде формат представляет из себя закодированную последовательность полей, состоящих из ключа и значения. В качестве ключа выступает номер, определённый для каждого поля сообщения в proto-файле.
##### Типы RPC

- Унарный (Unary RPC). 
Синхронный запрос клиента, который блокируются пока не будет получен ответ от сервера.
`rpc SayHello(HelloRequest) returns (HelloResponse){}`
- Серверный стрим (Server streaming RPC) 
при подключении клиента сервер открывает стрим и начинает отправлять сообщения.
`rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){}`
- Клиентский стрим (Client streaming RPC).
 То же самое, что и серверный, только клиент начинает стримить сообщения на сервер.
`rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {}`
- Двунаправленный стрим (Bidirectional streaming). 
Клиент инициализирует соединение, создаются два стрима. Сервер может отправить изначальные данные при подключении или отвечать на каждый запрос клиента по типу “пинг-понга”.
`rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){}`


##### Аутентификация 
Поддерживает два способа SSL/TLS и Token-based authentication with Google

##### Ошибки и их обработка
Списо ошибок который являються стандартными для go [здесь](https://github.com/grpc/grpc/blob/master/doc/statuscodes.md)

##### Синхронные и асинхронные вызовы

### Практика

Создание клиентского и серверного кода

`protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide`

