#ifndef LEXER_H
#define LEXER_H
#include "token.h"

typedef struct lexer lexer_t;

struct lexer {
	int numTokens;
	token_t* tokens;
	unsigned long len;
	char* file;
};

lexer_t* createLexer(char* filename);
void lex(lexer_t* lexer);
token_t* advanceToken(int* position, char* fileptr);
void deleteLexer(lexer_t* lexer);

#endif
