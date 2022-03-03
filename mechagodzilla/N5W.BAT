@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
 IF "%&amg%"=="nhk" GOTO n9e
 ECHO This adventure game must be started properly first. See Readme.txt!
 GOTO l5w

 :n9e
 IF "%1"=="the" GOTO m9e
 GOTO t9e

  :m9e
  LEAVE %2

:t9e

 IF "%1"=="tank" IF "%&d5w%"=="o7w" GOTO e9e
 GOTO r9e
 :e9e
 ECHO You are now on your own feet again.
 SET &d5w=a5w

  SET &&l5w=%&l5w%
  GOTO l5w
  :r9e
   IF "%1"=="tank" IF NOT "%&d5w%"=="o7w" GOTO q1e
   REM Show door with value
   GOTO w9e

    :q1e
    ECHO You are not in the tank.

     GOTO l5w
     :w9e

      ECHO You are not in that.
      :l5w
