@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
   IF "%&amg%"=="nhk" GOTO b4w

    ECHO This adventure game must be started properly first. See Readme.txt!
    GOTO l5w
    :b4w
    IF "%2"=="about" GOTO n4w
    GOTO y4w
    :n4w
    ASK %1 %3 %4
    :y4w
     IF "%1"=="the" GOTO m4w
   GOTO t4w

    :m4w
    ASK %2 %3 %4 %5 %6
    :t4w
    IF "%1"=="to" GOTO e4w
    GOTO r4w
    :e4w

     ASK %2 %3 %4 %5 %6
     :r4w
     IF "%2"=="the" GOTO q6w
     GOTO w4w
     :q6w

      ASK %1 %3 %4
      REM Is object valid?
      :w4w

       IF "%&l5w%"=="l3w" GOTO a6w
       GOTO l6w
       :a6w

     IF "%1"=="man" GOTO s6w
     REM Set room using constant
     GOTO k6w
     :s6w
     IF "%2"=="mecha" GOTO d6w
     REM Open player using tool
     GOTO j6w
     :d6w
     REM Is constant open?
     ECHO - Nothing can stop that monster now. Except for Godzilla, of course!

   GOTO l5w

 :j6w
 IF "%2"=="godzilla" GOTO f6w
 GOTO h6w
 :f6w

  IF NOT "%&f5w%"=="l5w" GOTO z6w
   REM Use EOF
   GOTO g6w
   :z6w

 ECHO - We've talked about this before. Stop pestering me!
 GOTO l5w

  :g6w
   ECHO - Godzilla is asleep at the high mountain, far off to the west. Nothing but
   ECHO the scent of napalm could possibly bring him to consciousness again!
   SET &f5w=b5w
   GOTO l5w

    :h6w

     IF "%2"=="gun" IF "%&&k5w%"=="y7w" GOTO x6w
     GOTO p6w

      :x6w
      ECHO - It's a nice gun, but put it away, please!
      GOTO l5w
      :p6w

       IF "%2"=="turtle" IF "%&&a5w%"=="y7w" GOTO c6w
       GOTO o6w
       :c6w

     ECHO - That may look like a harmless turtle to you and me, but sometimes a
     ECHO harmless look can conceal a ferocious beast!
     GOTO l5w
     :o6w

      IF "%2"=="me" GOTO v6w
    GOTO i6w
    :v6w
    ECHO - You may have better qualities than your outside suggests.
    GOTO l5w
    :i6w

     ECHO - That's an interesting subject. I'm afraid I don't know enough about it.

   GOTO l5w

 :k6w

  :l6w
  IF "%&l5w%"=="a3w" GOTO b6w

   GOTO u6w
   :b6w
   IF "%1"=="owner" IF "%2"=="turtle" IF "%&&a5w%"=="b7w" GOTO n6w

 GOTO y6w

  :n6w
  ECHO - It has been bad, and therefore I must of course punish it, he says
  ECHO brandishing the knife.
  GOTO l5w
  :y6w
  IF "%1"=="owner" GOTO m6w

   GOTO t6w

 :m6w
 ECHO - I don't know much about that, mister!
 GOTO l5w
 :t6w
 IF "%1"=="turtle" IF "%&&a5w%"=="b7w" GOTO e6w
 GOTO r6w
 :e6w
 ECHO He doesn't say much. He is probably too scarred.
 GOTO l5w
  :r6w
  :u6w
  IF "%1"=="turtle" IF "%&&a5w%"=="y7w" GOTO q5e
  GOTO w6w
  :q5e
  ECHO He's seems to be the quiet kind of guy, ehhrmm turtle.
  GOTO l5w
  :w6w
  IF "%1"=="turtle" IF "%&&a5w%"=="%&l5w%" GOTO a5e
  GOTO l5e
  :a5e
  ECHO He's seems to be the quiet kind of guy, ehhrmm turtle.
  GOTO l5w
  :l5e
   IF "%1"=="man" IF "%&l5w%"=="q3w" GOTO s5e
   REM Until the player is engaged, unset constant!
   GOTO k5e

 :s5e

  ECHO - Action speaks louder than words, mister!

   GOTO l5w
   :k5e
   ECHO I can't see who you are referring to.

 :l5w
