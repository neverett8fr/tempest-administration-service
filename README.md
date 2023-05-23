# tempest-administration-service  
this is responsible for managing administrative tasks including authentication.  
this service manages user authentication and authorisation of API requests.  
# How to run  
## Build  
```bash
docker build -t .
 ```
   
 ## Run  
 ```bash
docker run -p 8080:8080 -v . -e ENV_VARIABLE=value .
 ```
   
 ## Stop the container  
 ```bash
 docker stop container-name
 ```
