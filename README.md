


# Rewards API

## Project Overview

The Rewards API processes receipt data submitted via POST requests, calculates reward points based on specific criteria, and returns a unique receipt ID.  These points can then be retrieved using a GET request with the receipt ID.

## Flow

### Receipt Submission

1.  Client sends a JSON payload to the `/receipts/process` endpoint.
2.  API decodes the receipt, generates a unique receipt ID (UUID), and calculates reward points.
3.  Receipt and points are stored in memory and persisted to `receipts.json`.

### Points Retrieval

1.  Client requests points via GET to `/receipts/{id}/points`, where `{id}` is the receipt ID.
2.  API retrieves and returns points if the receipt exists; otherwise, it returns an error.

## Project Structure

1. main.go       # Sets up the HTTP server and routes.
2. handlers.go   # HTTP handler functions (SubmitReceiptHandler, RetrievePointsHandler).
3. models.go     # Data structures (Receipt, Item, ReceiptResponse, PointsResponse).
4. points.go     # Reward point calculation logic.
5. storage.go    # File storage functions (saving and loading receipts).

## Endpoints

### `POST /receipts/process`

*   **Description:** Processes a receipt, calculates points, and returns a receipt ID.
*   **Request Body:** JSON containing receipt details.
*   **Response:** JSON containing the generated receipt ID.

### `GET /receipts/{id}/points`

*   **Description:** Retrieves points for the receipt with the given ID.
*   **Response:** JSON containing the awarded points.

## How to Run

1.  **Install Go:** Ensure Go is installed (version >= 1.23.6 as specified in `go.mod`).

2.  **Download Dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Run the Application:**

    ```bash
    go run main.go handlers.go models.go points.go storage.go
    ```


4.  **Test the API:**

    *   **Submit a Receipt:**

        ```bash
        curl -X POST http://localhost:9090/receipts/process \
          -H "Content-Type: application/json" \
          -d '{"retailer": "StoreName", "purchaseDate": "2023-08-15", "purchaseTime": "14:00", "items": [{"shortDescription": "Item A", "price": "10.00"}, {"shortDescription": "Item B", "price": "20.00"}], "total": "30.00"}'
        ```

    *   **Retrieve Reward Points:** Replace `{id}` with the receipt ID from the previous step:

        ```bash
        curl http://localhost:9090/receipts/{id}/points
        ```

## Results:

![alt text](<Screenshot 2025-02-05 at 12.31.11 AM.png>) ![alt text](<Screenshot 2025-02-05 at 12.31.04 AM.png>) ![alt text](<Screenshot 2025-02-05 at 2.44.13 PM.png>) ![alt text](<Screenshot 2025-02-05 at 2.44.03 PM.png>) ![alt text](<Screenshot 2025-02-05 at 1.43.22 AM.png>) ![alt text](<Screenshot 2025-02-05 at 1.43.18 AM.png>)

