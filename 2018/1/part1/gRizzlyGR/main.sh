#!/bin/bash

result=0

for line in $(cat input)
do
	result=$(( result + line ))
done

echo $result