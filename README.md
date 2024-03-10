# chat-bot-data
 
 #### To Run this service we first need to run the Docker use the below command to run the docker 
 docker run --name chatbot -e POSTGRES_PASSWORD=aditya -d postgres

This service uses postgres Database to store the data 

and it contains the following Endpoints


#### user endpoints:

/user/login:

- GET Method 
- this takes 2 query param which are required params
  - Email
  - Password 
- It returns the data if it is present otherwise return empty data which means the user is not registered
<hr/>

/user/signup:
- POST Method 
- takes request body
   - firstName
   - lastName
   - dob
   - gender
   - email
   - phone 
   - password
- we need to provide them in the right format all the validations are taking place 
<hr/>

/user/{id}:
- GET method 
- here id refers to uuid which is generated in create 
- we need to provide the uuid and using that we can fetch the data 
- return 404 if id is not found 
<hr/>

/user/{id}:
- PATCH method
- we can update password 
- password is sent in the request body 
- id here is uuid which is unique to all the records 


#### query endpoints:

/chatbot:

- POST method
- The Request Body is :
  - question
  - solution 
- return the data with ID which uuid generated inside
<hr/>

/chatbot/{question}:
- GET Method 
- takes question as path param 
- since we are making search empty spaces between the words are converted to '-'
<hr/>

/chatbot:
- GET Method
- return all the data 

/chatbot/{question}:
- PATCH method 
- updates the count

#### we are using gin framework for writing the code and gorm for db 

#### About gin:
<hr/>

##### Features:

Fast
Radix tree based routing, small memory foot print. No reflection. Predictable API performance.

Middleware support
An incoming HTTP request can be handled by a chain of middlewares and the final action. For example: Logger, Authorization, GZIP and finally post a message in the DB.

Crash-free
Gin can catch a panic occurred during a HTTP request and recover it. This way, your server will be always available. As an example - it’s also possible to report this panic to Sentry!

JSON validation
Gin can parse and validate the JSON of a request - for example, checking the existence of required values.

Routes grouping
Organize your routes better. Authorization required vs non required, different API versions… In addition, the groups can be nested unlimitedly without degrading performance.

Error management
Gin provides a convenient way to collect all the errors occurred during a HTTP request. Eventually, a middleware can write them to a log file, to a database and send them through the network.

Rendering built-in
Gin provides an easy to use API for JSON, XML and HTML rendering.

Extendable
Creating a new middleware is so easy, just check out the sample codes.

#### About gorm:
<hr/>

##### Features:

Full-Featured ORM

Associations (has one, has many, belongs to, many to many, polymorphism, single-table inheritance)
Hooks (before/after create/save/update/delete/find)

Eager loading with Preload, Joins

Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point

Context, Prepared Statement Mode, DryRun Mode

Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer

SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery

Composite Primary Key, Indexes, Constraints

Auto Migrations

Logger

Extendable, flexible plugin API: Database Resolver (multiple databases, read/write splitting) / Prometheus…

Every feature comes with tests

Developer Friendly