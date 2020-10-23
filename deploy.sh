#!/usr/bin/env bash
# 确保脚本抛出遇到的错误
set -e



# 如果是发布到自定义域名
# echo 'www.example.com' > CNAME

git init
git add .
git commit -m 'upd |> update code'

git push
# 如果发布到 https://<USERNAME>.github.io
# git push -f git@github.com:<USERNAME>/<USERNAME>.github.io.git master

# 如果发布到 https://<USERNAME>.github.io/<REPO>
# git push -f git@github.com:yinian-R/vuepress-demo.git master:gh-pages

cd -