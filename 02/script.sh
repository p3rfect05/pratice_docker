#!/bin/sh
nginx -g 'daemon off;'
sed 's/h1/'\"$MESSAGE\"'/' /usr/share/nginx/html/index.html
