FROM nginx:1.19.0-alpine 

COPY . /nginxdrop

# Define working directory.
WORKDIR /etc/nginx

CMD \
    if [ "${APP_ENV}" = "production" ] ;\ 
    then \
    echo "PRODUCTION SETUP"; \
    cp /nginxdrop/nginx.conf /etc/nginx/nginx.conf; \
    nginx; \
    else \
    echo "DEVEOLPMENT SETUP"; \
    cp /nginxdrop/nginx.dev.conf /etc/nginx/nginx.conf; \
    nginx; \
    fi



# Define default command.
#CMD ["nginx"]

EXPOSE 80 
EXPOSE 443



