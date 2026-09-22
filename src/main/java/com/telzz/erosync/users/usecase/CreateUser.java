package com.telzz.erosync.users.usecase;

import java.util.Optional;
import java.util.UUID;

import org.springframework.stereotype.Service;

import com.telzz.erosync.users.UserRepository;
import com.telzz.erosync.users.domain.EmailExistException;
import com.telzz.erosync.users.domain.User;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class CreateUser {
    private final UserRepository repo;

    public record Command(
            String name,
            String email,
            String password) {
    }

    public User execute(Command cmd) {
        Optional<User> existingUser = repo.findByEmail(cmd.email());
        if (existingUser.isPresent()) {
            throw new EmailExistException();
        }

        User user = User.create(UUID.randomUUID().toString(), cmd.name(), cmd.email(), cmd.password());
        repo.create(user);
        return user;
    }
}
