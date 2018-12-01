rm -f sum_*; while true; do cat input;done|(t=0;while read x; do t=$[t$x]; echo $x $t;if [ -f sum_$t ]; then echo "ENDDDDDDDDDDDDDDDDDDDDDDDDD";break;fi;touch sum_$t;done )
