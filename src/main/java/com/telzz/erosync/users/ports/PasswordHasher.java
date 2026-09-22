package com.telzz.erosync.users.ports;

public interface PasswordHasher {
    String hash(String password);

    Boolean compareHash(String hash, String password);
}
