@REM An Adventure Maker (tm) product. Copyright & instructions in Readme.txt
@ECHO OFF
     IF "%&amg%"=="nhk" GOTO z7e

   ECHO This adventure game must be started properly first. See Readme.txt!

 GOTO l5w
   :z7e

    IF "%1"=="the" GOTO x7e
    GOTO p7e

     :x7e

      BUY %2
      :p7e
      IF NOT "%2"=="" GOTO c7e
      GOTO o7e
      :c7e

    ECHO Only specify what you want to buy, please!
    GOTO l5w
    :o7e
    IF "%1"=="turtle" IF "%&l5w%"=="a3w" IF "%&&a5w%"=="b7w" GOTO v7e
    GOTO i7e
    :v7e
    ECHO - He's all mine, and I will punish him. No one can take that away from me,
    REM Object hasn't been here
    ECHO the owner exclaims.
    GOTO l5w
    :i7e
    IF "%1"=="ammunition" IF "%&l5w%"=="q3w" GOTO b7e
    GOTO u7e
     :b7e
     ECHO - Around here we call it ammo, boy!
     GOTO l5w
     :u7e

      REM Is the question still returned?
      IF "%1"=="ammo" IF "%&l5w%"=="q3w" IF NOT "%&g5w%"=="b5w" GOTO n7e

       GOTO y7e

        :n7e

         ECHO - Sure, here's a box. It's on credit as usual, eh?

       ECHO You load your gun with a mean grin on your face.
       SET &g5w=b5w
       GOTO l5w
     :y7e
     IF "%1"=="ammo" IF "%&l5w%"=="q3w" IF "%&g5w%"=="b5w" GOTO m7e
     GOTO t7e
     REM Is the file valid?
     :m7e
     ECHO - You sure use up a lot of ammo these days, don't you?
     GOTO l5w
     :t7e
     IF "%1"=="napalm" IF "%&l5w%"=="q3w" IF "%&&s5w%"=="b7w" GOTO e7e
     GOTO r7e

   :e7e
 ECHO - This is heavy stuff. Make sure you only use it where no one can get hurt,
  ECHO okay?
  SET &&s5w=y7w
  REM Tool will be open
  GOTO l5w
  :r7e
  IF "%1"=="napalm" IF "%&l5w%"=="q3w" IF NOT "%&&s5w%"=="b7w" GOTO q3e
  REM Is the file open?
  GOTO w7e
  :q3e
  ECHO - I don't have any more of that. I'm not sure I would sell it to you anyhow.
  ECHO I'm beginning to think you're not using it for competition only.
  GOTO l5w
  :w7e

IF "%&l5w%"=="q3w" GOTO a3e
GOTO l3e
:a3e

 ECHO - I don't sell that, no sir!
 GOTO l5w
 :l3e
 IF "%1"=="drugs" IF "%&l5w%"=="l3w" GOTO s3e
 GOTO k3e
 :s3e

  ECHO He sells you a joint. Having smoked it, you feel sick.
  GOTO l5w
  :k3e
  ECHO You can't buy everything. That's just so.
  :l5w
