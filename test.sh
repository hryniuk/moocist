#!/bin/bash

for i in `cat test_slugs.txt`; do
    ./moocist template --coursera-slug $i &> /dev/null
    res=$?
    if [ $res -eq 0 ] ; then
        echo -e "${i}\t\033[32mOK\033[0m"
    else
        echo -e "${i}\t\033[31mNOK\033[0m"

    fi
done