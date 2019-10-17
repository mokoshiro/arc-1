#!/bin/sh

for f in $(find . -name \*\.pb\.go -not -path "./vendor/*"); do
    echo ${f}
    unlink ${f}
done;

for f in $(find . -name \*\.pb\.validate\.go -not -path "./vendor/*"); do
    echo ${f}
    unlink ${f}
done;

for f in $(find . -name \*\.mock\.go -not -path "./vendor/*"); do
    echo ${f}
    unlink ${f}
done;
