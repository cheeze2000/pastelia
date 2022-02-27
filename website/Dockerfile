FROM node:16
WORKDIR /pastelia
COPY . .

RUN npm ci
ARG API_URL
RUN VITE_API_URL=${API_URL} npm run build
RUN npm i sirv-cli

EXPOSE 3000
CMD ["npx", "sirv", "dist", "--single", "--host", "--port", "3000"]
