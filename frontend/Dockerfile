FROM mhart/alpine-node:12 as build
WORKDIR /frontend
COPY . .
RUN npm install
RUN npm run build

WORKDIR /frontend
EXPOSE 3000
CMD ["npm", "start"]