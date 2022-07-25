while true
do
echo First Make?[y/n]
read str
if [ $str = y ]; then
    echo input remote path
    read URL
    if [ -z "$URL"];then
    echo no URL.Please Input URL from Github.
    exit 0
    fi
    git init
    git branch -M main
    git remote add origin $URL
    git remote -v
    git add .
    git commit -m "created first commit."
    git push origin main
    break
fi
if [ $str = n ]; then
    echo input commit message
    read msg
    if [ -z "$msg" ];then
    echo no commit message.
    exit 0
    fi
    git add .
    git commit -m "$msg"
    git push origin main
    break
fi
done
exit 0