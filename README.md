## David Youn - Receipt Process Challenge

challenge link: https://github.com/fetch-rewards/receipt-processor-challenge

---

## Instructions to Run

1. After downloading, run `go run main.go`.
   <br> **note**: you may need to install google/uuid by running: `go get github.com/google/uuid`
   <br>           you may also need to install gorilla/mux by running: `go get github.com/gorilla/mux`

2. This will start running the server in localhost port `:8080`
3. Below I've provided two different ways to test below:
   - using Postman
   - using Rest Client plugin in VSCode

  ### 1. Using Postman
  This method is to use Postman to help test the API.
  #### POST
  1. Create a `POST` request with the path: `localhost:8080/receipts/process`
  2. In the request body, choose JSON and add details of the receipt you would like to process.
    <br> **EXAMPLE**
     <br> <img width="250" alt="Screenshot 2024-11-21 at 8 03 32 PM" src="https://github.com/user-attachments/assets/c767c8c4-f187-4422-aab4-7d2d6b22c069">
  3. This should return a body containing an id (a uuid to be specific). Make sure you copy this value.
  #### GET
  1. Create a `GET` request with the path: `localhost:8080/receipts/{id}/points`
     <br> **note**: replace `{id}` with the id value copied in step 3 above.
  2. This should return a body containing the number of points corresponding to the receipt with the id you wrote in the request.
     <br> **EXAMPLE**
      <br> <img width="250" alt="Screenshot 2024-11-21 at 8 08 51 PM" src="https://github.com/user-attachments/assets/103daa11-5921-4cd9-a1e2-a03cdc5a28bf">
  
  ### 2. Using Rest Client plugin in VSCode
  This VSCode plugin will make it so you can test in VSCode with the file [examples.rest](https://github.com/david-youn/receipt-processor/blob/main/examples.rest) provided.
  <img width="250" alt="Screenshot 2024-11-21 at 8 13 51 PM" src="https://github.com/user-attachments/assets/ca2f76b1-d048-4ef3-ac39-996a7173eb15">
  
  #### POST
  1. Above the `POST` request, there is a small, clickable text with the words `send request`. Clicking it will generate another window with the results of running the request.
  2. This should return a body containing an id value. Make sure to copy this.
     <br> **EXAMPLE**
  <br> <img width="500" alt="Screenshot 2024-11-21 at 8 17 22 PM" src="https://github.com/user-attachments/assets/9215a094-1d5f-44bb-b183-6dcfda4359e7">
  #### GET
  1. On line 58 of the examples.rest file, there should be a `GET` request that looks like `GET http://localhost:8080/receipts/{id}/points`.
  2. Replace `{id}` value with the id value copied in step 2 above.
  3. Above the `GET` request, there is a small, clickable text with the words `send request`. Clicking it will generate another window with the results of running the request.
  4. This should return a body containing the number of points corresponding to the receipt with the id you wrote in the request.
     <br> **EXAMPLE**
     <br> <img width="500" alt="Screenshot 2024-11-21 at 8 24 12 PM" src="https://github.com/user-attachments/assets/fe3407da-1a77-4807-9063-d38f12a0cb8d">
  
---
### Additional Notes:
- I've also printed the values to the console to show the point breakdowns.
<br> <img width="200" alt="Screenshot 2024-11-21 at 8 30 08 PM" src="https://github.com/user-attachments/assets/e5e103c2-5d68-46f0-89b7-2ee5e4e54d49">


---
### Unrelated Notes:
In case you make it this far: 
- This is my first time coding in Go, but wanted to show you that I am open to learning new tools and languages!
- Of course, this is not something I could've done without any help. I followed a few tutorials on youtube to understand the concept of setting up webservers in Go (have uploaded that work to git as well in case you're curious!).
- References I used:
- https://github.com/david-youn/GO-SERVER/tree/main
     <br> reference video: https://www.youtube.com/watch?v=eqvDSkuBihs
- https://github.com/david-youn/HTTP-SERVER
     <br> reference video: https://www.youtube.com/watch?v=5BIylxkudaE


Thanks for taking a look!
Hope to hear back,
David Youn
