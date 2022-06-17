bcrypt-password-checker
=======================

    Usage of bcrypt-password-checker:
    -cost int
            cost (aka rounds) (default 10)
    -file string
            file to scan for matching hashes, - for STDIN
    -hash string
            bcrypt hash
    -password string
            the password to check hash against

Installation
------------

    go install github.com/mxcu/bcrypt-password-checker@latest

Usage
-----

Hash a given password:

    # password from command line
    bcrypt-password-checker -password abc123xyz

    # read password from terminal ( -file != '-')
    bcrypt-password-checker

Check if password matches a given hash:

    bcrypt-password-checker -password abc123xyz -hash '$2a$10$HRw/ZX4iAZaBWnF.OYMYd.hOBsqeAxulYAHt/xuC/B17Ch5Ia16ji'

Check file for hashes, print out all hashes that match the given password:

    bcrypt-password-checker -password abc123xyz -file README.md

