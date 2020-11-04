#ifndef TOKEN_H
#define TOKEN_H

typedef struct token token_t;
typedef enum types types;

enum types {
	BEGINNING,
	END,
	COMMENT,
	IDKYET,
	NUMBER,
	IDENTIFIER,
	OPERATOR,
	LPAREN,
	RPAREN,
	LBRACK,
	RBRACK,
	COMMA,
	PERIOD,
	SEMICOLON,
	ADD,
	SUBTRACT,
	MULTIPLY,
	DIVIDE,
	SINGLEQUOTE,
	DOUBLEQUOTE,
	GREATERTHAN,
	LESSTHAN,
	EQUALS,
	NOTOP,
	INVALID
};

extern const char* typesAsStrings[];

struct token {
	types type;
	token_t* next;
	unsigned int len;
	char* val;
};

token_t* createToken(types type, char* token, int len);
void deleteTokens(token_t* head);

#endif
