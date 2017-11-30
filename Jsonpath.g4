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

dotExpr : DOTS 
        | DOTS (Identifier | STAR | Length )
        ;

selectExpr : '[' (starExpr | rangeExpr | itemsExpr | namesExpr | filterExpr | scriptExpr)? ']'
           ;

rangeExpr : (startIndex=INT)? ':' (endIndex=INT (':' step=INT)?)?
          ;

itemsExpr : INT (',' INT)*
          ;

namesExpr: QUOTED (',' QUOTED)*
        ;

starExpr : STAR
         ;

filterExpr : '?(' queryExpr ')'
           ;

scriptExpr : '(' queryExpr ')'
           ;

queryExpr : '@.' exists=Identifier
          | '@.' name=Identifier op=COMP value=(INT | NUMBER | QUOTED)
          | '@.' name=Identifier op='=~' value=REGEX
          ;

Identifier : [a-zA-Z][a-zA-Z0-9_]*
           ;

INT : '-'? ('0'..'9')+
    ;

NUMBER : INT ('.' ('0'..'9')+)?
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
