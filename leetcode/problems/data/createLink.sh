#!/usr/bin/env bash

## 因为题目太多
## 所以把题目每一百个放在子文件夹中
## 结果共有链表对象需要复制多分
## 就想到使用硬链接方式，在子文件夹中创建链接

problems=$(dirname $(cd `dirname $0`; pwd))
src=${problems}/data
files=(0.Array.go 0.ListNode.go)
subFolds=(100 200 300 400 500 600 700 800 900)
for subFold in ${subFolds[@]} ; do
    des=${problems}/${subFold}
    for file in ${files[@]} ; do
        srcFile=${src}/${file}
        desFile=${des}/${file}
        if [[ -a ${desFile} ]];then
            rm -f ${desFile}
            echo delete ${desFile} succeed
        fi;
        ln ${srcFile} ${desFile}
        echo link ${srcFile} and ${desFile} succeed

        desGit=${des}/.gitignore
        if [[ ! -a ${desGit} ]];then
            touch ${desGit}
        fi;
        sed -i 's/^'${file}'$//g' ${desGit} && sed -i /^[[:space:]]*$/d ${desGit} && echo ${file} >> ${desGit}
    done
done