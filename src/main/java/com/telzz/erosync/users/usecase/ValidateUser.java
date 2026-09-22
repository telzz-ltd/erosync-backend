package com.telzz.erosync.users.usecase;

import java.util.Optional;

import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import com.telzz.erosync.users.domain.InvalidCredentialException;
import com.telzz.erosync.users.domain.User;
import com.telzz.erosync.users.ports.UserRepository;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class ValidateUser {
    private final UserRepository repo;
    private final PasswordEncoder passwordHasher;

    public User execute(String email, String password) {
        Optional<User> user = repo.findByEmail(email);

        if (user.isPresent() && passwordHasher.matches(password, user.get().getPasswordHash())) {
            return user.get();
        }

        throw new InvalidCredentialException();
    }
}
