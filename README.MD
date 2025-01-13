# Cloud Resources Inventory Management System.
This service takes its messages from RabbitMQ and logs them to the console.

## Steps 
1. Do a `docker compose build` to build the docker images
2. Do a `docker compose up` to run it
3. You can seed the customers by calling the command `make seed-customers`. *NB: You have to have make installed or just* **run go run main.go seed CustomerSeeder** *straight*
4. You can seed the customers by calling the command `make eed-cloud-resources`. *NB: You have to have make installed or just* **go run main.go seed CloudResourceSeeder** *straight*
5. Make a POST call to teh following endpoint http://localhost:8080/v1/customers/:id/cloud-resources to attach resources to a customer.
```
Request Body Sample
{
    "resources": [
        "Compute Engine"
    ]
}
```
6. To view cloud resources attached to a user, you can call this endpoint GET call, http://localhost:8080/v1/customers/:id/cloud-resources .
```
Response Body Sample
{
    "status": true,
    "data": [
        {
            "id": "ff819ea0-18a2-4409-a550-5d6bc1868cfa",
            "Name": "Compute Engine",
            "Type": "VM",
            "Region": "us-central1"
        }
    ],
    "error": [],
    "message": "customer cloud resources"
}
```
7. To update a cloud resource you can call this endpoint PUT call,  http://localhost:8080/v1/cloud-resources/:id .
```
Request Body Sample
{
    "name" : "Firewall",
    "type" : "Security",
    "region" : "us-central1"  
}
```
8. You can use the command `make run-tests` to run the tests for this service.
*NB you have to be in the root directory and have make installed to run the tests*