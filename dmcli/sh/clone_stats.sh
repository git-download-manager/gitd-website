#!/bin/bash

# Before run: chmod a+x install.sh
# Local Variables
ARGS=("$@")
LOGFILE="${ARGS[0]}"
STATS_DATE=$(date '+%Y-%m-%d %H:%M')

echo ""
echo "$STATS_DATE Stats:"

# via: https://sysadmins.co.za/bash-script-to-parse-and-analyze-nginx-access-logs/
get_request_clone_repos_top_url() {
    echo ""
    echo "Top 10 Clone Repos:"
    echo "--------------------"

    cat $LOGFILE \
    | grep -o '"url":"[^"]*' | grep -o '[^"]*$' \
    | sort | uniq -c \
    | sort -rn \
    | head -10
}

get_request_clone_repos_top_hostname() {
    echo ""
    echo "Top 10 Hostname:"
    echo "--------------------"

    cat $LOGFILE \
    | grep -o '"hostname":"[^"]*' | grep -o '[^"]*$' \
    | sort | uniq -c \
    | sort -rn \
    | head -10
}

get_request_clone_repos_top_owner() {
    echo ""
    echo "Top 10 Repository Owner:"
    echo "--------------------"

    cat $LOGFILE \
    | grep -o '"owner":"[^"]*' | grep -o '[^"]*$' \
    | sort | uniq -c \
    | sort -rn \
    | head -10
}

get_request_clone_repos_top_repository() {
    echo ""
    echo "Top 10 Repository:"
    echo "--------------------"

    cat $LOGFILE \
    | grep -o '"repository":"[^"]*' | grep -o '[^"]*$' \
    | sort | uniq -c \
    | sort -rn \
    | head -10
}

# exec
get_request_clone_repos_top_url
get_request_clone_repos_top_hostname
get_request_clone_repos_top_owner
get_request_clone_repos_top_repository