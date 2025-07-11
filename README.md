# music player client for music player daemon.

This client app is intended for use on android cell phones.

# Installation on android cell phones.

## Requisites.

* a workstation to which you can connect your cell phone via usb cable.
  Give cell phone sufficient authorizations so that you can install software
  via usb cable.
* install go toolchain and fyne app according to
  [this instructions](https://docs.fyne.io/started/packaging).
* install java sdk if you do not yet have a keystore to sign our app.
  We will keytool from it later to create a keystore.
* install bundletool and android-ndk from Google's android's developing kit.
* provide an app icon for your android cell phone.
* we also need git so that we can checkout app's source code.
  Install it as well.

## Instructions.

* checkout app's source code

	git clone https://github.com/raumanzug/gr-mpc

* cd into `gr-mpc` subdirectory

	cd gr-mpc

* make sure that you have a keystore for signing our app.  To do this, use
  keytool:  Customize filename `gpsc.jks` and alias `gpsc` to your needs:

	keytool -genkey -v -keystore "${HOME}/gpsc.jks" -alias gpsc -keyalg RSA -keysize 2048 -validity 10000

* compile our app: `mpc.png` is image file of app icon.

	fyne release -icon mpc.png -os android --key-store "${HOME}/gpsc.jks" --key-name gpsc

* extract `.apks` file from `.aab` file that we created when compiling
  our app:

	bundletool build-apks --bundle=mpc.aab --output=mpc.apks --ks="${HOME}/gpsc.jks" --ks-key-alias=gpsc

* install our app via usb cable:

	bundletool install-apks --apks mpc.apks

