# ORM  - Object-Relational Mapping


In simpler terms:

        Database Tables are mapped to Objects/Structs.
        Rows in a table are mapped to Instances of those objects/structs.
        Columns in a table are mapped to Fields/Attributes of the objects/structs.


-----------------------------------------------------------------------
ORM	            Best For	                        Key Features
-----------------------------------------------------------------------
GORM	        General-purpose ORM	            Full-featured, easy to use
Ent	            Schema-first, type safety	    GraphQL support, code generation
Beego ORM	    Beego framework users	        Simple, supports raw SQL
Pop	            Buffalo framework users 	    Migrations, query builder
SQLBoiler	    Code generation, type safety	No reflection, highly performant
XORM	        High performance	            Caching, reverse engineering
-----------------------------------------------------------------------


without ORM 

    INSERT INTO users (name, email) VALUES ('John', 'john@example.com');

with ORM
    
    user := User{Name: "John", Email: "john@example.com"}
    db.Create(&user)



### Type Safety

without ORM 

    SELECT * FROM users WHERE id = 'abc'; // Runtime error: 'abc' is not a valid integer

with ORM

    var user User
    db.Where("id = ?", 123).First(&user) // Compile-time type checking



