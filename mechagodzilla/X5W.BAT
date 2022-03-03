@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
 IF "%&amg%"=="nhk" GOTO j8w
 ECHO This adventure game must be started properly first. See Readme.txt!
 GOTO l5w
 :j8w
 REM Path will be left
 IF "%1"=="the" GOTO h8w
 GOTO d8w
 :h8w
 EXAMINE %2
 :d8w

  SET &a5w=l5w
  IF "%&&l5w%"=="y7w" IF "%1"=="tank" GOTO g8w
  GOTO f8w
  :g8w
  ECHO It is a big, sturdy tank.
  IF "s3w"=="b5w" GOTO p8w
GOTO z8w

 :p8w
 ECHO There is some fuel, but no ammunition.
   :z8w

 IF "s3w"=="l5w" GOTO o8w
 GOTO x8w
 :o8w
 ECHO There is no fuel left, and no ammunition.
 :x8w
 SET &a5w=b5w
 :f8w
 IF "%&&l5w%"=="%&l5w%" IF "%1"=="tank" GOTO i8w
 GOTO c8w
 :i8w
 ECHO It is a big, sturdy tank.
 IF "s3w"=="b5w" GOTO u8w
 REM Is the constant engaged?
 GOTO v8w
 :u8w
 ECHO There is some fuel, but no ammunition.
:v8w
IF "s3w"=="l5w" GOTO y8w
 GOTO b8w
 :y8w

  ECHO There is no fuel left, and no ammunition.
  :b8w
  SET &a5w=b5w

   :c8w
   IF "%&&a5w%"=="y7w" IF "%1"=="turtle" GOTO t8w

 GOTO n8w
 :t8w
 ECHO It looks back at you with its big, glary eyes.
 SET &a5w=b5w
 :n8w
 IF "%&&a5w%"=="%&l5w%" IF "%1"=="turtle" GOTO r8w
  GOTO m8w

:r8w
ECHO It looks back at you with its big, glary eyes.
SET &a5w=b5w
:m8w
IF "%&&k5w%"=="y7w" IF "%1"=="gun" GOTO w8w
GOTO e8w
:w8w
ECHO It is big and heavy, but not too big and not too heavy. Being the most
 ECHO powerful handgun in the world, it has always served you well.
 IF "%&g5w%"=="b5w" GOTO l2w

  GOTO q2w

:l2w
ECHO It is loaded.
:q2w

 IF NOT "%&g5w%"=="b5w" GOTO k2w
 GOTO a2w
 :k2w
 ECHO It is empty.

  :a2w
   SET &a5w=b5w
 :e8w
 IF "%&&k5w%"=="%&l5w%" IF "%1"=="gun" GOTO j2w
 GOTO s2w
 :j2w
 ECHO It is big and heavy, but not too big and not too heavy. Being the most
 ECHO powerful handgun in the world, it has always served you well.
 IF "%&g5w%"=="b5w" GOTO h2w

GOTO d2w

 :h2w

  ECHO It is loaded.
  :d2w
  IF NOT "%&g5w%"=="b5w" GOTO g2w

   GOTO f2w
   :g2w
   ECHO It is empty.

    :f2w
  SET &a5w=b5w
  :s2w
  IF "%&&s5w%"=="y7w" IF "%1"=="napalm" GOTO p2w
  REM Is the room still returned?
  GOTO z2w

:p2w
 ECHO It's a small bottle containing a brownish substance. an ignition mechanism
 ECHO is fitted to the opening of the bottle, so that a flame will be emitted when
 ECHO the bottle is broken.

  SET &a5w=b5w

REM Is question engaged?
:z2w
IF "%&&s5w%"=="%&l5w%" IF "%1"=="napalm" GOTO o2w
GOTO x2w
:o2w

 REM Read sentence using constant
 ECHO It's a small bottle containing a brownish substance. an ignition mechanism
 ECHO is fitted to the opening of the bottle, so that a flame will be emitted when
 ECHO the bottle is broken.
 SET &a5w=b5w
 :x2w

   IF "%&a5w%"=="l5w" GOTO i2w
   GOTO c2w
   REM Show question
   :i2w

    ECHO I can't see what you are referring to.
    :c2w
    :l5w
