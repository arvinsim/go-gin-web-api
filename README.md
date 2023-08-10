# Go Gin Web API

## Experiments

### Connect Gin with Postgres
- Create initialization script for creating and seeding tables
- Get sample data from database
- Implement proper database migration
- Implement proper way to globalise DB object
 
### Process an uploaded image via 2 microservices
- Upload image to Gin endpoint
- Create microservices
  - Convert image to png via [imaging](https://github.com/disintegration/imaging)
  - Add watermark to image
- Connect via a message queue or rpc

### Connect Gin with NoSQL