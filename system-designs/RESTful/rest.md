REST (Representational State Transfer) is most common communication standard between computers over internet

Set of rules

[HTTP verb (GET/POST)] - [URI] - [Protocol (HTTP 1.1)]

200s - Successful
400s - Wrong input body
500s - Server unavailable

Idempotent 

REST should be stateless (two parties dont need to store info about each other)
Request - Response Cycle -> Independent from each other and easy to scale

Huge Data -> Paginate (using limit/offset, etc)
Versioning of APIs

Other API protocols includes GraphQL and gRPC
