# Dynomo

# Architecture 
<img width="802" alt="image" src="https://github.com/Dynomo-org/dynapgen/assets/7591715/64f0388f-1ebf-4e36-b10a-1b248079da22">


## API Schema
### End-user (to be consumed by dynomo app)
- **POST /app/build**
- **GET /app/build/status/{project_id}**

### Internal
- **POST /build/status**

  To update build status from worker to the main server
  
