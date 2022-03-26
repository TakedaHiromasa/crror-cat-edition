#include <stdio.h>

int     main(void)
{
        /* 変数宣言がない */

        for (i = 1; i <= 9; i++) {
          printf("◆ %d の段\n", i);
          for (j=1; j<=9; j++) {
                printf("%d x %d = %d\n", i, j, i * j);
          }
          printf("----------------\n", i);
        }
        return (0);
}