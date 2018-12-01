#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAXABS 10000000
char seen[MAXABS*2];

int main(void){
  memset(seen,0,MAXABS*2*sizeof(char));
  int t=0;
  for(int loop=0;;loop++){
    char sign;
    int abs;
    scanf("%c%i\n",&sign,&abs);
    if(sign=='+'){
      t=t+abs;
    }
    else if(sign=='-'){
      t=t-abs;
    }
    else{
      abort();
    }
    printf("%4i  %c%7i %7i\n",loop,sign,abs,t);
    int pos=MAXABS+t;
    if(seen[pos]==1){
      printf("FOUND\n");return 0;
    }
    seen[pos]=1;
  }
}
