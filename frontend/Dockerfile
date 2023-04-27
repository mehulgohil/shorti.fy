FROM node:alpine
WORKDIR /app

COPY ./frontend ./

RUN npm i

EXPOSE 3000

CMD ["npm", "run", "start"]