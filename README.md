# Receipt Processor

A simple receipt processor API to calculate and retrieve receipt points based on specific rules.

There is a Makefile in the root directory to easily build and run the project
- `make build`
- `make run`

You can also use the docker extension on vs code to do the same

Additionally you can use the below commands to build and run
- `docker build -t receipt-inator .`
- `docker run -p 8080:8080 receipt-inator`


---
## API Documentation

### 1. **POST /receipts/process**

#### **Description**
Submits a receipt for processing and returns a unique receipt ID.

#### **Request**
- **URL**: `/receipts/process`
- **Method**: `POST`
- **Body**:

  ```json
  {
      "retailer": "Target",
      "purchaseDate": "2022-01-01",
      "purchaseTime": "13:01",
      "items": [
          {
              "shortDescription": "Mountain Dew",
              "price": "1.99"
          },
          {
              "shortDescription": "Cheese",
              "price": "2.49"
          }
      ],
      "total": "4.48"
  }

### 2. GET /receipts/:id/points

#### **Description**
Fetches the points awarded for a previously processed receipt using its unique ID.


#### **Request**
- **URL**: `/receipts/:id/points`
- **Method**: `GET`
- **Path Parameter**:
  - `id` *(string)*: The unique ID of the receipt, returned by the `POST /receipts/process` endpoint.

##### **Example Request**
```bash
curl --location --request GET 'http://localhost:8080/receipts/123e4567-e89b-12d3-a456-426614174000/points'
