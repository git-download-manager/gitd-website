# Gitd Download Manager CLI
Cli tools for Gitdownloadmanager.com website

## Commands

Useful commands

```
build all services binary files with assets folders
./dmcli build go --config .build.prod.yaml

delete expires folders inside git download manager temp folder (for api service)
./dmcli clear --config .clear.prod.yaml

deploy app on the server side
./dmcli deploy --config .deploy.prod.yaml

clone repositories analysis everyday and collects stats datas (look at the data/cron folder)
./dmcli stats --config .stats.prod.yaml
```

### Stats Command Output
> Via: https://sysadmins.co.za/bash-script-to-parse-and-analyze-nginx-access-logs/

Sample Log Data:
```
{"level":"info","ts":1663438847.30647,"msg":"success.clone.new","url":"https://github.com/uretgec/my-product-hunt","hostname":"github.com","owner":"uretgec","repository":"my-product-hunt"}
```

Example Output:
```
2022-09-21 01:54 Stats:

Top 10 Clone Repos:
--------------------
  46 https://github.com/WebKit/WebKit
  29 https://github.com/matomo-org/matomo-for-wordpress
  12 https://gitlab.com/Mstrodl/sonic-builds
   8 https://github.com/uretgec/my-product-hunt
   8 https://github.com/WebKit/Documentation
   7 https://bitbucket.org/tiagoharris/url-shortener
   6 https://github.com/matomo-org/tag-manager
   5 https://github.com/matomo-org/device-detector
   4 https://github.com/matomo-org/matomo-sdk-ios
   4 https://github.com/matomo-org/matomo

Top 10 Hostname:
--------------------
 121 github.com
  14 gitlab.com
   9 bitbucket.org

Top 10 Repository Owner:
--------------------
  58 WebKit
  49 matomo-org
  12 Mstrodl
  10 uretgec
   7 tiagoharris
   3 github
   2 atlassian
   1 inkscape
   1 grafana
   1 gitlab-org

Top 10 Repository:
--------------------
  46 WebKit
  29 matomo-for-wordpress
  12 sonic-builds
   8 my-product-hunt
   8 Documentation
   7 url-shortener
   6 tag-manager
   5 device-detector
   4 matomo-sdk-ios
   4 matomo
```
