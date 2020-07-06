FROM nginx:1.19.0-alpine 

COPY . /nginxdrop

RUN \
    if [ "${APP_ENV}" = "production" ] ;\ 
    then \
    echo "PRODUCTION SETUP"; \
    cp /nginxdrop/nginx.conf /etc/nginx/nginx.conf; \
    else \
    echo "DEVEOLPMENT SETUP"; \
    cp /nginxdrop/nginx.dev.conf /etc/nginx/nginx.conf; \
    fi

#RUN /usr/local/nginx/sbin/nginx
#ENTRYPOINT ["tail", "-f", "/dev/null"]
CMD ["nginx", "-g", "daemon off;"]
#CMD ["tail", "-f", "/dev/null"]

EXPOSE 80 
EXPOSE 443



