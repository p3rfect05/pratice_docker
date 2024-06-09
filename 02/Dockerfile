FROM nginx:latest 

COPY ./index.html /usr/share/nginx/html/
COPY ./script.sh . 
ENV MESSAGE=LOLOLOLOL
RUN chmod +x ./script.sh
CMD ["/bin/sh", "-c", "sed -i 's/h1/'\"$MESSAGE\"'/' /usr/share/nginx/html/index.html && nginx -g 'daemon off;'"]

