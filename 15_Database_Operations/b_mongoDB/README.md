# MongoDB


- schema less

- mongodb equivalents for sql words
    tables -- > collections
    row   -- document
    db           BSON


- key value piair



{
  "_id": ObjectId("651f9b8e1c2d4e3f4a5b6c7d"),  // _id field is auto-generated
  "name": "John Doe",
  "age": 30,
  "address": {
    "city": "New York",
    "zip": "10001"
  },
  "hobbies": ["reading", "traveling"]
}


_id  Field


## Fields

Fields is key-value pair

### Data types supported 

    String
    Integer
    Boolean
    Double
    Decimal128
    ObjectId
    Date
    Array
    Embedded documents
    Binary data
    Null
    etc.


## Indexes

    for improve perofirmavce of queires 

### Types:

    Single Field Index
    Compound Index (on multiple fields)
    Multikey Index (for array fields)
    Text Index (for text search)
    Geospatial Index (for geospatial queries)
    Hashed Index (for sharding)


## Queries

### Query Operators:

    Comparison  : $eq, $ne, $gt, $lt, etc.
    Logical     : $and, $or, $not, $nor
    Array       : $in, $nin, $all, $elemMatch
    Element     : $exists, $type


select  * from users where age > 35
db.users.find({age: {$gt: 35}})


## Aggregation

- simiaalr to GROUP BY and aggregate functions

### pipelines 

    $match: Filters documents.
    $group: Groups documents by a specified key.
    $sort: Sorts documents.
    $project: Reshapes documents.
    $limit: Limits the number of documents.


db.orders.aggregate([
  { $match: { status: "completed" } },
  { $group: { _id: "$customerId", total: { $sum: "$amount" } } }
]);



## Transactions

    start transaction 
    klklkj
    commit
    end tarnsaction



    const session = db.getMongo().startSession();
    session.startTransaction();
    try {
        const users = session.getDatabase("test").users;
        const orders = session.getDatabase("test").orders;
        users.insertOne({ name: "Alice" });
        orders.insertOne({ userId: 1, item: "Book" });
        session.commitTransaction();
    } catch (error) {
        session.abortTransaction();
    } finally {
        session.endSession();
    }


#Archiectrure 


cluster of nodes (each node is database)
cooriidnate 
event consistency



## Sharding

: Sharding is a method for distributing data across multiple servers to handle large datasets and high throughput.

### components
    
    Shard: A single MongoDB instance that stores a subset of the data.

    Config Servers: Store metadata and configuration settings for the cluster.

    Query Router (mongos): Routes queries to the appropriate shard.


## Replication

-- having multiple copies and better syning them 

Replica set
    -- groupd of mongob instances mainaitining the same data
    -- availability 
    -- (Primary) Only one node among them will be accepting the writes


## CRUD operations 

    Create  -- INSERT   
    Read    -- select 
    Update  -- update table set col= value where condition
    Delete - dleete from table where condition


    Create: Insert documents into a collection using insertOne() or insertMany().
    Read: Retrieve documents using find() or findOne().
    Update: Modify documents using updateOne(), updateMany(), or replaceOne().
    Delete: Remove documents using deleteOne() or deleteMany().

## GridFS

    for storing and retriving large files 
        images, videos


## MongoDB Atlas

    -- managed service of mongoDB



## Working in Golang


Install Go MongoDB drivers:

    go get go.mongodb.org/mongo-driver/mongo
    go get go.mongodb.org/mongo-driver/mongo/options
    go get go.mongodb.org/mongo-driver/bson
    