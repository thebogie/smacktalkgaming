FROM node:latest

WORKDIR /frontend

COPY . /frontend

COPY .env.prod /frontend/.env


CMD npm install; npm run dev
#RUN npm run build; npm run start
#CMD tail -f /dev/null

EXPOSE 3000