package com.telzz.erosync.users.domain;

public class EmailExistException extends RuntimeException {
    public EmailExistException() {
        super("Email aready exist");
    }
}
