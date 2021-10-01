# Full Stack Project Template

backend -> go, fiber ( can switch framework )

frontend -> framework agnostinc

database -> postgres

## the .env file

So there are two .env files. 
- ./env 
  - used by docker-compose.yml (during production) and docker-compose.local.yml during local testing with nginx
- ./fiber_server/.env 
  - used during local development., without docker


## ./nginx folder


Holds an example of a nginx.conf file with reverse-proxy that connects the frontend and backend. Links to tutorials which explain what you need to do to add a domain and generate new ssl certs automatically

## How to run things

The current setup has 3 possible ways of running.
 - using docker-compose.yml for production
        
        # will create a postgres database with the variables that are mentioned in the .env file, 
        # build the frontend and serve with nginx
        # build the backend and serve as a single binary

        docker-compose up -d

        # to add domain name and ssl certs, check out the instructions in the ./nginx folder



 - using docker-compose.local.yml to test the production version. Nginx image connects the frontend and backend, so you can make api calls in the frontend from the same url (localhost).

        # will create a postgres database with the variables that are mentioned in the .env file + ovveride the PAGE_URL and BASE_URL variables, 
        # build the frontend and serve with nginx
        # build the backend and serve as a single binary

        docker-compose -f docker-compose.local.yml up -d

 - Running everything locally.

        Run the go server and frontend separately. 
        Delete everything except the Dockerfile and nginx.conf file in the ./frontend folder and create a new project


  ### Disclaimer
  *Code and setup takes heavy inspiration from this [repo](https://github.com/karlkeefer/pngr/tree/master/react)* 
