# Full Stack Project Template

backend -> go, fiber ( can switch framework )

frontend -> framework agnostinc

database -> postgres

# Notes

### the .env file

So there are two .env files. The one that is located in the root is the one that will be used by the docker-compose.yml (during production) and the one that is located in ./fiber_server folder will be used during local development.

### nginx folder

Holds an example of a nginx.conf file with reverse-proxy that connects the frontend and backend. Links to tutorials which explain what you need to do to add a domain and generate new ssl certs automatically
