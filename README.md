# Rotational (AKA shift) Encrypton / Decryption #

## Introduction ##

A rotational encryption / decryption implementation.  A play toy for your
enjoyment.

**Note: This is a toy.  Do not use it in production or any serious environment.**

### Motivation ##

This project was enspired by:

* A general interest in cryptography.
* History.  See: [Caesar Cipher](https://en.wikipedia.org/wiki/Caesar_cipher)
* Reading.  Example: 
    * An Introduction to Mathematical Cryptography, by Jeffrey Hoffstein, 
Jill Pipher, and Joseph H. Silverman.
    * A PDF might be available at: [Intro PDF](https://epdf.pub/an-introduction-to-mathematical-cryptography.html)
* Having some fun coding.

## Technical Notes ##

* Written in the go language, and useful only in that environment.
* Uses go modules.  You will need a module capable version of go (go1.11+).
* Hosted on github.

## Installaton ##

Use *git clone* to obtain a copy.  Not available using *go get*.

Your project will need to have an appropriate *go.mod* to use red.

## Character Set ##

The only supported character set is UTF8.  Errors are returned if an alphabet,
plaintext or ciphertext parameter contains invalid characters.

## Shift / Rotation Values ##

### Shift Value 0 ###

Supported.  You get what you ask for in terms of encryption.

### Shift Value a Multiple of Alphabet Length ###

Supported.  Treated like a shift value of 0.

### Shift Values Greater Than 0 ###

All values are supported.  Note that by default this is a *left shift* of the
alphabet.

### Shift Values Less Than 0 ###

All values are supported.  Note that by default this is a *right shift* of the
alphabet.

### Shift Values Left Versus Right ###

Default is a positive shift value uses a left shift of the alphabet.

And a negative shift value uses a right shift of the alphabet.

Default behavior can be changed by calling the constructor
with an appropriate instance of *NewParms*.

## Examples ##

Examples are found in the subdirectories of the *cmd* directory.
