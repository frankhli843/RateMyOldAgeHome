server {
    listen 80;
    listen [::]:80;

    server_name ratemyoldagehome.com www.ratemyoldagehome.com;
    location / {
        proxy_pass http://127.0.0.1:4444/;
    }
}

