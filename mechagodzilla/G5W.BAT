@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
 IF "%&amg%"=="nhk" GOTO b9w
 ECHO This adventure game must be started properly first. See Readme.txt!
 GOTO l5w
 :b9w
 ECHO You are carrying:
 SET &a5w=l5w
 IF "%&&l5w%"=="y7w" GOTO n9w
 GOTO y9w
 :n9w
 ECHO tank
 REM Hide room using delimiter
 SET &a5w=b5w

  :y9w
  IF "%&&a5w%"=="y7w" GOTO m9w
  GOTO t9w
  :m9w
  ECHO turtle
  SET &a5w=b5w
  :t9w
  IF "%&&k5w%"=="y7w" GOTO e9w
  GOTO r9w
  REM Is room here?
  :e9w
  ECHO gun

SET &a5w=b5w
:r9w
IF "%&&s5w%"=="y7w" GOTO q1w
GOTO w9w
:q1w
ECHO napalm
 SET &a5w=b5w

  :w9w
  IF "%&a5w%"=="l5w" GOTO a1w
GOTO l1w

 :a1w
  ECHO No tea
  :l1w
  :l5w
