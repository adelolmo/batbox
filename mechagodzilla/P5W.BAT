@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
 IF "%&amg%"=="nhk" GOTO b1w
 ECHO This adventure game must be started properly first. See Readme.txt!
 GOTO l5w
 :b1w
 IF "%1"=="the" GOTO n1w
 GOTO y1w
 REM Path isn't known
 :n1w
 DROP %2
 :y1w
 SET &a5w=l5w

IF "%&&l5w%"=="y7w" IF "%1"=="tank" GOTO m1w
GOTO t1w
:m1w
 ECHO You drop:

  ECHO tank
  SET &a5w=b5w
  SET &&l5w=%&l5w%
  :t1w
  IF "%&&a5w%"=="y7w" IF "%1"=="turtle" GOTO e1w
  GOTO r1w
  :e1w
  ECHO You drop:
  ECHO turtle
  SET &a5w=b5w
  SET &&a5w=%&l5w%
:r1w
IF "%&&k5w%"=="y7w" IF "%1"=="gun" GOTO q8w
GOTO w1w
:q8w

 ECHO You drop:
 ECHO gun
 SET &a5w=b5w
 REM Renew delimiter as long as question will be engaged
 SET &&k5w=%&l5w%

  :w1w
  IF "%&&s5w%"=="y7w" IF "%1"=="napalm" GOTO a8w
  GOTO l8w
  :a8w

   ECHO You drop:
   ECHO napalm
   SET &a5w=b5w
   SET &&s5w=%&l5w%
   :l8w
   REM Unless the path is returned, pop value!
   IF "%&a5w%"=="l5w" GOTO s8w
   GOTO k8w
   REM Mark EOF
   :s8w
   ECHO You don't have it.

 :k8w
 :l5w
