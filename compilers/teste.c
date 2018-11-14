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

int goto_label;

Param* newParam(int); // só um atalho

int main(void){
    printf("declaração de variáveis do main\n");
    AR* ARs = malloc(sizeof(AR));
    AR* ARlast   = ARs;
    AR* ARtopush = NULL;
    int func_ret_val = 0;

    printf("preparando para chamar func1\n");
    // push main AR
    ARtopush = malloc(sizeof(AR));
    {
        Param* p = newParam(4);
        ARtopush->params = p;
    }
    ARtopush->ret = func_ret1;
    goto_label = func_call;
    goto lbl_push_AR;

    lbl_func_ret1:{
        printf("retornou de func1\n");
        printf("preparando para chamar func2\n");
        // push main AR
        ARtopush = malloc(sizeof(AR));
        {
            Param* p = newParam(10);
            ARtopush->params = p;
        }
        ARtopush->ret    = func_ret2;
        goto_label = func_call;
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
        func_ret_val = x+1;
        goto_label = pop_AR;
        goto lbl_goto_switch;
    }

    lbl_push_AR:{
        printf("pushing AR...\n");
        ARtopush->dyn = ARlast;
        ARlast = ARtopush;
        goto lbl_goto_switch;
    }

    lbl_pop_AR:{
        printf("popping AR...\n");
        goto_label = ARlast->ret;
        ARlast = ARlast->dyn;
        // TODO: free params
        free(ARlast->next);
        ARlast->next = NULL;
        goto lbl_goto_switch;
    }
    
    lbl_goto_switch:{
        printf("switch(%d)=",goto_label);
        switch(goto_label){
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

Param* newParam(int value){
    Param* p = malloc(sizeof(AR));
    p->value = value;
    return p;
}