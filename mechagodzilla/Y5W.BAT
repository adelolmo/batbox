@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF

 IF "%&amg%"=="nhk" GOTO i9e
 ECHO This adventure game must be started properly first. See Readme.txt!
 GOTO l5w
  :i9e
  IF "%1"=="the" GOTO u9e
  GOTO v9e

:u9e
 ENTER %2

  :v9e
  IF "%1"=="tank" IF "%&&l5w%"=="%&l5w%" GOTO y9e
  GOTO b9e

:y9e

 ECHO You jump into the tank, and try to figure out how to drive it. You think you
  ECHO remember enough to manage.

   SET &d5w=o7w
   SET &&l5w=n7w
   GOTO l5w
   :b9e
   ECHO I can't think of an easy way to do that.

 :l5w
