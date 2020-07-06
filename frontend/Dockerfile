FROM node:latest

WORKDIR /frontend

COPY . /frontend

COPY .env /frontend/.env

CMD if [ "${APP_ENV}" = "production" ]; \
    then \
    echo "PRODUCTION SETUP"; \
    cp /frontend/.env.prod .env; \ 
    npm run build; \
    npm run start; \
    else \
    echo "DEVEOLPMENT SETUP"; \
    npm install; \
    npm run dev;  \
    fi

EXPOSE 3000