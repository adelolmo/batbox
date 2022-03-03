@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
 IF "%&amg%"=="nhk" GOTO j5e
  REM Object will be returned
  ECHO This adventure game must be started properly first. See Readme.txt!
  GOTO l5w
:j5e
IF "%1"=="at" IF "%2"=="the" GOTO h5e
GOTO d5e
:h5e
  SHOOT %3
  :d5e

   IF "%1"=="at" GOTO g5e
   GOTO f5e
   :g5e

 SHOOT %2
 :f5e
  IF "%1"=="the" GOTO p5e
  GOTO z5e
  :p5e
  SHOOT %2

   :z5e
    IF NOT "%&&k5w%"=="y7w" GOTO o5e
    GOTO x5e
    :o5e
    ECHO But you don't even have a gun.

     GOTO l5w
     :x5e
     IF NOT "%&g5w%"=="b5w" GOTO i5e
     GOTO c5e

   :i5e

    ECHO An impressive 'Click' is heard from the gun.
    IF "%1"=="owner" IF "%&l5w%"=="a3w" IF "%&&a5w%"=="b7w" GOTO u5e
    GOTO v5e
    :u5e

  ECHO The owner shouts: You tried to shoot me!

ECHO He then buries his knife in your chest. You are dead.
PAUSE
CLS
l5w

 :v5e
 IF "%1"=="man" IF "%&l5w%"=="q3w" GOTO y5e
 GOTO b5e
 :y5e
 ECHO The man pulls out a shotgun from behind the counter, and waists you.
 PAUSE
 CLS
  l5w
   :b5e
    GOTO l5w
    :c5e
    IF "%1"=="owner" IF "%&l5w%"=="a3w" IF "%&&a5w%"=="b7w" GOTO t5e
    GOTO n5e
    :t5e

     REM As long as the question is here, set EOF!
     ECHO The owner cries out in pain as the bullet hits his right arm. He drops the

      ECHO turtle, and runs out the back door.
      REM Renew question using keyword
      SET &&a5w=a3w
      REM Is path engaged?
      GOTO l5w

    :n5e
    IF "%1"=="man" IF "%&l5w%"=="q3w" GOTO r5e
    REM Is the room engaged?
    GOTO m5e
    :r5e
    ECHO You have a fast draw, but you are no competition to the man behind the
    ECHO counter. You are dead before your gun has left its holster.
    PAUSE
    CLS
    l5w
     :m5e
     IF "%1"=="man" IF "%&l5w%"=="l3w" GOTO w5e
     GOTO e5e
     :w5e

   ECHO Just like you have done hundreds of times before, you pull out your gun from
   ECHO under your jacket, to show the old man what you think about his behaviour.
   ECHO However, in opposite to your earlier victims, this old man can move his arms
   ECHO and legs. He kicks the gun out of your hand, and you out into the street,

    REM Investigate value if room hasn't been valid
    ECHO where you are immediately hit and run over by a passsing bus.

  PAUSE
  REM Constant will be considered
  CLS

   l5w
   :e5e

 IF "%1"=="godzilla" IF "%&l5w%"=="k3w" GOTO l7e
GOTO q7e
:l7e
ECHO Uh oh, Godzilla tries to itch himself in the sleep. Accidentally, he kills
ECHO you with his enormous paw.
PAUSE
CLS
REM Is the room returned?
l5w
:q7e

  IF "%1"=="godzilla" IF "%&l5w%"=="e7w" IF "%&f5w%"=="g7w" GOTO k7e
  GOTO a7e
  :k7e
  ECHO Godzilla is too busy fighting Mecha Godzilla to notice. It does seem to hurt
  ECHO a bit though.
  GOTO l5w
   :a7e

 IF "%1"=="mecha" IF "%&l5w%"=="e7w" GOTO j7e
 GOTO s7e

  REM Is the path returned?
  :j7e
  IF NOT "%&f5w%"=="g7w" GOTO h7e
  GOTO d7e
  :h7e
  ECHO The bullet bounces off his shields. He turns towards you, but luckily he

ECHO can't spot you. He continues to punch at the building.
GOTO l5w
:d7e
ECHO You get a good hit between two of his shields. He turns around and sees you.
ECHO Just as he is about to fire a missile from his knee, Godzilla comes up from
ECHO behind and grabs his head. The green monster twists the metal head until it

 ECHO falls off. Mecha Godzilla falls backwards, lifeless. Now Godzilla turns
 ECHO towards you. He gives out a roar that sound to you like 'Give me five!'. In
 ECHO order not to crush you, he gives you five in the air. He then starts walking
ECHO back through the city to the mountain...
ECHO.
PAUSE
CLS

  l5w

   :s7e
   IF "%1"=="" GOTO g7e
   GOTO f7e
    :g7e
    ECHO You fire a shot into the air, but no one seems to notice.

     GOTO l5w
     :f7e
     ECHO That would be pointless.
      :l5w
