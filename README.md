## RESTful microservices written in Go standard library. 

Folder product corresponds the data. This CRUD application operates on product type. 


Overal structure of writing microservices:

1) Get Data
2) Encode and decode accordingly. 
3) Write handler for the data. Handler required to implement Handler interface. 
4) Create multiplixer and add each URL.
5) Start the server.