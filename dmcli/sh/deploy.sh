#!/bin/bash

# Before run: chmod a+x deploy.sh

# Local Variables
ARGS=("$@")
IP_ADDR="${ARGS[0]}"
#IP_ADDR="$(ip route get 8.8.8.8 | head -1 | awk '{print $7}')"
DOMAIN="${ARGS[1]}"
HTTPADDR="${ARGS[2]}"
SERVICE_FOLDER="${ARGS[3]}"
SERVICE_FILENAME="${ARGS[4]}"
SUB_DOMAIN="${ARGS[5]}"

# Update .env file domain
ENVFILE=/usr/local/bin/$SERVICE_FOLDER/.env
if [ -f "$ENVFILE" ]; then
    echo "$ENVFILE exists."

    mkdir -p /usr/local/gitd-temp/

    sed -i "s|domain http://localhost:3001|domain https://${DOMAIN}|g" /usr/local/bin/$SERVICE_FOLDER/.env
    sed -i "s|domain http://localhost:3002|domain https://${SUB_DOMAIN}${DOMAIN}|g" /usr/local/bin/$SERVICE_FOLDER/.env
    sed -i "s|../gitd-temp/|/usr/local/gitd-temp/|g" /usr/local/bin/$SERVICE_FOLDER/.env

else
    echo "$ENVFILE does not exist."
fi

# Copy assets and templates (static files) to nginx root folder before create folder
STATICFOLDER=/usr/local/bin/$SERVICE_FOLDER/assets
if [ -d "$STATICFOLDER" ]; then
    echo "$STATICFOLDER folder does not exist."

    # Update main.js
    FILEMAINJS=/usr/local/bin/$SERVICE_FOLDER/assets/js/main.min.js
    if [ -f "$FILEMAINJS" ]; then
        echo "$FILEMAINJS exists."

        sed -i "s|http://localhost:3002|https://api.gitdownloadmanager.com|g" ${FILEMAINJS}

    else
        echo "$FILEMAINJS does not exist."
    fi
    
    mkdir -p /var/www/$SUB_DOMAIN$DOMAIN/public/assets

    echo "Assets folder copy to dest"
    cp -vr /usr/local/bin/$SERVICE_FOLDER/assets/* /var/www/$SUB_DOMAIN$DOMAIN/public/assets
fi

# ip route get 8.8.8.8 | head -1 | awk '{print $7}'

# Delete *.dev.env file
#rm /usr/local/bin/$SERVICE_FOLDER/*.env

# Syslog file and conf created
LOGFOLDER=/var/log/gitdm
if [ -d "$LOGFOLDER" ]; then
    echo "$LOGFOLDER exists."
else
    echo "$LOGFOLDER does not exist."

    # Log file and Cron Log File Folder
    mkdir -p /var/log/gitdm/ /var/log/gitdm/cron/
    chown syslog:adm /var/log/gitdm/
    chmod -R 0755 /var/log/gitdm/
fi

LOGCONFFILE=/etc/rsyslog.d/$SERVICE_FILENAME.conf
if [ -f "$LOGCONFFILE" ]; then
    echo "$LOGCONFFILE exists."
else
    echo "$LOGCONFFILE does not exist."

    # LogConf file
    cp -vr /usr/local/bin/templates/syslog.conf $LOGCONFFILE
    sed -i "s|SERVICENAME|${SERVICE_FILENAME}.service|g" $LOGCONFFILE
fi

# Symlink service unit file
FILE=/etc/systemd/system/$SERVICE_FILENAME.service
if [ -f "$FILE" ]; then
    echo "$FILE exists."
else
    echo "$FILE does not exist."

    # Symlink service unit file
    ln -sf /usr/local/bin/$SERVICE_FOLDER/$SERVICE_FILENAME.service /etc/systemd/system/$SERVICE_FILENAME.service
fi

# Systemctl register and start/reload
if systemctl is-active --quiet $SERVICE_FILENAME.service; then
    echo "$SERVICE_FILENAME service restart."

    systemctl restart $SERVICE_FILENAME.service
else
    echo "$SERVICE_FILENAME service enabled and started."

    systemctl enable --now $SERVICE_FILENAME.service
fi

#systemctl enable $SERVICE_FILENAME
#systemctl restart $SERVICE_FILENAME
echo "$SERVICE_FILENAME.service status"
systemctl --no-pager status $SERVICE_FILENAME.service

# Eğer service dosyalarında bir değişiklik olduysa çalışması gerek
#echo "Systemctl deamon reload"
#systemctl daemon-reload

# Syslog service has to restart
#echo "Systemctl rsyslog service restart"
#systemctl restart rsyslog.service

# Nginx service has to restart
#echo "Nginx service restart"
#systemctl restart nginx.service