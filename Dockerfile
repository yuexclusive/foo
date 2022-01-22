from alpine
workdir /app
add test .

# log
run mkdir log

# timezone
run sudo apk add tzdata
run sudo cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
run sudo echo "Asia/Shanghai" >  /etc/timezone


cmd ["/app/test"]