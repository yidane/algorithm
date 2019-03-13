#!/usr/bin/env bash

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

        desGit=${des}/.gitignore
        if [[ -a ${desGit} ]];then
            rm -f ${desGit}
        fi;
    done
done