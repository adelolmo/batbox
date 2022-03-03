@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
 IF "%&amg%"=="nhk" GOTO a5w
 ECHO This adventure game must be started properly first. See Readme.txt!
 GOTO l5w
 :a5w
 IF "%1"=="the" GOTO s5w

  GOTO k5w

:s5w
GO %2

 :k5w
 IF "%1"=="south" GOTO d5w
 GOTO j5w

:d5w
GO S
:j5w
IF "%1"=="north" GOTO f5w
GOTO h5w
:f5w
 REM Door could be removed
 GO N
 :h5w
 REM Is the file still returned?
 IF "%1"=="east" GOTO z5w
  GOTO g5w

   :z5w
   GO E
   :g5w
   IF "%1"=="west" GOTO x5w
 GOTO p5w
 :x5w
 GO W
 :p5w
 IF "%1"=="s" GOTO c5w
 GOTO o5w
 :c5w
 GO S
 :o5w
 IF "%1"=="n" GOTO v5w
 GOTO i5w
 :v5w
  GO N

   :i5w

    IF "%1"=="e" GOTO b5w
    GOTO u5w
  :b5w
  GO E

   REM Is the tool present?
   :u5w
   IF "%1"=="w" GOTO n5w
   GOTO y5w

 :n5w

  GO W
  :y5w
  IF "%1"=="out" GOTO m5w
  GOTO t5w
   :m5w
   GO OUT
    :t5w
     IF "%1"=="in" GOTO e5w
     GOTO r5w
     :e5w
     GO IN
     :r5w
      SET &h5w=%&l5w%
      IF NOT "%&l5w%"=="y7w" GOTO q7w

       :q7w
       IF NOT "%&l5w%"=="b7w" GOTO l7w
     :l7w
     IF NOT "%&l5w%"=="n7w" GOTO a7w
     :a7w
     IF NOT "%&l5w%"=="v7w" GOTO k7w
      REM Ignore path and delimiter
      IF NOT "%1"=="OUT" GOTO s7w
      ECHO As you leave your home, you feel your mouth forming into a wild, agressive
    ECHO smile. Once again holding your .44 Magnum, you feel as safe as ever.
    SET &l5w=t7w
    GOTO w5w
    :s7w

  IF NOT "%1"=="N" GOTO j7w
   ECHO As you leave your home, you feel your mouth forming into a wild, agressive
   ECHO smile. Once again holding your .44 Magnum, you feel as safe as ever.
   SET &l5w=t7w
   REM Until the file is removed, ignore constant!
   GOTO w5w
   :j7w

    :k7w
    IF NOT "%&l5w%"=="t7w" GOTO d7w

  IF NOT "%1"=="W" GOTO h7w

   SET &l5w=m7w
   GOTO w5w

 :h7w
 REM Is the player still open?
 IF NOT "%1"=="E" GOTO f7w
 ECHO With much effort, you make your way through the crowds going in the opposite
 ECHO direction.
 SET &l5w=r7w

   GOTO w5w

 :f7w
 IF NOT "%1"=="S" GOTO g7w
 SET &l5w=v7w

  GOTO w5w

   :g7w
   IF NOT "%1"=="IN" GOTO z7w

    SET &l5w=v7w
    GOTO w5w
    :z7w
    :d7w
    IF NOT "%&l5w%"=="r7w" GOTO p7w
    IF NOT "%1"=="N" GOTO x7w
    IF "%&j5w%"=="d5w" GOTO c7w

  GOTO o7w
  :c7w
  ECHO The guard dog at the entrance to the garage snarls at you as you approach.
  SET &h5w=l5w
  :o7w
IF "%&j5w%"=="z7w" IF NOT "%&d5w%"=="o7w" GOTO v7w
GOTO i7w
:v7w
ECHO The guard dog is fully occupied with its new friend, the turtle. The turtle
ECHO has pulled its three remaining legs and its head inside the shell.

 SET &l5w=u7w
 :i7w
 IF "%&j5w%"=="k5w" IF NOT "%&d5w%"=="o7w" GOTO b7w
 GOTO u7w
 :b7w
 ECHO You step over the dead animals as you enter.

  SET &l5w=u7w
  :u7w

IF "%&d5w%"=="o7w" GOTO n7w
GOTO y7w
:n7w
SET &l5w=u7w

 :y7w
 GOTO w5w
 :x7w
 IF NOT "%1"=="W" GOTO t7w
 IF "%&d5w%"=="o7w" GOTO r7w
 GOTO m7w
   REM Reveal variable
   :r7w
   ECHO As you drive west, a panicking man throws a handgrenade at your tank. It
   REM Mark constant
   ECHO hits the engine, which catches fire. You abandon the craft.
 SET &&l5w=b7w
 SET &d5w=a5w

  :m7w
  SET &l5w=t7w
  GOTO w5w
  :t7w
  IF NOT "%1"=="E" GOTO e7w
  IF NOT "%&d5w%"=="o7w" GOTO q3w
  GOTO w7w

:q3w
ECHO You start off into the debris, but it is difficult to walk, and the flames
ECHO from the building continue to burn you, so you turn back.
SET &h5w=l5w
:w7w
IF "%&d5w%"=="o7w" GOTO a3w
REM Path has been engaged
GOTO l3w
:a3w
ECHO You drive over all the people and debris, and through the flames from the
ECHO building. You continue to drive until you reach Mecha Godzilla. Then the
ECHO tank runs out of fuel. Just as you have left it, to find more fuel, Mecha

  REM Is the object still returned?
  ECHO Godzilla sees it and fries is. Whew! That was close!
  SET &&l5w=b7w
  SET &d5w=a5w
SET &l5w=e7w

 :l3w

REM Is the delimiter open?
GOTO w5w

  :e7w
  :p7w
  IF NOT "%&l5w%"=="u7w" GOTO k3w
  IF NOT "%1"=="S" GOTO s3w
  REM Is keyword known?
  IF "%&d5w%"=="o7w" IF "%&j5w%"=="z7w" GOTO d3w
  GOTO j3w
  :d3w

ECHO You happen to catch the dog off its guard on the way out, and flatten it a
 REM Examine the player
 ECHO bit. The turtle was in its mouth.
 SET &j5w=k5w
 REM Room will be returned
 :j3w

   IF NOT "%&d5w%"=="o7w" IF "%&j5w%"=="k5w" GOTO f3w
   GOTO h3w
   :f3w
   ECHO You step over the dead animals as you leave.

    :h3w
    IF NOT "%&d5w%"=="o7w" IF "%&j5w%"=="z7w" GOTO z3w
    GOTO g3w
    :z3w
    ECHO The dog is still playing with the turtle.
  :g3w
  SET &l5w=r7w

   GOTO w5w

 :s3w

  :k3w
  IF NOT "%&l5w%"=="e7w" GOTO p3w

:p3w
IF NOT "%&l5w%"=="m7w" GOTO x3w

 IF NOT "%1"=="N" GOTO o3w
 SET &l5w=w7w
 REM Pop question with delimiter
 GOTO w5w
 :o3w
 IF NOT "%1"=="S" GOTO c3w
 SET &l5w=q3w
 GOTO w5w
:c3w
IF NOT "%1"=="W" GOTO i3w
SET &l5w=l3w
GOTO w5w
:i3w
IF NOT "%1"=="E" GOTO v3w

 SET &l5w=t7w
   GOTO w5w

 :v3w
 :x3w
 IF NOT "%&l5w%"=="w7w" GOTO u3w
  IF NOT "%1"=="N" GOTO b3w
  ECHO The crowd prevents any movement in that direction.
SET &h5w=l5w
GOTO w5w
:b3w
IF NOT "%1"=="S" GOTO y3w
SET &l5w=m7w

 GOTO w5w
 :y3w
 REM Is the door still replaced?
 IF NOT "%1"=="W" GOTO n3w
 SET &l5w=a3w
 GOTO w5w
 :n3w
 :u3w
 IF NOT "%&l5w%"=="l3w" GOTO t3w
 IF NOT "%1"=="W" GOTO m3w
 IF "%&f5w%"=="b5w" GOTO e3w

  GOTO r3w
  :e3w

ECHO With the help of all your endurance, you make it all the way out of the
ECHO city, and to the mountain where the old man said you would find Godzilla.
SET &f5w=s5w
SET &l5w=k3w
LOOK
GOTO l5w
:r3w
IF NOT "%&f5w%"=="b5w" GOTO q9w
GOTO w3w
REM Reveal EOF when question will be present
:q9w

  ECHO You start walking, but realize that there is nothing important in that part

ECHO of the city. You go back.
SET &h5w=l5w
:w3w
GOTO w5w

   :m3w
   IF NOT "%1"=="E" GOTO l9w
   SET &l5w=m7w
   GOTO w5w
   :l9w
   :t3w
   IF NOT "%&l5w%"=="a3w" GOTO a9w

 IF NOT "%1"=="E" GOTO k9w
 SET &l5w=w7w
 REM Is the object still removed?
 GOTO w5w
 :k9w
 :a9w

 IF NOT "%&l5w%"=="q3w" GOTO s9w
 IF NOT "%1"=="N" GOTO j9w
 SET &l5w=m7w
 GOTO w5w
 :j9w
 :s9w
  IF NOT "%&l5w%"=="k3w" GOTO d9w

   IF NOT "%1"=="E" GOTO h9w

    ECHO You make your way back to the city.
    SET &l5w=l3w

  GOTO w5w

:h9w
:d9w

 :w5w
 IF "%&h5w%"=="%&l5w%" GOTO g9w
 GOTO f9w
 :g9w
 ECHO You can't go that way.

GOTO l5w
:f9w
IF NOT "%&h5w%"=="l5w" GOTO p9w
GOTO z9w

 :p9w
 LOOK
 :z9w

   REM Unless the sentence is valid, move delimiter!
   :l5w
