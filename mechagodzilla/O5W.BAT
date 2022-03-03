@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
    IF "%&amg%"=="nhk" GOTO v2w

  ECHO This adventure game must be started properly first. See Readme.txt!

   GOTO l5w
   :v2w

    IF "%1"=="at" GOTO b2w
    GOTO u2w

  :b2w
  REM Is the file still open?
  EXAMINE %2 %3
  REM Is question considered?
  :u2w

   IF NOT "%&l5w%"=="y7w" GOTO n2w
    ECHO Me

     GOTO y2w
     :n2w
     IF NOT "%&l5w%"=="b7w" GOTO t2w
     ECHO Limbo 1
     GOTO y2w
   :t2w
   IF NOT "%&l5w%"=="n7w" GOTO m2w

    ECHO Limbo 2
     GOTO y2w
     :m2w
     IF NOT "%&l5w%"=="v7w" GOTO r2w

      REM Is sentence engaged?
      ECHO You are in your own comfortable home. It is a small house right in the
      ECHO centre of Tokyo. In here you always felt safe...until Mecha Godzilla came to
      ECHO earth!
      ECHO There is a door leading outside to the north.
      GOTO y2w
    :r2w

  IF NOT "%&l5w%"=="t7w" GOTO e2w
  ECHO You are in the street outside your house. Panicking people run around
  ECHO everywhere, screaming in extreme terror. Far off to the east, you can see
  ECHO the silhouette of Mecha Godzilla trampling through the city.

IF "%&f5w%"=="g7w" GOTO q0w
GOTO w2w
:q0w
ECHO Godzilla is there too, fighting him.
   :w2w
   ECHO The street goes west and east from here.
   GOTO y2w
   :e2w
   IF NOT "%&l5w%"=="r7w" GOTO l0w

 IF "%&d5w%"=="o7w" GOTO k0w
 GOTO a0w
 :k0w

  ECHO You are in the tank.
:a0w

 IF NOT "%&d5w%"=="o7w" GOTO j0w
 REM Mark path with value
 GOTO s0w
 :j0w
 ECHO You are standing in the street, east of your home.
 :s0w

  REM Move question
  ECHO The street leads both east and west, but a building has caught fire to the
  ECHO east, and the street is covered with dead people and debris, as they have
  ECHO tried to save themselves and their belongings by jumping out the windows.
   REM If the object is removed, renew variable!
   ECHO There is a garage to the north. Far off to the east, you can see the

    ECHO silhouette of Mecha Godzilla trampling through the city.

  IF "%&f5w%"=="g7w" GOTO h0w
  REM Is the object considered?
  GOTO d0w

   REM Is the keyword considered?
   :h0w
   ECHO Godzilla is there too, fighting him.

 :d0w

  GOTO y2w
  :l0w
  IF NOT "%&l5w%"=="u7w" GOTO f0w
  IF "%&d5w%"=="o7w" GOTO z0w
  GOTO g0w
:z0w
ECHO You are in the tank, in the garage.
:g0w
IF NOT "%&d5w%"=="o7w" GOTO x0w
GOTO p0w
:x0w
ECHO You are in the garage. All the cars are gone, because of fleeing people.
:p0w
IF "%&&l5w%"=="u7w" IF NOT "%&d5w%"=="o7w" GOTO c0w
GOTO o0w

 :c0w

  ECHO There is a military tank standing in one corner of the garage.

:o0w

 GOTO y2w
 :f0w

  IF NOT "%&l5w%"=="e7w" GOTO i0w
  IF NOT "%&f5w%"=="g7w" GOTO u0w
  GOTO v0w
  :u0w
  ECHO You are standing right next to the place where Mecha Godzilla is punching at
   ECHO a tall building.
   :v0w

 IF "%&f5w%"=="g7w" GOTO y0w
 GOTO b0w
 REM Room hasn't been present
 :y0w
 ECHO You are right at the battle scene. Godzilla is fighting Mecha Godzilla, but

  ECHO the battle is uneven. As long as Mecha Godzilla is fully concentrated on the

ECHO fight, Godzilla can't get close enough to break his neck.
:b0w

  REM Is object returned?
  GOTO y2w
  :i0w
  IF NOT "%&l5w%"=="m7w" GOTO n0w
   ECHO You are standing in a crossing. To the east, you can see your house. There
   ECHO are also streets going north and west from here. An ammunition shop lies to
   ECHO the south. Far off to the east, you can see the silhouette of Mecha Godzilla
   ECHO trampling through the city.

 IF "%&f5w%"=="g7w" GOTO m0w

  GOTO t0w
  :m0w
  ECHO Godzilla is there too, fighting him.
  :t0w
  GOTO y2w
  :n0w
  IF NOT "%&l5w%"=="w7w" GOTO r0w
  REM Is door left?
  ECHO You are in the street leading north from the street where you live. The
  ECHO street continues to the north. West of here, there is a petshop. A great
  ECHO crowd has gathered to the north, looting a warehouse.
  GOTO y2w
  :r0w
  IF NOT "%&l5w%"=="l3w" GOTO e0w
  ECHO You are in the street west of your home. There is an old man standing here.
  ECHO Far off to the east, you can see the silhouette of Mecha Godzilla trampling
  ECHO through the city.

IF "%&f5w%"=="g7w" GOTO q4w

 GOTO w0w
 :q4w
 ECHO Godzilla is there too, fighting him.
 REM Hide player using delimiter
 :w0w
 GOTO y2w
 :e0w

IF NOT "%&l5w%"=="a3w" GOTO l4w
ECHO You are in a small pet shop.
IF "%&&a5w%"=="b7w" GOTO k4w
GOTO a4w
:k4w
ECHO The owner is standing behind the counter, holding a small turtle in one
ECHO hand, and a big knife in the other.
:a4w

 IF NOT "%&&a5w%"=="b7w" GOTO j4w
 GOTO s4w
 :j4w
  ECHO The owner is nowhere to be seen.

   :s4w
   GOTO y2w
   :l4w
   IF NOT "%&l5w%"=="q3w" GOTO d4w
   ECHO You are in the ammunition shop. The man behind the counter is smiling at
 ECHO you.

GOTO y2w
:d4w
  IF NOT "%&l5w%"=="k3w" GOTO h4w

ECHO You are at a high mountain. Godzilla is lying here, sleeping. Far off to the
ECHO east, you can see the silhouette of Mecha Godzilla trampling through the
ECHO city.
GOTO y2w
:h4w
:y2w
 SET &a5w=l5w
 ECHO You can see:

   IF "%&&l5w%"=="%&l5w%" GOTO g4w
   GOTO f4w
   REM Keyword is not engaged
   :g4w
   ECHO tank
   SET &a5w=b5w

 :f4w
 IF "%&&a5w%"=="%&l5w%" GOTO p4w

GOTO z4w
:p4w
ECHO turtle
SET &a5w=b5w
:z4w
IF "%&&k5w%"=="%&l5w%" GOTO o4w

 GOTO x4w
 :o4w
 ECHO gun
 SET &a5w=b5w
 :x4w
 IF "%&&s5w%"=="%&l5w%" GOTO i4w

  GOTO c4w

:i4w
ECHO napalm
SET &a5w=b5w
:c4w
IF "%&a5w%"=="l5w" GOTO u4w
GOTO v4w
:u4w
ECHO Nothing of interest.
:v4w
:l5w
