FROM node:16-alpine as build-stage
# define the place where all of the files are stored
WORKDIR /frontend
# Copy the package.json and install dependencies
COPY package*.json ./
RUN npm install
# Copy the rest of the files
COPY . .
# Build the project, meaning create a dist folder that will hold all of the files
RUN npm run build

FROM nginx:stable-alpine as production-stage
RUN rm -rf /usr/share/nginx/html/*
# copies the created dist folder to the /usr/share/nginx/html directory
# can check the content of the folder with 
#   docker exec -it frontend sh
#   cd /usr/share/nginx/html
#   ls -la
COPY --from=build-stage /frontend/dist /usr/share/nginx/html

# copy the defined nginx file to the appropriate directory
COPY ./.nginx/nginx.conf /etc/nginx/nginx.conf
CMD ["nginx", "-g", "daemon off;"]