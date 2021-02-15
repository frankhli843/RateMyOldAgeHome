#!/bin/bash
gitup;

echo "===============================  NGINX & SSL ==============================="
echo "------------ removing default files"
cd /etc/nginx/sites-enabled; rm default;
cd /etc/nginx/sites-available; rm default;
cd /var/www; rm -r html;
alias stopports='sudo kill -9 $(sudo lsof -t -i:80); sudo kill -9 $(sudo lsof -t -i:8000)'; stopports;
service nginx stop;

echo "============================  Full Stack =============================="
echo "------------ copying over configuraiton"
cp --remove-dest /root/RateMyOldAgeHome/nginx/ratemyoldagehome.com /etc/nginx/sites-available/ratemyoldagehome.com;
cd /etc/nginx/sites-available/ || exit;
pwd; ls; cat ratemyoldagehome.com;

echo "------------ create react directory"
cd /root/RateMyOldAgeHome/react-app; yarn; yarn build;

echo "------------ symbolically linking files"
ln -s /etc/nginx/sites-available/ratemyoldagehome.com /etc/nginx/sites-enabled/ratemyoldagehome.com;
cd /etc/nginx/sites-enabled/ || exit; pwd;
ls;
cat ratemyoldagehome.com;

echo "------------ ssl certificates"
certbot certonly --noninteractive --agree-tos -m lifrank1994@gmail.com --standalone -d ratemyoldagehome.com -d www.ratemyoldagehome.com

echo "------------ launch nginx!"
nginx -t;
service nginx start;
systemctl enable nginx;

echo "------------ relaunch go site"
cd /root/RateMyOldAgeHome || exit;
docker-compose up -d --build;

