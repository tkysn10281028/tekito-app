while true
do
echo build files?[y/n]
read str
if [ $str = y ]; then
# 自Mac環境でファイルアップロードするshell
cd ~/tekito-app/front
npm run build

cd ~/tekito-app
GOOS=linux GOARCH=amd64 go build

echo sending files to www.tekito-app.com...
scp -r -i ~/tkysn1028.pem ~/tekito-app/front/dist ec2-user@35.76.152.13:~/
scp -i ~/tkysn1028.pem ~/tekito-app/tekito-app ec2-user@35.76.152.13:~/
scp -i ~/tkysn1028.pem ~/desktop/mybatchTest.jar ec2-user@35.76.152.13:~/
echo sending file succeeded.
break
fi
if [ $str = n ]; then
    break
fi
done

while true
do
echo Do you want to ssh www.tekito-web.com?[y/n]
read str
if [ $str = y ]; then
    echo ssh to www.tekito-web.com
    ssh -i ~/tkysn1028.pem ec2-user@35.76.152.13
    break
fi
if [ $str = n ]; then
    break
fi
done
exit 0