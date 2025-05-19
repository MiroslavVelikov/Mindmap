FROM node:20-alpine AS build

WORKDIR /app

COPY frontend-app-mindmap/package*.json ./
RUN npm install

COPY frontend-app-mindmap/ ./
RUN npm run build --configuration=production

FROM nginx:alpine
COPY --from=build /app/dist/frontend-app-mindmap /usr/share/nginx/html
EXPOSE 80