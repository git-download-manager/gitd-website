#!/bin/bash

# Before run: chmod a+x deploy.sh

# Local Variables
ARGS=("$@")
IP_ADDR="${ARGS[0]}"
#IP_ADDR="$(ip route get 8.8.8.8 | head -1 | awk '{print $7}')"
DOMAIN="${ARGS[1]}"

# Update nginx conf IPADDR variable to real ipaddr
CONFFILE=/usr/local/bin/nginx-conf/sites-available/main_website.conf
if [ -f "$CONFFILE" ]; then
    echo "$CONFFILE exists."

    sed -i "s|IPADDR|${IP_ADDR}|g" $CONFFILE
    sed -i "s|DOMAIN_NAME|${DOMAIN}|g" $CONFFILE

else
    echo "$CONFFILE does not exist."
fi

# Nginx files copy to real directory
echo "nginx conf copy"
cp -vr /usr/local/bin/nginx-conf/* /etc/nginx/

# Eğer service dosyalarında bir değişiklik olduysa çalışması gerek
echo "Systemctl deamon reload"
systemctl daemon-reload

# Syslog service has to restart
echo "Systemctl rsyslog service restart"
systemctl restart rsyslog.service
systemctl --no-pager status rsyslog.service

# Nginx service has to restart
echo "Nginx service restart"
systemctl restart nginx.service
systemctl --no-pager status nginx.service