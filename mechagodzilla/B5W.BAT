@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
   IF "%&amg%"=="nhk" GOTO f9e
    ECHO This adventure game must be started properly first. See Readme.txt!
    GOTO l5w
    :f9e
    REM Is question left?
    IF "%1"=="the" IF "%3"=="to" GOTO z9e
  GOTO g9e
  :z9e

   REM Is the variable present?
   ASK %4 %2

 :g9e
 IF "%1"=="the" GOTO x9e
 GOTO p9e
 :x9e

  ASK %3 %2
  REM Is the tool considered?
  :p9e
  IF "%2"=="to" GOTO c9e
  GOTO o9e

   :c9e
   ASK %3 %1
   :o9e
   ASK %2 %1
   :l5w
