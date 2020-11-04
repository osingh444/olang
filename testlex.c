#include <stdlib.h>
#include <stdio.h>
#include "lexer.h"
#include "token.h"

void printTokens(lexer_t* lex) {
	token_t* curr = lex->tokens;
	while(curr != NULL) {
		printf("token value: %s, type of token: %s, length of token: %d\n", curr->val, typesAsStrings[curr->type], curr->len);
		curr = curr->next;
	}
}

int main() {
	char* t1 = "./test/olang/simple.ol";
	printf("lexing simple.ol\n");
	lexer_t* lex1 = createLexer(t1);
	printf("lexer file has length %d\n", lex1->len);
	lex(lex1);
	printTokens(lex1);
	deleteLexer(lex1);

	char* t2 = "./test/olang/operators.ol";
	printf("lexing operators.ol\n");
	lexer_t* lex2 = createLexer(t2);
	printf("lexer file has length %d\n", lex2->len);
	lex(lex2);
	printTokens(lex2);
	deleteLexer(lex2);

	exit(1);
}
