server {
    listen 80;
    listen [::]:80;

    server_name ratemyoldagehome.com www.ratemyoldagehome.com;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;
    proxy_ssl_server_name on;
    server_name ratemyoldagehome.com www.ratemyoldagehome.com;
    ssl_certificate /etc/letsencrypt/live/speakingfrankli.me/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/speakingfrankli.me/privkey.pem;

    root /var/www/parcel_blueprint/dist;
    index index.html index.htm index.nginx-debian.html;
    location / {
        try_files $uri /index.html =404;
    }
}

