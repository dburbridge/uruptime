# uruptime
Uptime report for Uptime Robot

docker build . -t uruptime

docker run -d -e "API_KEY=xxxx" -p 8080:8080 uruptime
