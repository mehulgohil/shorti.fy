# stage 1

FROM node:18.12.1-alpine as build
WORKDIR /app
COPY ./frontend .
RUN npm ci
RUN npm run build

# stage 2

FROM nginx:1.22.1-alpine as production
ENV NODE_ENV production
COPY --from=build /app/build /usr/share/nginx/html
COPY ./frontend/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]