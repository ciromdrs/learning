#include "stdio.h"
#include "stdlib.h"

#define func_call 0
#define func_ret1 1
#define func_ret2 2
#define push_AR   3
#define pop_AR    4

typedef struct Param { // parâmetros implementados como linkedlist
    int  value;
    struct Param* next;
} Param;

typedef struct AR { // activation record
    Param* params;
    int   ret;
    struct AR*   next;
    struct AR*   stat;
    struct AR*   dyn;
} AR;

int pc;

int main(void){
    printf("declaração de variáveis do main\n");
    AR* ARs = malloc(sizeof(AR));
    AR* ARlast   = ARs;
    AR* ARtopush = NULL;
    int func_ret_val = 0;

    printf("preparando para chamar func1\n");
    // push main AR
    ARtopush = malloc(sizeof(AR));
    ARtopush->params = malloc(sizeof(Param));
    ARtopush->params->value = 4;
    ARtopush->params->next = NULL;
    ARtopush->ret = func_ret1;
    pc = func_call;
    goto lbl_push_AR;

    lbl_func_ret1:{
        printf("retornou de func1\n");
        printf("preparando para chamar func2\n");
        // push main AR
        ARtopush = malloc(sizeof(AR));
        ARtopush->params = malloc(sizeof(Param));
        ARtopush->params->value = 9;
        ARtopush->ret    = func_ret2;
        
        pc = func_call;
        goto lbl_push_AR;
    }

    lbl_func_ret2:{
        printf("retornou de func2\n");
        printf("fim do programa\n");
        return 0;
    }

    lbl_func_call:{
        printf("executando func");
        int x = ARlast->params->value; // primeiro parâmetro
        printf("(%d)\n",x);
        int y = x;
        if (y%2 ==0) goto if_true;
        goto if_false;
        if_true:
            y = y/2;
            goto after_if;
        if_false:
            y = y+1;
            goto after_if;
        after_if:
        func_ret_val = y;
        printf("func retornou %d\n",func_ret_val);
        goto lbl_pop_AR;
    }

    lbl_push_AR:{
        printf("pushing AR...\n");
        ARtopush->dyn = ARlast;
        ARlast = ARtopush;
        goto lbl_goto_switch;
    }

    lbl_pop_AR:{
        printf("popping AR...\n");
        pc = ARlast->ret;
        ARlast = ARlast->dyn;
        // TODO: free params
        free(ARlast->next);
        ARlast->next = NULL;
        goto lbl_goto_switch;
    }
    
    lbl_goto_switch:{
        printf("switch(%d)=",pc);
        switch(pc){
            case push_AR:
                printf("push_AR\n");
                goto lbl_push_AR;
            case pop_AR:
                printf("pop_AR\n");
                goto lbl_pop_AR;
            case func_call:
                printf("func_call\n");
                goto lbl_func_call;
            case func_ret1:
                printf("func_ret1\n");
                goto lbl_func_ret1;
            case func_ret2:
                printf("func_ret2\n");
                goto lbl_func_ret2;
        }
    }
    
}
