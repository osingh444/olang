#include <stdlib.h>
#include <stdio.h>
#include "token.h"

const char* typesAsStrings[] = {"BEGINNING", "END", "COMMENT", "IDKYET", "NUMBER", "IDENTIFIER", "OPERATOR", "LPAREN", "RPAREN",
"LBRACK", "RBRACK", "COMMA", "PERIOD", "SEMICOLON", "ADD", "SUBTRACT", "MULTIPLY", "DIVIDE", "SINGLEQUOTE", "DOUBLEQUOTE",
"GREATERTHAN", "LESSTHAN", "EQUALS", "NOTOP", "INVALID"};

static void allocation_failed(void) {
  printf("Out of memory.\n");
  exit(0);
}

token_t* createToken(types type, char* token, int len) {
	token_t* t = malloc(sizeof(token_t));
	if(!t) {
		allocation_failed();
	}
	t->type = type;
	t->val = token;
	t->len = len;
	t->next = NULL;
	return t;
}

void deleteTokens(token_t* head) {
	while(head != NULL) {
		free(head->val);
		token_t* next = head->next;
		free(head);
		head = next;
	}
}
