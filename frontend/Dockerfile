FROM node:latest as frontend-build-stage
WORKDIR /frontend
COPY ./package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:latest as frontend-production-stage
COPY --from=frontend-build-stage /frontend/dist /usr/share/nginx/html