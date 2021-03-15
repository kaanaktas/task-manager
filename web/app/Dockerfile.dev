FROM node:14.12.0-alpine

WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./

RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

RUN npm install axios --silent
RUN npm install semantic-ui-react --silent

# add app
COPY . ./

# start app
CMD ["npm", "start"]

#docker build -t kaktas/task-manager-react .
#docker run -it -v ${PWD}:/app -v /app/node_modules -p 3000:3000 -e CHOKIDAR_USEPOLLING=true kaktas/task-manager-react