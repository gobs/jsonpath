grammar Jsonpath;

/*
 Limited set of JsonPath notation

 TODO
 - bracketnotation
 - Check against spec description (http://goessner.net/articles/JsonPath/), e.g. '..' for recursive descent
}
*/

jsonpath: '$' path EOF
        ;

path : nodeExpr+
     ;

nodeExpr : dotExpr
         | selectExpr
         ;

dotExpr : DOTS (Identifier | STAR)
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

queryExpr : queryExpr ('&&' queryExpr)+
          | queryExpr ('||' queryExpr)+
          | '*'
          | '@.' Identifier
          | '@.' Identifier COMP NUMBER
          | '@.length' OP INT
          | '@.' Identifier '=~' QUOTED
          ;

Identifier : [a-zA-Z][a-zA-Z0-9_]*
           ;

INT : '-'? ('0'..'9')+
    ;

NUMBER : '-'? ('0'..'9')+ ('.' ('0'..'9')+)?
       ;

QUOTED : '\'' (~'\'')* '\''
       ;

DOTS : '.'
     | '..'
     ;

STAR : '*'
     ;

OP : '+' | '-' | '/' | '*'
   ;

COMP : '>' | '>=' | '<' | '<=' | '==' | '!='
     ;

WS : [ \t\n\r]+ -> skip
   ;
