FROM node:16

WORKDIR /app

COPY package*.json ./


RUN npm i --production

COPY . ./

EXPOSE 5050

CMD ["npm start"]
