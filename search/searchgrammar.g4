grammar searchgrammar;


// Parser rules
clause: expr+;
expr
    : term              #termExpr
    | '(' expr ')'      #parensExpr
    | expr OR expr      #orExpr
    | expr AND expr     #andExpr
    | expr CMP expr     #cmpExpr
    | NOT expr          #notExpr
;

term
    : literal_number
    | reference
    | literal_string;

reference: STRING;
literal_number: INTEGER | DECIMAL;
literal_string: '"' STRING '"' | '\'' STRING '\'';



// Lexer rules
fragment A: ('a' | 'A');
fragment N: ('n' | 'N');
fragment D: ('d' | 'D');

fragment O: ('o' | 'O');
fragment R: ('r' | 'R');

fragment T: ('t' | 'T');

// Operators
OR: (O R) | '||';
AND: (A N D) | '&&';
NOT: (N O T) | '!';

// Literals
INTEGER: '-'? [0-9]+;
DECIMAL: '-'? [0-9]+ '.' [0-9]+;

// Comparisons
CMP: '=' | '==' | '!=' | '>' | '>=' | '<' | '<=';

fragment LOWERCASE  : [a-z] ;
fragment UPPERCASE  : [A-Z] ;
STRING
    // TODO: Support more special characters?
    : (LOWERCASE | UPPERCASE)+ (LOWERCASE | UPPERCASE | [0-9] | '_' | '-')*
    ;


// Skip whitespace
WS    : [ \t\r\n]+ -> skip ;