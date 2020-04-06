#!/bin/bash

failed=0

for i in `cat test_slugs.txt`; do
    ./moocist template --coursera-slug $i &> /dev/null
    res=$?
    url="https://www.coursera.org/learn/${i}"
    if [ $res -eq 0 ] ; then
        echo -e "${url}\t\033[32mOK\033[0m"
    else
        echo -e "${url}\t\033[31mNOK\033[0m"
        failed=1
    fi
done

if [ $failed -eq 1 ] ; then
    exit 1
fi
