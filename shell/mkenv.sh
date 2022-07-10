# 自Mac環境でファイルアップロードするshell
cd ~/tekito-app/front
npm run build

cd ~/tekito-app
GOOS=linux GOARCH=amd64 go build

echo sending files to www.tekito-app.com...
scp -r -i ~/tkysn1028.pem ~/tekito-app/front/dist ec2-user@www.tekito-app.com:~/
scp -r -i ~/tkysn1028.pem ~/tekito-app/db ec2-user@www.tekito-app.com:~/
scp -i ~/tkysn1028.pem ~/tekito-app/tekito-app ec2-user@www.tekito-app.com:~/


echo OK
while true
do
echo Do you want to ssh www.tekito-web.com?[y/n]
read str
if [ $str = y ]; then
    echo ssh to www.tekito-web.com
    ssh -i ~/tkysn1028.pem ec2-user@www.tekito-app.com
    break
fi
if [ $str = n ]; then
    break
fi
done
exit 0