# Modern Protocols: gRPC vs GraphQL vs REST

## REST (Representational State Transfer)

### Characteristics
- **HTTP-based**: Uses HTTP methods (GET, POST, PUT, DELETE).
- **Resource-oriented**: Each URL represents a resource (e.g., `/users/123`).
- **Stateless**: Each request is independent.
- **Text-based**: JSON or XML.

### Pros
- **Simple**: Easy to understand and implement.
- **Cacheable**: HTTP caching works out of the box.
- **Widely supported**: Every language has HTTP libraries.

### Cons
- **Over-fetching**: Client gets more data than needed (e.g., `GET /users/123` returns all user fields).
- **Under-fetching**: Client needs multiple requests (e.g., get user, then get user's posts).
- **Versioning**: API changes require versioning (`/v1/users`, `/v2/users`).

### Example
```
GET /users/123
Response: {"id": 123, "name": "Alice", "email": "alice@example.com", "age": 30}
```

---

## gRPC (Google Remote Procedure Call)

### Characteristics
- **HTTP/2-based**: Multiplexing, header compression, server push.
- **Binary protocol**: Uses **Protocol Buffers** (Protobuf) for serialization.
- **Strongly typed**: Schema defined in `.proto` files.
- **Bidirectional streaming**: Client and server can stream data.

### Pros
- **Fast**: Binary serialization is faster than JSON.
- **Type-safe**: Protobuf generates code in multiple languages.
- **Streaming**: Supports client streaming, server streaming, and bidirectional streaming.
- **Efficient**: HTTP/2 multiplexing reduces latency.

### Cons
- **Complexity**: Requires Protobuf definitions and code generation.
- **Not human-readable**: Binary format (harder to debug).
- **Browser support**: Limited (requires gRPC-Web proxy).

### Example: Protobuf Definition
```protobuf
syntax = "proto3";

service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}

message UserRequest {
  int32 id = 1;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

### Go Implementation
```go
// Server
func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
    return &pb.UserResponse{
        Id:    req.Id,
        Name:  "Alice",
        Email: "alice@example.com",
    }, nil
}

// Client
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
client := pb.NewUserServiceClient(conn)
resp, _ := client.GetUser(context.Background(), &pb.UserRequest{Id: 123})
fmt.Println(resp.Name) // Alice
```

---

## GraphQL

### Characteristics
- **Query language**: Clients specify exactly what data they need.
- **Single endpoint**: All queries go to `/graphql`.
- **Strongly typed**: Schema defined in GraphQL SDL (Schema Definition Language).

### Pros
- **No over-fetching**: Client requests only the fields it needs.
- **No under-fetching**: Client can request related data in a single query.
- **Introspection**: Clients can query the schema (self-documenting).
- **Versioning**: No need for versioning (add new fields without breaking old clients).

### Cons
- **Complexity**: Requires a GraphQL server and schema.
- **Caching**: Harder to cache (all requests go to `/graphql`).
- **N+1 problem**: Poorly written resolvers can cause many database queries (use DataLoader to fix).

### Example: GraphQL Query
```graphql
query {
  user(id: 123) {
    name
    email
    posts {
      title
    }
  }
}
```

**Response**:
```json
{
  "data": {
    "user": {
      "name": "Alice",
      "email": "alice@example.com",
      "posts": [
        {"title": "GraphQL is awesome"},
        {"title": "Learning Go"}
      ]
    }
  }
}
```

### Go Implementation (using `graphql-go`)
```go
type User struct {
    ID    int32
    Name  string
    Email string
}

var userType = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
        "id":    &graphql.Field{Type: graphql.Int},
        "name":  &graphql.Field{Type: graphql.String},
        "email": &graphql.Field{Type: graphql.String},
    },
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "user": &graphql.Field{
            Type: userType,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{Type: graphql.Int},
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                id := p.Args["id"].(int)
                return User{ID: int32(id), Name: "Alice", Email: "alice@example.com"}, nil
            },
        },
    },
})

schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
```

---

## Comparison Table

| Feature | REST | gRPC | GraphQL |
| :--- | :--- | :--- | :--- |
| **Protocol** | HTTP/1.1 or HTTP/2 | HTTP/2 | HTTP/1.1 or HTTP/2 |
| **Data Format** | JSON (text) | Protobuf (binary) | JSON (text) |
| **Schema** | Optional (OpenAPI) | Required (Protobuf) | Required (GraphQL SDL) |
| **Type Safety** | No | **Yes** | **Yes** |
| **Over-fetching** | ✅ Yes | ❌ No (request specific fields in Protobuf) | ❌ No |
| **Under-fetching** | ✅ Yes (multiple requests) | ❌ No (streaming) | ❌ No (nested queries) |
| **Caching** | **Easy** (HTTP caching) | Hard (binary, no URL-based caching) | Hard (single endpoint) |
| **Streaming** | No (unless SSE/WebSockets) | **Yes** (bidirectional) | No (unless subscriptions) |
| **Browser Support** | **Universal** | Limited (requires gRPC-Web) | **Universal** |
| **Performance** | Moderate | **High** (binary, HTTP/2) | Moderate |
| **Use Case** | Public APIs, CRUD apps | **Microservices, high-performance** | **Frontend-heavy apps, mobile** |

---

## When to Use Each

### REST
- **Public APIs** (easy for third-party developers).
- **Simple CRUD** applications.
- **Caching** is important.

### gRPC
- **Microservices** (internal service-to-service communication).
- **High performance** is critical (low latency, high throughput).
- **Streaming** is needed (e.g., real-time data).

### GraphQL
- **Frontend-heavy** applications (mobile, SPAs).
- **Flexible queries** (clients need different data).
- **Rapid iteration** (no need to version APIs).

---

## Interview Questions

### Q: When would you use gRPC over REST?
**A**: When performance is critical (e.g., microservices) and you need type safety, streaming, or HTTP/2 multiplexing. gRPC is faster due to binary serialization (Protobuf) and HTTP/2.

### Q: What's the main advantage of GraphQL over REST?
**A**: **No over-fetching or under-fetching**. Clients request exactly the data they need in a single query, reducing the number of requests.

### Q: What's the N+1 problem in GraphQL?
**A**: When resolving a list of objects, the server makes 1 query for the list + N queries for each object's related data. **Fix**: Use **DataLoader** to batch and cache database queries.

### Q: Why is caching harder with GraphQL?
**A**: All requests go to a single endpoint (`/graphql`), so URL-based HTTP caching doesn't work. You need application-level caching (e.g., Apollo Client cache).
