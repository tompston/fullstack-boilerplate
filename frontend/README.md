# Frontend folder

As the Dockerfile will copy the generated dist folder and serve it with nginx, the used framework does not matter. The only thing that the current setup needs is the

1. recognition of "npm run build" command to build the end version

2. generated build folder to be named "dist"

Or just change the Dockerfile if your config is different

## Starting a new project with _Vite_

    npm init @vitejs/app

## robots.txt

Create robots.txt file in the ./public folder  
link to a guide for more details   ->  http://www.robotstxt.org/robotstxt.html
