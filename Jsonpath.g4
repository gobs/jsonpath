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

queryExpr : '@.' exists=Identifier
          | '@.' name=Identifier op=COMP value=(NUMBER | QUOTED)
          | '@.' name=Identifier op='=~' value=REGEX
          ;

Identifier : [a-zA-Z][a-zA-Z0-9_]*
           ;

NUMBER : INT ('.' ('0'..'9')+)?
       ;

INT : '-'? ('0'..'9')+
    ;

QUOTED : '\'' (~'\'')* '\''
       ;

REGEX : '/' (~'/')* '/' ('i'|'u'|'m'|'S')*
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
