@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
   IF "%&amg%"=="nhk" GOTO j3e
   ECHO This adventure game must be started properly first. See Readme.txt!
   GOTO l5w

    :j3e
    IF "%1"=="the" IF "%3"=="at" IF "%4"=="the" GOTO h3e
    REM Is the object still present?
    GOTO d3e
    REM Examine constant
    :h3e

  THROW %2 %5
  :d3e
  IF "%1"=="the" IF "%3"=="at" GOTO g3e
  REM Show tool when room is not considered
  GOTO f3e
  :g3e
  THROW %2 %4
  :f3e

   IF "%1"=="the" GOTO p3e
   REM File isn't removed
   GOTO z3e
   :p3e
   THROW %2 %3
   :z3e
   IF "%2"=="at" IF "%3"=="the" GOTO o3e
   REM Is the sentence still valid?
   GOTO x3e
   :o3e

    THROW %1 %4
    :x3e

     IF "%2"=="at" GOTO i3e

   GOTO c3e
   :i3e
 THROW %1 %3

  :c3e
  IF "%1"=="napalm" GOTO u3e
  GOTO v3e
  :u3e
  IF NOT "%&&s5w%"=="y7w" GOTO y3e
  REM Reveal sentence using delimiter
  GOTO b3e
  :y3e
  ECHO You don't have it.
  GOTO l5w
  :b3e

IF "%2"=="godzilla" IF "%&l5w%"=="k3w" GOTO t3e
GOTO n3e
:t3e
ECHO With a nice throw, the bottle hits Godzilla and breaks. Godzilla wakes up
ECHO and sees you. The pain makes him mad, and so he fries you.
PAUSE
CLS
l5w
:n3e
SET &a5w=l5w
IF "%2"=="ground" IF "%&l5w%"=="k3w" GOTO r3e
 GOTO m3e
 :r3e

  SET &a5w=b5w
  :m3e
  IF "%2"=="mountain" IF "%&l5w%"=="k3w" GOTO w3e
  GOTO e3e

:w3e
SET &a5w=b5w
:e3e

  IF "%&a5w%"=="b5w" GOTO l9e

   GOTO q9e
   :l9e
   ECHO The napalm hits the ground, and starts burning. Soon the smell awakes
   ECHO Godzilla, who rizes to his feet. He then spots something in the distance

    ECHO that makes him mad. He rushes off towards the city. You decide to follow
    ECHO him.
    SET &&s5w=n7w
    SET &f5w=g7w
    SET &l5w=l3w
    ECHO.
    LOOK
    GOTO l5w
    :q9e
    ECHO That would only cause a mess.
    GOTO l5w

     :v3e
     IF "%1"=="turtle" GOTO k9e
     GOTO a9e

   :k9e
   IF NOT "%&&a5w%"=="y7w" GOTO j9e
   GOTO s9e
   :j9e
   ECHO You don't have it.
   GOTO l5w

 :s9e
 IF "%2"=="dog" IF "%&l5w%"=="r7w" GOTO h9e
 GOTO d9e
:h9e
ECHO You throw the little turtle at the guard dog, which starts to play with it.
ECHO They seem to get along pretty well.
SET &&a5w=n7w
SET &j5w=z7w
GOTO l5w
:d9e
ECHO Why would you throw away the turtle like that? What are you trying to

 ECHO acheive?
 GOTO l5w
 :a9e

ECHO Why? It doesn't seem sensible!

 :l5w
