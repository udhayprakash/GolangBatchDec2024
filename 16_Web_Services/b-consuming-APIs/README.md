http  status code
=================
1xx: Informational Responses

    100 Continue: The server received the initial part of the request and the client should proceed.
    101 Switching Protocols: The server agrees to switch protocols as requested by the client.
    102 Processing (WebDAV): The server is processing the request but no response is yet available.

2xx: Successful Responses

    200 OK  : The request succeeded, and the response contains the requested resource or confirmation.
    201 Created: The request succeeded, and a new resource was created.
    202 Accepted: The request was accepted but has not been processed yet.
    203 Non-Authoritative Information: The returned meta-information is not from the origin server.
    204 No Content: The request succeeded, but there is no content to return.
    205 Reset Content: The client should reset the view or form for the resource.
    206 Partial Content: The server is delivering part of the resource due to a range header.

3xx: Redirection

    300 Multiple Choices: There are multiple options for the resource.
    301 Moved Permanently: The resource has been moved permanently to a new URL.
    302 Found: The resource is temporarily located at a different URL.
    303 See Other: The client should retrieve the resource using a GET request at a different URL.
    304 Not Modified: The resource has not been modified since the last request.
    307 Temporary Redirect: The resource is temporarily at a new URL, and the client should use the original method (e.g., POST or GET).
    308 Permanent Redirect: The resource has been moved permanently, and the client should use the same method.


4xx: Client Errors

    400 Bad Request: The request cannot be processed due to malformed syntax.
    401 Unauthorized: Authentication is required and has failed or not been provided.
    402 Payment Required: Reserved for future use (e.g., digital payments).
    403 Forbidden: The server understands the request but refuses to authorize it.
    404 Not Found: The requested resource could not be found.
    405 Method Not Allowed: The HTTP method is not allowed for the resource.
    406 Not Acceptable: The resource cannot fulfill the request with the provided Accept headers.
    407 Proxy Authentication Required: Authentication is required through a proxy.
    408 Request Timeout: The client took too long to send the request.
    409 Conflict: There is a conflict with the current state of the resource.
    410 Gone: The resource is no longer available.
    411 Length Required: The server requires the Content-Length header.
    412 Precondition Failed: A precondition in the request headers was not met.
    413 Payload Too Large: The request payload is too large for the server to handle.
    414 URI Too Long: The URI is too long for the server to process.
    415 Unsupported Media Type: The media type is not supported by the server.
    416 Range Not Satisfiable: The range specified cannot be satisfied.
    417 Expectation Failed: The Expect header cannot be fulfilled by the server.
    418 I'm a teapot: A playful code from RFC 2324, indicating a teapot is unable to brew coffee.
    429 Too Many Requests: The user has sent too many requests in a given time.

5xx: Server Errors

    500 Internal Server Error: A generic error occurred on the server.
    501 Not Implemented: The server does not recognize the method or lacks capability.
    502 Bad Gateway: The server received an invalid response from an upstream server.
    503 Service Unavailable: The server is overloaded or under maintenance.
    504 Gateway Timeout: The upstream server failed to send a response in time.
    505 HTTP Version Not Supported: The HTTP version is not supported by the server.
    507 Insufficient Storage (WebDAV): The server cannot store the representation.
    508 Loop Detected (WebDAV): The server detected an infinite loop while processing the request.


CRUD operations & http Methods
=========

GET	
    (Read) Retrieve data from a server.
    Fetch resources (e.g., /users/1).

POST
    (Create)Send data to the server to create a resource.
    Submit forms or create a resource (e.g., /users).

PUT
    (Update) Send data to the server to update a resource or create it if it doesn't exist.	
    Update user data (e.g., /users/1).

PATCH
    (Update) Send partial data to update a resource.	
    Update specific fields (e.g., /users/1).

DELETE
    (Delete)  Delete a resource on the server.	
    Remove a resource (e.g., /users/1).

HEAD
    (Read) Retrieve headers for a resource, similar to GET, but without the response body.
    Check if a resource exists.

OPTIONS
    (Read) Retrieve supported HTTP methods for a resource.	
    Check allowed methods for an endpoint.

TRACE
    (Read) Echo back the received request, primarily used for debugging.	
    Test request path for debugging.

CONNECT
    Establish a tunnel to the server, typically used with proxies or SSL/TLS connections.	
    Proxy tunneling (e.g., HTTPS via proxies).



protocol - http
interface - 
    SOAP API  : (GET , POST)
    REST API  :  all methods supported (server side developer need to implement)
        restfulness of restapi
    GraphQL



- REST
---------
- REST stands for REpresentational State Transfer
- It's principles are:
	- It should be stateless
	- It should access all the resources from the server using only URI
	- It does not have inbuilt encryption
	- It does not have session
	- It uses one and only one protocol - HTTP
	- For CRUD operations, it should use HTTP methods like
		GET  - to retrieve one or more records
		POST - to create record
		PUT  - to update an existing record
		PATCH- to update a specific column in existing record
		DELETE- to delete a record

		HEAD - response identical to GET, but without response body
		CONNECT- establishes a tunnel to server identifier by the target resource
		OPTIONS- used to describe the communication options for the target resource.
		TRACE- performs a message loop-back test along the path to the target resource.
		PATCH- used to apply partial modifications to a resource

	- response can be JSON or XML, atom, OData etc.


	Idempotence - If something is idempotent, then no matter how many times you do it, the result will always be the same.
	1. POST is NOT idempotent.
	2. GET, PUT, DELETE, HEAD, OPTIONS and TRACE are idempotent.
	Nulipotent- GET is so because nothing is added or changed, except logging.



- PATCH operations
	add		: Adds a new value to the resource.
	remove	: Removes a value from the resource.
	replace	: Replaces the value of a property.
	move	: Moves a property to a new location.
	copy	: Copies a property to a new location.
	test	: Tests whether a property has a particular value.


- Richardson-maturity-model
	- to know level of RESTful your API is.
	- https://developers.redhat.com/blog/2017/09/13/know-how-restful-your-api-is-an-overview-of-the-richardson-maturity-model


http protocols
    1.1
    1.2


Authentication 
    Access to the system

    1) http Basic
        Oauth2
            AWs cognito  -> ldap/active directory
            Social Oauth 
                Github, google, facebook, microsoft,,...
        JWT (Json web token)


Authorization
    level of access 
    RBAC - role based access


## Authentication

### 1. Token-Based Authentication:

    - JSON Web Tokens (JWT)
    - OAuth 2.0 and OpenID Connect
    - Classification: Stateless Authentication

### 2. Session-Based Authentication:

    - Session-based authentication with cookies
    - Classification: Server-Side Authentication

### 3. Basic Authentication:

    - HTTP Basic Authentication
    - Classification: Stateless Authentication

### 4. Digest Authentication:

    - HTTP Digest Authentication
    - Classification: Stateless Authentication

### 5. API Key Authentication:

    - API Key-based authentication
    - Classification: Stateless Authentication

### 6. Two-Factor Authentication (2FA):

    - Time-based One-Time Password (TOTP)
    - SMS-based OTP

### 7. Classification: Multi-Factor Authentication

    - Social Authentication:
    - Integration with social login providers (e.g., Google, Facebook, Twitter)
    - Classification: External Authentication

### 8. LDAP Authentication:

    - Integration with Lightweight Directory Access Protocol (LDAP) servers
    - Classification: External Authentication

### 9. Single Sign-On (SSO):

    - Integration with SSO providers (e.g., SAML, OAuth, OpenID Connect)
    - Classification: External Authentication

### 10. Role-Based Access Control (RBAC):

    - User roles and permissions management
    - Classification: Authorization Mechanism

### 11. Multi-Factor Authentication (MFA):

    - Combining multiple authentication factors (e.g., password, OTP, biometrics)
    - Classification: Multi-Factor Authentication

### 12. Custom Authentication:

    - Implementing custom authentication mechanisms specific to your application
    - Classification: Custom Authentication