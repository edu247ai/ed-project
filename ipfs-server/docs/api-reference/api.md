# API Reference

## Base URL
http://localhost:8082

Code

## Endpoints

### POST /upload
Upload raw bytes.

**Request**
```json
{
  "data": "base64-encoded-bytes"
}
Response

json
{
  "cid": "<sha256>"
}
GET /download?cid=<cid>
Download stored content.

Response

json
{
  "cid": "<cid>",
  "size": 12345,
  "data": "base64-encoded-bytes"
}
POST /replicate?cid=<cid>
Replicate block to this node.

GET /dashboard
Live metrics dashboard.

Authentication
All endpoints except /dashboard require:

Code
X-API-Key: <key>