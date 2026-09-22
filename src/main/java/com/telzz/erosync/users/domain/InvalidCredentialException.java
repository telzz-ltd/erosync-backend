package com.telzz.erosync.users.domain;

public class InvalidCredentialException extends RuntimeException {
    public InvalidCredentialException() {
        super("invalid credentials");
    }
}
