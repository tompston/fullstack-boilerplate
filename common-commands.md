### running docker-compose + updating containers

    docker-compose down
    docker-compose build
    docker system prune -f
    docker-compose up -d
    docker image prune -a

### Other commands

    docker stats $(docker ps | awk '{if(NR>1) print $NF}')          # check resource usage (CPU / MEM)
    SELECT version();                                               # chech the current postgres version (from psql shell)
    chmod +x SCRIPT_NAME.sh && ./SCRIPT_NAME.sh                     # allow execution of .sh scripts

    docker system df        # check the amount of disk space + general info
    docker-compose ps       # check name - command state ports
    docker image ls         # check images
    netstat -tln            # check port stuff
    free -m                 # check free memory

### Removing Docker stuff

    # volumes
    docker volume ls
    docker volume rm NAME NAME

    # containers
    docker container ls -a
    docker container rm NAME NAME

    # everything
    docker-compose down
    docker system prune -a -f
    docker rm -f $(docker ps -q)

    # unused volumes
    docker volume prune

### Rebuild only one image

    docker-compose up -d --no-deps --build <service_name>

    # full example could look like this
    docker-compose down
    docker-compose up -d --no-deps --build frontend
    docker-compose up -d

### cd into a docker container to check stuff

    docker exec -it <container_name> <entrypint>

    # examples
    docker exec -it gofiber_server bash
    docker exec -it frontend sh

### cd into postgres docker container

    # change the uppercase names to your defined variables in the root .env file
    docker exec -it db psql -d POSTGRES_DB -U POSTGRES_USER

### Other notes

    WORKDIR command in Dockerfile will create the dir if it doesn’t exist yet

    cd /usr/share/nginx/html        # directory that holds the dist folder
    cd etc/nginx/nginx.conf         # directory that holds the nginx.conf file
