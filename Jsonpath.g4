grammar Jsonpath;

//
// Derived from https://github.com/stevenalexander/antlr4-jsonpath-grammar
//

jsonpath: '$' path EOF
        ;

path : nodeExpr+
     ;

nodeExpr : dotExpr
         | selectExpr
         ;

dotExpr : DOTS (Identifier | STAR | Length)
        ;

selectExpr : '[' (starExpr | rangeExpr | itemsExpr | nameExpr | filterExpr | scriptExpr)? ']'
           ;

rangeExpr : (startIndex=INT)? ':' (endIndex=INT (':' step=INT)?)?
          ;

itemsExpr : INT (',' INT)*
          ;

nameExpr: QUOTED
        ;

starExpr : STAR
         ;

filterExpr : '?(' queryExpr ')'
           ;

scriptExpr : '(' queryExpr ')'
           ;

queryExpr : '@.' Identifier
          | '@.' Identifier COMP NUMBER
          | '@.length' OP INT
          | '@.' Identifier '=~' REGEX
          ;

Identifier : [a-zA-Z][a-zA-Z0-9_]*
           ;

INT : '-'? ('0'..'9')+
    ;

NUMBER : '-'? ('0'..'9')+ ('.' ('0'..'9')+)?
       ;

QUOTED : '\'' (~'\'')* '\''
       ;

REGEX : QUOTED
      | '/' (~'/')* '/' [imsU]*
      ;

DOTS : '.'
     | '..'
     ;

STAR : '*'
     ;

Length : 'length()'
       ;

OP : '+' | '-' | '/' | '*'
   ;

COMP : '>' | '>=' | '<' | '<=' | '==' | '!='
     ;

WS : [ \t\n\r]+ -> skip
   ;
