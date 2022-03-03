@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
  IF "%&amg%"=="nhk" GOTO k1w
  ECHO This adventure game must be started properly first. See Readme.txt!

GOTO l5w

 :k1w
 IF "%1"=="the" GOTO j1w
 GOTO s1w
 :j1w
 TAKE %2
 :s1w
 IF "%1"=="tank" GOTO h1w
 GOTO d1w
 :h1w

   ECHO It is far too heavy.
   GOTO l5w
 :d1w

  SET &a5w=l5w
  IF "%&&l5w%"=="%&l5w%" IF "%1"=="tank" GOTO g1w
  GOTO f1w
:g1w
ECHO You take:
 ECHO tank
 SET &a5w=b5w
 SET &&l5w=y7w
 :f1w
 IF "%&&a5w%"=="%&l5w%" IF "%1"=="turtle" GOTO p1w

  GOTO z1w

   :p1w
   ECHO You take:
   ECHO turtle
   SET &a5w=b5w
   SET &&a5w=y7w
   :z1w
   IF "%&&k5w%"=="%&l5w%" IF "%1"=="gun" GOTO o1w

 GOTO x1w
 REM When the sentence is here, open delimiter!
 :o1w
 ECHO You take:

ECHO gun
SET &a5w=b5w
SET &&k5w=y7w

:x1w
IF "%&&s5w%"=="%&l5w%" IF "%1"=="napalm" GOTO i1w

 REM Is delimiter replaced?
 GOTO c1w
 :i1w
 ECHO You take:

  ECHO napalm
SET &a5w=b5w
SET &&s5w=y7w
:c1w

 IF "%&a5w%"=="l5w" GOTO u1w
 GOTO v1w
 :u1w
 REM Is value here?
 ECHO You can't see it here.
 :v1w
 :l5w
