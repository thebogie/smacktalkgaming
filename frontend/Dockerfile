FROM node:latest

WORKDIR /frontend

COPY . /frontend

COPY .env.prod /frontend/.env

CMD if [ ${APP_ENV} = production ]; \
    then \
    npm run build; \
    npm run start; \
    else \
    npm install; \
    npm run dev;  \
    fi

EXPOSE 3000